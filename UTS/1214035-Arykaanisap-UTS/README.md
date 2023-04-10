```
1. PROSES
```

- Week 4 (Pembuatan serta pemilihan studi kasus, disini saya menggunakan tema penggajian dimana di dalam mogo.db terdapat databases yang berjudul  Penggejian_db, dimana  di dalamnya terdapat 7 collection, dan salah satu collection yang berjudul uang didalamnya terdapat 17 jumlah data.)
- Week5 (Membuat boilerplate yang sudah berhasil tersabnung ke heroku dan menampikan json pada browser)
- Week6 (Membuat sebuah frontend dimana data yang ditampilkan adalah data yang terdapat pada week4)
- UTS membuat fronend atau menampilkan taampilan yang berbeda (tidak boleh sama dengaan template tugas kemarin)
- Link sumber template frontend (https://tailwindcomponents.com/components/dashboard)

Template Awal
![Screenshot (365)](https://user-images.githubusercontent.com/97519820/230965052-c227b3f8-81a7-4fb8-abca-18aa085f3683.png)


Pada project Boilerplate yang terdapat pada VScode pada terminal-nya ketikkan perintah:

```
go get github.com/ArykaAnisaP/arykaanisap
```

import package pada controller/coba.go untuk alias boleh bebas, contoh di bawah ini
![Screenshot 2023-04-11 000208](https://user-images.githubusercontent.com/97519820/230965291-998fa829-2140-426a-b333-d6b6bb970cc4.png)


Tambahkan Fungsi get all pada controller/coba.go 
```
func GetAllUang(c *fiber.Ctx) error {
	ps := inimodule.GetAllUang(config.Ulbimongoconn, "uang")
	return c.JSON(ps)
}
```

selanjutnya pada folder url tepatnya file url.go tambahkan path baru 
```
page.Get("/uang", controller.GetAllUang)
```

Selanjutnya adalah:

```
go mod tidy //Untuk merapikan depedensi
go run main.go //untuk mencoba run pada local


(ss link vscode dan uang )

Lakukan pengecekan apakah file yang akan kita push benar dan sesuai dengan menggunakan perintah: 

```
git status
```

jika terdapat banyak tulisan error tidak usah panik lanjutkan langkah selanjutnya yaitu:

```
git add .
git status
```

maka nanti tulisan yang tadinya merah akan berubah jadi hijau semua. Setelah itu lakukan commit dan push

```
git commit -m "Boleh Diisi Apapun Alias Bebas"
```

Jika berhasil nanti akan keluar link, nah yang kita lakukan adalah Buka link tautan nomor 2 dari bawah ke browser apapun yang ingin anda pakai. Jika sudah muncul berarti proses berhasil

langkah selanjutnya adalah Membuat boilerplate framework frontend

Buat folder baru dengan nama (fe_uts)
di dalamnya buat agi beberapa folder, agar tidak bingung ikutilah seperti gambar dibawah ini
![Screenshot 2023-04-11 005352](https://user-images.githubusercontent.com/97519820/230966193-16a519d6-4f11-4ec2-8be2-db66dcf7db7a.png)

Kita akan mulai mengisi file di dalam folder yang telah kita buat.
a. file fetch.js
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```

b. file config/url.js
export let urlAPI = "https://aryka.herokuapp.com/uang"; //Link disesuaikan dengan heroku kalian
```

c. file controller/get.js

```
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    let content = isiTabel
      .replace("#NAMA#", value.biodata.nama)
      .replace("#NOHP#", value.biodata.phone_number)
      .replace("#JABATAN#", value.biodata.jabatan)
      .replace("#LOKASI#", value.absensi.location)
      .replace("#TANGGAL#", value.absensi.waktug.tanggal)
      .replace("#GAJIPOKOK#", value.gaji_pokok)
      .replace("#JAMMASUK#", value.bonus.jam_masuk)
      .replace("#JAMKELUAR#", value.bonus.jam_keluar)
      .replace("#WARNA#", getRandomColor())
      .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
  }
  ```
//yang di dalam kurung di sesuaikan dengan data kalian masing-masing.
```

d. file template/table.js

```
export let isiTabel = 
`
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <input class="w-4 h-4 bg-white rounded" type="checkbox">
            </div>
            <div class="w-auto p-2">
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">RYKA</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JABATAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-right">#LOKASI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#TANGGAL#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#GAJIPOKOK#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JAMMASUK#</p>
        <p class="text-xs font-medium text-coolGray-500">#JAMKELUAR#</p>
        </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>`
```

2. URL Heroku
   https://aryka.herokuapp.com/uang
   
3. URL Frontend
   https://arykaanisap.github.io/fe_uts/template/
   
4. SCREENSHOT pada MongoDB
  ![Screenshot (366)](https://user-images.githubusercontent.com/97519820/230967216-40af6f3d-5067-4cd0-baa2-57df82321c5c.png)
  ![Screenshot (373)](https://user-images.githubusercontent.com/97519820/230967189-0b5266bc-0357-4534-8af0-927abe9f2a47.png)
  ![Screenshot (368)](https://user-images.githubusercontent.com/97519820/230967213-47d2442e-6360-44f0-8107-f589318458c0.png)
  ![Screenshot (369)](https://user-images.githubusercontent.com/97519820/230967211-13e0fdea-ea12-4250-bded-73b89ad7c37d.png)
  ![Screenshot (370)](https://user-images.githubusercontent.com/97519820/230967205-1eb86d8d-41cd-4dba-af1a-cf6494e12a79.png)
  ![Screenshot (371)](https://user-images.githubusercontent.com/97519820/230967199-1d4610ac-d43e-48e1-8cf4-34f70042e275.png)
  ![Screenshot (372)](https://user-images.githubusercontent.com/97519820/230967196-4b9f5fc9-1cd7-4b68-acaa-b1115312d2a3.png)






   
