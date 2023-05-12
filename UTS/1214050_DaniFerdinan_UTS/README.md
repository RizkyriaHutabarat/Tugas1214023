# UJIAN TENGAH SEMESTER PEMROG 3

|                |                                     |
| -------------- | ----------------------------------- |
| NPM            | 1214050                             |
| NAMA           | DANI FERDINAN                       |
| PRODI/KELAS    | D4TI/2B                             |
| MATAKULIAH     | PEMOGRAMAN III                      |
| DOSEN PENGAMPU | INDRA RIKSA HERLAMBANG M.KOM., SFPC |
| HARI & TANGGAL | RABU, 12 APRIL 2023                 |
|                |                                     |

# PENJELASAN ALUR PROSES

## SOURCE CODE DAN PENJELASAN TAHAPAN
Pertama kita cek data yang telah dibuat sebelumnya. Berikut merupakan URL web service yang telah dibuat sebelumnya.

- https://ws-dani.herokuapp.com/dhs-all

- https://ws-dani.herokuapp.com/mahasiswa-all

- https://ws-dani.herokuapp.com/dosen-all

- https://ws-dani.herokuapp.com/matkul-all

kita cek cukup dengan browser

- DHS 
- Mata Kuliah
- Mahasiswa
- Dosen

Dapat dilihat data yang akan ditampilkan di frontend terdapat semua.

Selanjutnya buat template sesuai pertemuan 6, yang berisi folder `js` dan `template`.

![image](https://user-images.githubusercontent.com/55969069/231115709-5da8b3c4-cdd1-4927-8dd3-eba96a29f1da.png)

Jangan Lupa pada boiler plate atau projek web service nya diatur terlebih dahulu. pada file `config/cors.go` tambahkan domain pages kita di variabel `origins` 
```go
"https://daniferdinandall.github.io"
```
dan atur variabel 'Cors' menjadi seperti ini
```go
var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```

Kembali lagi pada Projek Frontend

pada folder template isikan template html untuk file `index.html`, `mahasiswa.html`, `dosen.html` dan `matkul.html`.

![image](https://user-images.githubusercontent.com/55969069/231115470-870c2699-e645-4530-ac46-e62423f358e6.png)

Lalu pada setiap template isikan Tabel seperti berikut.

- Isi `index.html`

```html
<table id="iniTabel" class="min-w-full text-left text-sm font-medium">
    <thead class="border-b font-medium dark:border-neutral-500">
        <tr>
            <th scope="col" class="px-6 py-4">Nama & NPM</th>
            <th scope="col" class="px-6 py-4">PROGRAM STUDI & FAKULTAS</th>
            <th scope="col" class="px-6 py-4">Matkul & Nilai 1</th>
            <th scope="col" class="px-6 py-4">Matkul & Nilai 2</th>
            <th scope="col" class="px-6 py-4">Matkul & Nilai 3</th>
            <th scope="col" class="px-6 py-4">Matkul & Nilai 4</th>
        </tr>
    </thead>
    <tbody>
        
    </tbody>
</table>
```

- Isi `mahasiswa.html`

```html
<table id="iniTabel" class="min-w-full text-left text-sm font-medium">
    <thead class="border-b font-medium dark:border-neutral-500">
        <tr>
            <th scope="col" class="px-6 py-4">NPM</th>
            <th scope="col" class="px-6 py-4">Nama</th>
            <th scope="col" class="px-6 py-4">Fakultas</th>
            <th scope="col" class="px-6 py-4">Program Studi</th>
            <th scope="col" class="px-6 py-4">Dosen Wali</th>
        </tr>
    </thead>
    <tbody>

    </tbody>
</table>
```

- Isi `dosen.html`

```html
<table id="iniTabel" class="min-w-full text-left text-sm font-medium">
    <thead class="border-b font-medium dark:border-neutral-500">
        <tr>
            <th scope="col" class="px-6 py-4">Kode Dosen</th>
            <th scope="col" class="px-6 py-4">Nama</th>
            <th scope="col" class="px-6 py-4">Phone Number</th>
        </tr>
    </thead>
    <tbody>

    </tbody>
</table>
```

- Isi `matkul.html`

```html
<table id="iniTabel" class="min-w-full text-left text-sm font-medium">
    <thead class="border-b font-medium dark:border-neutral-500">
        <tr>
            <th scope="col" class="px-6 py-4">Kode Matkul</th>
            <th scope="col" class="px-6 py-4">Nama Matkul</th>
            <th scope="col" class="px-6 py-4">Dosen Pengampu</th>
            <th scope="col" class="px-6 py-4">SKS</th>
        </tr>
    </thead>
    <tbody>

    </tbody>
</table>
```

foldet `img` dalam folder `template` digunakan untuk menyimpan gambar yang digunakan pada template.

lalu pada tabel di setiap template html diberi id dengan nama iniTabel.

![image](https://user-images.githubusercontent.com/55969069/231116124-b3029bc5-5823-4ba0-979c-5ce5870dff27.png)

setelah itu didalam folder `template` buat 3 folder yaitu `config`, `controller` dan `temp`.

didalam folder `config` di isi `url.js` yang isi nya seperti ini

```javascript
export let urlAPIdhs = "https://ws-dani.herokuapp.com/dhs-all";
export let urlAPImatkul = "https://ws-dani.herokuapp.com/matkul-all";
export let urlAPIdosen = "https://ws-dani.herokuapp.com/dosen-all";
export let urlAPImahasiswa = "https://ws-dani.herokuapp.com/mahasiswa-all";
```
lalu pada folder `temp` buat folder dengan nama `table.js`.

```javascript
export let isiTabel =
    `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#NAMA#</p>
            <p class="text-xs font-medium text-coolGray-500">#NPM#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#PROGRAM_STUDI#</p>
            <p class="text-xs font-medium text-coolGray-500">#FAKULTAS#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
            </div>
        </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL1#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI1#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL2#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI2#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL3#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI3#</p>
        </div>
    </td>
</tr>
`
export let isiTabelDosen =
    `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4">#KODE#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#PHONE#</td>
</tr>
`
export let isiTabelMatkul =
    `
    <tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">#KODE#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#DOSEN#</td>
    <td class="whitespace-nowrap px-6 py-4">#SKS#</td>
  </tr>
  `

export let isiTabelMahasiswa = `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">#NPM#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#FAKULTAS#</td>
    <td class="whitespace-nowrap px-6 py-4">#PROGRAM_STUDI#</td>
    <td class="whitespace-nowrap px-6 py-4">#DOSEN_WALI#</td>
</tr>
`
```

Lalu untuk folder `controller` tambahkan 4 file yang berguna untuk memasukan data kedalam tabel di tamplate html. file yang dibuat adalah 

- `get.js` untuk data DHS

- `getMahasiswa.js` untuk data Mahsiswa

- `getDosen.js` untuk data Dosen

- `getMatkul.js` untuk data Mata Kuliah

isi dari `get.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
        isiTabel.replace("#NAMA#", value.mahasiswa.nama)
            .replace("#NPM#", value.mahasiswa.npm)
            .replace("#PROGRAM_STUDI#", value.mahasiswa.program_studi?value.mahasiswa.program_studi.nama:"#PROGRAM_STUDI#")
            .replace("#FAKULTAS#", value.mahasiswa.fakultas?value.mahasiswa.fakultas.nama:"#FAKULTAS#")
            .replace("#MATKUL#", value.mata_kuliah?value.mata_kuliah[0].nama:"#MATKUL#")
            .replace("#NILAI#", value.mata_kuliah?value.mata_kuliah[0].nilai:"#NILAI#")
            .replace("#MATKUL1#", value.mata_kuliah?value.mata_kuliah[1].nama:"#MATKUL1#")
            .replace("#NILAI1#", value.mata_kuliah?value.mata_kuliah[1].nilai:"#NILAI1#")
            .replace("#MATKUL2#", value.mata_kuliah?value.mata_kuliah[2].nama:"#MATKUL2#")
            .replace("#NILAI2#", value.mata_kuliah?value.mata_kuliah[2].nilai:"#NILAI2#")
            .replace("#MATKUL3#", value.mata_kuliah?value.mata_kuliah[3].nama:"#MATKUL3#")
            .replace("#NILAI3#", value.mata_kuliah?value.mata_kuliah[3].nilai:"#NILAI3#")
            // .replace("#WARNA#", getRandomColor())
            // .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
}
```

isi dari `getMahasiswa.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelMahasiswa } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
        isiTabelMahasiswa.replace("#NAMA#", value.nama)
            .replace("#NPM#", value.npm)
            .replace("#PROGRAM_STUDI#", value.program_studi?value.program_studi.nama:"#PROGRAM_STUDI#")
            .replace("#FAKULTAS#", value.fakultas?value.fakultas.nama:"#FAKULTAS#")
            .replace("#DOSEN_WALI#", value.dosen_wali?value.dosen_wali.nama:"#DOSEN_WALI#");
    addInner("iniTabel", content);
}
```

isi dari `getDosen.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelDosen } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
    isiTabelDosen.replace("#KODE#", value.kode_dosen)
            .replace("#NAMA#", value.nama)
            .replace("#PHONE#", value.phone_number?value.phone_number:"#PHONE#")
    addInner("iniTabel", content);
}
```

isi dari `getMatkul.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelMatkul } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah ? value.mata_kuliah[1].nama : "1")
    let content =
        isiTabelMatkul.replace("#KODE#", value.kode_matkul)
            .replace("#NAMA#", value.nama)
            .replace("#DOSEN#", value.dosen ? value.dosen.nama : "#DOSEN#")
            .replace("#SKS#", value.sks ? value.sks : "#SKS#")
    addInner("iniTabel", content);
}
```

Lalu pada folder `js` tambahkan 4 file untuk mendapatkan data dari API dan menempatkannya pada tabel dengan fungsi yang telah dibuat pada folder `controller`. file yang dibuat adalah:

- `fetch.js` untuk data DHS

- `fetchMahasiswa.js` untuk data Mahsiswa

- `fetchDosen.js` untuk data Dosen

- `fetchMatkul.js` untuk data Mata Kuliah

isi dari `fetch.js`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPIdhs } from "./config/url.js";
get(urlAPIdhs, isiTablePresensi);
```

isi dari `fetchMahasiswa.js`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/getMahasiswa.js";
import { urlAPImahasiswa } from "./config/url.js";
get(urlAPImahasiswa, isiTablePresensi);

```

isi dari `fetchDosen.js`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/getDosen.js";
import { urlAPIdosen } from "./config/url.js";
get(urlAPIdosen, isiTablePresensi);

```

isi dari `fetchMatkul.js`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/getMatkul.js";
import { urlAPImatkul } from "./config/url.js";
get(urlAPImatkul, isiTablePresensi);

```

Terakhir kita panggil file Fetch pada setiap template Html. file fetch tersebut kita panggil dengan tambahan param `type="module"`.
file tersebut kita panggil pada tag `head`.

- Pada file `index.html`

```javascript
<script type="module" src="../js/fetch.js"></script>
```

- Pada file `mahasiswa.html`

```javascript
<script type="module" src="../js/fetchMahasiswa.js"></script>
```

- Pada file `dosen.html`

```javascript
<script type="module" src="../js/fetchDosen.js"></script>
```

- Pada file `matkul.html`

```javascript
<script type="module" src="../js/fetchMatkul.js"></script>
```

Dan Terakhir push repositori dan setting github pages nya.

## URL HEROKU
https://ws-dani.herokuapp.com/dhs-all

https://ws-dani.herokuapp.com/mahasiswa-all

https://ws-dani.herokuapp.com/dosen-all

https://ws-dani.herokuapp.com/matkul-all

## URL GITHUB PAGES FRONTEND
pages: https://daniferdinandall.github.io/fe_uts/template

repo: https://github.com/daniferdinandall/fe_uts

## SCREENSHOOT SETIAP COLLECTION PADA MONGO
### DHS
![image](https://user-images.githubusercontent.com/55969069/231034237-3aa10da5-49ce-464c-8915-c369aaa419d6.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231034342-00c2eaf7-fb77-408b-ba3a-7bdc50ca3cd0.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231034293-fbed9c88-55dd-4163-a50d-5f0342feec3e.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231034419-406ef1cc-95b0-4b50-becc-af056f13f5f6.png)

## SCREENSHOOT APLIKASI FRONTEND UNTUK SETIAP DATA/COLLECTION
### DHS
![image](https://user-images.githubusercontent.com/55969069/231035230-a6f618cc-a4e5-4e0a-9c8d-f0776ecf3d8b.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231035275-0839db16-65c5-4c44-9152-2a94eb931664.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231035385-85683c12-994a-4341-adca-fb4cda4ad0a0.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231035340-c198898f-7d52-4904-b291-584802fe0ade.png)


## SCREENSHOOT POSTMAN UNTUK SETIAP DATA/COLLECTION
### DHS
![image](https://user-images.githubusercontent.com/55969069/231034658-5a73313d-0b49-4ffb-9469-4360fe081eff.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231034695-77db7087-5281-407e-a4cf-e7f4034d448d.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231034722-3e50f782-87fb-4b10-854a-b6fd75858609.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231034763-f277099e-d3fa-4fa4-ac7e-cea70587d33f.png)
