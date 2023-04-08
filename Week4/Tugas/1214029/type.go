package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tamu struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
}

type Waktu struct {
	Jam      string   `bson:"jam,omitempty" json:"jam,omitempty"`	
	Hari     string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Tanggal  string   `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Location struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_lokasi    string             `bson:"nama_lokasi,omitempty" json:"nama,omitempty"`
}