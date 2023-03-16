package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertProyek1(t *testing.T) {
	biodata_mahasiswa := Mahasiswa{
		NPM:   1214049,
		Nama:  "Auliana Fahrian",
		Kelas: "2B",
		Jurusan: Jurusan{
			Nama: "Teknik Informatika",
		},
		Prodi: Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}
	dosen_pembimbing := Dosen{
		Nama: "Dimas",
		Prodi: Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}

	dosen_penguji := Dosen{
		Nama: "Dani",
		Prodi: Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}
	Judul := "Perangkat Keras"
	tanggal_sidang := "2023-12-12"
	hasil := InsertProyek1(biodata_mahasiswa, dosen_pembimbing, dosen_penguji, Judul, tanggal_sidang)
	fmt.Println(hasil)
}

func TestGetProyek1FromNPM(t *testing.T) {
	npm := 1214049
	hasil := GetProyek1FromNPM(npm)
	fmt.Println(hasil)
}
