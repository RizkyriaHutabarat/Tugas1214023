package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Admin        []Admin            `bson:"admin,omitempty" json:"admin,omitempty"`
	Hari_gaji    []string           `bson:"hari_gaji,omitempty" json:"hari_gaji,omitempty"`
}

type Admin struct {
	Nama         string   `bson:"nama,omitempty" json:"nama,omitempty"`
	Email        string   `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number string   `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Hari         []string `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Gaji struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tanggal       primitive.DateTime `bson:"Tanggal,omitempty" json:"Tanggal,omitempty"`
	Gaji_pokok    string             `bson:"Gaji_pokok,omitempty" json:"Gaji_pokok,omitempty"`
	Bonus         string             `bson:"Bonus,omitempty" json:"Bonus,omitempty"`
	Tunjangan_pph string             `bson:"Tunjangan_pph,omitempty" json:"Tunjangan_pph,omitempty"`
	Potongan      string             `bson:"Potongan,omitempty" json:"Potongan,omitempty"`
	Biodata       Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
