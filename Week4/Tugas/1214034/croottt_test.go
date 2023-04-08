package main

import (
	"fmt"
	"testing"
)

func TestInsertNilai(t *testing.T) {
    nama_ms := "Kai D Jayson"
    npm := 1214034
    phone_number := 2129119
    kategori := Matakuliah{
        Nama_MK      : "Etika & Manajemen Profesi IT", 
        SKS        : 2,
        Jam_masuk  : "8.40",
        Jam_keluar : "10.20",
        Hari : "Selasa",
    }
    biodata := Jurusan{
        Nama_j : "D4 Teknik Informatika",
    }
    hasil:=InsertNilai(nama_ms ,npm ,phone_number , kategori, biodata )
    fmt.Println(hasil)
}

func TestGetMahasiswaFromNPM(t *testing.T) {
    npm := 1214034
    biodata:=GetMahasiswaFromNPM(npm)
    fmt.Println(biodata)
}