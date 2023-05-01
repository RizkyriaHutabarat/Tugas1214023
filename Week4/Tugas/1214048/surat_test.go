package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertSurat(t *testing.T) {
	NoSurat := 12143113
	Status := Status{
		Kode_status:  12,
		Status_surat: "Terkirim",
	}
	Perihal := "surat tugas"
	Kodepos := Kodepos{
		Kode_pos:    33412,
		Nama_daerah: "Bandung",
	}
	Pengirim := Pengirim{
		Nama_pengirim: "ucok",
		Tanggal_kirim: "12 Januari 2023",
	}
	Penerima := Penerima{
		Nama_penerima:  "Dani",
		Tanggal_terima: "20 februari 2023",
	}

	hasil := InsertSurat(NoSurat, Status, Perihal, Kodepos, Pengirim, Penerima)
	fmt.Println(hasil)
}

func TestGetSuratFromNoSurat(t *testing.T) {
	NoSurat := 12143113
	surat := GetSuratFromNoSurat(NoSurat)
	fmt.Println(surat)
}
