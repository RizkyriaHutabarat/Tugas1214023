package intership

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"Nama,omitempty" json:"Nama,omitempty"`
	Antrian_sidang string             `bson:"Antrian_sidang,omitempty" json:"Antrian_sidang,omitempty"`
	Bimbingan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Hari_sidang    string           `bson:"hari_sidang,omitempty" json:"hari_sidang,omitempty"`
}

type JamKuliah struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Intership struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm    string          `bson:"Npm,omitempty" json:"Npm,omitempty"`
	Kelas    string           `bson:"Kelas,omitempty" json:"Kelas,omitempty"`
	Kampus     string             `bson:"Kampus,omitempty" json:"Kampus,omitempty"`
	Judul string             `bson:"Judul,omitempty" json:"Judul,omitempty"`
	Tanggal       primitive.DateTime `bson:"Tanggal,omitempty" json:"Tanggal,omitempty"`
	
	Biodata      Mahasiswa           `bson:"Biodata Mahasiswa,omitempty" json:"Biodata Mahasiswa,omitempty"`
}



type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
