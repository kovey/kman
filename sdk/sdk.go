package sdk

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const version = "1.0.0"

type Meta struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

type Config struct {
	Addrs           []string
	Username        string
	Password        string
	CachePath       string
	Namespace       string
	UseCacheWhenErr bool
}

type WatchFunc func(Meta)

type sdk struct {
	namespace       string
	version         string
	prefix          string
	config          Config
	cli             *clientv3.Client
	watchers        []*watcher
	wait            sync.WaitGroup
	caches          *caches
	useCacheWhenErr bool
}

func (s *sdk) init(conf Config) error {
	etcdConfig := clientv3.Config{
		Endpoints:   conf.Addrs,
		Username:    conf.Username,
		Password:    conf.Password,
		DialTimeout: 30 * time.Second,
	}
	s.namespace = conf.Namespace
	s.caches.path = conf.CachePath
	s.caches.load(s.namespace)
	s.caches.get(s.namespace)
	s.useCacheWhenErr = conf.UseCacheWhenErr
	s.prefix = fmt.Sprintf("/ko/configs/%s", s.namespace)
	cli, err := clientv3.New(etcdConfig)
	if err != nil {
		return err
	}

	s.cli = cli
	return nil
}

func (s *sdk) key(key string) string {
	return fmt.Sprintf("%s/%s", s.prefix, key)
}

func (s *sdk) realKey(key string) string {
	return strings.ReplaceAll(key, s.prefix+"/", "")
}

func (s *sdk) get(ctx context.Context, key string) (*Meta, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("sdk not init")
	}
	resp, err := s.cli.Get(ctx, s.key(key), clientv3.WithFirstKey()...)
	if err != nil {
		if s.useCacheWhenErr {
			if meta := s.caches.get(s.namespace).get(key); meta != nil {
				return meta, nil
			}
		}

		return nil, err
	}

	if resp.Count == 0 {
		return &Meta{Key: key}, nil
	}

	kv := resp.Kvs[0]
	meta := &Meta{Key: s.realKey(string(kv.Key)), Value: string(kv.Value)}
	s.caches.get(s.namespace).update(meta.Key, meta.Value)
	return meta, nil
}

func (s *sdk) all(ctx context.Context, key string) ([]Meta, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("sdk not init")
	}
	resp, err := s.cli.Get(ctx, s.key(key), clientv3.WithPrefix())
	if err != nil {
		if s.useCacheWhenErr {
			if meta := s.caches.get(s.namespace).get(key); meta != nil {
				return []Meta{*meta}, nil
			}
		}
		return nil, err
	}

	var res = make([]Meta, len(resp.Kvs))
	for index, kv := range resp.Kvs {
		res[index] = Meta{Key: s.realKey(string(kv.Key)), Value: string(kv.Value)}
		s.caches.get(s.namespace).update(res[index].Key, res[index].Value)
	}

	return res, nil
}

func (s *sdk) close() {
	for _, w := range s.watchers {
		w.shutdown()
	}

	if s.cli != nil {
		s.cli.Close()
	}

	s.caches.save()
	s.wait.Wait()
}

func (s *sdk) listen(key string, create, update, delete WatchFunc) {
	w := newWatcher(s.key(key), create, update, delete)
	s.watchers = append(s.watchers, w)
	s.wait.Add(1)
	go s._listen(w)
}

func (s *sdk) _listen(w *watcher) {
	defer s.wait.Done()
	w.listen(s.namespace)
}

var s = &sdk{version: version, wait: sync.WaitGroup{}, caches: newCaches()}

func Init(config Config) error {
	return s.init(config)
}

func Get(ctx context.Context, key string) (*Meta, error) {
	return s.get(ctx, key)
}

func All(ctx context.Context, key string) ([]Meta, error) {
	return s.all(ctx, key)
}

func Listen(key string, create, update, delete WatchFunc) {
	s.listen(key, create, update, delete)
}

func Version() string {
	return s.version
}

func Close() {
	s.close()
}
