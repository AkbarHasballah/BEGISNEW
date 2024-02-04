package module

import (
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/handler"
	"github.com/aiteung/module/model"
)

func Caller(Modulename string, Pesan model.IteungMessage) (reply string) {
	switch Modulename {
	case "tokengrup":
		reply = handler.TokenGroup(Pesan)
	}
	return
}

func CallAndSend(Modulename string, Pesan model.IteungMessage, URLApiWa string) (resp atmessage.Response, errormessage string) {
	reply := Caller(Modulename, Pesan)
	var msg = model.GowaNotif{
		User:     Pesan.Chat_number,
		Server:   Pesan.Chat_server,
		Messages: reply,
	}
	resp, errormessage = SendToGoWAAPI(msg, URLApiWa)
	return
}
