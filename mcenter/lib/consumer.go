// Copyright 2017 jungle Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// date     : 2018/4/12 17:24
// author   : caimmy@hotmail.com

package lib

import (
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

/**
承载消息回调处理器的结构体
 */

type MsgConsumer struct {
	Host 					string 				// 消息服务器的地址
	Topic 					string				// 关注消息的主题
	Channel					string				// 处理消息的管道
	CallbackHandler			nsq.Handler			// 处理消息的结构体

	inRunning 				bool				// 消费worker是否在运行中
	lock 					sync.Mutex
}

func (this *MsgConsumer) RegisterCallbackHandler(handler nsq.Handler) {
	this.CallbackHandler = handler
}

func (this *MsgConsumer) startConsumption() {
	consumer_cfg := nsq.NewConfig()
	consumer_cfg.LookupdPollInterval	= time.Second
	consumer, err := nsq.NewConsumer(this.Topic, this.Channel, consumer_cfg)
	if err == nil {
		consumer.SetLogger(nil, 0)
		consumer.AddHandler(this.CallbackHandler)

		if err = consumer.ConnectToNSQD(this.Host); err != nil {
			panic(err)
		}
		this.inRunning = true
		<- consumer.StopChan
		consumer.DisconnectFromNSQD(this.Host)
		this.inRunning = false
	}
}

func (this *MsgConsumer) Run() error {
	this.lock.Lock()
	if !this.inRunning {
		go this.startConsumption()
	}
	return nil
}


/**
构造一个消息消费器
 */
func GetMsgConsumerViaNSQD(host string, topic string, channel string, handler nsq.Handler) *MsgConsumer {
	msg_consumer := MsgConsumer{host, topic, channel, handler, false, sync.Mutex{}}
	return &msg_consumer
}