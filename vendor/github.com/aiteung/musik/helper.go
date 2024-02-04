package musik

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"go.mau.fi/whatsmeow"
)

func GetIPaddress() string {

	resp, err := http.Get("https://icanhazip.com/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func GetQR(wc *whatsmeow.Client) {
	var err error
	if wc.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := wc.GetQRChannel(context.Background())
		err = wc.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR code:", evt.Code)
				WriteQR(evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
				WriteQR(evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = wc.Connect()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client Connected")
		WriteQR("connected")
	}
}

func WriteQR(qr_string string) {
	qrdata := []byte(qr_string)
	err := os.WriteFile("session/qr.data", qrdata, 0777)
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func ReadQR() (qr_string string) {
	data, err := os.ReadFile("session/qr.data")
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
