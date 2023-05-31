package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "085718177810"
	checkin := "masuk"
	biodata := Karyawan{
		Nama : "Drake",
		Phone_number : "085718177810",
		Jabatan : "Rakyat Biasa",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "085718177810"
	biodata:=GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

