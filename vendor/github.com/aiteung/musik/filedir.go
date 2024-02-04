package musik

import (
	"log"
	"os"
)

func CreateFolderifNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

}
