package etcd

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/kovey/cli-go/env"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var cli *clientv3.Client

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
	return nil
}

func Put(ctx context.Context, key, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return cli.Put(ctx, key, value, opts...)
}

func Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return cli.Get(ctx, key, opts...)
}

func Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	return cli.Delete(ctx, key, opts...)
}

func Compact(ctx context.Context, rev int64, opts ...clientv3.CompactOption) (*clientv3.CompactResponse, error) {
	return cli.Compact(ctx, rev, opts...)
}

func Close() {
	if cli == nil {
		return
	}

	cli.Close()
}
