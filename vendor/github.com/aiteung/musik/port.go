package musik

import "os"

func Dangdut() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		port = ":" + port
	}
	return
}

func Koplo() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	} else if port[0:1] != ":" {
		port = ":" + port
	}
	return
}

func Tarling() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return
}
