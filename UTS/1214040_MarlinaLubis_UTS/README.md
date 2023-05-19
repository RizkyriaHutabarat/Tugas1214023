# 1. Tahapan dan Penjelasan

Sebelum sampai pada tahap ini, sebelumnya sudah dibuat backend tentang Bap yang terdiri dari 5 colection dimana setiap collection berisi 20 data.

Kemudian membuat sebuah boilerplate dan sudah terhubung dengan heroku serta dapat menampilan data backend pada browser.

Setelah data pada backend berhasil ditampilkan, maka selanjutnya kita membuat frontend dengan menampilkan data yang ada di backend.

Sebelumnya kita sudah membuat repo digithub kita dengan format nama fe_bap.

Kemudian pada boilerplate yang sudah dibuat sebelumnya, lakukan perintah go get untuk memanggil package yang sudah dibuat sebelumnya.

```bash
go get github.com/MarlinaLubis/be_bap
```

Selanjutnya import package pada controller/coba.go seperti berikut ini
inimodullatihan "github.com/MarlinaLubis/be_bap/module"

Lalu merapikan repidensi dengan perintah

```bash
go mod tidy
```

Kemudian jalankan dengan perintah
```bash
go run main.go
```

Maka akan di tampilkan data seprti berikut
![Screenshot (1936)](https://user-images.githubusercontent.com/110896535/231118171-9227de48-92f5-43b9-83a3-3e3ca8f6edad.png)

Setelah data sudah tampil, maka lakukan perintah berikut

```bash
git status
git add.
git status
git commit -m "bebas isi apa aja"
git push
git push heroku main
```

Selanjutnya pada repo frontend yang sudah di buat sebelumnya, buat folder js, template, README, dan LICENSE. Lalu pada folder js terdiri dari 3 folder yaitu, temp, controller, dan config.

Pada folder js, tambahkan fetch.js dengan perintah

```bash
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTableBap);
```

Kemudian pada controller/url.js

```bash
export let urlAPI = "https://marlinapp.herokuapp.com/bap";
```

Kemudian pada folder temp/table.js

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
                <p class="text-xs font-semibold text-coolGray-800">#JUDUL#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#NAMAMAHASISWA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#PRODI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#HASILREVISI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#DOSEN#</th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`;
```

Kemudian isi pada folder controler/get.js

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { addInner } from "https://bukulapak.github.io/element/process.js";
import {
  getRandomColor,
  getRandomColorName,
} from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "./table.js";
let urlAPI = "https://marlinapp.herokuapp.com/bap";
get(urlAPI, isiTableBap);
function isiTableBap(results) {
  results.forEach(isiRow);
}
function isiRow(value) {
  let content = isiTabel
    .replace("#JUDUL#", value.judul)
    .replace("#NAMAMAHASISWA#", value.nama_mahasiswa)
    .replace("#PRODI#", value.prodi)
    .replace("#HASILREVISI#", value.hasil_revisi)
    .replace("#DOSEN#", value.dosen)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}
```

Pada setiap bagian yang akan diisi dengan data memiliki penanda pagar (#) dan diakhiri dengan tanda pagar(#) juga adalah variabel yang akan kita replace dengan data dari API.

---

# 2. URL Heroku

https://lubisapp.herokuapp.com/bap

---

# 3. URL Frontend

https://MarlinaLubis.github.io/fe_bap/template/

---

# 4. Screenshoot MongoDB

- bap
  ![Screenshot (1937)](https://user-images.githubusercontent.com/110896535/231116890-263a8cd1-65fe-4976-91f0-f69f3b734897.png)

- dosen
  ![Screenshot (1938)](https://user-images.githubusercontent.com/110896535/231115814-a0e6b350-0109-49e5-96d8-c86cdb1b8022.png)

- jam sidang
  ![Screenshot (1939)](https://user-images.githubusercontent.com/110896535/231116362-352125c8-61f1-4b3e-8467-5fcc67349d1c.png)

- mahasiswa
  ![Screenshot (1940)](https://user-images.githubusercontent.com/110896535/231116616-01a0eebf-412a-4302-80be-85569dec3803.png)

- presensi
  ![Screenshot (1941)](https://user-images.githubusercontent.com/110896535/231116729-5f76263a-d6d8-427e-b028-f637c8d1c272.png)

---

# 5. Screenshot Postman

![postman](https://user-images.githubusercontent.com/110896535/231117354-d9b6ad23-f239-4899-870f-5b1686c62a90.png)

---

# 6. Screeshot Frontend

![fe1](https://user-images.githubusercontent.com/110896535/231117821-9ae1da91-115a-469a-9bf3-6d847f41b252.png)

![fe2](https://user-images.githubusercontent.com/110896535/231117963-7134d913-a455-4d44-9a94-89610b416596.png)

---
