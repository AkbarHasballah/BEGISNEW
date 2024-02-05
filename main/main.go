package main

import (
	"fmt"
	"net/http"

	BEGIS "github.com/AkbarHasballah/GISNEW"
)

func main() {
	http.HandleFunc("/", HelloHTTP)
	http.ListenAndServe(":3000", nil)
}
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call GeoIntersects and check for errors
	// Tulis respons Anda ke Writer seperti yang Anda lakukan sebelumnya.
	response := BEGIS.PostGeoIntersects("publickey", "MONGOCONNSTRINGENV", "MigrasiData", "JsonMongo", r)
	fmt.Fprintf(w, response)
}
