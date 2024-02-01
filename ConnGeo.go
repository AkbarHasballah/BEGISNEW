package BEGIS
import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func PostPoint(MONGOCONNSTRINGENV *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return InsertOneDoc(MONGOCONNSTRINGENV, collection, pointdata)
}

func PostLinestring(MONGOCONNSTRINGENV *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return InsertOneDoc(MONGOCONNSTRINGENV, collection, linestringdata)
}

func PostPolygon(MONGOCONNSTRINGENV *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return InsertOneDoc(MONGOCONNSTRINGENV, collection, polygondata)
}

// Read

func GetAllBangunan(MONGOCONNSTRINGENV *mongo.Database, collname string) []GeoJson {
	return GetAllDoc[[]GeoJson](MONGOCONNSTRINGENV, collname)
}

func GeoIntersects(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetGeoIntersectsDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func GeoWithin(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Polygon) string {
	return GetGeoWithinDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func Near(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetNearDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func NearSphere(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetNearSphereDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func Box(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Polyline) string {
	return GetBoxDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func Center(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetCenterDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func CenterSphere(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetCenterSphereDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func MaxDistance(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetMaxDistanceDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

func MinDistance(MONGOCONNSTRINGENV *mongo.Database, collname string, coordinates Point) string {
	return GetMinDistanceDoc(MONGOCONNSTRINGENV, collname, coordinates)
}

// Update

// Delete

func DeleteGeojson(mMONGOCONNSTRINGENV *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return DeleteOneDoc(mMONGOCONNSTRINGENV, collname, filter)
}
