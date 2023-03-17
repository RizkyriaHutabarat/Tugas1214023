package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kemahasiswaan struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Identitas_Mahasiswa Mahasiswa          `bson:"identitas,omitempty" json:"identitas,omitempty"`
	Status_Keuangan     Keuangan           `bson:"status_keuangan,omitempty" json:"status_keuangan,omitempty"`
	Nilai_Mahasiswa     Nilai              `bson:"nilai_mhs,omitempty" json:"nilai_mhs,omitempty"`
}

type Mahasiswa struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm     string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama    string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Prodi   ProgramStudi       `bson:"prodi,omitempty" json:"prodi,omitempty"`
	Jurusan string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Kelas   string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
}

type ProgramStudi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Prodi string             `bson:"kode_prodi,omitempty" json:"kode_prodi,omitempty"`
	Nama_Prodi string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Keuangan struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata          Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Total_Pembayaran float32            `bson:"total_pembayaran,omitempty" json:"total_pembayaran,omitempty"`
}

type Nilai struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata     Mahasiswa          `bson:"nama,omitempty" json:"nama,omitempty"`
	Matakuliah  Matakuliah         `bson:"matakuliah,omitempty" json:"matakuliah,omitempty"`
	Nilai_Angka int                `bson:"nilai_angka,omitempty" json:"nilai_angka,omitempty"`
	Nilai_Huruf string             `bson:"nilai_huruf,omitempty" json:"nilai_huruf,omitempty"`
}

type Matakuliah struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Nama_Dosen  string             `bson:"dosen,omitempty" json:"dosen,omitempty"`
}
