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

func GeoIntersects(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial)([]FullGeoJson, error){
	return GetGeoIntersectsDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry",geospatial)
}
func GeoWithin(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetGeoWithinDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}
func Near(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetNearDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}

func NearSphere(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetNearSphereDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}

func Box(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetBoxDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}

func Center(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetCenterDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}

func CenterSphere(MONGOCONNSTRINGENV *mongo.Database, collname string, geospatial Geospatial) ([]FullGeoJson, error) {
	return GetCenterSphereDoc[FullGeoJson](MONGOCONNSTRINGENV, collname, "geometry", geospatial)
}

// Update

// Delete

func DeleteGeojson(mMONGOCONNSTRINGENV *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return DeleteOneDoc(mMONGOCONNSTRINGENV, collname, filter)
}
