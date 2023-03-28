package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertGaji(t *testing.T) {
	gatot := "3500000"
	bon := "1250000"
	pph := "250000"
	pot := "500000"
	biodata := Karyawan{
		Nama:         "Aryka Anisa",
		Phone_number: "6281278239012",
		Jabatan:      "Teknisi Komputer",
	}
	hasil := InsertGaji(gatot, bon, pph, pot, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromJabatan(t *testing.T) {
	jabatan := "Teknisi Komputer"
	biodata := GetKaryawanFromJabatan(jabatan)
	fmt.Println(biodata)
}
