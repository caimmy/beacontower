package main

import (
	"fmt"
	"os"
	"github.com/nsqio/go-nsq"
)

var (
	NSQD_ADDR = "127.0.0.1:4150"
)

func sendMsg(content string)  {
	producer, err := nsq.NewProducer(NSQD_ADDR, nsq.NewConfig())
	if err == nil {
		err = producer.Publish("nsq_demo", []byte(content))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("send msg success")
		}
	}
}

func main()  {
	if len(os.Args) == 2 {
		sendMsg(os.Args[1])
	}
}