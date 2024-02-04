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

	geointersects, err := GeoIntersects(mconn, collname, geospatial)
	if err != nil {
		response.Message = "GetGeoInterDOc Error Coyz: " + err.Error()
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(geointersects)
	if result == "" {
		response.Message = "Geojson yang bersinggungan dengan geometry anda adalah" + result
	}
	response.Message = "Berhasil input data"
	return GCFReturnStruct(geointersects)
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
		response.Message = "Ga ada Geojson Di dalam polygon coyz: " + err.Error()
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(geowithin)
	if result == "" {
		response.Message = "Geojson yang berada di polygon nya  adalah" + result
	}
	response.Message = "Berhasil input data"
	return GCFReturnStruct(geowithin)
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
		response.Message = " Tidak terdapat geojson yang berdekatan pada koordinat anda " + err.Error()
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(near)
	if result == "" {
		response.Message = "Geojson yang terdekat dari koordinat anda adalah" + result
	}
	return GCFReturnStruct(near)
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

	nearsphere, err:= NearSphere(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson yang berdekatan pada koordinat anda " + err.Error()
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(nearsphere)
	if result == "" {
		response.Message = "Geojson yang terdekat dari koordinat anda adalah" + result
	}
	return GCFReturnStruct(nearsphere)
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
	box, err:= Box(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson point di box anda " + err.Error()
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(box)
	if result == "" {
		response.Message = "Geojson point yang berada diarea dari box anda adalah" + result
	}
	return GCFReturnStruct(box)
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
		response.Message = " Tidak terdapat geojson point didalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius,'f',-1,64)
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(center)
	if result == "" {
		response.Message = "Geojson point yang berada diarea dari box anda adalah" + strconv.FormatFloat(geospatial.Radius,'f', -1, 64)+ "adalah" +result
	}
	return GCFReturnStruct(center)
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
	centersphere, err:= CenterSphere(mconn, collname, geospatial)
	if err != nil {
		response.Message = " Tidak terdapat geojson point didalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius,'f',-1,64)
		return GCFReturnStruct(response)
	}
	result := GeojsonNameString(centersphere)
	if result == "" {
		response.Message = "Geojson point yang berada diarea dari box anda adalah" + strconv.FormatFloat(geospatial.Radius,'f', -1, 64)+ "adalah" +result
	}
	return GCFReturnStruct(centersphere)
}

func AmbilDataGeojson(MONGOCONNSTRINGENV, dbname, collname string) []GeoJson {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunan(mconn, collname)
	return datagedung
}
