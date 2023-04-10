# Penjelasan Pembuatan Frontend Kuesioner

## Langkah 1. Backend Kuesioner

Sebelum membuat frontend, pertama-tama saya membuat backend kuesioner terlebih dahulu yang didalamnya ada beberapa file seperti type.go, controller.go, dan dev_test.go

1. Pada type.go, saya membuat 6 struct yang berbeda yaitu (responden, kuesioner, lokasi, question, answer, dan survey).
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230852847-d2001229-1002-4cd2-9181-5311011fc520.png)
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230853076-27c5207a-9503-4936-81a4-8d7c16e967ff.png)

2. Pada controller.go, saya membuat fungsi dan menggunakan environment variabel di golang. Dan di dalam controller ini menambahkan koneksi mongodb.
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230854018-d2dd64e8-9268-42b5-b8db-7c05c9d8fa81.png)

3. Pada dev_test.go, disini tempat untuk melakukan testing dari semua fungsi yang dibuat dengan cara menginput dan memanggil data secara tertentu.
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230854438-400add7f-5d31-4244-805a-bc4bd8493896.png)

4. Kemudian, setelah data yang di test masuk ke MongoDB Compass, selanjutnya saya melakukan listing untuk publish package.

```sh
  git tag
  git tag v0.0.7
  git push origin --tags
  go list -m github.com/stefanymelda/be_kuesioner@v0.0.7
```

![Screenshot](https://user-images.githubusercontent.com/94597289/230855653-c02ef17c-d20f-49b4-b5bd-6d472bb5839a.png)

## Langkah 2. Boiler Plate Kuesioner

1. Pada boiler plate yang sudah ada, saya melakukan go get repository be_kuesioner terlebih dahulu

```go
go get github.com/stefanymelda/be_kuesioner
```

2. Kemudian, melakukan import package pada coba.go
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230859081-2923bf09-643a-4ad9-8315-30dd1529bc00.png)

3. Selanjutnya, menambahkan fungsi get all data kuesioner tanpa filter

```go
func GetAllKuesioner(c *fiber.Ctx) error {
	ps := inimodule.GetAllKuesioner(config.Ulbimongoconn, "kuesioner")
	return c.JSON(ps)
}
```

4. Pada file url.go saya menambahkan path

```go
page.Get("/kuesioner", controller.GetAllKuesioner) //menampilkan seluruh data kuesioner
```

Kemudian dirapihkan depedensinya menggunakan

```go
go mod tidy
```

5. Setelah itu, saya melakukan coba run main

```go
go run main.go
```

maka akan muncul link pada terminal, dimana ketika link di klik akan berpindah ke browser lalu saya menambahkan link tersebut dengan endpoint /kuesioner
![Screenshot](https://user-images.githubusercontent.com/94597289/230866585-d40fdb93-f149-4f2b-b090-daf9220e150c.png)

6. Jika semuanya sudah berjalan, saya melakukan commit and push ke github dan heroku.

- pertama saya melakukan git remote untuk mengetahui apakah sudah mengarah ke repository github (stefanyapp) dan heroku ( jaehyun)
  ![Screenshot](https://user-images.githubusercontent.com/94597289/230868281-023bb827-3732-4083-9a55-0a3ae2b2cc66.png)

- kedua saya melakukan git status untuk mengetahui apakah file yang akan di push sudah sesuai. Dimana output tulisannya masih merah

```sh
git status
```

selanjutnya saya git add dan saya git status lagi untuk mengecek outputnya telah berhasil di add yang ditandai dengan berubah warna menjadi hijau.

```sh
git add .
git status
```

- Setelah semuanya hijau, saya lanjut untuk commit and push dengan menambahkan komentar yang bebas

```sh
git commit -m "push kuesioner ya"
```

- Kemudian, saya push dan push heroku main

```sh
git push
git push heroku main
```

7. Setelah di deployment ke heroku, maka akan menampilkan link heroku. Lalu di klik dan data yang ada pada MongodbCompass akan tampil dengan url heroku yang telah saya buat.
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230870627-b4728780-c153-46cc-bc6b-d0e74805d01e.png)

## Langkah 3. Frontend Kuesioner dengan Setting Github Pages

1. Pertama-tama, terlebih dahulu saya membuat repository baru dengan nama fe_kuesioner dan melakukan git bash.

```
git clone https://github.com/stefanymelda/fe_kuesioner.git
```

2. Berikutnya saya memasukkan file-file yang ada pada folder TEMP. Setelah itu saya membuka file index.html dan menambahkan tag untuk memanggil file js.
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230872905-66c7fb52-d5d2-42bc-a020-9b4349a6494b.png)

3. Langkah selanjutnya, saya melakukan setting github pages dengan:

- membuat file fetch.js dan mengambil data dari backend kuesioner

```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { setInner } from "https://bukulapak.github.io/element/process.js";
let urlAPI = "https://jaehyun.herokuapp.com/kuesioner";
get(urlAPI, isiTableKuesioner);
function isiTableKuesioner(results) {
  console.log(results);
  results.forEach(isiRow);
}
function isiRow(value) {
  console.log(value);
}
```

kemudian commit and push, lalu saya refresh dengan fn+shift+f5 dan tampilannya seperti gambar berikut
![Screenshot](https://user-images.githubusercontent.com/94597289/230874802-caea9cba-4a10-47c9-823d-9978f42e9168.png)

4. Pada fe_kuesioner ini saya mengubah dan menambahkan beberapa codingan dalam setiap file.

- File table.js dimana data pada file ini akan terpanggil di index.html pada bagian table yang menggunakan id

```js
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
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JENISKELAMIN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#EMAIL#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#USIA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#STATUS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#LOKASI#</th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`;
```

- file fetch.js

```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTableKuesioner } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTableKuesioner);
```

- file url.js

```js
export let urlAPI = "https://jaehyun.herokuapp.com/kuesioner";
```

- file get.js

```js
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTableKuesioner(results) {
  results.forEach(isiRow);
}
function isiRow(value) {
  let content = isiTabel
    .replace("#NAMA#", value.biodata.nama)
    .replace("#NOHP#", value.biodata.phone_number)
    .replace("#JENISKELAMIN#", value.biodata.jenis_kelamin)
    .replace("#EMAIL#", value.biodata.email)
    .replace("#USIA#", value.biodata.usia)
    .replace("#STATUS#", value.status)
    .replace("#LOKASI#", value.location)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}
```

5. Setelah semuanya sudah selesai, saya merapihkan struktur penempatan setiap folder.
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230878052-c15cd940-e66d-478e-8ab7-bd73e383564a.png)

6. Kemudian, sesudah saya merapihkan dan mempercantik tampilan githubpages nya. Saya lakukan commit and push lagi dan melakukan refresh kembali dengan fn+shift+f5. Maka tampilan akhir dari frontend dengan setting github pages adalah seperti gambar berikut ini :
   ![Screenshot](https://user-images.githubusercontent.com/94597289/230883675-63ebdf2c-c246-4895-8d49-19523c7ace72.png)

## URL HEROKU KUESIONER

https://jaehyun.herokuapp.com/kuesioner

## URL Github Pages Frontend

https://stefanymelda.github.io/fe_kuesioner/template/

## Screenshot Collection dan Isi MongoDB

- Collection Jampengisian
  ![jampengisian](https://user-images.githubusercontent.com/94597289/230881898-81091b8b-6bb3-4bd1-b7c2-4695df158420.png)

- Collection Kuesioner
  ![kuesioner](https://user-images.githubusercontent.com/94597289/230882081-0f42a263-f634-4964-9049-02c06686bf5b.png)

- Collection Lokasi
  ![lokasi](https://user-images.githubusercontent.com/94597289/230882304-f8b81438-6352-4f9a-85e3-8c5584e2f7e5.png)

- Collection Responden
  ![responden](https://user-images.githubusercontent.com/94597289/230882524-5e1451e0-228f-4886-828a-ea566a0b9257.png)

- Collection Survey
  ![survey](https://user-images.githubusercontent.com/94597289/230882707-1f77c488-0cea-49a6-b9a3-137e5a7b2a49.png)

## Screenshot Aplikasi Postman

![postman](https://user-images.githubusercontent.com/94597289/230883160-21bd4a57-3234-4a49-883c-0e6a15a74a2e.png)

## Screenshot Frontend Github Pages

Hasil Akhir dari Aplikasi yang telah dibuat ke frontend dan di setting dengan github pages.
![frontend](https://user-images.githubusercontent.com/94597289/230883330-c50c57d6-f0a9-40c1-bcde-c53b397cb2e5.png)
