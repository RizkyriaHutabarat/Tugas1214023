package intership

import (
	"fmt"
	"testing"
)

func TestInsertIntership(t *testing.T) {
	npm := "1194031"
	kelas := "D4 TI 4A"
	kampus := "POLITEKNIK POS INDONESIA"
	judul := "APA AJA"
	biodata := Mahasiswa{
		Nama:         "R Bimantoro Putra",
		Antrian_sidang: "Tahap 3",
		Hari_sidang:      "Senin",
	}
	hasil := InsertIntership(npm,kelas,kampus,judul,biodata)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "R Bimantoro Putra"
	biodata := GetKaryawanFromNama(nama)
	fmt.Println(biodata)
}