package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NPM			 int 	            `bson:"npm,omitempty" json:"npm,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type Matakuliah struct {
	Nama_MK	   string   			`bson:"nama_mk,omitempty" json:"nama_mk,omitempty"`
	SKS        int      			`bson:"sks,omitempty" json:"sks,omitempty"`
	Jam_masuk  string   			`bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   			`bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Hari       string 				`bson:"hari,omitempty" json:"hari,omitempty"`
}

type Nilai struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tugas1    	 int            	`bson:"tugas1,omitempty" json:"tugas1,omitempty"`
	Tugas2     	 int            	`bson:"tugas2,omitempty" json:"tugas2,omitempty"`
	Tugas3     	 int             	`bson:"tugas3,omitempty" json:"location,omitempty"`
	Tugas4 		 int             	`bson:"tugas4,omitempty" json:"tugas4,omitempty"`
	Tugas5     	 int				`bson:"tugas5,omitempty" json:"tugas5,omitempty"`
	UTS      	 int             	`bson:"uts,omitempty" json:"uts,omitempty"`
	UAS      	 int             	`bson:"uas,omitempty" json:"uas,omitempty"`
	Kategori	 Matakuliah			`bson:"kategori,omitempty" json:"kategori,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
