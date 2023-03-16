package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertKuesioner(t *testing.T) {
	long := -6.873165194466392
	lat := 107.57589801606773
	lokasi := "ULBI"
	email := "juwitastefany13@gmail.com"
	status := "selesai"
	biodata := Responden{
		Nama : "Juwita Stefany",
		Email : "juwitastefany13@gmail.com",
		Jenis_kelamin : "perempuan",
	}
	hasil:=InsertKuesioner(long ,lat , lokasi , email , status , biodata )
	fmt.Println(hasil)
}

func TestGetRespondenFromEmail(t *testing.T) {
	email := "juwitastefany13@gmail.com"
	biodata:=GetRespondenFromEmail(email)
	fmt.Println(biodata)
}

