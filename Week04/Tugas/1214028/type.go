package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nasabah struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Address      string             `bson:"address,omitempty" json:"address,omitempty"`
	Penagih      []Penagih          `bson:"penagih,omitempty" json:"penagih,omitempty"`
	Hari_tagihan []string			`bson:"hari_tagihan,omitempty" json:"hari_tagihan,omitempty"`
	
}

type Penagih struct {
	Nama         		string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Gmt        			int      			`bson:"gmt,omitempty" json:"gmt,omitempty"`
	Phone_number 		string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email		 		string             `bson:"email,omitempty" json:"email,omitempty"`
	Hari      			[]string 			`bson:"hari,omitempty" json:"hari,omitempty"`
}

type  Tagihan struct {
	ID           		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Total_tagihan   	string      		`bson:"total_tagihan,omitempty" json:"total_tagihan,omitempty"`
	Deskripsi 			string   			`bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Status				string				`bson:"status,omitempty" json:"status,omitempty"`
	Tanggal_jatuhtempo  primitive.DateTime  `bson:"tanggal_jatuhtempo,omitempty" json:"tanggal_jatuhtempo,omitempty"`
	Biodata      		Nasabah            `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
