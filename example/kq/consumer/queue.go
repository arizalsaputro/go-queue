package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var c kq.KqConf
	conf.MustLoad("config.yaml", &c)

	q := kq.MustNewQueue(c, kq.WithHandle(func(ctx context.Context, key string, value []byte) error {
		fmt.Printf("%s => %s\n", key, value)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
