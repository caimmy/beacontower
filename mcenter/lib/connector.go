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

// date     : 2018/4/12 14:13
// author   : caimmy@hotmail.com

package lib

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type BeaconioMsgQueuePublisher interface {
	Init()
	Append(topic string, content []byte) error
	Close()
}

type NsqMsgQueue struct {
	Host 			string
	producer 		*nsq.Producer
}

func (this *NsqMsgQueue) Init() {
	var err error
	this.producer, err = nsq.NewProducer(this.Host, nsq.NewConfig())
	if err != nil {
		log.Fatal("err occur: ", err)
	}
}

/**
向消息队列中发布消息
 */
func (this *NsqMsgQueue) Append(topic string, content []byte) error {
	return this.producer.Publish(topic, content)
}

func (this *NsqMsgQueue) Close() {
	this.producer.Stop()
}

/**
构造发布消息的连接器
 */
func GetMsgQueuePublisher(msg_type string, msg_server string) BeaconioMsgQueuePublisher {
	var ret_if BeaconioMsgQueuePublisher
	ret_if = nil
	switch msg_type {
	case "nsq":
		msg_connector := NsqMsgQueue{msg_server, nil}
		msg_connector.Init()
		ret_if = &msg_connector
	}
	return ret_if
}