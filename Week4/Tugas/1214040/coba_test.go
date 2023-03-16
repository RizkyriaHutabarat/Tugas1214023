package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Kampus"
	phonenumber := "682435678623"
	checkin := "masuk"
	biodata := Mahasiswa{
		Nama : "Marlina Lubis",
		Phone_number : "6284623996425",
		Jurusan : "D4 Teknik Informatika",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phonenumber := "682435678623"
	biodata:=GetMahasiswaFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

