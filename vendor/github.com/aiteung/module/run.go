package module

import (
	"github.com/aiteung/module/model"
)

func Run(Pesan model.IteungMessage, DBIface model.IteungDBConfig, URLApiWa string, HostIteungV1 string, apikeyv1 string) (Modulename string) {
	NormalizeAndTypoCorrection(&Pesan.Message, DBIface.MongoConn, DBIface.TypoCollection)
	if IsIteungCall(Pesan) {
		Modulename = GetModuleName(Pesan, DBIface.MongoConn, DBIface.ModuleCollection)
	}
	if Modulename != "" {
		CallAndSend(Modulename, Pesan, URLApiWa)
	} else {
		IteungV1(HostIteungV1, apikeyv1, Pesan)
	}

	return
}
