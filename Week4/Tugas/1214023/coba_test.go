package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertProfile(t *testing.T) {
	phone := "6281219882869"
	almt := "perum bpa"
	log := "berhasil"
	biodata := People{
		Nama : "Rizkyria",
		Phone_number : "6281219882869",
		Tanggal_lahir: "18 November 2003",
	}
	hasil:=InsertProfile(phone , almt , log , biodata )
	fmt.Println(hasil)
}

func TestPeopleFromPhoneNumber(t *testing.T) {
	phonenumber := "6281219882869"
	biodata:=GetPeopleFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

