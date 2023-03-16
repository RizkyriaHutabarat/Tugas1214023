package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah Saya"
	phonenumber := "681219882869"
	checkin := "masuk"
	biodata := Karyawan{
		Nama : "Rizkyria",
		Phone_number : "6281219882869",
		Jabatan : "Dosen",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata:=GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

