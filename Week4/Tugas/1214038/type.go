package presensiMahasiswa

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Npm           int             	 `bson:"npm,omitempty" json:"npm,omitempty"`
	Kelas         string           	 `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Jurusan       string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type JadwalKuliah struct {
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Hari       string   `bson:"hari,omitempty" json:"hari,omitempty"`
}

type MataKuliah struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_matkul  	int                `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama_matkul  	string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Sks          	int         	   `bson:"sks,omitempty" json:"sks,omitempty"`
	Dosen_pengajar  string 			   `bson:"dosen_pengajar,omitempty" json:"dosen_pengajar,omitempty"`
	Jadwal_kuliah   JadwalKuliah       `bson:"jadwal_kuliah,omitempty" json:"jadwal_kuliah,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Mata_kuliah  MataKuliah         `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
}

