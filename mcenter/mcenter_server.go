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

// date     : 2018/4/10 16:23
// author   : caimmy@hotmail.com

package main

import (
	"time"
	"fmt"
	"github.com/caimmy/beacontower/rpcdatas"
	"net/rpc"
	"net"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"sync"
	"github.com/caimmy/beacontower/mcenter/lib"
	"github.com/caimmy/beacontower/mcenter/mconsumers"
)

var (
	NSQ_SERVER = "192.168.5.17:4150"
	NSQ_LOOKUP = "192.168.5.17:4161"

	msg_publisher lib.BeaconioMsgQueuePublisher

	once sync.Once
)

func continue_print() {
	m := 1
	for ; ;  {
		time.Sleep(5 * time.Second)
		fmt.Println(m)
	}
}



type MessageChannel struct {
	publisher 			lib.BeaconioMsgQueuePublisher
}

/**
推送消息入栈
 */
func (mc * MessageChannel) PushData(md *rpcdatas.MsgInjection, reply *int) error {
	err := msg_publisher.Append(md.Sender, md.Content)
	if err == nil {
		*reply = 0
	} else {
		*reply = 1
	}
	return err
}

func init()  {
	fmt.Println("module initialized...")
	once.Do(func() {
		msg_publisher = lib.GetMsgQueuePublisher("nsq", NSQ_SERVER)
	})
}

func main()  {

	startMsgConsumerOnWebsocket()

	msg_channel_service := new(MessageChannel)
	rpc.Register(msg_channel_service)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error : " , e)
	}
	http.Serve(l, nil)
}

func startMsgConsumerOnWebsocket() {
	worker1 := lib.GetMsgConsumerViaNSQD(NSQ_SERVER, "websocket", "dispatch", nil)
	worker1.RegisterCallbackHandler(&mconsumers.WebsocketMsgWorker{})
	worker1.Run()
}