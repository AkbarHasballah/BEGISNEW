package BEGIS

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CreatetGeojsonPoint(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
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

func PostGeoIntersects(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	geointersect, err := GeoIntersects(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson yang berdekatan pada koordinat anda " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(geointersect) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(geointersect)
	result = "Geojson yang bersinggungan dengan geometry adalah " + result + ". Jumlah geojson terdekat: " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}
func GeojsonNameString(geojson []FullGeoJson) (result string) {
	var names []string
	for _, geojson := range geojson {
		names = append(names, geojson.Properties.Name)
	}
	result = strings.Join(names, ", ")
	return result
}

func PostGeoWithin(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	geowithin, err := GeoWithin(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson yang terdapat pada polygon anda " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(geowithin) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(geowithin)
	result = "Geojson yang terdapat dalam polygon anda adalah " + result + ". Jumlah geojsonnya : " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}

func PostNear(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection2dsphere(MONGOCONNSTRINGENV, dbname, collname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	near, err := Near(mconn, collname, geospatial)
	if err != nil {
		response.Message = "Tidak terdapat geojson yang berdekatan pada koordinat anda " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(near) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(near)
	result = "Geojson yang terdekat dari koordinat anda adalah " + result + ". Jumlah geojson terdekat: " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}

func PostNearSphere(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection2dsphere(MONGOCONNSTRINGENV, dbname, collname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	nearsphere, err := NearSphere(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson yang berdekatan pada koordinat anda " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(nearsphere) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(nearsphere)
	result = "Geojson yang terdekat dari koordinat anda adalah " + result + ". Jumlah geojson terdekat: " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}
func PostBox(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	box, err := Box(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson point di box anda " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(box) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(box)
	result = "Geojson point yang berada diarea dari box anda adalah " + result + ". Jumlah geojsonnya: " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}

func PostCenter(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geospatial Geospatial
	var response BeriPesan
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	center, err := Center(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson point didalam lingkaran dengan radius " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(center) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(center)
	result = "Geojson point yang berada diarea dari lingkaran anda adalah" + result + ". Jumlah geojson nya ada;ah : " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}

func PostCenterSphere(publickey, MONGOCONNSTRINGENV, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var response BeriPesan
	var geospatial Geospatial
	err := json.NewDecoder(r.Body).Decode(&geospatial)
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

	// Check if the user has admin or user privileges
	if tokenrole != "admin" && tokenrole != "user" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	centerSphere, err := CenterSphere(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson point didalam lingkaran dengan radius " + err.Error()
		return GCFReturnStruct(response)
	}

	numNearby := len(centerSphere) // Menghitung jumlah geojson yang terdekat

	result := GeojsonNameString(centerSphere)
	result = "Geojson point yang berada diarea dari lingkaran anda adalah" + result + ". Jumlah geojson nya ada;ah : " + strconv.Itoa(numNearby)
	return GCFReturnStruct(result)
}

func AmbilDataGeojson(MONGOCONNSTRINGENV, dbname, collname string) []GeoJson {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunan(mconn, collname)
	return datagedung
}
