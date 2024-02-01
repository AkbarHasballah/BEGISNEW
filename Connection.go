package BEGIS

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AkbarHasballah/GISNEW/models"
	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IsAuthorized(username string, mconn *mongo.Client, collectionname string) bool {
	// Implement logic to check if the user is authorized based on the given criteria
	// For example, you might query a MongoDB collection to check if the user has the required permissions.

	// Placeholder example (you need to replace this with your actual authorization logic)
	// In this example, we assume there is a MongoDB collection named 'permissions' where user roles are stored.
	// You need to modify this based on your actual data structure and authorization logic.
	// For demonstration purposes, this example allows any user with any role.
	filter := bson.M{"username": username}
	result := mconn.Database("InformasiWisataBandung").Collection("Users").FindOne(context.Background(), filter)

	return result.Err() == nil // Assume the user is authorized if there are no errors in the query
}

func SetConnection(MONGOCONNSTRINGENVSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENVSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func SetConnection2dsphere(mongoenv, dbname, collname string) *mongo.Database {
	var DBmongoinfo = models.DBInfo2{
		DBString:       os.Getenv(mongoenv),
		DBName:         dbname,
		CollectionName: collname,
	}
	return helpers.Create2dsphere(DBmongoinfo)
}

func GetAllBangunanLineString(MONGOCONNSTRINGENV *mongo.Database, collection string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](MONGOCONNSTRINGENV, collection)
	return lokasi
}

func GetAllProduct(MONGOCONNSTRINGENV *mongo.Database, collection string) []Product {
	product := atdb.GetAllDoc[[]Product](MONGOCONNSTRINGENV, collection)
	return product
}

func GetNameAndPassowrd(MONGOCONNSTRINGENV *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](MONGOCONNSTRINGENV, collection)
	return user
}

func GetAllUser(MONGOCONNSTRINGENV *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](MONGOCONNSTRINGENV, collection)
	return user
}

func GetAllContent(MONGOCONNSTRINGENV *mongo.Database, collection string) []Content {
	content := atdb.GetAllDoc[[]Content](MONGOCONNSTRINGENV, collection)
	return content
}

//	func GetAllUser(MONGOCONNSTRINGENV *mongo.Database, collection string) []User {
//		user := atdb.GetAllDoc[[]User](MONGOCONNSTRINGENV, collection)
//		return user
//	}
func CreateNewUserRole(MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, userdata)
}
func usernameExists(MONGOCONNSTRINGENVSTRINGENV, dbname string, userdata User) bool {
	mconn := SetConnection(MONGOCONNSTRINGENVSTRINGENV, dbname).Collection("Users")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

func DeleteUser(MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(MONGOCONNSTRINGENV, collection, filter)
}
func ReplaceOneDoc(MONGOCONNSTRINGENV *mongo.Database, collection string, filter bson.M, userdata User) interface{} {
	return atdb.ReplaceOneDoc(MONGOCONNSTRINGENV, collection, filter, userdata)
}
func FindUser(MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](MONGOCONNSTRINGENV, collection, filter)
}

func FindUserUser(MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](MONGOCONNSTRINGENV, collection, filter)
}

func IsPasswordValid(MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](MONGOCONNSTRINGENV, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}

// product

func CreateNewProduct(MONGOCONNSTRINGENV *mongo.Database, collection string, productdata Product) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, productdata)
}

func InsertUserdata(MONGOCONNSTRINGENV *mongo.Database, username, role, password string) (InsertedID interface{}) {
	req := new(User)
	req.Username = username
	req.Password = password
	req.Role = role
	return InsertOneDoc(MONGOCONNSTRINGENV, "user", req)
}
func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// gis function

// content
func CreateNewContent(MONGOCONNSTRINGENV *mongo.Database, collection string, contentdata Content) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, contentdata)
}

func DeleteContent(MONGOCONNSTRINGENV *mongo.Database, collection string, contentdata Content) interface{} {
	filter := bson.M{"id": contentdata.ID}
	return atdb.DeleteOneDoc(MONGOCONNSTRINGENV, collection, filter)
}

func ReplaceContent(MONGOCONNSTRINGENV *mongo.Database, collection string, filter bson.M, contentdata Content) interface{} {
	return atdb.ReplaceOneDoc(MONGOCONNSTRINGENV, collection, filter, contentdata)
}

func CreateNewBlog(MONGOCONNSTRINGENV *mongo.Database, collection string, blogdata Blog) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, blogdata)
}

func FindContentAllId(MONGOCONNSTRINGENV *mongo.Database, collection string, contentdata Content) Content {
	filter := bson.M{"id": contentdata.ID}
	return atdb.GetOneDoc[Content](MONGOCONNSTRINGENV, collection, filter)
}

func GetAllBlogAll(MONGOCONNSTRINGENV *mongo.Database, collection string) []Blog {
	blog := atdb.GetAllDoc[[]Blog](MONGOCONNSTRINGENV, collection)
	return blog
}

func GetIDBlog(MONGOCONNSTRINGENV *mongo.Database, collection string, blogdata Blog) Blog {
	filter := bson.M{"id": blogdata.ID}
	return atdb.GetOneDoc[Blog](MONGOCONNSTRINGENV, collection, filter)
}

func AuthenticateUserAndGenerateToken(privateKeyEnv string, MONGOCONNSTRINGENV *mongo.Database, collection string, userdata User) (string, error) {
	// Cari pengguna berdasarkan nama pengguna
	username := userdata.Username
	password := userdata.Password
	userdata, err := FindUserByUsername(MONGOCONNSTRINGENV, collection, username)
	if err != nil {
		return "", err
	}

	// Memeriksa kata sandi
	if !CheckPasswordHash(password, userdata.Password) {
		return "", errors.New("Password salah") // Gantilah pesan kesalahan sesuai kebutuhan Anda
	}

	// Generate token untuk otentikasi
	tokenstring, err := watoken.Encode(username, os.Getenv(privateKeyEnv))
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

func FindUserByUsername(MONGOCONNSTRINGENV *mongo.Database, collection string, username string) (User, error) {
	var user User
	filter := bson.M{"username": username}
	err := MONGOCONNSTRINGENV.Collection(collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// create login using Private

func CreateComment(MONGOCONNSTRINGENV *mongo.Database, collection string, commentdata Comment) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, commentdata)
}

func DeleteComment(MONGOCONNSTRINGENV *mongo.Database, collection string, commentdata Comment) interface{} {
	filter := bson.M{"id": commentdata.ID}
	return atdb.DeleteOneDoc(MONGOCONNSTRINGENV, collection, filter)
}

func UpdatedComment(MONGOCONNSTRINGENV *mongo.Database, collection string, filter bson.M, commentdata Comment) interface{} {
	filter = bson.M{"id": commentdata.ID}
	return atdb.ReplaceOneDoc(MONGOCONNSTRINGENV, collection, filter, commentdata)
}

func GetAllComment(MONGOCONNSTRINGENV *mongo.Database, collection string) []Comment {
	comment := atdb.GetAllDoc[[]Comment](MONGOCONNSTRINGENV, collection)
	return comment
}

func PostLineString(MONGOCONNSTRINGENV *mongo.Database, collection string, commentdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, commentdata)
}

func PostLinestring1(MONGOCONNSTRINGENV *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, linestringdata)
}

func GetByCoordinate(MONGOCONNSTRINGENV *mongo.Database, collection string, linestringdata GeoJsonLineString) GeoJsonLineString {
	filter := bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.GetOneDoc[GeoJsonLineString](MONGOCONNSTRINGENV, collection, filter)
}

func DeleteLinestring(MONGOCONNSTRINGENV *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	filter := bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.DeleteOneDoc(MONGOCONNSTRINGENV, collection, filter)
}

func UpdatedLinestring(MONGOCONNSTRINGENV *mongo.Database, collection string, filter bson.M, linestringdata GeoJsonLineString) interface{} {
	filter = bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.ReplaceOneDoc(MONGOCONNSTRINGENV, collection, filter, linestringdata)
}

func PostPolygone(MONGOCONNSTRINGENV *mongo.Database, collection string, polygonedata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(MONGOCONNSTRINGENV, collection, polygonedata)
}
func FindNearestRoad(mconn *mongo.Database, collectionname string, coordinates []float64) GeoJsonLineString {
	// Gunakan query $near di MongoDB untuk mencari jalan terdekat
	filter := bson.M{
		"geometry.coordinates": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates,
				},
			},
		},
	}

	var result GeoJsonLineString
	err := mconn.Collection(collectionname).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Fungsi untuk mencari jalur dari jalan awal ke jalan akhir
// Fungsi untuk mencari jalur dari jalan awal ke jalan akhir
func FindRoute(mconn *mongo.Database, collectionname string, startGeometry, endGeometry GeometryLineString) []GeoJsonLineString {
	var result []GeoJsonLineString

	// Mencari jalan berdasarkan geometri awal
	startRoadFilter := bson.M{
		"geometry.coordinates": bson.M{
			"$near": bson.M{
				"$geometry": startGeometry,
			},
		},
	}

	var startRoad GeoJsonLineString
	err := mconn.Collection(collectionname).FindOne(context.Background(), startRoadFilter).Decode(&startRoad)
	if err != nil {
		log.Fatal(err)
	}

	// Mencari jalan berdasarkan geometri akhir
	endRoadFilter := bson.M{
		"geometry.coordinates": bson.M{
			"$near": bson.M{
				"$geometry": endGeometry,
			},
		},
	}

	var endRoad GeoJsonLineString
	err = mconn.Collection(collectionname).FindOne(context.Background(), endRoadFilter).Decode(&endRoad)
	if err != nil {
		log.Fatal(err)
	}

	// Mengembalikan hasil pencarian
	result = append(result, startRoad, endRoad)

	return result
}

// GEografis Fix Takis

func GetAllDoc[T any](MONGOCONNSTRINGENV *mongo.Database, collection string) (doc T) {
	ctx := context.TODO()
	cur, err := MONGOCONNSTRINGENV.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
	}
	return
}

func GetGeoIntersectsDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GeoIntersects: %v\n", err)
	}
	return "Koordinat anda bersinggungan dengan " + doc.Properties.Name
}
func GetGeoWithinDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Polygon) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GeoWithin: %v\n", err)
	}
	return "Koordinat anda berada di " + doc.Properties.Name
}

func GetNearDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": 1000,
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("Near: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}

func GetNearSphereDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": 1000,
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("NearSphere: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}
func GetBoxDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Polyline) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$box": coordinates.Coordinates,
			},
		},
	}

	var docs []FullGeoJson
	cur, err := MONGOCONNSTRINGENV.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Box: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return result
}
func GetCenterDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$center": []interface{}{coordinates.Coordinates, 0.003},
			},
		},
	}

	var docs []FullGeoJson
	cur, err := MONGOCONNSTRINGENV.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Box: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return result
}

func GetCenterSphereDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{coordinates.Coordinates, 0.00003},
			},
		},
	}

	var docs []FullGeoJson
	cur, err := MONGOCONNSTRINGENV.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Box: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return result
}
func GetMaxDistanceDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": coordinates.Max,
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("Near: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}
func GetMinDistanceDoc(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$minDistance": coordinates.Min,
			},
		},
	}
	var doc FullGeoJson
	err := MONGOCONNSTRINGENV.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("Near: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}
func DeleteOneDoc(MONGOCONNSTRINGENV *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := MONGOCONNSTRINGENV.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteOneDoc: %v\n", err)
	}
	return
}

func Create2dsphere(mconn atdb.DBInfo) (MONGOCONNSTRINGENV *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}

	// Mengecek apakah indeks sudah ada
	collection := client.Database(mconn.DBName).Collection(mconn.x)
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		fmt.Printf("Error listing indexes: %v", err)
	}

	indexExists := false
	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			fmt.Printf("Error decoding index: %v", err)
		}
		if index["name"] == "geometry_2dsphere" {
			indexExists = true
			break
		}
	}

	// Membuat indeks jika belum ada
	if !indexExists {
		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "geometry", Value: "2dsphere"},
			},
		}

		_, err = client.Database(mconn.DBName).Collection(mconn.DBName).Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			fmt.Printf("Error creating geospatial index: %v", err)
		}
	}
	return client.Database(mconn.DBName)
}
