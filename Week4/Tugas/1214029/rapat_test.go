package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertWaktu(t *testing.T) {
	jam := "10.00"
	hari := "Jumat"
	tanggal := "10/03/2023"
	hasil:=InsertWaktu(jam, hari, tanggal )
	fmt.Println(hasil)
}

func TestGetTamuFromJam(t *testing.T) {
	jam := "10.00"
	hari:=GetTamuFromJam(jam)
	fmt.Println(hari)
}