package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Surat struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	No_surat      int                `bson:"no_surat,omitempty" json:"no_surat,omitempty"`
	Kepala_surat  string             `bson:"kepala_surat,omitempty" json:"kepala_surat,omitempty"`
	Alamat_tujuan string             `bson:"alamat_tujuan,omitempty" json:"alamat_tujuan,omitempty"`
	Kode_pos      int                `bson:"kode_pos,omitempty" json:"kode_pos,omitempty"`
	Pengirim      Pengirim           `bson:"pengirim,omitempty" json:"pengirim,omitempty"`
	Penerima      Penerima           `bson:"penerima,omitempty" json:"penerima,omitempty"`
}

type Penerima struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_pengirim string             `bson:"nama_pengirim,omitempty" json:"nama_pengirim,omitempty"`
	Tanggal_masuk string             `bson:"jadwal_surat,omitempty" json:"jadwal_surat,omitempty"`
}

type Pengirim struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_penerima  string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Tanggal_keluar string             `bson:"tanggal_keluar,omitempty" json:"tanggal_keluar,omitempty"`
}
