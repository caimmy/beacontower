package main

import (
	"github.com/nsqio/go-nsq"
	"fmt"
	"time"
)

const (
	nsq_server = "127.0.0.1:4150"
	nsq_lookup = "127.0.0.1:4161"
)

var (
	exit_cnt = 0
	consumer_server *nsq.Consumer
)

type ConsumerT struct{}

func publishMsg(msg string) {
	producer, err := nsq.NewProducer(nsq_server, nsq.NewConfig())
	if err == nil {
		err = producer.Publish("test_abc", []byte(msg))
		fmt.Println(err)
	}
}

func (c *ConsumerT) HandleMessage(msg *nsq.Message) error {
	recv_msg := string(msg.Body)
	fmt.Println("receive: ", msg.NSQDAddress, "message: ", recv_msg)
	exit_cnt += 1
	fmt.Printf("recved %d times\n", exit_cnt)
	if exit_cnt == 3 {
		consumer_server.StopChan <- 9527
	}
	return nil
}

func consumeMsg(topic string, channel string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	var err error
	consumer_server, err = nsq.NewConsumer(topic, channel, cfg)
	if (err == nil) {
		consumer_server.SetLogger(nil, 0)
		consumer_server.AddHandler(&ConsumerT{})
	}
	if err = consumer_server.ConnectToNSQLookupd(nsq_lookup); err != nil {
		panic(err)
	}
	m := <- consumer_server.StopChan
	fmt.Println(m)
	fmt.Println("consumer exited!")
}

func main() {
	//publishMsg("hello fouth")
	consumeMsg("nsq_demo", "dododo")
}
