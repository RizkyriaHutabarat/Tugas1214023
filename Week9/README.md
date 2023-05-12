# Insert Pada Aplikasi Frontend

Masih menggunakan repo yang sama pada Chapter06 yaitu repo untuk frontend

# Backend Package
Buka project backend pada VScode atau GoLand
Pada folder module/{nama file}.go tambahkan code berikut (jika function sudah ada, replace saja)
```go
func InsertPresensi(db *mongo.Database, col string, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
		"longitude":    long,
		"latitude":     lat,
		"location":     lokasi,
		"phone_number": phonenumber,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"checkin":      checkin,
		"biodata":      biodata,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
```
Kode di atas berfungsi untuk menyisipkan data presensi karyawan ke dalam database MongoDB. Fungsi ini menerima beberapa parameter seperti db yang merepresentasikan koneksi ke database MongoDB, col yang merepresentasikan nama koleksi atau tabel di dalam database, long dan lat yang merupakan koordinat longitude dan latitude dari lokasi presensi, lokasi yang merepresentasikan nama lokasi presensi, phonenumber yang merepresentasikan nomor telepon karyawan, checkin yang merepresentasikan waktu check-in, dan biodata yang berisi data diri karyawan seperti nama, jabatan, jam kerja, dan hari kerja.

Fungsi InsertPresensi akan membuat sebuah dokumen baru di dalam koleksi col dengan menggunakan data yang diterima sebagai parameter. Dokumen ini memiliki beberapa field seperti longitude, latitude, location, phone_number, datetime, checkin, dan biodata. Field datetime akan diisi dengan waktu saat ini menggunakan format UTC. Setelah dokumen berhasil disisipkan ke dalam koleksi, fungsi akan mengembalikan _id dari dokumen yang baru disisipkan sebagai insertedID. Jika terjadi kesalahan dalam proses penyisipan, maka fungsi akan mengembalikan pesan error melalui variabel err.

Kemudian pada file {namafile}_test.go isikan code berikut untuk melakukan testing insert data
```go
func TestInsertPresensi(t *testing.T) {
	var jamKerja1 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "08:00",
		Jam_keluar: "16:00",
		Gmt:        7,
		Hari:       []string{"Senin", "Rabu", "Kamis"},
		Shift:      1,
		Piket_tim:  "Piket A",
	}
	var jamKerja2 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "09:00",
		Jam_keluar: "17:00",
		Gmt:        7,
		Hari:       []string{"Sabtu"},
		Shift:      2,
		Piket_tim:  "",
	}

	long := 98.345345
	lat := 123.561651
	lokasi := "New York"
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:        "George Best",
		PhoneNumber: "6284564562",
		Jabatan:     "Football Player",
		Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:  []string{"Senin", "Selasa"},
	}
	insertedID, err := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
```

Rapihkan dependensi
```go
go mod tidy
```

Jalankan testing
```go
go test -run ^TestInsertPresensi$
```
Kemudian lakukan git tag dan go list agar package terupdate, tutor [disini](https://github.com/indrariksa/WS/tree/main/Week4#publish-package)

# Boilerplate
Buka project boilerplate pada VScode atau GoLand, pada terminal ketikkan perintah go get package backend yang sebelumnya kita update
```go
go get {url package backend di pkg.go.dev}

contoh : go get github.com/indrariksa/be_presensi
```

Kemudian pada file coba.go di folder controller kita tambahkan fungsi InsertData
```go
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodul.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
```
inimodul dan inimodel merupakan alias dari import package backend, bisa saja berbeda dengan kalian
![image](https://user-images.githubusercontent.com/26703717/228791775-9a76c3a7-10db-443b-af9e-0ff07a39f80d.png)

Function InsertData di atas digunakan untuk memproses data yang dikirimkan dari sebuah request POST untuk melakukan insert data ke dalam database MongoDB dengan menggunakan method InsertPresensi yang ada di package inimodul. Data yang di-insert adalah data presensi yang terdiri dari informasi lokasi (longitude, latitude, location), nomor telepon (phone_number), waktu check-in (checkin), dan biodata karyawan (biodata). Jika proses insert data berhasil dilakukan, maka function akan mengembalikan status OK dan pesan berhasil disimpan beserta inserted ID. Namun, jika terjadi error selama proses insert data, function akan mengembalikan status Internal Server Error dan pesan error tersebut.

Pada file url.go di folder url kita tambah path baru
```go
page.Post("/ins", controller.InsertData)
```
Kode di atas mendefinisikan sebuah rute POST pada path /ins yang akan di-handle oleh fungsi InsertData yang berada pada package controller. Ketika request POST dengan path /ins diterima, Fiber akan menjalankan fungsi InsertData pada package controller.

Jangan lupa selalu menjalankan perintah
```go
go mod tidy
```

Kemudian lakukan commit dan push pada repo github dan heroku.

Jika sudah, kita coba testing terlebih dahulu via postman sebelum kita implementasi pada frontend. Pada Postman gunakan method POST, masukkan endpoint heroku, kemudian pilih Body->raw->JSON isikan dengan kode berikut
```json
{
    "longitude" : 91.32313,
    "latitude" : 125.541278,
    "location" : "Setiabudhi",
    "phone_number" : "628451231876",
    "checkin" : "MASUK",
    "biodata" : {
        "nama": "David Ginola",
		"phone_number": "628451231876",
		"jabatan": "CEO",
		"jam_kerja": [{
                "durasi":     8,
                "jam_masuk":  "08:00",
                "jam_keluar": "16:00",
                "gmt":        7,
                "hari":       ["Senin", "Rabu", "Kamis"],
                "shift":      1,
                "piket_tim":  "Piket A"
            }],
		"hari_kerja":  ["Senin", "Kamis"]
	}
}
```
![image](https://user-images.githubusercontent.com/26703717/228794604-b6f8f72b-8a58-4fde-938b-d005fe190f9a.png)

Klik send, maka response terlihat seperti berikut
![image](https://user-images.githubusercontent.com/26703717/228794786-6f293384-db6a-4d1e-affe-822cf0a5540d.png)

Pada tahap ini post data sudah aman, selanjutnya kita akan melakukan post data pada frontend

# Frontend
## post.js
Hal yang pertama kali kita lakukan adalah membuka file insert.html dan menambahkan tag di dalam dan akhir section body untuk memanggil file post.js.
```html
<script type="module" src="../js/controller/post.js"></script>
```
![image](https://user-images.githubusercontent.com/26703717/228773157-b31b2fda-5909-48b1-a3d3-5e07f2d1e824.png)

Selanjutnya kita buat file post.js di dalam folder js/controller yang berisi.
```js
import { postData } from "https://bukulapak.github.io/api/process.js";
import { onClick, getValue } from "https://bukulapak.github.io/element/process.js";
import { Urlpost, AmbilResponse} from "../config/url_post.js";


function pushData(){
    var hari_kerja = getValue("hari_kerja");

    let data = {
        longitude : parseFloat(getValue("longitude")),
        latitude : parseFloat(getValue("latitude")),
        location : getValue("location"),
        phone_number : getValue("phone_number"),
        checkin : getValue("checkin"),
        biodata : {
            nama : getValue("nama"),
            jabatan : getValue("jabatan"),
            jam_kerja : [{
                durasi : parseInt(getValue("durasi")),
                jam_masuk : getValue("jam_masuk"),
                jam_keluar : getValue("jam_keluar"),
            }],
            hari_kerja : hari_kerja.split(",")
        }
    }
    postData(Urlpost, data, AmbilResponse);

}

onClick("button", pushData);
```
Kode di atas mendefinisikan sebuah fungsi pushData() yang akan dipanggil saat tombol di klik. Pada saat tombol di klik, fungsi ini akan mengambil nilai dari beberapa elemen pada halaman web menggunakan fungsi getValue() yang didefinisikan pada file "https://bukulapak.github.io/element/process.js". Nilai-nilai tersebut kemudian digunakan untuk membentuk sebuah objek data dan dikirim ke sebuah URL melalui fungsi postData() yang didefinisikan pada file "https://bukulapak.github.io/api/process.js". Jadi kalian tidak perlu lagi menuliskan ulang untuk kode tersebut karena kita sudah melakukan import. Bisa dilihat ada import pada repo bukulapak salah satunya yaitu fungsi postData, bisa kalian [cek url](https://bukulapak.github.io/api/process.js)

![image](https://user-images.githubusercontent.com/26703717/228774721-e6af8caf-ca60-425a-8f53-977753063e59.png)

Fungsi kode di atas adalah untuk mengirim data dalam format JSON ke server menggunakan metode POST melalui API fetch() pada JavaScript.Kode tersebut mengambil tiga parameter:
* target_url : alamat URL tujuan untuk mengirimkan data.
* datajson : objek JavaScript yang akan dikirim dalam format JSON.
* responseFunction : fungsi yang akan dijalankan setelah menerima respons dari server.

Dalam fungsi postData(), terdapat pengaturan header request seperti content-type, metode POST, body data yang akan dikirim, dan pengaturan redirect. Kemudian, kode akan menggunakan metode fetch() untuk mengirim request ke server. Setelah menerima respons dari server, fungsi responseFunction akan dijalankan dengan parameter result yang berisi data respons dari server dalam format JSON.

Setelah data berhasil dikirim, fungsi AmbilResponse() akan dipanggil untuk mengolah respon yang diterima dari server.

## url_post.js
Selanjutnya kita buat file url_post.js di dalam folder js/config yang berisi.
```js
export let urlPOST = "https://ws-ulbi.herokuapp.com/ins"

export function AmbilResponse(result) {
    console.log(result); //menampilkan response API pada console
    alert(result.message); //menampilkan response API pada alert
    window.location.reload(); //reload halaman setelah klik ok pada alert
}
```
* urlPOST: Variabel ini menyimpan URL endpoint dari API yang akan dipanggil saat melakukan POST request.
* AmbilResponse(result): Fungsi ini merupakan callback function yang akan dipanggil setelah mendapatkan respons dari API. Fungsi ini akan melakukan tiga hal:
  * Menampilkan respons API pada console melalui console.log(result).
  * Menampilkan respons API pada alert melalui alert(result.message).
  * Mereload halaman setelah pengguna menekan OK pada alert melalui window.location.reload().

## Inisialisasi id
Cari element dari input dan buttonnya kemudian kita beri id
* pada tag <button> beri id="button"
* pada tag <input> beri id sesuai dengan json pada mongoDB

![image](https://user-images.githubusercontent.com/26703717/228778229-2ec1a71c-630c-4b54-9ecc-87ae75e3eda5.png)
    
Kemudian kita commit dan push, kemudian tunggu hingga github pages sudah terdeploy dengan baik
    
Lakukan insert data seperti berikut
![image](https://user-images.githubusercontent.com/26703717/228781661-b8fc9741-649f-434d-a70e-29c9238308b4.png)

Klik button Save, maka akan muncul alert
![image](https://user-images.githubusercontent.com/26703717/228781857-7eb73846-02d8-4efa-a984-bee15ee50553.png)

Data yang diinputkan berhasil masuk
![image](https://user-images.githubusercontent.com/26703717/228782155-9be0020a-5416-4e7e-b403-fb52d4939ed6.png)


