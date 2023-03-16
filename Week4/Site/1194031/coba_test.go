package NPM

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 170.5
	lat := 65.5
	lokasi := "Kampus"
	phonenumber := "6287870491280"
	checkin := "OYO"
	biodata := Karyawan{	
		Nama:         "Rhmtptraaa",
		Phone_number: "6287870491280",
		Jabatan:      "Apa aja",
	}
	hasil := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "1"
	biodata := GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}
