package main

import (
	"fmt"
	"local-pubsub/pubsub/localpubsub"
	"log"
	"time"
)

func main() {
	start := time.Now()
	pubsub := localpubsub.NewPubSub()
	pubsub.Run()
	sub1 := pubsub.Subscribe("topic_1")
	sub2 := pubsub.Subscribe("topic_1")
	_ = pubsub.Subscribe("topic_2")

	_ = pubsub.Publish("topic_1", localpubsub.NewMessage("Hello Loc Pham"))

	go func() {
		for {
			data := <-sub1
			fmt.Println("data sub 1", data.Data.(string))
		}
	}()

	go func() {
		for {
			data := <-sub2
			fmt.Println("data sub 2", data.Data.(string))
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Printf("pubsub %+v \n", pubsub)

	log.Println("time", time.Since(start))
}
