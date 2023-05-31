package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jurusan struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_j       string             `bson:"nama_j,omitempty" json:"nama_j,omitempty"`
}

type Matakuliah struct {
    Nama_MK    string               `bson:"nama_mk,omitempty" json:"nama_mk,omitempty"`
    SKS        int                  `bson:"sks,omitempty" json:"sks,omitempty"`
    Jam_masuk  string               `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
    Jam_keluar string               `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
    Hari       string               `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Mahasiswa struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Nama_ms      string             `bson:"nama_ms,omitempty" json:"nama_ms,omitempty"`
    NPM          int                `bson:"npm,omitempty" json:"npm,omitempty"`
    Phone_number int                `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
    Kategori     Matakuliah         `bson:"kategori,omitempty" json:"kategori,omitempty"`
    Biodata      Jurusan            `bson:"biodata,omitempty" json:"biodata,omitempty"`
}