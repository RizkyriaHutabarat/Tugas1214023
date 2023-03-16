package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertSurat(t *testing.T) {
	NoSurat := 12143113
	Header := "surat tugas"
	Tujuan := "surabaya"
	Kodepos := 12146
	Pengirim := Pengirim{
		Nama_penerima:  "ucok",
		Tanggal_keluar: "12 Januari 2023",
	}
	Penerima := Penerima{
		Nama_pengirim: "Dani",
		Tanggal_masuk: "20 februari 2023",
	}

	hasil := InsertSurat(NoSurat, Header, Tujuan, Kodepos, Pengirim, Penerima)
	fmt.Println(hasil)
}

func TestGetSuratFromNoSurat(t *testing.T) {
	NoSurat := 12143113
	surat := GetSuratFromNoSurat(NoSurat)
	fmt.Println(surat)
}
