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

func InsertGaji(gatot string, bon string, pph string, pot string, biodata Karyawan) (InsertedID interface{}) {
	var gaji Gaji
	gaji.Gaji_pokok = gatot
	gaji.Bonus = bon
	gaji.Tunjangan_pph = pph
	gaji.Potongan = pot
	gaji.Tanggal = primitive.NewDateTimeFromTime(time.Now().UTC())
	gaji.Biodata = biodata
	return InsertOneDoc("tugas_db", "gaji", gaji)
}

func GetKaryawanFromJabatan(jabatan string) (staf Gaji) {
	karyawan := MongoConnect("tugas_db").Collection("gaji")
	filter := bson.M{"jabatan": jabatan}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromJabatan: %v\n", err)
	}
	return staf
}
