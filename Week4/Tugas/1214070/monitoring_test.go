package monitoring

import (
	"fmt"
	"testing"
)

func TestInsertOrangTua(t *testing.T) {
	nama_ot := "Rini"
	phone_number := "0888321157026"
	anak := Mahasiswa{
		Nama : "Farel",
		NPM : 1214070,
		Phone_number : "083821157026",
		
	}
	hasil:=InsertOrangTua(nama_ot , phone_number , anak )
	fmt.Println(hasil)
}

func TestGetTemaFromPhoneNumber(t *testing.T) {
	phone_number := "0888321157026"
	anak :=GetTemaFromPhoneNumber(phone_number)
	fmt.Println(anak)
}

