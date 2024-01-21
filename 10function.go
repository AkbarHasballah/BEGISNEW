package BEGIS

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GeoIntersects(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request (koordinat untuk membentuk poligon)
	coordinatesStr := r.URL.Query().Get("coordinates")
	coordinates := strings.Split(coordinatesStr, ",")
	if len(coordinates) < 6 {
		fmt.Println("Koordinat untuk membentuk poligon tidak valid.")
		return
	}

	// Buat GeoJSON Polygon dari parameter koordinat
	polygon := bson.M{
		"type":        "Polygon",
		"coordinates": [][]interface{}{},
	}

	// Konversi koordinat menjadi format yang sesuai untuk GeoJSON
	for i := 0; i < len(coordinates); i += 2 {
		lng, errLng := strconv.ParseFloat(coordinates[i], 64)
		lat, errLat := strconv.ParseFloat(coordinates[i+1], 64)
		if errLng != nil || errLat != nil {
			fmt.Println("Koordinat tidak valid.")
			return
		}
		polygon["coordinates"] = append(polygon["coordinates"].([][]interface{}), []interface{}{lng, lat})
	}

	// Buat filter untuk geoIntersects query
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": polygon,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geoIntersects query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geoIntersects query:", results)
}

func GeoWithin(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	DBString := os.Getenv(MONGOCONNSTRINGENV)
	if DBString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(DBString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request (koordinat untuk membentuk poligon)
	coordinatesStr := r.URL.Query().Get("coordinates")
	coordinates := strings.Split(coordinatesStr, ",")
	if len(coordinates) < 6 {
		fmt.Println("Koordinat untuk membentuk poligon tidak valid.")
		return
	}

	// Buat GeoJSON Polygon dari parameter koordinat
	polygon := bson.M{
		"type":        "Polygon",
		"coordinates": [][]interface{}{},
	}

	// Konversi koordinat menjadi format yang sesuai untuk GeoJSON
	for i := 0; i < len(coordinates); i += 2 {
		lng, errLng := strconv.ParseFloat(coordinates[i], 64)
		lat, errLat := strconv.ParseFloat(coordinates[i+1], 64)
		if errLng != nil || errLat != nil {
			fmt.Println("Koordinat tidak valid.")
			return
		}
		polygon["coordinates"] = append(polygon["coordinates"].([][]interface{}), []interface{}{lng, lat})
	}

	// Buat filter untuk geoWithin query
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": polygon,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geoWithin query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geoWithin query:", results)
}

func Near(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []interface{}{longitude, latitude},
	}

	// Buat filter untuk near query
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": point,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan near query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil near query:", results)
}
func NearSphere(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []interface{}{longitude, latitude},
	}

	// Buat filter untuk nearSphere query
	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": point,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan nearSphere query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil nearSphere query:", results)
}
func Box(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request
	latitude1 := r.URL.Query().Get("latitude1")
	longitude1 := r.URL.Query().Get("longitude1")
	latitude2 := r.URL.Query().Get("latitude2")
	longitude2 := r.URL.Query().Get("longitude2")

	// Buat GeoJSON Box dari parameter koordinat
	box := bson.M{
		"type": "Polygon",
		"coordinates": [][]float64{
			{parseCoordinate(longitude1), parseCoordinate(latitude1)},
			{parseCoordinate(longitude2), parseCoordinate(latitude1)},
			{parseCoordinate(longitude2), parseCoordinate(latitude2)},
			{parseCoordinate(longitude1), parseCoordinate(latitude2)},
			{parseCoordinate(longitude1), parseCoordinate(latitude1)},
		},
	}

	// Buat filter untuk box query
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": box,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan box query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil box query:", results)
}

// Fungsi bantuan untuk mengonversi string koordinat ke float64
func parseCoordinate(coordStr string) float64 {
	coord, err := strconv.ParseFloat(coordStr, 64)
	if err != nil {
		fmt.Printf("Kesalahan saat mengonversi koordinat: %v\n", err)
	}
	return coord
}
func Center(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	radius := r.URL.Query().Get("radius")

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{parseCoordinate(longitude), parseCoordinate(latitude)},
	}

	// Buat filter untuk center query
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{point["coordinates"], ParseRadius(radius) / 6371}, // Radius dalam radian
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan center query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil center query:", results)
}

// Fungsi bantuan untuk mengonversi string koordinat ke float64

// Fungsi bantuan untuk mengonversi string radius ke float64
func ParseRadius(radiusStr string) float64 {
	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		fmt.Printf("Kesalahan saat mengonversi radius: %v\n", err)
	}
	return radius
}
func GeometryFix(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{parseCoordinate(longitude), parseCoordinate(latitude)},
	}

	// Buat filter untuk geometri queryt
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": point,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geometri query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geometri query:", results)
}
func MaxDistance(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dan jarak dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	distanceStr := r.URL.Query().Get("distance")

	// Konversi jarak ke float64
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		fmt.Printf("Kesalahan saat mengonversi jarak: %v\n", err)
		return
	}

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{parseCoordinate(longitude), parseCoordinate(latitude)},
	}

	// Buat filter untuk geometri query dengan operasi $maxDistance
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry":    point,
				"$maxDistance": distance,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geometri query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geometri query dengan $maxDistance:", results)
}
func MinDistance(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dan jarak dari request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	distanceStr := r.URL.Query().Get("distance")

	// Konversi jarak ke float64
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		fmt.Printf("Kesalahan saat mengonversi jarak: %v\n", err)
		return
	}

	// Buat GeoJSON Point dari parameter koordinat
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{parseCoordinate(longitude), parseCoordinate(latitude)},
	}

	// Buat filter untuk geometri query dengan operasi $minDistance
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry":    point,
				"$minDistance": distance,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geometri query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geometri query dengan $minDistance:", results)
}
func Polygon(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) {
	// Ambil nilai lingkungan MongoDB Connection String
	connString := os.Getenv(MONGOCONNSTRINGENV)
	if connString == "" {
		fmt.Println("MongoDB Connection String tidak ditemukan.")
		return
	}

	// Set up koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("Kesalahan saat membuat klien MongoDB: %v\n", err)
		return
	}

	// Tunggu hingga koneksi tersambung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Kesalahan saat menghubungkan ke MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Pilih database dan koleksi
	db := client.Database(dbname)
	collection := db.Collection(collectionname)

	// Ambil parameter koordinat dari request (format: "longitude1,latitude1,longitude2,latitude2,...")
	coordinatesStr := r.URL.Query().Get("coordinates")

	// Split string koordinat menjadi array float64
	coordinates := ParseCoordinatesPolygon(coordinatesStr)

	// Buat GeoJSON Polygon dari parameter koordinat
	polygon := bson.M{
		"type":        "Polygon",
		"coordinates": [][]float64{coordinates},
	}

	// Buat filter untuk geometri query dengan operasi $geoWithin
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": polygon,
			},
		},
	}

	// Lakukan query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("Kesalahan saat menjalankan geometri query: %v\n", err)
		return
	}
	defer cursor.Close(ctx)

	// Menggunakan cursor untuk melakukan sesuatu, misalnya mengambil hasil query
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Printf("Kesalahan saat mengambil hasil query: %v\n", err)
		return
	}

	// Proses hasil query sesuai kebutuhan
	fmt.Println("Hasil geometri query dengan $geoWithin (Polygon):", results)
}

// Fungsi bantuan untuk mengonversi string koordinat menjadi array float64
func ParseCoordinatesPolygon(coordStr string) []float64 {
	var coordinates []float64
	coords := strings.Split(coordStr, ",")
	for _, coord := range coords {
		val, err := strconv.ParseFloat(coord, 64)
		if err != nil {
			fmt.Printf("Kesalahan saat mengonversi koordinat buos: %v\n", err)
		}
		coordinates = append(coordinates, val)
	}
	return coordinates
}
