package tugas

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	lokasi := "Chemical Lab no.310"
	phonenumber := "62120049978"
	checkin := "masuk"
	biodata := Peneliti{
		Nama:         "Salsabila Irbah",
		Phone_number: "62120049978",
		Jabatan:      "Chemical Analyst",
	}
	hasil := InsertPresensi(lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetPenelitiFromPhoneNumber(t *testing.T) {
	phonenumber := "62120049978"
	biodata := GetPenelitiFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}
