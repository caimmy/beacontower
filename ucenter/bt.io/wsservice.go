package bt_io

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}, EnableCompression: true}

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