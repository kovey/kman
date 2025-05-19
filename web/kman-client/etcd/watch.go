package etcd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kovey/debug-go/debug"
	"github.com/kovey/kman/kman-client/cache"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type watcher struct {
	cancel    context.CancelFunc
	prefix    string
	namespace string
	c         clientv3.WatchChan
}

func newWatcher(namespace string) *watcher {
	return &watcher{namespace: namespace, prefix: fmt.Sprintf("/ko/configs/%s", namespace)}
}

func (w *watcher) listen() {
	ctx, cancel := context.WithCancel(context.Background())
	w.c = cli.Watch(ctx, w.prefix, clientv3.WithPrefix())
	w.cancel = cancel
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-w.c:
			if data.Canceled {
				return
			}

			for _, ev := range data.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					key := strings.ReplaceAll(string(ev.Kv.Key), w.prefix+"/", "")
					value := string(ev.Kv.Value)
					if err := cache.Add(w.namespace, key, value); err != nil {
						debug.Erro("flush key[%s], value[%s] failure: %s", key, value, err)
					}
				}
			}
		}
	}
}

func (w *watcher) shutdown() {
	if w.cancel == nil {
		return
	}

	w.cancel()
}

type watchers struct {
	ws   map[string]*watcher
	wait sync.WaitGroup
}

func newWatchers() *watchers {
	return &watchers{ws: make(map[string]*watcher), wait: sync.WaitGroup{}}
}

func (w *watchers) watch() {
	namespaces := strings.Split(os.Getenv("CONFIG_NAMESPACE"), ",")
	for _, ns := range namespaces {
		ns = strings.Trim(ns, " ")
		wt := newWatcher(ns)
		w.wait.Add(1)
		w.ws[ns] = wt
		go w.listen(wt)
	}
	w.wait.Wait()
}

func (w *watchers) listen(wt *watcher) {
	defer w.wait.Done()
	wt.listen()
}

func (w *watchers) shutdown() {
	for _, wt := range w.ws {
		wt.shutdown()
	}
}
