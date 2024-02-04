package module

import (
	"fmt"
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atdb"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/helper"
	"github.com/aiteung/module/model"
	"github.com/aiteung/musik"
	"go.mongodb.org/mongo-driver/mongo"
)

func IteungModuleCall(WAIface model.IteungWhatsMeowConfig, DBIface model.IteungDBConfig) (Modulename string, Pesan model.IteungMessage) {
	Pesan = Whatsmeow2Struct(WAIface)
	NormalizeAndTypoCorrection(&Pesan.Message, DBIface.MongoConn, DBIface.TypoCollection)
	if IsIteungCall(Pesan) {
		Modulename = GetModuleName(Pesan, DBIface.MongoConn, DBIface.ModuleCollection)
	}
	return
}

func Whatsmeow2Struct(WAIface model.IteungWhatsMeowConfig) (im model.IteungMessage) {
	im.Phone_number = helper.GetPhoneNumber(WAIface)
	im.Chat_number = WAIface.Info.Chat.User
	im.Chat_server = WAIface.Info.Chat.Server
	im.Alias_name = WAIface.Info.PushName
	im.Message = helper.GetMessage(WAIface.Message)
	im.From_link = helper.GetStatusFromLink(WAIface)
	if im.From_link {
		im.From_link_delay = helper.GetFromLinkDelay(WAIface.Message)
	}
	im.Filename, im.Filedata = helper.GetFile(WAIface.Waclient, WAIface.Message)
	im.Longitude, im.Latitude, im.LiveLoc = helper.GetLongLat(WAIface.Message)
	if WAIface.Info.Chat.Server == "g.us" {
		groupInfo, err := WAIface.Waclient.GetGroupInfo(WAIface.Info.Chat)
		if err != nil {
			fmt.Println("cek err : ", err)
		}
		if groupInfo != nil {
			im.Group = groupInfo.GroupName.Name + "@" + WAIface.Info.Chat.User
			im.Group_name = groupInfo.GroupName.Name
			im.Group_id = WAIface.Info.Chat.User
		}
		im.Is_group = true
	}
	return
}

func IsIteungCall(im model.IteungMessage) bool {
	if (strings.Contains(im.Message, "teung") && im.Chat_server == "g.us") || (im.Chat_server == "s.whatsapp.net") {
		return true
	} else {
		return false
	}
}

func GetModuleName(im model.IteungMessage, MongoConn *mongo.Database, ModuleCollection string) (modulename string) {
	modules := atdb.GetAllDoc[[]model.Module](MongoConn, ModuleCollection)
	for _, mod := range modules {
		complete, _ := musik.IsMatch(strings.ToLower(im.Message), mod.Keyword...)
		if complete {
			modulename = mod.Name
		}
	}
	return
}

func SendToIteungAPI(pesan model.IteungMessage, urltarget string) (response atmessage.Response, errormessage string) {
	response, errormessage = atapi.PostStruct[atmessage.Response](pesan, urltarget)
	return
}

func SendToGoWAAPI(pesan model.GowaNotif, urltarget string) (response atmessage.Response, errormessage string) {
	response, errormessage = atapi.PostStruct[atmessage.Response](pesan, urltarget)
	return
}

func DocumentSendToGoWAAPI(pesan model.IteungMessage, urltarget string) (response atmessage.Response, errormessage string) {
	response, errormessage = atapi.PostStruct[atmessage.Response](pesan, urltarget)
	return
}
