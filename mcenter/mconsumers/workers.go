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

// date     : 2018/4/12 17:59
// author   : caimmy@hotmail.com

package mconsumers

import (
	"github.com/nsqio/go-nsq"
	"fmt"
)

type WebsocketMsgWorker struct{}

func (c *WebsocketMsgWorker) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}