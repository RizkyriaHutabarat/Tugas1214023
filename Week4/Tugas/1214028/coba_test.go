package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertTagihan(t *testing.T) {
	tota := "900000"
	des := "Tagihan Listrik Kos"
	stat := "Lunas"
	biodata := Nasabah{
		Nama : "Fedhira Syaila",
		Phone_number : "62812345128342",
		Address : "Bandung",
	}
	hasil:=InsertTagihan(tota , des , stat , biodata )
	fmt.Println(hasil)
}

func TestGetNasabahFromNama(t *testing.T) {
	nama := "Fedhira Syaila"
	biodata:=GetNasabahFromNama(nama)
	fmt.Println(biodata)
}

