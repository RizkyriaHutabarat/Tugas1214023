# 1. Source Code dan Penjelasan

Sebelumnya telah dibuat untuk backend pada week 4 dengan tema Nilai yang memiliki 8 collection seperti yang terlihat pada bagian [no 4](#4-skrinsut-mongodb) dengan jumlah data sebanyak 20.

Lalu pada week 5 juga sudah dibuat sebuah boilerplate yang sudah menyambung dengan heroku dan bisa menampilkan data json dari backend pada browser.

Setelah itu pada week 6 dibuat frontend dan sudah menampilkan data dari backend pada week 4.

Selanjutnya untuk UTS dibuatkan frontend dengan template yang berbeda dari sebelumnya.

Berikut adalah alur proses menampilkan data dari backend pada mongodb ke frontend.

Pada boilerplate dilakukan get pada terminal untuk memanggil pakage yang sudah masuk ke [pkg.go.dev](https://pkg.go.dev/github.com/Febriand1/Nilai) sebelumnya.

```bash
go get github.com/febriand1/nilai
```

Jika sudah berhasil, import pakage pada `controller/coba.go` dengan menginisialisasikan pakage seperti dibawah ini.

```go
inimodel "github.com/Febriand1/Nilai/Model"
inimodul "github.com/Febriand1/Nilai/Module"
```

Setelah dilakukan import, tambahkan fungsi berikut untuk memanggil fungsi dari repositori backend yang sudah masuk ke [pkg.go.dev](https://pkg.go.dev/github.com/Febriand1/Nilai).

```go
func GetAllNilai(c *fiber.Ctx) error {
	ps := inimodul.GetAllNilai(config.Ulbimongoconn, "nilai")
	return c.JSON(ps)
}
```

Lalu rapikan dependensi dengan mengetikkan perintah berikut pada terminal dan jalankan.

```bash
go mod tidy
go run main.go
```

Akan muncul link pada terminal, dan buka link tersebut lalu tambahkan `/nilai` pada belakang link untuk menampilkan data dari nilai.
![link](https://user-images.githubusercontent.com/110885840/230781058-548e1713-65c0-4c79-b12c-fba75614e283.png)

![link2](https://user-images.githubusercontent.com/110885840/230781061-4428aec6-4a1d-418d-8f82-b37a9b3fd52b.png)

Jika data sudah tampil, langkah selanjutnya lakukan push ke heroku dengan mengetikkan perintah berikut pada terminal.

```bash
git status
git add.
git status
git commit -m "bebas isi apa aja"
git push
git push heroku main
```

Lalu pada terminal akan menampilkan link heroku, klik link tersebut dan tambahkan `/nilai` pada belakang link untuk menampilkan data dari nilai seperti pada bagian [no 2](#2-url-heroku).

![link3](https://user-images.githubusercontent.com/110885840/230784743-be4a28ff-df7f-4a5e-82bb-1484195c28cf.png)

![link4](https://user-images.githubusercontent.com/110885840/230784745-6c26bf11-c8fe-4b43-b054-fb0579165f7a.png)

Setelah itu lanjutkan ke frontend dengan membuatkan repositori baru yang berisi foler `js`, `template`, dan file `README` dan `LICENSE`. Pada folder `js` berisi folder `config`, `controller`, dan `temp` dan berisi file `fetch.js`. Pada folder `config` berisi file `url.js`, `controller` berisi file `get.js`, dan `temp` berisi file `table.js`. Lalu pada folder `template` berisi folder `css` dari tailwind dan file `index.html`. File `index.html` berisi template untuk membuat interface dan menampilkannya. Untuk source code dari templatenya ada [disini](https://github.com/Febriand1/frontend_uts/blob/main/template/index.html).

Berikut isi dari folder `js`.

1. Pada `fetch.js`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```

Dilakukan import dan menyatakan lokasi dari yang akan diimport tersebut seperti diatas.

Jika di jalankan akan terjadi error pada console, karena `var Cors = cors.Config` pada `cors.go` pada boilerplate belun di ganti.

```go
var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```

Dan pada `var origins = []string` di tambahkan github pages kita agar github pages kita biasa diakses untuk frontend.

```go
"https://febriand1.github.io"

```

2. Pada `config/url.js`

```javascript
export let urlAPI = "https://ws-nilai.herokuapp.com/nilai";
```

Disini di panggil url heroku yang sudah menampilkan data tadi seperti pada [no 2](#2-url-heroku).

3. Pada `temp/table.js`

```javascript
export let isiTabel = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <input class="w-4 h-4 bg-white rounded" type="checkbox">
            </div>
            <div class="w-auto p-2">
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NPM#</p>
                <p class="text-xs font-medium text-coolGray-500">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
        <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#MATAKULIAH#</p>
        <p class="text-xs font-medium text-coolGray-500">#DOSEN#</p>
        <p class="text-xs font-medium text-coolGray-500">#NOHPD#</p>
        </div>
    </th>
        <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JAMMASUK#</p>
        <p class="text-xs font-medium text-coolGray-500">#JAMKELUAR#</p>
        <p class="text-xs font-medium text-coolGray-500">#HARI#</p>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#HADIR#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#TUGAS1#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#TUGAS2#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#TUGAS3#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#TUGAS4#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#TUGAS5#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#UTS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#UAS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#GRADE#</th>

    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`;
```

Dibuatkan source code untuk tempat menampung data yang di panggil dari `get.js`, atau juga bisa di katakan semua yang memiliki tanda "#" salah satu contohnya `#NAMA#` akan di ganti dengan data yang ada pada mongodb. Dan data tersebut akan tampil didalam table yang di buat pada `index.html`.

4. Pada `controller/get.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
  results.forEach(isiRow);
  console.log(results);
}
function isiRow(value) {
  let content = isiTabel
    .replace("#NPM#", value.absensi.biodata.npm)
    .replace("#NAMA#", value.absensi.biodata.nama)
    .replace("#NOHP#", value.absensi.biodata ? value.absensi.biodata.phonenumber : "#NOHP#")
    .replace("#MATAKULIAH#", value.kategori.nama_mk)
    .replace("#DOSEN#", value.kategori.pengampu ? value.kategori.pengampu.namadosen : "#DOSEN#")
    .replace("#NOHPD#", value.kategori.pengampu.phonenumberd)
    .replace("#JAMMASUK#", value.kategori.jadwal ? value.kategori.jadwal.jammasuk : "#JAMMASUK#")
    .replace("#JAMKELUAR#", value.kategori.jadwal.jamkeluar)
    .replace("#HARI#", value.kategori.jadwal.hari)
    .replace("#HADIR#", value.absensi.jumlahkehadiran)
    .replace("#TUGAS1#", value.alltugas.tugas1)
    .replace("#TUGAS2#", value.alltugas.tugas2)
    .replace("#TUGAS3#", value.alltugas.tugas3)
    .replace("#TUGAS4#", value.alltugas.tugas4)
    .replace("#TUGAS5#", value.alltugas.tugas5)
    .replace("#UTS#", value.uts)
    .replace("#UAS#", value.uts)
    .replace("#GRADE#", value.grade.namagrade)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}
```

Dilakukan import dan melakukan replace atau mengganti semua yang memiliki "#" salah satu contohnya `#NAMA#` dengan data dari mongodb dengan memanggil data tersebut sesuai dengan dimana letak data tersebut seperti contoh source code diatas yaitu `value.absensi.biodata.nama` akan mengganti `#NAMA#` dengan data yang berada pada `biodata` yang di dalamnya ada `nama`.

Lalu untuk `value.kategori.jadwal ? value.kategori.jadwal.jammasuk : "#JAMMASUK#"` ini adalah sebuah kondisi dimana `#JAMMASUK#` akan di ganti dengan data yang berada pada `kategori` yang di dalamnya ada `jadwal` dan jika di dalamnya lagi ada `jammasuk` maka akan menampilkan data dari `jammasuk`, tetapi jika tidak ditemukan `jammasuk` tidak akan menampilkan data dan akan menghasilkan **undefined**.

Lalu data tersebut akan ditampilkan pada pages github, dapat dilihat pada [no 3](#3-url-frontend), dan untuk skrinsutnya pada [no 6](#6-skrinsut-frontend).

---

# 2. URL Heroku

https://ws-nilai.herokuapp.com/nilai

---

# 3. URL Frontend

https://febriand1.github.io/frontend_uts/template/

---

# 4. Skrinsut MongoDB

- dosen

  ![1](https://user-images.githubusercontent.com/110885840/230623929-dd879ec2-b46b-4950-b42d-56c364d57c68.png)

- grade

  ![2](https://user-images.githubusercontent.com/110885840/230623934-0b5f9efc-6c3f-404c-8690-0093b0cecefd.png)

- mahasiswa

  ![3](https://user-images.githubusercontent.com/110885840/230623940-e2aee8c2-6658-467d-a013-28c81ae69701.png)

- matakuliah

  ![4](https://user-images.githubusercontent.com/110885840/230623942-4eb35ca5-3f84-4f5b-9250-f3daabb6c2de.png)

- nilai

  ![5](https://user-images.githubusercontent.com/110885840/230624329-3c6cebf7-71d2-493e-9e24-282c252b0d03.png)

- presensi

  ![6](https://user-images.githubusercontent.com/110885840/230623944-f7906a74-67fd-4dfc-8421-27cab83945ee.png)

- tugas

  ![7](https://user-images.githubusercontent.com/110885840/230623948-81294ca7-b4a1-4258-bb0c-819d8a60846d.png)

- waktu

  ![8](https://user-images.githubusercontent.com/110885840/230623953-8556d033-7dec-43f7-b0cf-a19ba1479321.png)

---

# 5. Skrinsut Postman

![pm](https://user-images.githubusercontent.com/110885840/230626115-ec544786-17d0-4f14-b663-edfba55cf738.png)

---

# 6. Skrinsut Frontend

![fe 1](https://user-images.githubusercontent.com/110885840/230782512-c5965248-0fe8-49fe-ad16-baae8a290799.png)

![fe 2](https://user-images.githubusercontent.com/110885840/230782509-5c22fd3d-1181-49b6-a9c7-7970569a6573.png)

---
