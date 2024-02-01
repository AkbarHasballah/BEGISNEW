package BEGIS

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/whatsauth/wa"
	"go.mongodb.org/mongo-driver/bson"
)

var savedToken string

func Otorisasi(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	var response Userkredensial
	var auth User
	response.Status = false

	// Extract token from the request header
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token values
	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)

	// Create User struct with the decoded username
	auth.Username = tokenusername

	// Check if decoding results are valid
	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Successful token decoding and user validation
	response.Message = "Berhasil decode token"
	response.Status = true
	response.Data.Username = tokenusername
	response.Data.Role = tokenrole

	return GCFReturnStruct(response)
}

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllUser(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func LoginHandler(token, privatekey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	var response BeriPesan
	response.Status = false

	// Establish MongoDB connection
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Decode user data from the request body
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)

	// Check for JSON decoding errors
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the entered password is not valid
	if !IsPasswordValid(mconn, collname, datauser) {
		response.Message = "Password Salah"
		return GCFReturnStruct(response)
	}

	// Retrieve user details
	user := FindUser(mconn, collname, datauser)

	// Prepare and encode token
	tokenstring, tokenerr := Encode(user.Username, user.Role, os.Getenv(privatekey))
	if tokenerr != nil {
		response.Message = "Gagal encode token: " + tokenerr.Error()
		return GCFReturnStruct(response)
	}

	// Successful login
	response.Status = true
	response.Token = tokenstring
	response.Message = "Berhasil login"

	// Send a WhatsApp message notifying the user about the successful login
	var nama = user.Username
	var nohp = user.No_whatsapp
	dt := &wa.TextMessage{
		To:       nohp,
		IsGroup:  false,
		Messages: nama + " berhasil login\nNikmati Web Wisata di kota bandung\nIni token Untuk melanjutkan Proses selanjutnya yah  " + "\n" + tokenstring,
	}
	atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv(token), dt, "https://api.wa.my.id/api/send/message/text")

	return GCFReturnStruct(response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GCFLoginAfterCreatee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, return a success message or some other response.
		return "Login successful"

	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}

func GCFLoginAfterCreateee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, construct and return the GCFReturnStruct.
		response := CreateResponse(true, "Berhasil Login", userdata)
		return GCFReturnStruct(response) // Return GCFReturnStruct directly
	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}
func GCFLoginAfterCreateeee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		return err.Error()
	}
	if IsPasswordValid(mconn, collectionname, userdata) {
		// Password is valid, return a success message or some other response.
		return GCFReturnStruct(userdata)
	} else {
		// Password is not valid, return an error message.
		return "Password Salah"
	}
}

func GCFCreteCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	err := json.NewDecoder(r.Body).Decode(&commentdata)
	if err != nil {
		return err.Error()
	}

	if err := CreateComment(mconn, collectionname, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Succes Create Comment", commentdata))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Create Comment", commentdata))
	}
}

func GCFGetAllComment(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datacomment := GetAllComment(mconn, collectionname)
	if datacomment != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Comment", datacomment))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Comment", datacomment))
	}
}
func GFCUpadatedCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	err := json.NewDecoder(r.Body).Decode(&commentdata)
	if err != nil {
		return err.Error()
	}

	if err := UpdatedComment(mconn, collectionname, bson.M{"id": commentdata.ID}, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Updated Comment", commentdata))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Updated Comment", commentdata))
	}
}

func GCFDeletedCommnet(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var commentdata Comment
	if err := json.NewDecoder(r.Body).Decode(&commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(false, "Failed to process request", commentdata))
	}

	if err := DeleteComment(mconn, collectionname, commentdata); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Successfully deleted comment", commentdata))
	}

	return GCFReturnStruct(CreateResponse(false, "Failed to delete comment", commentdata))
}

func GCFCreatePostLineStringg(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}

	// Mengambil nilai header PASETO dari permintaan HTTP
	pasetoValue := r.Header.Get("PASETOPRIVATEKEYENV")

	// Disini Anda dapat menggunakan nilai pasetoValue sesuai kebutuhan Anda
	// Misalnya, menggunakannya untuk otentikasi atau enkripsi.
	// Contoh sederhana menambahkan nilainya ke dalam pesan respons:
	response := GCFReturnStruct(geojsonline)
	response += " PASETO value: " + pasetoValue

	PostLinestring(mconn, collection, geojsonline)
	return response
}

func GCFCreatePostLineString(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collection, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func GCFDeleteLineString(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	if err := DeleteLinestring(mconn, collectionname, dataline); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Delete LineString", dataline))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Delete LineString", dataline))
	}
}

func GCFUpdateLinestring(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	if err := UpdatedLinestring(mconn, collectionname, bson.M{"properties.coordinates": dataline.Geometry.Coordinates}, dataline); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Update LineString", dataline))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Update LineString", dataline))
	}
}
func validateToken(token string) bool {
	return token == savedToken
}

func GCFCreateLineStringgg(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	// Get the token from the Authorization header
	token := r.Header.Get("Authorization")

	if token == "" {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Token is missing", nil))
	}

	// Handling Authorization by validating the token
	if validateToken(token) {
		// Perform the action (e.g., PostLinestring)
		err := PostLinestring(mconn, collectionname, dataline)
		if err != nil {
			return GCFReturnStruct(CreateResponse(false, "Data Sukses terinput Coy", nil))
		}
		// Success
		return GCFReturnStruct(CreateResponse(true, "Success: LineString created", dataline))
	}

	return GCFReturnStruct(CreateResponse(false, "Unauthorized: Invalid token", nil))
}

func GCFCreatePolygone(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body1
	var datapolygone GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapolygone)
	if err != nil {
		return err.Error()
	}

	// Handling Authorization
	if err := PostPolygone(mconn, collectionname, datapolygone); err != nil {
		// Success
		return GCFReturnStruct(CreateResponse(true, "Success Create Polygone", datapolygone))
	} else {
		// Failure
		return GCFReturnStruct(CreateResponse(false, "Failed Create Polygone", datapolygone))
	}
}

func GCFlineStingCreate(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collection, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func GCFlineStingCreatea(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, "Bad Request: Invalid JSON", nil))
	}

	// Checking Secret Header
	secretHeader := r.Header.Get("Secret")
	expectedSecret := os.Getenv("SECRET")

	if secretHeader != expectedSecret {
		log.Printf("Unauthorized: Secret header does not match. Expected: %s, Actual: %s", expectedSecret, secretHeader)
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// Handling Authorization
	PostLinestring(mconn, collection, geojsonline)

	return GCFReturnStruct(CreateResponse(true, "Success: LineString created", geojsonline))
}

func GCFCreatePolygonee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var datapolygone GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapolygone)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, "Bad Request: Invalid JSON", nil))
	}

	// Checking Secret Header
	secretHeader := r.Header.Get("Secret")
	expectedSecret := os.Getenv("SECRET")

	if secretHeader != expectedSecret {
		log.Printf("Unauthorized: Secret header does not match. Expected: %s, Actual: %s", expectedSecret, secretHeader)
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// Handling Authorization
	if err := PostPolygone(mconn, collectionname, datapolygone); err != nil {
		log.Printf("Error creating polygon: %v", err)
		return GCFReturnStruct(CreateResponse(false, "Failed Create Polygone", nil))
	}

	log.Println("Success: Polygon created")
	return GCFReturnStruct(CreateResponse(true, "Success Create Polygone", datapolygone))
}
