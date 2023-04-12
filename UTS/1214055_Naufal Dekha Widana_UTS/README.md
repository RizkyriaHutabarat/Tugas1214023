# Penjelasan Langkah-Langkah Pengerjaan

Disini saya akan membuat frontend berdasarkkan semua data yang dibuat pada week 4

Untuk langkah pertama, memeriksa yang ada pada boilerplate heroku dari aplikasi yang sudah di deploy
![heroku](https://user-images.githubusercontent.com/94073120/231399560-772291fc-09e4-46f1-8f91-ea8ddd1708d3.png)

kita cek data yang mana saja yang di tampilkan pada frontend, disini saya akan menampilkan biodata
dari peneliti dengan url dari https://penelitian-tugas.herokuapp.com/semua

jika sudah menentukan data yang mana yang akan dipanggil, maka langkah kedua adalah membuat sebuah repositori
baru yang berisi folder js, template, file LICENSE, dan juga file README seperti berikut :

1. ![repos1](https://user-images.githubusercontent.com/94073120/231402443-9948dbfd-e775-42f5-8104-641fac3fd653.png)

2. ![repos2](https://user-images.githubusercontent.com/94073120/231402516-a57a3bc8-22cf-45a2-bd92-4aa5802a29ca.png)

3. ![repos3](https://user-images.githubusercontent.com/94073120/231402551-f6bd964c-94ea-4b21-83dc-101fc6ef708e.png)

langkah selanjutnya adalah kita siapkan sebuah template yang akan kita jadikan sebuah template

setelah itu kita harus mengubah terlebih dahulu mengubah isi dari file cors.go yang ada pada latihan sebelumnya
seperti berikut

package config
```go
import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://auth.ulbi.ac.id",
	"https://sip.ulbi.ac.id",
	"https://euis.ulbi.ac.id",
	"https://home.ulbi.ac.id",
	"https://alpha.ulbi.ac.id",
	"https://dias.ulbi.ac.id",
	"https://iteung.ulbi.ac.id",
	"https://whatsauth.github.io",
	"https://naufaldekha002.github.io",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```
dimana pada variable origin ditambahkan url dari github pages yang telah kita buat untuk
menampilkan dari halaman untuk menampilkan data nantinya

setelah itu kita buat sebuah file pada folder config yang ada di dalam folder js,
export let urlAPI = "https://penelitian-tugas.herokuapp.com/semua";
yang mana file ini akan di isi dengan link dari heroku yang akan digunakan untuk pemanggilan
data agar bisa di tampilkan nantinya

lalu buat sebuah file pada folder controller dengan format get.js
yang isinya seperti berikut.

```js
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
    console.log(results);
}
function isiRow(value) {
    let content =
        isiTabel.replace("#NAMA#", value.biodata.nama)
            .replace("#PHONE NUMBER#", value.biodata.phone_number)
            .replace("#JABATAN#", value.biodata.jabatan)
            .replace("#JADWAL PENELITIAN#", value.biodata.jadwalpenelitian)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
}
```
yang mana untuk file get.js berfungsi untuk mengisi data yang kita inginkan 
dari boilerplate heroku

lalu pada folter temp, buat lah sebuah dengan format table.js, yang mana isi dari
file tersebut adalah sebuah string yang nanti isinya adalah data-data yang akan disikan 
dari url heroku

```html
export let isiTabel = 
`
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
            </div>
        </div>
    </th> <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#PHONE NUMBER#</p>
        </div>
    </th>

     <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JABATAN#</p>
        </div>
    </th>

     <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JADWAL PENELITIAN#</p>
        </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
```
setiap kolom harus diganti nama sesuai dengan nama struct yang akan diambil, dan isi dari file
diambil dari file index.html yang telah dibuat. Tapi pada akhir sebelum tutup dari head </head> kita
tambahkan dulu script

```js
<script type="module" src="../js/fetch.js"></script>
```
dan isi dari file fetch.js adalah method yang digunakan untuk mengambil data dari backend yang sudah dibuat
sebelumnya

```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```
lalu setelah itu coba lakukan test untuk mengecek apakah data yang kita inginkan telah
ditampilkan pada github page yang telah dibuat

# URL Heroku
https://penelitian-tugas.herokuapp.com/semua

# MongoDB
1. 
![peneliti](https://user-images.githubusercontent.com/94073120/231402920-4dfc8aad-f20c-45f5-a19a-15d3b1401eb3.png)

2. 
![hasil teliti](https://user-images.githubusercontent.com/94073120/231402742-6d1854ef-88d6-4d16-8c8e-9af0a2638cd0.png)

3. 
![hasil](https://user-images.githubusercontent.com/94073120/231402803-c91cd2da-e8c2-405b-ba03-098dbfbb37c0.png)

4. 
![jadwal](https://user-images.githubusercontent.com/94073120/231402870-0b387594-fc0c-4aa8-8138-347e36d2c325.png)

5. 
![lab](https://user-images.githubusercontent.com/94073120/231402896-ee4fb248-50dd-4809-a402-6a17c1e7b241.png)

# Postman
![postman](https://user-images.githubusercontent.com/94073120/231402613-4c9af6a1-f1d8-4e56-ba94-0fad2c680fd9.png)

# Tampilan Frontend
![frontend](https://user-images.githubusercontent.com/94073120/231402977-188f9bc4-c00b-47ae-a39f-dd39f0722494.png)