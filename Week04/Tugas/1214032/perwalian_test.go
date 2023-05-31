package perwalian

import (
	"fmt"
	"testing"
)

func TestInsertWaktu(t *testing.T) {
	jam := "9 Pagi"
	hari := 	"Selasa"
	tanggal := "15 Maret 2023"
	hasil:=InsertWaktu(jam , hari , tanggal )
	fmt.Println(hasil)
}

func TestGetMahasiswaFromJam(t *testing.T) {
	jam := "9 pagi"
	biodata:=GetMahasiswaFromJam(jam)
	fmt.Println(biodata)
}

