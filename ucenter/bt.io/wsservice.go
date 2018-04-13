package bt_io

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"net/rpc"
	"sync"
	"github.com/caimmy/beacontower/rpcdatas"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}, EnableCompression: true}

	rpc_client *rpc.Client

	once sync.Once
)

func init() {
	once.Do(func() {
		var err error
		rpc_client, err = rpc.DialHTTP("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("rpc client initialize failure : ", err)
		}
	})
	fmt.Println("rpc client initialize success")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err == nil {
		for {
			mt, message, err := conn.ReadMessage()
			log.Println("read ", mt, message)
			if err != nil {
				log.Println("read error: ", err)
				break
			}
			log.Println("recv: ", string(message))
			wx_push_data := &rpcdatas.MsgInjection{"websocket", message}
			var ret_push int
			rpc_client.Call("MessageChannel.PushData", wx_push_data, &ret_push)
			fmt.Println("RPC CALL RESULT : ", ret_push)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write: ", err)
				break
			}
		}
	} else {
		fmt.Println(err)
	}
}