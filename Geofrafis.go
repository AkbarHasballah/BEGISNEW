package BEGIS

import (
	"encoding/json"
	"net/http"
	"os"

)

func CreatetGeojsonPoint(MONGOCONNSTRINGENV,publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	response.Status = false
	var geojsonpoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	// Otorisasi
	var auth User
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header token tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token to get user details

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)
	auth.Username = tokenusername

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	// Insert data user
	response.Status = true
	PostPoint(mconn, collname, geojsonpoint)
	response.Message = "Berhasil input data"
	return GCFReturnStruct(response)
}

func MembuatGeojsonPolyline(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geojsonpolyline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonpolyline)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	var auth User
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header token tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token to get user details

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)
	auth.Username = tokenusername

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	// Insert data user
	PostLinestring(mconn, collname, geojsonpolyline)
	response.Message = "Berhasil input data"
	return GCFReturnStruct(response)
}

func MembuatGeojsonPolygon(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geojsonpolygon GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	var auth User
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header token tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token to get user details

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)
	auth.Username = tokenusername

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	// Insert data user
	PostPolygon(mconn, collname, geojsonpolygon)
	response.Message = "Berhasil input data"
	return GCFReturnStruct(response)
}

func PostGeoIntersects(publickey, MONGOCONNSTRINGENV, dbname, collname string ,r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var coordinate Point
	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	var auth User
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header token tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token to get user details

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)
	auth.Username = tokenusername

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	geointersects := GeoIntersects(mconn, collname, coordinate)
	response.Message = "Berhasil input data"
	return GCFReturnStruct(geointersects)
}

func PostGeoWithin(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var coordinate Polygon
	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	var auth User
	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header token tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Decode token to get user details

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)
	auth.Username = tokenusername

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user account exists
	if !usernameExists(MONGOCONNSTRINGENV, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	geowithin := GeoWithin(mconn, collname, coordinate)
	response.Message = "Berhasil input data"
	return GCFReturnStruct(geowithin)
}

func PostNear(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection2dsphere(MONGOCONNSTRINGENV, dbname, collname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	near := Near(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: near})
}

func PostNearSphere(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection2dsphere(MONGOCONNSTRINGENV, dbname, collname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	nearsphere := NearSphere(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: nearsphere})
}

func PostBox(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var coordinate Polyline
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	box := Box(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: box})
}

func PostCenter(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	box := Center(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: box})
}

func PostCenterSphere(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	box := CenterSphere(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: box})
}

func PostMaxDistance(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	box := MaxDistance(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: box})
}

func PostMinDistance(publickey, MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var coordinate Point
	err := c.BindJSON(&coordinate)
	if err != nil {
		c.JSON(http.StatusBadRequest, Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return GCFReturnStruct(response)
	}
	// Otorisasi
	Otorisasi(publickey)(c)
	if c.IsAborted() {
		return GCFReturnStruct(response)
	}
	role := c.GetString("role")
	// Cek role
	if role != "owner" {
		if role != "dosen" {
			c.JSON(http.StatusUnauthorized, Pesan{Status: false, Message: "Anda tidak memiliki akses"})
			c.Abort()
			return GCFReturnStruct(response)
		}
	}

	box := MinDistance(mconn, collname, coordinate)
	c.JSON(http.StatusOK, Pesan{Status: true, Message: box})
}

func AmbilDataGeojson(MONGOCONNSTRINGENV, dbname, collname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunan(mconn, collname)
	c.JSON(http.StatusOK, datagedung)
}
