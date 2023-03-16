package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Sekolah"
	phonenumber := "089647129890"
	checkin := "masuk"
	biodata := Karyawan{
		Nama : "Dimas Ardianto",
		Phone_number : "089647129890",
		Jabatan : "Mahasiswa",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "0896647129890"
	biodata:=GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

