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

// date     : 2018/1/5 11:13
// author   : caimmy@hotmail.com

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/caimmy/beacontower/bt.io/wservice"
	"flag"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main()  {
	fmt.Println("bt.io started!")
	log.SetFlags(0)

	http.HandleFunc("/echo", wservice.EchoServer)
	log.Fatal(http.ListenAndServe(*addr, nil))

}

