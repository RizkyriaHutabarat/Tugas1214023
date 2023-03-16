package main

import (
	"fmt"
	"testing"
)

func TestInsertNilai(t *testing.T) {
	tugas1 := 80
	tugas2 := 70
	tugas3 := 50
	tugas4 := 90
	tugas5 := 100
	uts	   := 70
	uas    := 70
	kategori := Matakuliah{
		Nama_MK	  : "Algoritma", 
		SKS        : 3,
		Jam_masuk  : "7.00",
		Jam_keluar : "10.00",
		Hari : "Jumat",
	}
	biodata := Mahasiswa{
		Nama : "Budiman",
		NPM : 115611,
		Phone_number : "628456456211",
	}
	hasil:=InsertNilai(tugas1 ,tugas2 ,tugas3 ,tugas4 ,tugas5, uts, uas, kategori, biodata )
	fmt.Println(hasil)
}

func TestGetMahasiswaFromUAS(t *testing.T) {
	uas := 70
	biodata:=GetMahasiswaFromUAS(uas)
	fmt.Println(biodata)
}

