# UJIAN TENGAH SEMESTER PEMROGRAMAN 3
|  |  |
| ------ | ------ |
| NPM | 1214031 |
| NAMA | ERDITO NAUSHA ADAM |
| PRODI/KELAS | D4TI/2B |
| MATAKULIAH | PEMOGRAMAN III |
| DOSEN PENGAMPU | INDRA RIKSA HERLAMBANG M.KOM., SFPC |
| HARI & TANGGAL | RABU, 12 APRIL 2023 |
|  |  |

## DAFTAR ISI
* [ALUR PROSES PENGERJAAN](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#a--alur-proses-pengerjaan)
* [URL HEROKU](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#b--url-heroku)
* [URL FRONTEND](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#c--url-frontend)
* [URL PACKAGE & REPOSITORY](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#d--url-package--repository)
* [SS MONGODB](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#e--ss-mongodb)
* [SS POSTMAN DAN FRONTEND](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#f--ss-postman-dan-frontend)
* [TECH](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#g-tech)


## A . ALUR PROSES PENGERJAAN
#### 1. PERSIAPAN
* Menyiapkan API

  Langkah pertama adalah menyiapkan API, untuk mendapatkan ini, yang harus dilakukan adalah sama seperti latihan dan tugas sebelumnya, setelah membuat function `GetAll` pada backend seperti berikut [be_pmb](https://github.com/erditona/be_pmb/blob/main/module/controller.go) kemudian diupload ke [pkg.go.dev](https://pkg.go.dev/github.com/erditona/be_pmb), dengan cara ketik command dibawah ini di terminal/ git bash:
  ```bash
  git tag
  git tag v0.0.9 //sesuaikan dengan versi yang terbaru (terus bertambah)
  git push origin --tags
  go list -m github.com/erditona/be_pmb@v0.0.9
  ```
  jika sudah berhasil diupload akan seperti gambar dibawah ini:
  <img width="959" alt="pkg go dev" src="https://user-images.githubusercontent.com/91595733/230923608-364f3b95-57b8-424c-a395-759247cb074f.png">
  
  Langkah kedua adalah buka boilerplate dan ketik command dibawah ini
  ```go
  go get github.com/erditona/be_pmb
  ```
  Setelah itu kita panggil atau import modulnya dan lakukan inisialisasi seperti gambar berikut, masukan atau taruh di file `coba.go` pada folder controller, lebih jelasnya dapat dilihat di [boilerplate/controller](https://github.com/erditona/ws-dito/blob/main/controller/coba.go)
  
  <img width="303" alt="inisialisasi modul" src="https://user-images.githubusercontent.com/91595733/230930092-429cfd7c-2c15-4fca-9b79-527b41bd691f.png">
  
  selanjutnya masih di file `coba.go` kita panggil function `GetAll` yang telah dibuat tadi, dan codenya adalah seperti dibawah ini,  
  ```go
    //GetAllFunction
    func GetAllPendaftaran(c *fiber.Ctx) error {
      ps := module.GetAllPendaftaran(config.Ulbimongoconn,"pendaftaran_maba")
      return c.JSON(ps)
    }

    func GetAllJurusan(c *fiber.Ctx) error {
      ps := module.GetAllJurusan(config.Ulbimongoconn,"daftar_jurusan")
      return c.JSON(ps)
    }

    func GetAllSekolah(c *fiber.Ctx) error {
      ps := module.GetAllSekolah(config.Ulbimongoconn,"daftar_sekolah")
      return c.JSON(ps)
    }

    func GetAllCamaba(c *fiber.Ctx) error {
      ps := module.GetAllCamaba(config.Ulbimongoconn,"daftar_camaba")
      return c.JSON(ps)
    }
  ```
  
  Langkah ketiga adalah membuat link untuk mengaktifkan function tadi, dan codenya adalah seperti dibawah ini, masukan atau taruh di file `url.go` pada folder url, lebih jelasnya dapat dilihat di [boilerplate/url](https://github.com/erditona/ws-dito/blob/main/url/url.go)
  ```go
    page.Get("/pendaftaran", controller.GetAllPendaftaran) 
    page.Get("/jurusan", controller.GetAllJurusan) 
    page.Get("/sekolah", controller.GetAllSekolah) 
    page.Get("/camaba", controller.GetAllCamaba) 
  ```
  
  Langkah keempat test terlebih dahulu di lokal bahwa function tersebut bisa berjalan, sebelum itu rapikan dengan `go mod tidy` dan setelah itu running dengan ketik seperti dibawah pada terminal/gitbash
  ```go
  go run main.go
  ```
  Klik link yang muncul dan coba panggil link yang telah dibuat tadi seperti gambar berikut jika berhasil
  
  <img width="960" alt="Lokal-Pendaftaran" src="https://user-images.githubusercontent.com/91595733/230931785-fb0489fe-817c-4cea-911f-31f9b7ba48ab.png">
  
  Selanjutnya kita tambahkan url https://erditona.github.io pada `var origins` di file [cors.go](https://github.com/erditona/ws-dito/blob/main/config/cors.go). Dan ubah `var Cors` menjadi seperti berikut:
  
  ```go
  var Cors = cors.Config{
    AllowOrigins:     strings.Join(origins[:], ","),
    AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
    AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
    ExposeHeaders:    "Content-Length",
    AllowCredentials: true,
  }
  ```
  
  Jika sudah berjalan lancar semua setiap link, selanjutnya lakukan commit & push pada repo github dan juga heroku. Dan API pun jadi dapat dilihat di [URL HEROKU](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#b--url-heroku)
  
* Menyiapkan Template
  
  Pertama download template di [templateku](https://github.com/erditona/TemplateFE/tree/main/Notus.js), selanjutnya masukan template yang telah didownload tadi pada repository [`FrontEnd`](https://github.com/erditona/fe_pmb), disini saya tidak membuat baru karena dilatihan dan tugas sebelumnya sudah membuat repository ini, jadi tinggal memasukan template tadi.
  
  <img width="214" alt="addTemplate" src="https://user-images.githubusercontent.com/91595733/230934004-4b654363-397a-467b-89c0-a3fe5fd93192.png">
  
  Jangan lupa aktifkan *github pages* di pengaturan.

#### 2. EDIT JS

 Untuk folder JS berisi folder dan file seperti gambar berikut:
 
 <img width="215" alt="js-fileFolder" src="https://user-images.githubusercontent.com/91595733/231220136-37c62250-5de8-4261-96c9-4764cc870b4c.png">
 
 * Edit File `url.js` dengan kode seperti berikut
 ```js
 export let urlAPI = "https://ws-dito.herokuapp.com/pendaftaran";
 export let urlAPIJurusan = "https://ws-dito.herokuapp.com/jurusan";
 export let urlAPISekolah = "https://ws-dito.herokuapp.com/sekolah";
 export let urlAPICamaba = "https://ws-dito.herokuapp.com/camaba";
 ```
 
 * Edit File `get.js` dengan kode seperti berikut
 ```js
 import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTable, isiTable2, isiTable3, isiTable4 } from "../temp/table.js";

export function isiTablePendaftaran(results) {
  results.forEach(isiRow);
}
function isiRow(value) {
  let content = isiTable
    .replace("#KDPENDAFTAR#", value.kdpendaftar)
    .replace("#NAMA#", value.biodata.nama)
    .replace("#NOHP#", value.biodata.phone_number)
    .replace("#SEKOLAH#", value.asalsekolah.nama)
    .replace("#NOHPSEKOLAH#", value.asalsekolah.phone_number)
    .replace("#JURUSAN#", value.jurusan.nama)
    .replace("#JENJANG#", value.jurusan.jenjang)
    .replace("#JALUR#", value.jalur)
    .replace("#ALULBI#", value.alulbi)
    .replace("#ALJURUSAN#", value.aljurusan)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}

export function isiTableJurusan(results) {
  results.forEach(isiRow2);
}
function isiRow2(value) {
  let content = isiTable2
    .replace("#KDJURUSAN#", value.kdjurusan)
    .replace("#NAMA#", value.nama)
    .replace("#JENJANG#", value.jenjang)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("TabelJurusan", content);
}

export function isiTableSekolah(results) {
  results.forEach(isiRow3);
}
function isiRow3(value) {
  let content = isiTable3
    .replace("#KDSEKOLAH#", value.kdsekolah)
    .replace("#NAMA#", value.nama)
    .replace("#NOHP#", value.phone_number)
    .replace("#ALAMAT#", value.alamat)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("TabelSekolah", content);
}

export function isiTableCamaba(results) {
  results.forEach(isiRow4);
}
function isiRow4(value) {
  let content = isiTable4
    .replace("#KTP#", value.ktp)
    .replace("#NAMA#", value.nama)
    .replace("#NOHP#", value.phone_number)
    .replace("#ALAMAT#", value.alamat)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("TabelCamaba", content);
}

 ```
 
 * Edit File `teble.js` dengan kode seperti berikut
 ```js
 //tabelPendaftaran
export let isiTable = `
  <tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
      <div class="flex items-center -m-2">
        <div class="w-auto p-2">
          <input class="w-4 h-4 bg-white rounded" type="checkbox" />
        </div>
        <div class="w-auto p-2">
            <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
        </div>
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#KDPENDAFTAR#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
          <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#SEKOLAH#</p>
          <p class="text-xs font-medium text-coolGray-500">#NOHPSEKOLAH#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#JURUSAN#</p>
          <p class="text-xs font-medium text-coolGray-500">#JENJANG#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#JALUR#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#ALULBI#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#ALJURUSAN#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
  </tr>
`;

//tabelJurusan
export let isiTable2 = `
  <tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
      <div class="flex items-center -m-2">
        <div class="w-auto p-2">
          <input class="w-4 h-4 bg-white rounded" type="checkbox" />
        </div>
        <div class="w-auto p-2">
            <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
        </div>
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#KDJURUSAN#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#JENJANG#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
  </tr>
`;

//tabelSekolah
export let isiTable3 = `
  <tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
      <div class="flex items-center -m-2">
        <div class="w-auto p-2">
          <input class="w-4 h-4 bg-white rounded" type="checkbox" />
        </div>
        <div class="w-auto p-2">
            <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
        </div>
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#KDSEKOLAH#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NOHP#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#ALAMAT#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
  </tr>
`;
//tabelCamaba
export let isiTable4 = `
  <tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
      <div class="flex items-center -m-2">
        <div class="w-auto p-2">
          <input class="w-4 h-4 bg-white rounded" type="checkbox" />
        </div>
        <div class="w-auto p-2">
            <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
        </div>
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#KTP#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#NOHP#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-center">
        <div class="w-auto p-2">
          <p class="text-xs font-semibold text-coolGray-800">#ALAMAT#</p>
        </div>
      </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
  </tr>
`;
 ```
 
* Edit File `fetch.js` dengan kode seperti berikut
 ```js
 import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePendaftaran, isiTableJurusan, isiTableSekolah, isiTableCamaba } from "./controller/get.js";
import { urlAPI, urlAPIJurusan, urlAPISekolah, urlAPICamaba } from "./config/url.js";

get(urlAPI, isiTablePendaftaran);
get(urlAPIJurusan, isiTableJurusan);
get(urlAPISekolah, isiTableSekolah);
get(urlAPICamaba, isiTableCamaba);
 ```

#### 3. EDIT HTML

* Membuat tabel pertama yaitu tabel pendaftaran
 
 Jangan lupa tambahkan kode dibawah pada setiap halaman table, untuk mengambil/ mengaktifkan api yg kita telah buat
 ```js
 <script type="module" src="../../../js/fetch.js"></script>
 ```
 
 Untuk Kode html di tabel pendaftaran yaitu sebagai berikut
 ```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="theme-color" content="#000000" />
    <link rel="shortcut icon" href="../../assets/img/favicon.ico" />
    <link rel="apple-touch-icon" sizes="76x76" href="../../assets/img/apple-icon.png" />
    <link rel="stylesheet" href="../../assets/vendor/@fortawesome/fontawesome-free/css/all.min.css" />
    <link rel="stylesheet" href="../../assets/styles/tailwind.css" />
    <link rel="preconnect" href="https://fonts.gstatic.com" />
    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&display=swap" />
    <link rel="stylesheet" href="../../../template/css/tailwind/tailwind.min.css" />
    <title>Table | Pendaftaran</title>
    <script src="../../../js/main.js"></script>
    <script type="module" src="../../../js/fetch.js"></script>
  </head>
  <body class="text-blueGray-700 antialiased">
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root">
      <nav class="md:left-0 md:block md:fixed md:top-0 md:bottom-0 md:overflow-y-auto md:flex-row md:flex-nowrap md:overflow-hidden shadow-xl bg-white flex flex-wrap items-center justify-between relative md:w-64 z-10 py-4 px-6">
        <div class="md:flex-col md:items-stretch md:min-h-full md:flex-nowrap px-0 flex flex-wrap items-center justify-between w-full mx-auto">
          <button class="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent" type="button" onclick="toggleNavbar('example-collapse-sidebar')">
            <i class="fas fa-bars"></i>
          </button>
          <a class="md:block text-left md:pb-2 text-blueGray-600 mr-0 inline-block whitespace-nowrap text-xl uppercase font-bold p-4 px-0" href="../../index.html"> PMB - ULBI </a>
          <ul class="md:hidden items-center flex flex-wrap list-none">
            <li class="inline-block relative">
              <a class="text-blueGray-500 block py-1 px-3" href="#pablo" onclick="openDropdown(event,'notification-dropdown')"><i class="fas fa-bell"></i></a>
              <div class="hidden bg-white text-base z-50 float-left py-2 list-none text-left rounded shadow-lg min-w-48" id="notification-dropdown">
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Another action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Something else here</a>
                <div class="h-0 my-2 border border-solid border-blueGray-100"></div>
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Seprated link</a>
              </div>
            </li>
            <li class="inline-block relative">
              <a class="text-blueGray-500 block" href="#pablo" onclick="openDropdown(event,'user-responsive-dropdown')"
                ><div class="items-center flex">
                  <span class="w-12 h-12 text-sm text-white bg-blueGray-200 inline-flex items-center justify-center rounded-full"
                    ><img alt="..." class="w-full rounded-full align-middle border-none shadow-lg" src="../../assets/img/team-1-800x800.jpg"
                  /></span></div
              ></a>
              <div class="hidden bg-white text-base z-50 float-left py-2 list-none text-left rounded shadow-lg min-w-48" id="user-responsive-dropdown">
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Another action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Something else here</a>
                <div class="h-0 my-2 border border-solid border-blueGray-100"></div>
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Seprated link</a>
              </div>
            </li>
          </ul>
          <div
            class="md:flex md:flex-col md:items-stretch md:opacity-100 md:relative md:mt-4 md:shadow-none shadow absolute top-0 left-0 right-0 z-40 overflow-y-auto overflow-x-hidden h-auto items-center flex-1 rounded hidden"
            id="example-collapse-sidebar"
          >
            <div class="md:min-w-full md:hidden block pb-4 mb-4 border-b border-solid border-blueGray-200">
              <div class="flex flex-wrap">
                <div class="w-6/12">
                  <a class="md:block text-left md:pb-2 text-blueGray-600 mr-0 inline-block whitespace-nowrap text-sm uppercase font-bold p-4 px-0" href="../../index.html"> PMB - ULBI </a>
                </div>
                <div class="w-6/12 flex justify-end">
                  <button type="button" class="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent" onclick="toggleNavbar('example-collapse-sidebar')">
                    <i class="fas fa-times"></i>
                  </button>
                </div>
              </div>
            </div>
            <form class="mt-6 mb-4 md:hidden">
              <div class="mb-3 pt-0">
                <input
                  type="text"
                  placeholder="Search"
                  class="border-0 px-3 py-2 h-12 border border-solid border-blueGray-500 placeholder-blueGray-300 text-blueGray-600 bg-white rounded text-base leading-snug shadow-none outline-none focus:outline-none w-full font-normal"
                />
              </div>
            </form>
            <!-- Divider -->
            <hr class="my-4 md:min-w-full" />
            <!-- Heading -->
            <h6 class="md:min-w-full text-blueGray-500 text-xs uppercase font-bold block pt-1 pb-4 no-underline">Admin Layout Pages</h6>
            <!-- Navigation -->

            <ul class="md:flex-col md:min-w-full flex flex-col list-none">
              <li class="items-center">
                <a href="./dashboard.html" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-tv mr-2 text-sm text-blueGray-300"></i>
                  Dashboard
                </a>
              </li>

              <li class="items-center">
                <a href="#" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-tools mr-2 text-sm text-blueGray-300"></i>
                  Settings
                </a>
              </li>

              <li class="items-center">
                <a href="./table-pendaftaran.html" class="text-xs uppercase py-3 font-bold block text-pink-500 hover:text-pink-600">
                  <i class="fas fa-table mr-2 text-sm opacity-75"></i>
                  Table Pendaftaran
                </a>
              </li>

              <li class="items-center">
                <a href="./table-camaba.html" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-table mr-2 text-sm text-blueGray-300"></i>
                  Table Camaba
                </a>
              </li>

              <li class="items-center">
                <a href="./table-sekolah.html" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-table mr-2 text-sm text-blueGray-300"></i>
                  Table Sekolah
                </a>
              </li>

              <li class="items-center">
                <a href="./table-jurusan.html" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-table mr-2 text-sm text-blueGray-300"></i>
                  Table Jurusan
                </a>
              </li>

              <li class="items-center">
                <a href="#" class="text-xs uppercase py-3 block text-blueGray-700 hover:text-blueGray-500">
                  <i class="fas fa-map-marked mr-2 text-sm text-blueGray-300"></i>
                  Maps
                </a>
              </li>
            </ul>

            <!-- Divider -->
            <hr class="my-4 md:min-w-full" />
            <!-- Heading -->
            <h6 class="md:min-w-full text-blueGray-500 text-xs uppercase font-bold block pt-1 pb-4 no-underline">Auth Layout Pages</h6>
            <!-- Navigation -->

            <ul class="md:flex-col md:min-w-full flex flex-col list-none md:mb-4">
              <li class="items-center">
                <a href="#" class="text-blueGray-700 hover:text-blueGray-500 text-xs uppercase py-3 block">
                  <i class="fas fa-fingerprint text-blueGray-300 mr-2 text-sm"></i>
                  Login
                </a>
              </li>

              <li class="items-center">
                <a href="#" class="text-blueGray-700 hover:text-blueGray-500 text-xs uppercase py-3 block">
                  <i class="fas fa-clipboard-list text-blueGray-300 mr-2 text-sm"></i>
                  Register
                </a>
              </li>
            </ul>

            <!-- Divider -->
            <hr class="my-4 md:min-w-full" />
            <!-- Heading -->
            <h6 class="md:min-w-full text-blueGray-500 text-xs uppercase font-bold block pt-1 pb-4 no-underline">No Layout Pages</h6>
            <!-- Navigation -->

            <ul class="md:flex-col md:min-w-full flex flex-col list-none md:mb-4">
              <li class="items-center">
                <a href="#" class="text-blueGray-700 hover:text-blueGray-500 text-xs uppercase py-3 block">
                  <i class="fas fa-newspaper text-blueGray-300 mr-2 text-sm"></i>
                  Landing Page
                </a>
              </li>

              <li class="items-center">
                <a href="#" class="text-blueGray-700 hover:text-blueGray-500 text-xs uppercase py-3 block">
                  <i class="fas fa-user-circle text-blueGray-300 mr-2 text-sm"></i>
                  Profile Page
                </a>
              </li>
            </ul>

            <!-- Divider -->
            <hr class="my-4 md:min-w-full" />
            <!-- Heading -->
            <h6 class="md:min-w-full text-blueGray-500 text-xs uppercase font-bold block pt-1 pb-4 no-underline">Documentation</h6>
            <!-- Navigation -->
            <ul class="md:flex-col md:min-w-full flex flex-col list-none md:mb-4">
              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/js/colors/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fas fa-paint-brush mr-2 text-blueGray-300 text-base"></i>
                  Styles
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/js/alerts/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-css3-alt mr-2 text-blueGray-300 text-base"></i>
                  CSS Components
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/angular/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-angular mr-2 text-blueGray-300 text-base"></i>
                  Angular
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/js/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-js-square mr-2 text-blueGray-300 text-base"></i>
                  Javascript
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/nextjs/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-react mr-2 text-blueGray-300 text-base"></i>
                  NextJS
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/react/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-react mr-2 text-blueGray-300 text-base"></i>
                  React
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/svelte/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fas fa-link mr-2 text-blueGray-300 text-base"></i>
                  Svelte
                </a>
              </li>

              <li class="inline-flex">
                <a href="https://www.creative-tim.com/learning-lab/tailwind/vue/overview/notus" target="_blank" class="text-blueGray-700 hover:text-blueGray-500 text-sm block mb-4 no-underline">
                  <i class="fab fa-vuejs mr-2 text-blueGray-300 text-base"></i>
                  VueJS
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
      <div class="relative md:ml-64 bg-blueGray-50">
        <nav class="absolute top-0 left-0 w-full z-10 bg-transparent md:flex-row md:flex-nowrap md:justify-start flex items-center p-4">
          <div class="w-full mx-autp items-center flex justify-between md:flex-nowrap flex-wrap md:px-10 px-4">
            <a class="text-white text-sm uppercase hidden lg:inline-block font-semibold" href="./dashboard.html">Dashboard</a>
            <form class="md:flex hidden flex-row flex-wrap items-center lg:ml-auto mr-3">
              <div class="relative flex w-full flex-wrap items-stretch">
                <span class="z-10 h-full leading-snug font-normal absolute text-center text-blueGray-300 absolute bg-transparent rounded text-base items-center justify-center w-8 pl-3 py-3"><i class="fas fa-search"></i></span>
                <input
                  type="text"
                  placeholder="Search here..."
                  class="border-0 px-3 py-3 placeholder-blueGray-300 text-blueGray-600 relative bg-white bg-white rounded text-sm shadow outline-none focus:outline-none focus:ring w-full pl-10"
                />
              </div>
            </form>
            <ul class="flex-col md:flex-row list-none items-center hidden md:flex">
              <a class="text-blueGray-500 block" href="#pablo" onclick="openDropdown(event,'user-dropdown')">
                <div class="items-center flex">
                  <span class="w-12 h-12 text-sm text-white bg-blueGray-200 inline-flex items-center justify-center rounded-full"
                    ><img alt="..." class="w-full rounded-full align-middle border-none shadow-lg" src="../../assets/img/team-1-800x800.jpg"
                  /></span>
                </div>
              </a>
              <div class="hidden bg-white text-base z-50 float-left py-2 list-none text-left rounded shadow-lg min-w-48" id="user-dropdown">
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Another action</a
                ><a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Something else here</a>
                <div class="h-0 my-2 border border-solid border-blueGray-100"></div>
                <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-blueGray-700">Seprated link</a>
              </div>
            </ul>
          </div>
        </nav>
        <!-- Header -->
        <div class="relative bg-pink-600 md:pt-32 pb-32 pt-12">
          <div class="px-4 md:px-10 mx-auto w-full">
            <div>
              <!-- Card stats -->
              <div class="flex flex-wrap">
                <div class="w-full lg:w-6/12 xl:w-3/12 px-4">
                  <div class="relative flex flex-col min-w-0 break-words bg-white rounded mb-6 xl:mb-0 shadow-lg">
                    <div class="flex-auto p-4">
                      <div class="flex flex-wrap">
                        <div class="relative w-full pr-4 max-w-full flex-grow flex-1">
                          <h5 class="text-blueGray-400 uppercase font-bold text-xs">Pendaftar</h5>
                          <span class="font-semibold text-xl text-blueGray-700"> 22 </span>
                        </div>
                        <div class="relative w-auto pl-4 flex-initial">
                          <div class="text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 shadow-lg rounded-full bg-red-500">
                            <i class="far fa-chart-bar"></i>
                          </div>
                        </div>
                      </div>
                      <!-- <p class="text-sm text-blueGray-400 mt-4">
                        <span class="text-emerald-500 mr-2"> <i class="fas fa-arrow-up"></i> 3.48% </span>
                        <span class="whitespace-nowrap"> Since last month </span>
                      </p> -->
                    </div>
                  </div>
                </div>
                <div class="w-full lg:w-6/12 xl:w-3/12 px-4">
                  <div class="relative flex flex-col min-w-0 break-words bg-white rounded mb-6 xl:mb-0 shadow-lg">
                    <div class="flex-auto p-4">
                      <div class="flex flex-wrap">
                        <div class="relative w-full pr-4 max-w-full flex-grow flex-1">
                          <h5 class="text-blueGray-400 uppercase font-bold text-xs">Camaba</h5>
                          <span class="font-semibold text-xl text-blueGray-700"> 22 </span>
                        </div>
                        <div class="relative w-auto pl-4 flex-initial">
                          <div class="text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 shadow-lg rounded-full bg-pink-500">
                            <i class="fas fa-users"></i>
                          </div>
                        </div>
                      </div>
                      <!-- <p class="text-sm text-blueGray-400 mt-4">
                        <span class="text-orange-500 mr-2"> <i class="fas fa-arrow-down"></i> 1.10% </span>
                        <span class="whitespace-nowrap"> Since yesterday </span>
                      </p> -->
                    </div>
                  </div>
                </div>
                <div class="w-full lg:w-6/12 xl:w-3/12 px-4">
                  <div class="relative flex flex-col min-w-0 break-words bg-white rounded mb-6 xl:mb-0 shadow-lg">
                    <div class="flex-auto p-4">
                      <div class="flex flex-wrap">
                        <div class="relative w-full pr-4 max-w-full flex-grow flex-1">
                          <h5 class="text-blueGray-400 uppercase font-bold text-xs">Sekolah</h5>
                          <span class="font-semibold text-xl text-blueGray-700"> 21 </span>
                        </div>
                        <div class="relative w-auto pl-4 flex-initial">
                          <div class="text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 shadow-lg rounded-full bg-orange-500">
                            <i class="fas fa-chart-pie"></i>
                          </div>
                        </div>
                      </div>
                      <!-- <p class="text-sm text-blueGray-400 mt-4">
                        <span class="text-red-500 mr-2"> <i class="fas fa-arrow-down"></i> 3.48% </span>
                        <span class="whitespace-nowrap"> Since last week </span>
                      </p> -->
                    </div>
                  </div>
                </div>
                <div class="w-full lg:w-6/12 xl:w-3/12 px-4">
                  <div class="relative flex flex-col min-w-0 break-words bg-white rounded mb-6 xl:mb-0 shadow-lg">
                    <div class="flex-auto p-4">
                      <div class="flex flex-wrap">
                        <div class="relative w-full pr-4 max-w-full flex-grow flex-1">
                          <h5 class="text-blueGray-400 uppercase font-bold text-xs">Jurusan</h5>
                          <span class="font-semibold text-xl text-blueGray-700"> 15 </span>
                        </div>
                        <div class="relative w-auto pl-4 flex-initial">
                          <div class="text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 shadow-lg rounded-full bg-orange-500">
                            <i class="fas fa-chart-pie"></i>
                          </div>
                        </div>
                      </div>
                      <!-- <p class="text-sm text-blueGray-400 mt-4">
                        <span class="text-red-500 mr-2"> <i class="fas fa-arrow-down"></i> 3.48% </span>
                        <span class="whitespace-nowrap"> Since last week </span>
                      </p> -->
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="px-4 md:px-10 mx-auto w-full -m-24">
          <div class="flex flex-wrap mt-4">
            <div class="w-full mb-12 px-4">
              <div class="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded bg-pink-900 text-pink-300">
                <div class="rounded-t mb-0 px-4 py-3 border-0">
                  <div class="flex flex-wrap items-center">
                    <div class="relative w-full px-4 max-w-full flex-grow flex-1">
                      <h3 class="font-semibold text-lg text-white">Tabel Pendaftaran</h3>
                    </div>
                  </div>
                </div>
                <div class="block w-full overflow-x-auto">
                  <table class="items-center w-full bg-pink-900 border-collapse" id="iniTabel">
                    <thead>
                      <tr>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-pink-800 text-pink-300 border-pink-700">Kode Pendaftaran</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nama & No HP</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Asal Sekolah</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Pilihan Jurusan</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Jalur</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Alasan Pilih Ulbi</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Alasan Pilih Jurusan</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700"></th>
                      </tr>
                    </thead>
                  </table>
                </div>
              </div>
            </div>
          </div>
          <footer class="block py-4">
            <div class="container mx-auto px-4">
              <hr class="mb-4 border-b-1 border-blueGray-200" />
              <div class="flex flex-wrap items-center md:justify-between justify-center">
                <div class="w-full md:w-4/12 px-4">
                  <div class="text-sm text-blueGray-500 font-semibold py-1 text-center md:text-left">
                    Copyright © <span id="get-current-year"></span>
                    <a href="https://www.creative-tim.com?ref=njs-settings" class="text-blueGray-500 hover:text-blueGray-700 text-sm font-semibold py-1"> Creative Tim </a>
                  </div>
                </div>
                <div class="w-full md:w-8/12 px-4">
                  <ul class="flex flex-wrap list-none md:justify-end justify-center">
                    <li>
                      <a href="https://www.creative-tim.com?ref=njs-settings" class="text-blueGray-600 hover:text-blueGray-800 text-sm font-semibold block py-1 px-3"> Creative Tim </a>
                    </li>
                    <li>
                      <a href="https://www.creative-tim.com/presentation?ref=njs-settings" class="text-blueGray-600 hover:text-blueGray-800 text-sm font-semibold block py-1 px-3"> About Us </a>
                    </li>
                    <li>
                      <a href="http://blog.creative-tim.com?ref=njs-settings" class="text-blueGray-600 hover:text-blueGray-800 text-sm font-semibold block py-1 px-3"> Blog </a>
                    </li>
                    <li>
                      <a href="https://github.com/creativetimofficial/notus-js/blob/main/LICENSE.md?ref=njs-settings" class="text-blueGray-600 hover:text-blueGray-800 text-sm font-semibold block py-1 px-3"> MIT License </a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </footer>
        </div>
      </div>
    </div>
    <script src="https://unpkg.com/@popperjs/core@2/dist/umd/popper.js"></script>
    <script type="text/javascript">
      /* Make dynamic date appear */
      (function () {
        if (document.getElementById("get-current-year")) {
          document.getElementById("get-current-year").innerHTML = new Date().getFullYear();
        }
      })();
      /* Sidebar - Side navigation menu on mobile/responsive mode */
      function toggleNavbar(collapseID) {
        document.getElementById(collapseID).classList.toggle("hidden");
        document.getElementById(collapseID).classList.toggle("bg-white");
        document.getElementById(collapseID).classList.toggle("m-2");
        document.getElementById(collapseID).classList.toggle("py-3");
        document.getElementById(collapseID).classList.toggle("px-6");
      }
      /* Function for dropdowns */
      function openDropdown(event, dropdownID) {
        let element = event.target;
        while (element.nodeName !== "A") {
          element = element.parentNode;
        }
        Popper.createPopper(element, document.getElementById(dropdownID), {
          placement: "bottom-start",
        });
        document.getElementById(dropdownID).classList.toggle("hidden");
        document.getElementById(dropdownID).classList.toggle("block");
      }
    </script>

    <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>
    <!-- <script src="../js/charts-demo.js"></script> -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
  </body>
</html>

 ```
 
 * Membuat Tabel Camaba, kita cukup merubah bagian tablenya saja, menjadi seperti berikut
 ```html
 <table class="items-center w-full bg-transparent border-collapse" id="TabelCamaba">
                    <thead>
                      <tr>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-pink-800 text-pink-300 border-pink-700">No Ktp</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nama</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nomor Telepon</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Alamat</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700"></th>
                      </tr>
                    </thead>
                  </table>
 ```
 
 * Membuat Tabel Sekolah, kita cukup merubah bagian tablenya saja, menjadi seperti berikut
  ```html
 <table class="items-center w-full bg-transparent border-collapse" id="TabelSekolah">
                    <thead>
                      <tr>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-pink-800 text-pink-300 border-pink-700">Kode Sekolah</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nama Sekolah</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nomor Telepon</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Alamat</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700"></th>
                      </tr>
                    </thead>
                  </table>
 ```
 
 * Membuat Tabel Jurusan, kita cukup merubah bagian tablenya saja, menjadi seperti berikut
 ```html
 <table class="items-center w-full bg-transparent border-collapse" id="TabelJurusan">
                    <thead>
                      <tr>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-pink-800 text-pink-300 border-pink-700">Kode Jurusan</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Nama Jurusan</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700">Jenjang jurusan</th>
                        <th class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-center bg-pink-800 text-pink-300 border-pink-700"></th>
                      </tr>
                    </thead>
                  </table>
 ``` 
 
 Jangan lupa merubah judul pada setiap file, sesuaikan dengan isi table.

#### 4. FINAL
* Setelah github pages aktif kita bisa langsung mencobanya dengan memanggil url yang terdapat di tabel [URL FRONTEND](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#c--url-frontend)
* Jika berhasil maka akan muncul seperti gambar yang terdapat pada screenshot tampilan [FRONTEND](https://github.com/erditona/Pemrog3WebService/blob/main/UTS/1214031_ERDITO%20NAUSHA%20ADAM_UTS/README.md#f--ss-postman-dan-frontend)


## B . URL HEROKU
|  |  |
| ------ | ------ |
| PENDAFTARAN | https://ws-dito.herokuapp.com/pendaftaran |
| CAMABA | https://ws-dito.herokuapp.com/camaba |
| SEKOLAH | https://ws-dito.herokuapp.com/sekolah |
| JURUSAN | https://ws-dito.herokuapp.com/jurusan |
|  |  |

## C . URL FRONTEND
|  |  |
| ------ | ------ |
| DASHBOARD | https://erditona.github.io/fe_pmb/pmb/pages/admin/dashboard.html |
| TABEL PENDAFTARAN | https://erditona.github.io/fe_pmb/pmb/pages/admin/table-pendaftaran.html |
| TABEL CAMABA | https://erditona.github.io/fe_pmb/pmb/pages/admin/table-camaba.html |
| TABEL SEKOLAH | https://erditona.github.io/fe_pmb/pmb/pages/admin/table-sekolah.html |
| TABEL JURUSAN | https://erditona.github.io/fe_pmb/pmb/pages/admin/table-jurusan.html |
|  |  |

## D . URL PACKAGE & REPOSITORY
|  |  |
| ------ | ------ |
| PACKAGE | https://pkg.go.dev/github.com/erditona/be_pmb |
| FRONTEND | https://github.com/erditona/fe_pmb |
| BACKEND | https://github.com/erditona/be_pmb |
|  |  |

## E . SS MONGODB
1. PENDAFTARAN
<img width="960" alt="mongo-pendaftaran" src="https://user-images.githubusercontent.com/91595733/230714361-dbe1df22-7c25-4270-b6b5-97f3cd86d959.png">

2. CAMABA
<img width="960" alt="mongo-camaba" src="https://user-images.githubusercontent.com/91595733/230714382-8893f194-799d-4e60-a652-5d6f9be1bb14.png">

3. SEKOLAH
<img width="960" alt="mongo-sekolah" src="https://user-images.githubusercontent.com/91595733/230714388-ab0c97c8-092e-418c-978b-1386040aa01e.png">

4. JURUSAN
<img width="960" alt="mongo-jurusan" src="https://user-images.githubusercontent.com/91595733/230714399-bc79c093-c01e-40e6-8d74-bf84944d1faf.png">

## F . SS POSTMAN DAN FRONTEND
1. PENDAFTARAN
<img width="960" alt="fe-tabel-pendaftaran" src="https://user-images.githubusercontent.com/91595733/230919483-fee4fe5b-0dc9-4c39-8706-3e459d73fb51.png">

<img width="960" alt="postman-pendaftaran" src="https://user-images.githubusercontent.com/91595733/230714421-d2797378-62b6-4fe4-b103-4a939bec5b7b.png">

2. CAMABA
<img width="960" alt="fe-tabel-camaba" src="https://user-images.githubusercontent.com/91595733/230714489-40b6ba88-ce97-4afb-9ccf-d72fac7cc67d.png">

<img width="960" alt="postman-camaba" src="https://user-images.githubusercontent.com/91595733/230714484-824731ad-eb79-4b45-9e3e-c20692652417.png">

3. SEKOLAH
<img width="960" alt="fe-tabel-sekolah" src="https://user-images.githubusercontent.com/91595733/230714505-dcc5f0f8-1376-4a08-aae8-700c08ae1b92.png">

<img width="960" alt="postman-sekolah" src="https://user-images.githubusercontent.com/91595733/230714507-9ec84a1c-5248-429e-876d-2eaa321373b0.png">

4. JURUSAN
<img width="960" alt="fe-tabel-jurusan" src="https://user-images.githubusercontent.com/91595733/230714518-2b60a403-50c5-48e8-93d7-8b9dbb943136.png">

<img width="959" alt="postman-jurusan" src="https://user-images.githubusercontent.com/91595733/230714521-9d04f3ba-4515-4329-9148-1fb2a7686bc5.png">

## G. TECH
![My Skills](https://skills.thijs.gg/icons?i=go,js,html,tailwind,=light) 

![Build Status](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=Postman&logoColor=white)
![Build Status](https://img.shields.io/badge/VSCode-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white)
