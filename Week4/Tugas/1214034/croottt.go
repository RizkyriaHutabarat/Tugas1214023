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

func InsertNilai(nama_ms string,npm int, phone_number int, kategori Matakuliah ,biodata Jurusan) (InsertedID interface{}) {
    var mahasiswa Mahasiswa
    mahasiswa.Nama_ms = nama_ms
    mahasiswa.NPM = npm
    mahasiswa.Phone_number = phone_number
    mahasiswa.Kategori = kategori
    mahasiswa.Biodata = biodata
    return InsertOneDoc("db_mahasiswa", "mahasiswa", mahasiswa)
}

func GetMahasiswaFromNPM(npm int) (data Mahasiswa) {
    Jurusan := MongoConnect("db_mahasiswa").Collection("Mahasiswa")
    filter := bson.M{"npm": npm}
    err := Jurusan.FindOne(context.TODO(), filter).Decode(&data)
    if err != nil {
        fmt.Printf("getMahasiswaFromNPM: %v\n", err)
    }
    return data
}