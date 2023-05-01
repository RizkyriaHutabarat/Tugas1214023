package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Surat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Status_surat Status             `bson:"status,omitempty" json:"status,omitempty"`
	No_surat     int                `bson:"no_surat,omitempty" json:"no_surat,omitempty"`
	Perihal      string             `bson:"perihal,omitempty" json:"perihal,omitempty"`
	Kode_pos     Kodepos            `bson:"kodepos,omitempty" json:"kodepos,omitempty"`
	Pengirim     Pengirim           `bson:"pengirim,omitempty" json:"pengirim,omitempty"`
	Penerima     Penerima           `bson:"penerima,omitempty" json:"penerima,omitempty"`
}

type Disposisi struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tgl_disposisi int                `bson:"kode_pos,omitempty" json:"kode_pos,omitempty"`
	No_surat      Surat              `bson:"no_surat,omitempty" json:"no_surat,omitempty"`
	Penerima      Penerima           `bson:"penerima,omitempty" json:"penerima,omitempty"`
	Status_surat  Status             `bson:"status,omitempty" json:"status,omitempty"`
}
type Kodepos struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_pos    int                `bson:"kode_pos,omitempty" json:"kode_pos,omitempty"`
	Nama_daerah string             `bson:"nama_daerah,omitempty" json:"nama_dareah,omitempty"`
}

type Status struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_status  int                `bson:"kode_status,omitempty" json:"kode_status,omitempty"`
	Status_surat string             `bson:"nama_daerah,omitempty" json:"nama_dareah,omitempty"`
}

type Penerima struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_penerima  string             `bson:"nama_pengirim,omitempty" json:"nama_pengirim,omitempty"`
	Alamat         string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tanggal_terima string             `bson:"jadwal_surat,omitempty" json:"jadwal_surat,omitempty"`
}

type Pengirim struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_pengirim string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tanggal_kirim string             `bson:"tanggal_keluar,omitempty" json:"tanggal_keluar,omitempty"`
}
