package main

import (
	"github.com/caimmy/jungle"
	"github.com/caimmy/beacontower/ucenter/bt.io"
)

type WebEntryJungle struct {
	jungle.JungleController
}

func (j *WebEntryJungle) Get()  {
	j.RenderPartial("ws_client.html", nil)
}

func main()  {
	entry_controller := &WebEntryJungle{}
	jungle.Router("/", entry_controller)
	jungle.WebsocketRouter("/echo", bt_io.EchoHandler)
	jungle.StaticFilePath = "D:\\caimmy\\code\\golang\\src\\github.com\\caimmy\\beacontower\\ucenter\\static"
	jungle.TemplatesPath = "D:\\caimmy\\code\\golang\\src\\github.com\\caimmy\\beacontower\\ucenter\\templates"
	jungle.Run("localhost:8000")
}
