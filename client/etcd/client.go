package etcd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kovey/cli-go/env"
	"github.com/kovey/kman/client/cache"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var cli *clientv3.Client
var wts = newWatchers()

func Init() error {
	timeout, _ := env.GetInt("ETCD_TIMEOUT")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(os.Getenv("ETCD_ENDPOINTS"), ","),
		DialTimeout: time.Duration(timeout) * time.Second,
		Username:    os.Getenv("ETCD_USERNAME"),
		Password:    os.Getenv("ETCD_PASSWORD"),
	})
	if err != nil {
		return err
	}

	cli = c
	if err := cache.Init(); err != nil {
		return err
	}

	return _init()
}

func _init() error {
	for _, ns := range strings.Split(os.Getenv("CONFIG_NAMESPACE"), ",") {
		ns = strings.Trim(ns, " ")
		prefix := fmt.Sprintf("/ko/configs/%s", ns)
		resp, err := Get(context.Background(), prefix, clientv3.WithPrefix())
		if err != nil {
			return err
		}

		fmt.Println(prefix, resp.Kvs)
		for _, kv := range resp.Kvs {
			key := strings.ReplaceAll(string(kv.Key), prefix+"/", "")
			cache.Add(ns, key, string(kv.Value))
			cache.AddToFile(ns, key, string(kv.Value))
		}

	}

	return cache.Parse()
}

func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return cli.Get(ctx, key, opts...)
}

func Watch() {
	wts.watch()
}

func Close() {
	wts.shutdown()
	if cli == nil {
		return
	}

	cli.Close()
}
