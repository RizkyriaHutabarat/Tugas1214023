package presensiMahasiswa

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func InsertPresensi(checkin string, biodata Mahasiswa, matkul MataKuliah) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	presensi.Mata_kuliah = matkul
	return InsertOneDoc("tugas_db", "presensi", presensi)
}

func GetMahasiswaFromNpm(npm int) (mhs Presensi) {
	mahasiswa := MongoConnect("tugas_db").Collection("presensi")
	filter := bson.M{"biodata.npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("getMahasiswaFromNpm: %v\n", err)
	}
	return mhs
}

func GetMatkulFromKodeMatkul(kode_matkul int) (mk Presensi) {
	matkul := MongoConnect("tugas_db").Collection("presensi")
	filter := bson.M{"mata_kuliah.kode_matkul": kode_matkul}
	err := matkul.FindOne(context.TODO(), filter).Decode(&mk)
	if err != nil {
		fmt.Printf("getMatkulFromKodeMatkul: %v\n", err)
	}
	return mk
}
