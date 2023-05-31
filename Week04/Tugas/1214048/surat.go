package namapackage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

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

func InsertSurat(NoSurat int, Status Status, Perihal string, Kodepos Kodepos, pengirim Pengirim, penerima Penerima) (InsertedID interface{}) {
	var surat Surat
	surat.No_surat = NoSurat
	surat.Status_surat = Status
	surat.Perihal = Perihal
	surat.Kode_pos = Kodepos
	surat.Pengirim = pengirim
	surat.Penerima = penerima
	return InsertOneDoc("tes_db", "surat", surat)
}

func GetSuratFromNoSurat(No_surat int) (ns Surat) {
	surat := MongoConnect("tes_db").Collection("surat")
	filter := bson.M{"no_surat": No_surat}
	err := surat.FindOne(context.TODO(), filter).Decode(&ns)
	if err != nil {
		fmt.Printf("getSuratFromNoSurat: %v\n", err)
	}
	return ns
}
