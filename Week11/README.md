# Dokumentasi Golang Fiber menggunakan Swagger

## Intro
Swagger adalah sebuah alat yang digunakan untuk mendokumentasikan dan menguji API. Dengan menggunakan Swagger, kita dapat membuat dokumentasi API yang lengkap dan interaktif secara otomatis berdasarkan kode sumber aplikasi.

Golang Fiber adalah sebuah framework web yang ringan dan cepat untuk bahasa pemrograman Go. Fiber menyediakan fitur-fitur yang kuat untuk membangun aplikasi web, termasuk dukungan untuk routing, middleware, dan pengelolaan respons HTTP yang efisien.

Menggabungkan Swagger dengan Golang Fiber memungkinkan kita untuk secara otomatis menghasilkan dokumentasi API yang komprehensif dan mudah diakses. Dengan menggunakan anotasi atau tag khusus dalam kode sumber aplikasi, kita dapat menentukan detail rute API, metode yang didukung, parameter yang diterima, dan respons yang dihasilkan.

Beberapa langkah yang umumnya dilakukan untuk menggunakan Swagger dengan Golang Fiber adalah sebagai berikut:

1. Instalasi dependensi: Pertama, kita perlu menginstal paket-paket terkait Swagger dan Golang Fiber. Dalam kasus ini, kita dapat menggunakan package `swaggo` untuk Swagger menggunakan Golang Fiber.
2. Anotasi kode sumber: Pada titik-titik yang sesuai dalam kode sumber aplikasi Golang Fiber, kita dapat menambahkan anotasi atau tag Swagger untuk mendefinisikan rute API, parameter, respons, dan informasi lainnya. Ini biasanya dilakukan menggunakan komentar pada fungsi handler atau struktur data yang terlibat.
3. Menghasilkan dokumentasi: Setelah anotasi telah ditambahkan pada kode sumber (Golang), kita dapat menggunakan tools `go-swagger` untuk menghasilkan file Swagger JSON atau YAML yang berisi dokumentasi lengkap tentang API. Tools ini nantinya akan membaca anotasi dan menghasilkan file Swagger berdasarkan informasi tersebut.
4. Mengintegrasikan dengan Golang Fiber: Selanjutnya, kita perlu mengintegrasikan file Swagger yang dihasilkan dengan aplikasi Golang Fiber. Dalam hal ini, kita dapat menyediakan endpoint khusus pada aplikasi yang akan digunakan untuk menyajikan file Swagger kepada pengguna.
5. Uji API dan dokumentasi: Setelah aplikasi berjalan, kita dapat mengakses dokumentasi API yang dihasilkan melalui endpoint yang telah ditentukan. Dari sana, kita dapat menjelajahi rute API, mencoba permintaan, dan melihat respons yang diharapkan.

Dengan menggunakan Swagger dan Golang Fiber, kita dapat meningkatkan efisiensi dalam mengembangkan dan memelihara API. Dokumentasi yang dihasilkan secara otomatis memudahkan pengguna atau pengembang lain untuk memahami dan menggunakan API yang kita buat. Selain itu, Swagger juga memungkinkan kita untuk menguji API dengan mudah dan memberikan kesempatan untuk berinteraksi dengan endpoint-endpoint yang ada secara langsung dari dalam dokumentasi.

## Membuat Dokumentasi
1. Yang pertama dilakukan adalah silahkan install terlebih dahulu swagger pada project Boilerplate masing-masing. Buka project Boilerplate, kemudian klik [disini](https://github.com/gofiber/swagger) untuk instalasi fiber pada terminal. Dan lakukan sampai tahap 4 (namun pada tahap `Canonical example` tidak perlu membuat file main.go , cukup menambahkan anotasi @title hingga @BasePath pada bagian atas function main)

<p align="center">
  <img src="https://github.com/indrariksa/WS/assets/26703717/57ec7e90-a191-4a0f-be1e-516f63f3028a">
</p>

  > Anotasi Swagger di atas adalah metadata yang digunakan untuk mengonfigurasi dan mendokumentasikan API menggunakan Swagger pada framework Golang Fiber. Berikut adalah penjelasan dari setiap anotasi tersebut:
  > * @title: Menentukan judul atau nama dari API yang akan didokumentasikan. Pada contoh di atas, judul API adalah "TES SWAG".
  > * @version: Menentukan versi dari API yang akan didokumentasikan. Pada contoh di atas, versi API adalah "1.0".
  > * @description: Memberikan deskripsi atau penjelasan singkat tentang API yang akan didokumentasikan.
  > * @contact.name: Menyediakan informasi tentang nama kontak untuk dukungan API.
  > * @contact.url: Menyediakan URL yang mengarah ke sumber daya yang memberikan dukungan atau informasi tambahan tentang API. Pada contoh di atas, URL GitHub (https://github.com/indrariksa) digunakan.
  > * @contact.email: Menyediakan alamat email untuk kontak dukungan API.
  > * @host: Menentukan host atau server tempat API dijalankan. Pada contoh di atas, API dijalankan di host "ws-ulbi.herokuapp.com".
  > * @BasePath: Menentukan basis path atau awalan URL dari semua rute API. Pada contoh di atas, basis path adalah "/" yang berarti tidak ada awalan tambahan.
  > * @schemes: Menentukan skema protokol yang didukung oleh API. Pada contoh di atas, API mendukung skema HTTPS dan HTTP.
  > Dengan menggunakan anotasi Swagger ini, kalian dapat mengkonfigurasi dan mendokumentasikan API Anda secara lebih terstruktur dan informatif menggunakan Swagger.

2. Selanjutnya adalah membuat routing swagger pada file url.go
```go
page.Get("/docs/*", swagger.HandlerDefault)
```
Di atas kita akan menampilkan swagger pada endpoint `docs` menggunakan method GET

3. Jika sudah, kita harus melakukan perintah swag init kembali untuk menggenerate `Swagger Specification`
```sh
swag init
```
![image](https://github.com/indrariksa/WS/assets/26703717/f3af61e9-53b1-4af9-9c7a-3404256479e9)

Jika berhasil, kalian akan melihat 3 file yang terupdate di dalam folder docs. File-file ini adalah:
* docs.go => Diperlukan untuk menghasilkan SwaggerUI.
* swagger.json => Spesifikasi Swagger dalam format file json.
* swagger.yaml => Spesifikasi Swagger dalam format file yaml.

4. Sekarang lakukan push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs
![image](https://github.com/indrariksa/WS/assets/26703717/13a9dc7b-0d4c-4c62-afec-cd235b246b15)

5. Dan kita berhasil menampilkan swagger pada project kita. Namun disini kita belum membuat Dokumentasi lengkapnya (seperti method Get,Post,Put,Delete)

### Dokumentasi GET (Get All dan Get By ID)
1. Buka Project Backend, kemudian kita cari struct pada folder model/type.go, kemudian kita copy semua code di atas dan paste pada Project Boilerplate di folder controller dan buat file bernama struct.go (nama package samakan dengan file coba.go yang ada pada file controller)
![image](https://github.com/indrariksa/WS/assets/26703717/b7cea6bc-a7ee-4556-b442-68efc64dd7d3)

2. Pada project Boilerplate tambahkan anotasi berikut pada file coba.go di atas function GetAllPresensi
![image](https://github.com/indrariksa/WS/assets/26703717/3271310b-f15c-4eb0-8d57-7d0a546533cc)

```sh
// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
```
> Anotasi di atas digunakan untuk mendokumentasikan endpoint `GetAllPresensi` yang mengambil semua data presensi.
> - `// @Summary Get All Data Presensi.`: Memberikan ringkasan singkat tentang apa yang dilakukan oleh endpoint ini, yaitu mengambil semua data presensi.
> - `// @Description Mengambil semua data presensi.`: Menyediakan deskripsi lebih rinci tentang fungsi endpoint ini.
> - `// @Tags Presensi`: Menyertakan tag "Presensi" untuk mengelompokkan endpoint ini bersama dengan endpoint-endpoint terkait yang memiliki tag yang sama.
> - `// @Accept json`: Menginformasikan bahwa endpoint ini dapat menerima permintaan dengan tipe konten json.
> - `// @Produce json`: Menginformasikan bahwa endpoint ini akan menghasilkan respons dengan tipe konten json.
> - `// @Success 200 {object} Presensi`: Menyatakan bahwa jika operasi berhasil (respons kode status 200), respons yang dikembalikan akan berbentuk objek dengan struktur yang ditentukan oleh tipe `Presensi`.
> - `// @Router /presensi [get]`: Menentukan rute URL untuk endpoint ini, yaitu `/presensi` dengan metode HTTP GET.

3. Lakukan perintah `swag init` pada terminal. Seharusnya terdapat error seperti di bawah
![image](https://github.com/indrariksa/WS/assets/26703717/2d7b5e8f-b505-4e51-874e-20f93325c0a0)

Pada Datetime buat saja komentar `//` seperti gambar di bawah atau bisa dihapus (karena struct pada Boilerplate hanya digunakan sebagai contoh response pada Swagger)
![image](https://github.com/indrariksa/WS/assets/26703717/a8b1a585-e921-469f-9b58-636eece0c26c)

4. Lakukan perintah `swag init` kembali, seperti gambar di bawah tidak terdapat error
![image](https://github.com/indrariksa/WS/assets/26703717/84fdb500-4824-4b6c-af45-bb0805ab222a)

5. Push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs
![image](https://github.com/indrariksa/WS/assets/26703717/3b40cd47-62d4-4b13-a2f2-fb67df9f5cea)

Bisa dilihat contoh response di bawah didapatkan dari struct yang kita buat di Boilerplate
![image](https://github.com/indrariksa/WS/assets/26703717/919a2ef3-483f-4973-9865-468a81980dc2)

Untuk ujicoba menggunakan SwaggerUI bisa klik pada Tombol `Try it out` kemudian pilih `Execute`, bisa dilihat data berhasil ditampilkan dengan response code 200
![image](https://github.com/indrariksa/WS/assets/26703717/fe81ec2a-32fb-4417-b6bf-c75d39bba602)

6. Sampai disini kita berhasil membuat contoh dokumentasi dengan method GET (Get All)
7. Selanjutnya kita akan membuat untuk mendapatkan data berdasarkan parameter ID
8. Buka file coba.go kemudian tambah anotasi berikut di atas function GetPresensiID
![image](https://github.com/indrariksa/tes_ws/assets/26703717/c9cdb0e5-cb9c-429f-b16e-4f4b7e654373)

```sh
// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
```
> Anotasi Swagger pada kode di atas adalah sebagai berikut:
> - @Summary, @Description, @Tags, @Accept, @Produce sudah di jelaskan sebelumnya
> - @Param: Mendefinisikan parameter yang dibutuhkan oleh operasi, yaitu id yang merupakan bagian dari path dan memiliki tipe data string. Parameter ini wajib diisi (required) dan akan digunakan untuk mengambil data presensi berdasarkan ID.
> - @Success: Menyatakan bahwa jika operasi berhasil (respons kode status 200), respons yang dikembalikan akan berbentuk objek dengan struktur yang ditentukan oleh tipe `Presensi`.
> - @Failure: Mendefinisikan respons gagal dari operasi ini dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request), 404 (Not Found), dan 500 (Internal Server Error).
> - @Router: Menentukan rute URL untuk endpoint ini, yaitu `/presensi/{id}` dengan metode HTTP GET.

9. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs (pada gambar di bawah terlihat sudah ada untuk Get By ID)
![image](https://github.com/indrariksa/tes_ws/assets/26703717/c3886818-664f-4a66-ba73-6337312a4a43)

Untuk ujicoba menggunakan SwaggerUI bisa klik pada Tombol `Try it out` masukkan ID kemudian pilih `Execute`, bisa dilihat data berhasil ditampilkan dengan response code 200
![image](https://github.com/indrariksa/WS/assets/26703717/64936e25-d17c-4490-b84f-ab45c20793c0)

Jika kondisi gagal seperti berikut yang menghasilkan response code 400 dengan memasukkan ID misal `123`
![image](https://github.com/indrariksa/WS/assets/26703717/f78bab3e-ad62-430f-b8ec-544849c889a8)

### Dokumentasi POST
1. Buka file coba.go kemudian tambah anotasi berikut di atas function InsertData
```sh
// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
```
> Anotasi Swagger pada kode di atas adalah sebagai berikut:
- `@Summary, @Description, @Tags, @Accept, @Produce sudah di jelaskan sebelumnya
- `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi, yaitu `request` yang merupakan body payload untuk data presensi. Parameter ini memiliki tipe data `Presensi` dan wajib diisi (required).
- `@Success`: Mendefinisikan respons sukses, yaitu kode status 200 (OK) dan objek yang dihasilkan adalah tipe data `Presensi`.
- `@Failure`: Mendefinisikan respons gagal dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request) dan 500 (Internal Server Error).
- `@Router`: Mendefinisikan endpoint URL, yaitu "/ins" dengan metode POST.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs (pada gambar di bawah terlihat sudah ada untuk method POST)
![image](https://github.com/indrariksa/WS/assets/26703717/b314d44a-6344-486a-b70f-ed153763fe9e)

3. Namun bisa dilihat ketika kita memilih `Try it out`, contoh data RAW nya hanya menampilkan tipe data seperti int(0) dan string, untuk memudahkannya kita dapat mengubah sedikit code pada struct.go (project boilerplate). Dengan menambah `example:"Contoh"`
![image](https://github.com/indrariksa/WS/assets/26703717/24c5298a-34d7-403f-b907-fbb999a4cbe6)
Atau bisa copy paste di bawah
```go
type Karyawan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty" example:"MASUK"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
```

4. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /doc. Pada method POST dan GET contoh response code pun akan berubah mengikuti struct yang ada
![image](https://github.com/indrariksa/WS/assets/26703717/6a63521e-1762-4dc6-a47c-8edaa6b06d70)

5. Klik pada Tombol `Try it out` maka contoh RAW data pun akan mengikuti struct `Presensi`
6. Kemudian kalian coba execute dengan RAW data yang ada, seharusnya kalian mendapat response code 500 seperti di bawah
![image](https://github.com/indrariksa/WS/assets/26703717/3216ea91-2d11-41ea-b8b9-80ffb5fe465c)
7. Coba selesaikan masalah di atas, apa saja yang harus dilakukan, jika sudah bisa boleh angkat tangan

### Dokumentasi PUT
1. Buka file coba.go kemudian tambah anotasi berikut di atas function UpdateData
```sh
// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
```
> Anotasi Swagger pada kode di atas hampir sama seperti POST hanya berbeda pada:
> - `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi. Terdapat dua parameter dalam contoh ini. Pertama adalah `id` yang diambil dari path URL, dan harus diisi (required). Kedua adalah `request` yang merupakan body payload untuk data presensi. Parameter ini memiliki tipe data `Presensi` dan juga wajib diisi (required).
> - `@Router`: Mendefinisikan endpoint URL , yaitu "/upd/{id}" dengan metode PUT.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs (pada gambar di bawah terlihat sudah ada untuk method PUT)
![image](https://github.com/indrariksa/WS/assets/26703717/096b6150-d141-4804-9faf-9bdfe9104363)

3. Kemudian kalian ujicoba perintah PUT ini, jika berhasil akan menampilkan response code 200
![image](https://github.com/indrariksa/WS/assets/26703717/1786aa2d-3331-4742-b1fa-29bc20deef57)

4. Kemudian jika tidak ada perubahan apapun ketika melakukan update data maka akan menampilkan response code 500
![image](https://github.com/indrariksa/WS/assets/26703717/2cc1b56f-bd21-4579-9a76-58b217bcaf4d)

### Dokumentasi DELETE
1. Buka file coba.go kemudian tambah anotasi berikut di atas function DeletePresensiByID
```sh
// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
```
> Anotasi Swagger pada kode di atas adalah sebagai berikut:
> - `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi. Terdapat satu parameter dalam contoh ini, yaitu `id` yang diambil dari path URL, dan harus diisi (required).
> - `@Success`: Mendefinisikan respons sukses, yaitu kode status 200 (OK).
> - `@Failure`: Mendefinisikan respons gagal dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request) dan 500 (Internal Server Error).
> - `@Router`: Mendefinisikan endpoint URL, yaitu "/delete/{id}" dengan metode DELETE.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint /docs (pada gambar di bawah terlihat sudah ada untuk method DELETE)
![image](https://github.com/indrariksa/WS/assets/26703717/b5fa5c42-0015-4343-9115-0c8da5511a27)

3. Kemudian kalian ujicoba perintah DELETE ini, jika berhasil akan menampilkan response code 200
![image](https://github.com/indrariksa/WS/assets/26703717/c19447a1-8808-42a8-a825-71ffda449a30)

4. Kemudian jika format id tidak sesuai maka akan menampilkan response code 400
![image](https://github.com/indrariksa/WS/assets/26703717/7e60236b-69e8-41db-9c84-e6279c91099b)

5. Kemudian jika id tidak terdapat pada database maka akan menampilkan response code 500
![image](https://github.com/indrariksa/WS/assets/26703717/36486951-6d68-45b8-917b-fdc9d4738de8)

---
*UNTUK MELIHAT DOKUMENTASI LENGKAP GOLANG SWAGGER https://github.com/swaggo/swag#getting-started*
