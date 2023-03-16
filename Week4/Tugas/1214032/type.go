package perwalian

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type Waktu struct {
	Jam      string   `bson:"jam,omitempty" json:"jam,omitempty"`	
	Hari     string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Tanggal  string   `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Location struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_lokasi    string             `bson:"nama_lokasi,omitempty" json:"nama_lokasi,omitempty"`
}