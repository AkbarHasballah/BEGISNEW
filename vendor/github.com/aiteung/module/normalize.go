package module

import (
	"regexp"

	"github.com/aiteung/atdb"
	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NormalizeAndTypoCorrection(message *string, MongoConn *mongo.Database, TypoCollection string) {
	typos := atdb.GetAllDoc[[]model.Typo](MongoConn, TypoCollection)
	for _, typo := range typos {
		re := regexp.MustCompile(`(?i)` + typo.From + ``)
		*message = re.ReplaceAllString(*message, typo.To)
	}

}
