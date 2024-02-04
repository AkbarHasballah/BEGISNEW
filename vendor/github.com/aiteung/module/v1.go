package module

import (
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
)

func IteungV1(IteungIPAddress string, apikey string, Pesan model.IteungMessage) {
	var im atmessage.IteungMessage
	im.Phone_number = Pesan.Phone_number
	im.Group_name = Pesan.Group_name + "@" + Pesan.Chat_number
	im.Alias_name = Pesan.Alias_name
	im.Messages = Pesan.Message
	if Pesan.Is_group {
		im.Is_group = "true"
	} else {
		im.Is_group = "false"
	}
	im.Filename = Pesan.Filename
	im.Filedata = Pesan.Filedata
	im.Latitude = Pesan.Latitude
	im.Longitude = Pesan.Longitude
	im.Api_key = apikey
	SendToIteungV1API(IteungIPAddress, im)
}

func SendToIteungV1API(IteungIPAddress string, pesan atmessage.IteungMessage) (response atmessage.IteungRespon, errormessage string) {
	urltarget := "http://" + IteungIPAddress + "/iteung/chatbot"
	response, errormessage = atapi.PostStruct[atmessage.IteungRespon](pesan, urltarget)
	return
}
