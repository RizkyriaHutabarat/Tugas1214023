package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPendaftaran(t *testing.T) {
	kdpendaftar := 1
	biodata := Camaba{
		Ktp : 320132321321,
		Nama : "Dito Adam",
		Phone_number : "085718177810",
		Address : "Parongpong, Kab. Bandung Barat",
	}
	asalsekolah := DaftarSekolah{
		KDSekolah : 01,
		Nama : "SMK Negeri 1 Cirebon",
		Phone_number : "085718172053",
		Address : "Jl.Perjuangan, Kota Cirebon",
	}
	jurusan := Jurusan{
		KDJurusan : "D4TI",
		Nama : "SMK Negeri 1 Cirebon",
		Jenjang : "Diploma 4",
	}
	jalur := "Rapot"
	alulbi := "Universitas Internasional"
	aljurusan := "Sedang trand"
	hasil:=InsertPendaftaran(kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
	fmt.Println(hasil)
}

func TestInsertDaftarCamaba(t *testing.T) {
	ktp := 232312312312
	nama := "Dito"
	phone_number := "085725722483"
	alamat := "Kota Bandung"
	hasil:=InsertDaftarCamaba(ktp, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarSekolah(t *testing.T) {
	kodesklh := 4
	nama := "SMA Negeri 1 Bandung"
	phone_number := "085725722483"
	alamat := "Kota Bandung"
	hasil:=InsertDaftarSekolah(kodesklh, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarJurusan(t *testing.T) {
	kodejurusan := "D3TI"
	nama := "Teknik Informatika"
	jenjang := "Diploma 3"
	hasil:=InsertDaftarJurusan(kodejurusan, nama, jenjang)
	fmt.Println(hasil)
}

// test getFunction

func TestGetPendaftaranFromKTP(t *testing.T) {
	ktp := 320132321321
	pendaftar:=GetPendaftaranFromKTP(ktp)
	fmt.Println(pendaftar)
}

func TestGetCamabaFromPhoneNumber(t *testing.T) {
	phonenumber := "085725722483"
	camaba:=GetCamabaFromPhoneNumber(phonenumber)
	fmt.Println(camaba)
}

func TestGetDaftarSekolahFromKDSekolah(t *testing.T) {
	kdsekolah := 1
	daftar_sekolah:=GetDaftarSekolahFromKDSekolah(kdsekolah)
	fmt.Println(daftar_sekolah)
}

func TestGetJurusanFromKDJurusan(t *testing.T) {
	kdjurusan := "D3TI"
	daftar_jurusan:=GetJurusanFromKDJurusan(kdjurusan)
	fmt.Println(daftar_jurusan)
}



