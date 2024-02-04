package musik

import (
	"log"
	"os"
)

func GenerateStaticAssets(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		CloneRepo("https://github.com/InformaticsResearchCenter/static/archive/refs/heads/master.zip")
		err = os.Rename("static-master", "static")
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove("master.zip")
		if err != nil {
			log.Fatal(err)
		}
	}
}
