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

type ConsumerT struct{}

func publishMsg(msg string) {
	producer, err := nsq.NewProducer(nsq_server, nsq.NewConfig())
	if err == nil {
		err = producer.Publish("test_abc", []byte(msg))
		fmt.Println(err)
	}
}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive: ", msg.NSQDAddress, "message: ", string(msg.Body))
	return nil
}

func consumeMsg(topic string, channel string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if (err == nil) {
		c.SetLogger(nil, 0)
		c.AddHandler(&ConsumerT{})
	}
	if err = c.ConnectToNSQLookupd(nsq_lookup); err != nil {
		panic(err)
	}
}

func main() {
	//publishMsg("hello fouth")
	consumeMsg("test_abc", "dododo")
	for {
		time.Sleep(time.Second * 10)
	}
}
