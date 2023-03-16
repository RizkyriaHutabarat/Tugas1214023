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

func InsertProyek1(mahasiswa Mahasiswa, dosen_pembimbing Dosen, dosen_penguji Dosen, judul string, tanggal_sidang string) (InsertedID interface{}) {
	var proyek Proyek1
	proyek.Biodata_mahasiswa = mahasiswa
	proyek.Dosen_pembimbing = dosen_pembimbing
	proyek.Dosen_penguji = dosen_penguji
	proyek.Judul = judul
	proyek.Tanggal_sidang = tanggal_sidang
	return InsertOneDoc("proyek", "proyek1", proyek)
}

func GetProyek1FromNPM(npm int) (staf Proyek1) {
	proyek1 := MongoConnect("proyek").Collection("proyek1")
	filter := bson.M{"biodata_mahasiswa.npm": npm}
	fmt.Print(npm)
	err := proyek1.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("GetProyek1FromNPM: %v\n", err)
	}
	return staf
}
