package namapackage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertTagihan(tota string, des string, stat string, biodata Nasabah) (InsertedID interface{}) {
	var tagihan Tagihan
	tagihan.Total_tagihan = tota
	tagihan.Deskripsi = des
	tagihan.Status = stat
	tagihan.Tanggal_jatuhtempo = primitive.NewDateTimeFromTime(time.Now().UTC())
	tagihan.Biodata = biodata
	return InsertOneDoc("file_db", "tagihan", tagihan)
}

func GetNasabahFromNama(nama string) (sistem Tagihan) {
	nasabah := MongoConnect("file_db").Collection("tagihan")
	filter := bson.M{"nama": nama}
	err := nasabah.FindOne(context.TODO(), filter).Decode(&sistem)
	if err != nil {
		fmt.Printf("getNasabahFromNama: %v\n", err)  
	}
	return sistem
}
