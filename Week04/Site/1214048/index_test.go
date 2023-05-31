package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "083116281797"
	checkin := "masuk"
	biodata := Karyawan{
		Nama:         "Ade Candra",
		Phone_number: "08456456216",
		Jabatan:      "Ketua MPR RI",
	}
	hasil := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "083116281797"
	biodata := GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}
