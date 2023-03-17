package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertKemahasiswaan(t *testing.T) {
	identitas := Mahasiswa{
		Npm:  "1214054",
		Nama: "Dimas Ardianto",
		Prodi: ProgramStudi{
			Kode_Prodi: "012021",
			Nama_Prodi: "D4",
		},
		Jurusan: "Teknik Informatika",
		Kelas:   "2B",
	}
	status_keuangan := Keuangan{
		Total_Pembayaran: 7700000,
	}
	nilai_mhs := Nilai{
		Matakuliah: Matakuliah{
			Nama_Matkul: "Pemrograman 3",
			Nama_Dosen:  "Indra Riksa",
		},
		Nilai_Angka: 90,
		Nilai_Huruf: "A",
	}
	hasil := InsertKemahasiswaan(identitas, status_keuangan, nilai_mhs)
	fmt.Println(hasil)
}

func TestGetKemahasiswaanFromNpm(t *testing.T) {
	npm := "1214054"
	identitas := TestGetKemahasiswaanFromNPM(npm)
	fmt.Println(identitas)
}
