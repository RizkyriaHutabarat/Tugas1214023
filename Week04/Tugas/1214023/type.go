package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tanggal_lahir    string         `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Media_sosial	[]MediaSosial	`bson:"media_sosial,omitempty" json:"media_sosial,omitempty"`
}

type Pendidikan  struct {
	Pendidikan_terakhir    string      `bson:"pendidikan_terakhir,omitempty" json:"pendidikan_terakhir,omitempty"`
	Tahun_lulus  int   `bson:"tahun_lulus,omitempty" json:"tahun_lulus,omitempty"`
	Bidang_studi string   `bson:"bidang_studi,omitempty" json:"bidang_studi,omitempty"`
	Institusi_pendidikan string `bson:"institusi_pendidikan,omitempty" json:"institusi_pendidikan,omitempty"`
}

type MediaSosial struct {
	Instagram    string            `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Facebook     string            `bson:"facebook,omitempty" json:"facebook,omitempty"`
	Twitter     string             `bson:"twitter,omitempty" json:"twitter,omitempty"`
}

type Profile struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Login      string             `bson:"login,omitempty" json:"login,omitempty"`
	Biodata      People           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}


