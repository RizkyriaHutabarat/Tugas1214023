package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pimpinan struct {
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Gaji         []Gaji				`bson:"gaji,omitempty" json:"gaji,omitempty"`
	Karyawan	 Karyawan           `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
	Bonus        Jam_Kerja          `bson:"Bonus,omitempty" json:"Bonus,omitempty"`
	Biodata      Karyawan           `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
	Presensi     Presensi           `bson:"presensi,omitempty" json:"presensi,omitempty"`
}

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Hari_gaji    []string           `bson:"hari_gaji,omitempty" json:"hari_gaji,omitempty"`
	Presensi	 Presensi	        `bson:"presensi,omitempty" json:"presensi,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Lokasi		 Lokasi				`bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Waktu		 Waktu				`bson:"waktu,omitempty" json:"waktu,omitempty"`
}

type Waktu struct {
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Jam        string   `bson:"jam,omitempty" json:"jam,omitempty"`
	Tanggal	   string   `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Jam_Kerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`

}

type Gaji struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tanggal       primitive.DateTime `bson:"Tanggal,omitempty" json:"Tanggal,omitempty"`
	Gaji_pokok    string             `bson:"Gaji_pokok,omitempty" json:"Gaji_pokok,omitempty"`
	Bonus         Jam_Kerja           `bson:"Bonus,omitempty" json:"Bonus,omitempty"`
	Tunjangan_pph string             `bson:"Tunjangan_pph,omitempty" json:"Tunjangan_pph,omitempty"`
	Potongan      string             `bson:"Potongan,omitempty" json:"Potongan,omitempty"`
	Biodata       Karyawan           `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
	Presensi      Presensi           `bson:"presensi,omitempty" json:"presensi,omitempty"`
}

type Bendahara struct {
	Nama         string   `bson:"nama,omitempty" json:"nama,omitempty"`
	Email        string   `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number string   `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Hari         string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Biodata 	 Karyawan `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
	Gaji		 Gaji 	  `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
