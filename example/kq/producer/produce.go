package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
)

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	pusher := kq.NewPusher([]string{
		"127.0.0.1:9093",
	}, "kq")

	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		<-ticker.C

		count := rand.Intn(100)
		m := message{
			Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
			Value:   fmt.Sprintf("%d,%d", round, count),
			Payload: fmt.Sprintf("%d,%d", round, count),
		}
		body, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}

		//if err := pusher.Push(context.Background(), string(body)); err != nil {
		//	log.Fatal(err)
		//}

		headers := []kafka.Header{
			{
				Key:   "retry-count",
				Value: []byte("2"),
			},
		}

		if err := pusher.KPush(kq.WithHeaders(context.Background(), headers), "test", body); err != nil {
			log.Fatal(err)
		}
	}

	cmdline.EnterToContinue()
}
