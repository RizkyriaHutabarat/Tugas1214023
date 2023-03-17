# Deployment Aplikasi Backend

Disini kita akan melakukan :

1. Mendaftarkan diri ke github student pack
2. Menyatukan backend menjadi satu repositori
3. Menyatukan frontend menjadi satu repositori
4. Testing deployment boiler plate dari framework iteung
5. Mengatasi error dan kendala masalah umum yang terjadi pada deployment

## Github Student Pack

Student pack berguna untuk mendapatkan beberapa benefit gratis dari github. Untuk mendapatkan student pack. kita mendaftarkan diri dahulu ke alamat https://education.github.com/pack
![image](https://user-images.githubusercontent.com/26703717/225259718-36086fc1-a687-4f43-b3f6-6bbee263efca.png)

Disana terlihat beberapa Experiences yang bisa kita coba secara gratis atau free trial. 
untuk mendaftarkan diri klik Explore Student Programs. Kemudian pilih Student.
![image](https://user-images.githubusercontent.com/26703717/225260384-b6e0db26-f59a-4bff-9faf-08c1522aaffb.png)

Kemudian pilih email kampus ULBI, jika belum memiliki akun email ULBI, bisa meminta ke TIK (Rofi) di ruang Audit Lt.1. 
Selanjutnya adalah melengkapi keterangan mahasiswa dengan melakukan upload KTM, yang bisa diminta ke Iteung.
![image](https://user-images.githubusercontent.com/26703717/225260830-09361909-a5a2-46d8-887f-d664a2c8ebfd.png)

## Heroku Student offer
Heroku merupakan layanan cloud yang berguna untuk deployment atau pemasangan aplikasi berbasis web yang kita buat. Cara pemasangan aplikasi di Heroku cukup mudah, tinggal hubungkan saja dengan repositori github yang sudah ada.
Buat akun heroku (Sign Up) terlebih dahulu, install juga aplikasi : Google Authenticator https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=en&pli=1 atau SalesForce Authenticator https://play.google.com/store/apps/details?id=com.salesforce.authenticator&hl=en&pli=1
*Saya menggunakan SalesForce Authenticator

![image](https://user-images.githubusercontent.com/26703717/225572670-5cfd9499-19c8-4e0a-9f23-9544835e330b.png)

Untuk menikmati layanan heroku secara gratis hubungkan dengan akun github anda yang sudah diaktifkan student pack pada link https://www.heroku.com/github-students
![image](https://user-images.githubusercontent.com/26703717/225280553-aa5c500e-84d2-440b-9adf-073115fa4f12.png)

Kemudian klik Get student offer. Gunakan email kampus anda, dan lanjutkan untuk mendaftarkan diri. Jika dibutuhkan verifikasi pembayaran dengan kartu kredit, bisa menggunakan Virtual Credit Card atau Virtual Debit Card atau Debit Card kartu dari Bank Digital yang bisa didapatkan dengan install aplikasi di handphone. 

Tampilan jika sudah mendaftar Heroku
![image](https://user-images.githubusercontent.com/26703717/225329409-f3d9307b-172f-4655-adc4-8e38358115f2.png)

Tunggu beberapa menit atau jam hingga credit pada heroku bertambah seperti gambar di bawah
![image](https://user-images.githubusercontent.com/26703717/225510369-9da2a2e8-2099-40b4-bcf4-11f3a5e51fdb.png)

## Dashboard Heroku dan Heroku cli
Sebelum mulai kita subscribe eco dyno pada heroku, di menu Billing (catatan : JIKA CREDIT HEROKU MASIH $0 jangan lanjut dulu ke tahap ini)
![image](https://user-images.githubusercontent.com/26703717/225517047-8c697835-602d-4ef5-ac8d-2ea39729d6ef.png)

Setelah login, masuk ke laman https://dashboard.heroku.com/apps. Maka akan muncul list aplikasi yang sudah kita buat. Klik New dan pilih Create new app untuk melakukan deployment aplikasi baru kita.
![image](https://user-images.githubusercontent.com/26703717/225284159-bdd4a4d0-32e7-4887-ad16-a2a9b0f203c7.png)

Masukkan nama aplikasi kita, dan pilih lokasi server kita apakah amerika atau eropa. Kemudian klik Create app.
![image](https://user-images.githubusercontent.com/26703717/225514366-5108a90d-5d83-487c-9d70-2588a656b387.png)

Lakukan instalasi Heroku CLI, untuk menghubungkan komputer kita dengan server heroku. Link Instalasi https://devcenter.heroku.com/articles/heroku-cli
![image](https://user-images.githubusercontent.com/26703717/225332095-7fff6f88-40e1-4014-bbcf-a1d9831324a8.png)

Pilih 64-bit installer jika laptop kalian 64 bit
![image](https://user-images.githubusercontent.com/26703717/225332966-01c6bdfd-9ed7-4ed5-bb0c-24d5eb80f2c7.png)

Lanjutkan sampai selesai langkah instalasinya

![image](https://user-images.githubusercontent.com/15622730/224493999-d208a079-df02-4bcf-b6f6-618c52414d54.png)

Buka git bash dan ketik heroku login, maka akan muncul button login heroku di browser, klik saja login. 
![image](https://user-images.githubusercontent.com/26703717/225334281-53356654-fff8-47b0-b7fc-7c9d7a0b6f3b.png)

Kemudian di gitbash akan ada tulisan Logged in as .....
![image](https://user-images.githubusercontent.com/26703717/225334590-fdbef3bf-969b-4449-9079-754413e31cfa.png)

Kita bisa juga melakukan add SSH key ke heroku kita dengan klik gambar profile di pojok kanan atas > Account Settings > SSH Keys
![image](https://user-images.githubusercontent.com/26703717/225334824-96e4121a-9f95-4f62-92b7-e8bf4699c991.png)

## Deployment Boiler Plate

Disini kita akan mencoba testing deployment ke Heroku. Aplikasi web yang akan dilakukan deploymeny adalah Boiler Plate iTeung yang berada di repo https://github.com/Bukulapak/iteung

![image](https://user-images.githubusercontent.com/26703717/225514611-b89f63c6-ea3a-473c-b95e-26233c0269e0.png)
Kita lakukan fork ke repositori kita, kita beri nama sesuai dengan nama aplikasi di heroku (contoh di atas : ws-ulbi). Kemudian, lakukan clone repo ke komputer kita.
![image](https://user-images.githubusercontent.com/26703717/225514710-82ee3868-c781-438b-a9ba-474ffb28702f.png)

Setelah di clone menggunakan git bash kemudian masuk ke direktori repo di PC kita. Lakukan add remote heroku sesuai nama aplikasi yang sudah kita buat di heroku dengan perintah
```sh
cd ws-ulbi/

heroku git:remote -a ws-ulbi
```
dimana ws-ulbi adalah nama aplikasi kita di heroku.

![image](https://user-images.githubusercontent.com/26703717/225514828-45f95900-dd63-4f34-b871-13ec739d2ae4.png)

Kita cek dulu apakah remote kita sudah sesuai (mengarah ke repositori github kita dan heroku dengan nama app ws-ulbi)
```sh
git remote -v
```
![image](https://user-images.githubusercontent.com/26703717/225514937-85f58f4c-4d99-4133-86ab-b22a530fa7cd.png)

Sekarang buka code editor (VScode atau GoLand)
Selanjutnya kita harus melakukan go mod init, karena terlihat belum ada file gomod dan go.sum di folder repo. Kita lakukan dulu init kemudian go mod tidy dan berhasil.
```sh
go mod init github.com/{username github}/{nama repo}

go mod tidy
```
![image](https://user-images.githubusercontent.com/26703717/225515213-e593866f-0d65-4886-bdfa-4e7c9a15da7c.png)

Muncul error berikutnya. Tenang saja jangan panik. Karena error kita jadi belajar. terlihat tiba tiba terdeteksi package iteung, ini berarti kode program di dalam file masih import nama package yang lama. Tinggal kita ganti semua importnya dengan nama package yang di deklarasikan pada saat kita go mod init tadi.
Kita buka VS Code, kita buka terminal pada bagianPROBLEMS akan mengeluarkan error yang serupa. Selesaikan satu persatu di masing-masing file yang muncul problems.

![image](https://user-images.githubusercontent.com/26703717/225350358-1e2ec3e6-46f7-4f77-895f-ae448ae7ef6a.png)

Masalah selesai dengan mengganti import menuju direktori awal kita melakukan go mod init di ketiga file yaitu main.go, controller.go dan url.go.

![image](https://user-images.githubusercontent.com/26703717/225714926-079be597-a37e-4873-9232-092de17865cd.png)

Jika terdapat error pada controller.go line 22 ubah RunModule menjadi RunModuleLegacy
![image](https://user-images.githubusercontent.com/26703717/225715134-58dc071b-f169-41f7-ae90-d2785c0c4901.png)

Sebelum melakukan push kita cek dulu apakah remote kita sudah sesuai (mengarah ke repositori github kita dan heroku dengan nama app ws-ulbi)
```sh
git remote -v
```
![image](https://user-images.githubusercontent.com/26703717/225715293-cf61b1f6-e9ad-41ae-9ebc-1bdcafb431d4.png)
Pada gambar di atas terlihat bahwa remote sudah sesuai mengarah ke repo kita dan ws-ulbi pada heroku

Sekarang kita tinggal add commit dan push ke github dan heroku kembali.
Tapi, sebelum itu kita cek dulu apakah file yang akan di push sudah sesuai
```sh
git status
```
![image](https://user-images.githubusercontent.com/26703717/225515577-ab43ef4f-c166-4f5c-a4b5-b58e0faf73cf.png)

Jika sudah kita coba git add
```sh
git add .
git status
```
![image](https://user-images.githubusercontent.com/26703717/225515983-174d1d02-daed-4438-bd9e-07252a228d75.png)

Jika sudah hijau semua selanjutnya adalah commit dan push
```sh
git commit -m "KOMENTAR"
```

![image](https://user-images.githubusercontent.com/26703717/225516218-4ea1fa3e-0ce5-4a22-989b-19681600e067.png)
![image](https://user-images.githubusercontent.com/26703717/225516282-d152d463-853e-4f65-822e-b76f46ac2951.png)

Terlihat damai dan tentram akhirnya kita berhasil melakukan deployment ke heroku
![image](https://user-images.githubusercontent.com/26703717/225516506-029c69cf-9c5a-486e-aefb-326e2f937b0a.png)

Oke kita coba buka url aplikasi kita, dicontohkan di heroku CLI disini https://ws-ulbi.herokuapp.com/ kita coba buka apakah sesuai dengan yang diharapkan.
![image](https://user-images.githubusercontent.com/26703717/225516583-0b051e7e-2737-497e-bc09-05f6f28016e7.png)

Dan kita menemukan error selanjutnya. kita coba jalankan heroku logs --tail pada git bash atau terminal
![image](https://user-images.githubusercontent.com/26703717/225516763-62751fca-2147-416f-ab02-52ac556f22da.png)

Katanya Code H14 No Web processes running, berdasarkan hasil penelusuran https://stackoverflow.com/questions/41804507/h14-error-in-heroku-no-web-processes-running hal ini terjadi karena kita belum melakukan set web dynos di aplikasi kita pada dashboard heroku. Oleh karena itu kita kunjungi kembali dashboard heroku kita.

![image](https://user-images.githubusercontent.com/26703717/225517175-5cceb4b6-499b-47b7-a953-7497db45fafc.png)

Pastikan sudah tertera tulisan "Included in Eco subscription" seperti gambar di atas. Pada bagian Menu Resources, dan pada bagian main bin/ws-ulbi kita klik tombol tanda pensil, kita geser tombol menjadi aktif dan klik tombol confirm

![image](https://user-images.githubusercontent.com/26703717/225517488-1e25ed4b-5242-4e24-a8fe-ee988ffa8c1d.png)

Dan akhirnya web sudah bisa diakses.

![image](https://user-images.githubusercontent.com/26703717/225518335-4d11e657-724f-41ca-9d81-abb89a288d6e.png)

## Iteung Boiler Plate

### Aktifasi Prefork

![image](https://user-images.githubusercontent.com/11188109/223247620-782c1571-d0e8-4f2a-abd4-89a52f457d69.png)

Sekarang aplikasi sudah berjalan dengan baik. Akan tetapi ada beberapa hal yang harus di konfigurasi kembali. Yang paling pertama untuk dilakukan konfigurasi dari boiler plate iteung adalah melakukan aktifasi Prefork. Kita bisa lihat di heroku logs, terlihat Prefork statusnya masih Disabled. Kita buka file main.go tambahkan config.Iteung pada fungsi fiber.New()

```go
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
```

Kemudian kita buat file config.go di dalam folder config yang berisi :

```go
package config

import "github.com/gofiber/fiber/v2"

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Iteung",
	AppName:       "Message Router",
}
```

Simpan add, commit push kembali ke repo dan heroku. Terlihat di heroku logs, prefork sudah aktif

![image](https://user-images.githubusercontent.com/11188109/223252919-aacc55b5-897d-4b38-a55c-3a54827b6dff.png)

### Free MariaDB JawsDB

Kita akan menambahkan add on MariaDB dari heroku. Pertama dari dashboard kita klik Find more add-ons dan pilih JawsDB Maria https://elements.heroku.com/addons/jawsdb-maria

![image](https://user-images.githubusercontent.com/26703717/225518941-c6078362-899a-4acb-b5f1-0ea5c4efa7a4.png)

Klik Install JawsDB Maria

![image](https://user-images.githubusercontent.com/26703717/225519104-ab0d00a5-d812-4637-8908-d1ebd0fcc562.png)

Pilih nama aplikasi kita di heroku dan klik Submit Order Form

![image](https://user-images.githubusercontent.com/26703717/225519280-c5266b1b-770e-497e-826c-84bbf236b942.png)

JawsDB akan terinstall, untuk memakainya kita hanya butuh Connection String yang bisa di dapatkan dengan klik JawsDB Maria di bagian Add-ons

![image](https://user-images.githubusercontent.com/26703717/225519487-62f1d233-38cf-4bbf-9aa4-964ffa59b931.png)

![image](https://user-images.githubusercontent.com/26703717/225519764-226a3a37-6f57-40e2-9f3f-1758b2334699.png)

### Edit Environtment Variabel

Jika pada windows kita melakukan edit di environtment variabel. Begitu juga heroku memiliki fitur tersebut yang bernama config vars. Tinggal kita masuk ke Dashboard bagian Settings > Config Vars > Reveal Config Vars.

![image](https://user-images.githubusercontent.com/26703717/225520354-fee0411d-9c18-44cf-967f-98570004793e.png)

Kita masukkan beberapa environment yang didapat pada folder config dari boiler plate iteung seperti :

file db.go
```env
ITEUNGBEV1
MONGOSTRING
MARIASTRINGAKADEMIK
```
file token.go
```env
PUBLICKEY
PRIVATEKEY
```
file cors.go
```env
INTERNALHOST
PORT
```
file api.go
```env
URLAPIWABUTTON
```

Untuk saat ini kita tambahkan MARIASTRINGAKADEMIK dan MONGOSTRING terlebih dahulu contoh :
```env
MARIASTRINGAKADEMIK=username:password@tcp(ao9moanwus0rjiex.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/
MONGOSTRING=mongodb+srv://username:password@cluster0.wghp85v.mongodb.net/test
```
![image](https://user-images.githubusercontent.com/26703717/225520322-d4937d5a-bbea-47a8-9678-887a66e022b0.png)

Edit dan tambahkan konfigurasi nama database yang dipakai pada file db.go folder config (untuk DBName pada Mongo diubah saja menjadi tes_db seperti di pertemuan ke-4)

```go
var DBUlbimariainfo = atdb.DBInfo{
	DBString: MariaStringAkademik,
	DBName:   "NAMADBMASINGMASING",
}

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "NAMA DB PADA MONGO COMPASS",
}

var Ulbimariaconn = atdb.MariaConnect(DBUlbimariainfo)

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)
```
![image](https://user-images.githubusercontent.com/26703717/225521525-35277006-a658-453f-8ae3-f6a7e65dc50d.png)

Jika terdapat error seperti diatas maka modif Connection string menjadi seperti berikut :
```env
MARIASTRINGAKADEMIK=username:password@tcp(xxxxxxxxxxxxxxxx.xxxxxxxxxxxxxxxx.us-east-1.rds.amazonaws.com:3306)/
```
Maka akan seperti ini
![image](https://user-images.githubusercontent.com/26703717/225521340-80c9b861-9442-471c-b0e9-ba7ddf9221ca.png)


### Memanggil package golang

Dari package yang sudah dibuat pada chapter sebelumnya kita akan coba panggil dari aplikasi boiler plate iteung yang sudah di deploy di heroku. Sebagai contoh kita akan mencoba memanggil package musik di https://github.com/aiteung/musik hal yang pertama kali adalah kita membuka file url.go di folder url kita tambahkan baris di dalam func Web sebuah baris kode yang memanggil controller.Homepage yaitu fungsi yang akan kita buat :


pada terminal ketikkan perintah go get package yang akan digunakan
```sh
go get github.com/aiteung/musik
```

pada file url.go
```go
page.Get("/", controller.Homepage) //ujicoba panggil package musik
```
![image](https://user-images.githubusercontent.com/26703717/225522437-705f04f8-d018-4579-af47-8545a80fdae6.png)

kemudian kita buat file coba.go di folder controller kita tambahkan fungsi Homepage
```go
func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
```
![image](https://user-images.githubusercontent.com/26703717/225530005-0f8f4aa2-a642-4672-abca-92a2d349a678.png)

Jangan lupa selalu menjalankan perintah 
```sh
go mod tidy
```
kemudian lakukan commit dan push pada repo github dan heroku.

Jika sudah, kita buka alamat aplikasi dan terlihat ip address dari host heroku
![image](https://user-images.githubusercontent.com/26703717/225522344-04810ba4-47ce-4bd1-91c4-cd0f75e71a49.png)

### Mongodb.com Network Access

Agar aplikasi di heroku bisa mengakses mongodb menggunakan seluruh IP Address

![image](https://user-images.githubusercontent.com/11188109/223330948-06b8493d-a9b7-4f19-86ba-06c35d6aaa96.png)

Setelah itu, silahkan masukkan data dummy pada collection presensi dengan mongo compass pada minggu sebelumnya

### Deployment Package

Pada chapter sebelumnya. Package yang sudah dibuat bisa kita panggil di controller. Cukup dengan 3 tahap, yaitu :
1. Definisikan alamat URL untuk akses package beserta nama fungsi di controller yang akan kita buat pada file url.go
   ```go
   page.Get("/presensi", controller.GetPresensi)
   ```
2. Go get package yang akan digunakan melalui terminal di vscode
   ```sh
   go get github.com/aiteung/presensi
   ```
3. Buat fungsi di file coba.go
   ```go
   func GetPresensi(c *fiber.Ctx) error {
        ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
        return c.JSON(ps)
   }
   ```
   ![image](https://user-images.githubusercontent.com/26703717/225530968-9343f59f-7184-4cd6-b232-ef5caedab797.png)
   
Jangan lupa selalu menjalankan perintah 
```sh
go mod tidy
```
Kemudian add, commit push ke repo dan heroku. Kemudian kita akses melaui url yang kita deklarasikan di file url.go. Untuk method GET bisa menggunakan browser saja, untuk POST menggunakan POSTMAN

![image](https://user-images.githubusercontent.com/26703717/225527396-66eba03d-7dd8-43ef-8320-53be5500b6ef.png)

## Kumpulkan 
Buat file README.md dan masukkan link repo boilerplate masing-masing mahasiswa dan url heroku di folder Week5/Site/{NPM}

