package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Bengkel"
	phonenumber := "68122221814"
	checkin := "masuk"
	biodata := Karyawan{
		Nama : "Kai D Jayson",
		Phone_number : "2129119",
		Jabatan : "Pimpinan",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "081386780708"
	biodata:=GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

