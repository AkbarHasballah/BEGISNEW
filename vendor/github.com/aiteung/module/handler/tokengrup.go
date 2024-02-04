package handler

import (
	"os"

	"github.com/aiteung/module/model"
	"github.com/whatsauth/watoken"
)

func TokenGroup(Pesan model.IteungMessage) string {
	group_id := Pesan.Group_id
	token, _ := watoken.EncodeforHours(group_id, os.Getenv("PRIVATEKEY"), 4380)
	reply := "Hai.. hai.. ini dia token group kaka :\n" + token + "\nsimpan dan catat baik baik ya.. jangan sampai ilang lho berlaku selama 6 bulan tokennya"
	return reply
}
