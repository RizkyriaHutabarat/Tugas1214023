package main

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

func InsertNilai(tugas1 int,tugas2 int, tugas3 int, tugas4 int, tugas5 int, uts int, uas int, kategori Matakuliah ,biodata Mahasiswa) (InsertedID interface{}) {
	var nilai Nilai
	nilai.Tugas1 = tugas1
	nilai.Tugas2 = tugas2
	nilai.Tugas3 = tugas3
	nilai.Tugas4 = tugas4
	nilai.Tugas5 = tugas5
	nilai.UTS = uts
	nilai.UAS = uas
	nilai.Kategori = kategori
	nilai.Biodata = biodata
	return InsertOneDoc("db_nilai", "nilai", nilai)
}

func GetMahasiswaFromUAS(uas int) (pencapaian Nilai) {
	Mahasiswa := MongoConnect("db_nilai").Collection("Nilai")
	filter := bson.M{"uas": uas}
	err := Mahasiswa.FindOne(context.TODO(), filter).Decode(&pencapaian)
	if err != nil {
		fmt.Printf("getMahasiswaFromUAS: %v\n", err)
	}
	return pencapaian
}