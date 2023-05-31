package presensiMahasiswa

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	checkin := "masuk"
	biodata := Mahasiswa{
		Nama : "Erdito Nausha Adam",
		Npm : 1214031,
		Kelas : "2B",
		Jurusan : "D4 Teknik Informatika",
	}
	matkul := MataKuliah{
		Nama_matkul : "Pemrograman 3",
		Kode_matkul : 21711,
		Jadwal_kuliah : JadwalKuliah{
			Jam_masuk : "13:30",
			Jam_keluar : "17:00",
			Hari : "Jum'at",
		},
	}
	hasil:=InsertPresensi(checkin , biodata, matkul )
	fmt.Println(hasil)
}

func TestGetMahasiswaFromNpm(t *testing.T) {
	npm := 1214038
	biodata:=GetMahasiswaFromNpm(npm)
	fmt.Println(biodata)
}
func TestGetMatkulFromKodeMatkul(t *testing.T) {
	kodeMk := 21711
	biodata:=GetMatkulFromKodeMatkul(kodeMk)
	fmt.Println(biodata)
}

