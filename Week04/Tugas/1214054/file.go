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

func InsertKemahasiswaan(identitas Mahasiswa, status_keuangan Keuangan, nilai_mhs Nilai) (InsertedID interface{}) {
	var kemahasiswaan Kemahasiswaan
	kemahasiswaan.Identitas_Mahasiswa = identitas
	kemahasiswaan.Status_Keuangan = status_keuangan
	kemahasiswaan.Nilai_Mahasiswa = nilai_mhs
	return InsertOneDoc("kemahasiswaan_db", "kemahasiswaan", kemahasiswaan)
}

func TestGetKemahasiswaanFromNPM(npm string) (staf Kemahasiswaan) {
	mahasiswa := MongoConnect("kemahasiswaan_db").Collection("kemahasiswaan")
	filter := bson.M{"identitas.npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKemahasiswaanFromNpm: %v\n", err)
	}
	return staf
}
