package main

import (
	"github.com/caimmy/jungle"
	"github.com/caimmy/beacontower/ucenter/bt.io"
)


type WebEntryJungle struct {
	jungle.JungleController
}

func (j *WebEntryJungle) Get()  {
	j.Echo("<h2>welcome to BeaconTower!</h2>", true)
}

func main()  {
	entry_controller := &WebEntryJungle{}
	jungle.Router("/", entry_controller)
	jungle.WebsocketRouter("/echo", bt_io.EchoHandler)
	jungle.Run("localhost:8000")
}
