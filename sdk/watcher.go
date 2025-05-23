package sdk

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type watcher struct {
	cancel context.CancelFunc
	c      clientv3.WatchChan
	key    string
	create WatchFunc
	delete WatchFunc
	update WatchFunc
}

func newWatcher(key string, create, update, delete WatchFunc) *watcher {
	return &watcher{key: key, create: create, update: update, delete: delete}
}

func (w *watcher) listen(namespace string) {
	ctx, cancel := context.WithCancel(context.Background())
	w.c = s.cli.Watch(ctx, w.key, clientv3.WithPrefix())
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
					key := s.realKey(string(ev.Kv.Key))
					value := string(ev.Kv.Value)
					if ev.IsCreate() && w.create != nil {
						s.caches.get(namespace).add(key, value)
						w.create(Meta{Key: key, Value: value})
					} else if ev.IsModify() && w.update != nil {
						s.caches.get(namespace).update(key, value)
						w.update(Meta{Key: key, Value: value})
					}
				case clientv3.EventTypeDelete:
					if w.delete != nil {
						key := s.realKey(string(ev.Kv.Key))
						s.caches.get(namespace).delete(key)
						w.delete(Meta{Key: key, Value: string(ev.Kv.Value)})
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
