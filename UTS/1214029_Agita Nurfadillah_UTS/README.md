Pada week 4, 5, dan 6 kita sudah membuat frontend, backend dan menampilkan frontend & backend. Selanjutnya pada UTS menampilkan hasil pada week 4,5 dan 6 dengan template yang berbeda.

1.Langkah pertama kita membuat repositori baru (pada github ada +new)

2.Selanjutnya git clone pada git bash here (klik kanan > Git Bash Here) lalu masukan link repo yang sudah kita buat dengan code :

```
git clone https://github.com/agitanurfd/uts_rapat.git
```

3.Setelah melakukan git clone pada repo yang sudah dibuat, selanjutnya buka file index.html dengan menambahkan tag untuk memanggil file js/fetch.js dengan code :

```
<script type="module" src="../js/fetch.js"></script>
```

4.Untuk pengambilan data dari backend kita membuat file fetch.js dengan code : (jangan lupa lakukan commit and push)

```
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi, TabelJamRapat, TabelLokasi, TabelRuangan, TabelTamu, TabelUniversitas } from "./controller/get.js";
import { urlAPI, urlAPIJamRapat, urlAPILokasi, urlAPIRuangan, urlAPITamu, urlAPIUniversitas } from "./config/url.js";
get(urlAPI, isiTablePresensi);
get(urlAPIJamRapat, TabelJamRapat);
get(urlAPILokasi, TabelLokasi);
get(urlAPIRuangan, TabelRuangan);
get(urlAPITamu, TabelTamu );
get(urlAPIUniversitas, TabelUniversitas);
```

5.Masukkan data ke dalam tabel di html, dengan code :
(jangan lupa lakukan commit and push)

- index.html

```
<table class="w-full" id="iniTabel">
```

-jamrapat.html

```
<table class="w-full" id="iniTabelJamRapat">
```

-lokasi.html

```
<table class="w-full" id="iniTabelLokasi">
```

- ruangan.html

```
<table class="w-full" id="iniTabelRuangan">
```

- tamu.html

```
<table class="w-full" id="iniTabelTamu">
```

-universitas.html

```
<table class="w-full" id="iniTabelUniversitas">
```

6.Lalu buat file tabel.js di folder temp, dengan code :

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
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#LOKASI#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#TAMU#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#NO_HP_TAMU#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#JURUSAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#DATETIME#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
export let isiTabelJamRapat =
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
                <p class="text-xs font-semibold text-coolGray-800">#DURASI#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JAM_RAPAT#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#HARI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#TANGGAL#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
export let isiTabelLokasi =
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
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#ALAMAT#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
export let isiTabelRuangan =
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
                <p class="text-xs font-semibold text-coolGray-800">#NO_RUANGAN#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
export let isiTabelTamu =
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
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#PHONE_NUMBER#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JABATAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JAM_KERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#HARI_KERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
export let isiTabelUniversitas =
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
                <p class="text-xs font-semibold text-coolGray-800">#JURUSAN#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
   
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
```

7.Membuat folder config dengan membuat file url.js untuk menampilkan data pada template yang sudah di push ke heroku dengan code :

```
export let urlAPI = "https://agita.herokuapp.com/all";
export let urlAPIJamRapat = "https://agita.herokuapp.com/all-jamrapat";
export let urlAPILokasi = "https://agita.herokuapp.com/all-lokasi";
export let urlAPIRuangan = "https://agita.herokuapp.com/all-ruangan";
export let urlAPITamu = "https://agita.herokuapp.com/all-tamu";
export let urlAPIUniversitas = "https://agita.herokuapp.com/all-universitas";
```

8.Jangan lupa mebuat folder controller dengan mebuat file get.js untuk memanggil isi tabel nya dengan code :

```
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel, isiTabelJamRapat, isiTabelLokasi, isiTabelRuangan, isiTabelTamu, isiTabelUniversitas } from "../temp/tabel.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
export function TabelJamRapat(results) {
    results.forEach(isiRowJamRapat);
}
export function TabelLokasi(results) {
    results.forEach(isiRowLokasi);
}
export function TabelRuangan(results) {
    results.forEach(isiRowRuangan);
}
export function TabelTamu(results) {
    results.forEach(isiRowTamu);
}
export function TabelUniversitas(results) {
    results.forEach(isiRowUniversitas);
}

function isiRow(value) {
    let content =
        isiTabel.replace("#LOKASI#", value.lokasi)
            .replace("#NOHP#", value.phone_number)
            .replace("#TAMU#", value.biodata.nama)
            .replace("#NO_HP_TAMU#", value.biodata.phone_number)
            .replace("#JURUSAN#", value.prodi.jurusan)
            .replace("#DATETIME#", value.datetime)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
}
function isiRowJamRapat(value) {
    let content =
        isiTabelJamRapat.replace("#DURASI#", value.durasi)
            .replace("#JAM_RAPAT#", value.jam_rapat)
            .replace("#HARI#", value.hari)
            .replace("#TANGGAL#", value.tanggal)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabelJamRapat", content);
}
function isiRowLokasi(value) {
    let content =
        isiTabelLokasi.replace("#NAMA#", value.nama)
            .replace("#ALAMAT#", value.alamat)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabelLokasi", content);
}
function isiRowRuangan(value) {
    let content =
        isiTabelRuangan.replace("#NO_RUANGAN#", value.no_ruangan)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabelRuangan", content);
}
function isiRowTamu(value) {
    let content =
        isiTabelTamu.replace("#NAMA#", value.nama)
            .replace("#PHONE_NUMBER#", value.phone_number)
            .replace("#JABATAN#", value.jabatan)
            .replace("#JAM_KERJA#", value.jam_kerja)
            .replace("#HARI_KERJA#", value.hari_kerja)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabelTamu", content);
}
function isiRowUniversitas(value) {
    let content =
        isiTabelUniversitas.replace("#JURUSAN#", value.jurusan)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabelUniversitas", content);
}
```

Link heroku :
https://agita.herokuapp.com/all
https://agita.herokuapp.com/all-jamrapat
https://agita.herokuapp.com/all-lokasi
https://agita.herokuapp.com/all-ruangan
https://agita.herokuapp.com/all-tamu
https://agita.herokuapp.com/all-universitas

Link frontend : https://agitanurfd.github.io/uts_rapat/template/

Screenshoot MongoDB
Jam rapat :
![Jam Rapat](https://user-images.githubusercontent.com/104063079/230986197-3680c105-776e-4d86-ab92-f130dd7f54f9.png)

Lokasi :
![Lokasi](https://user-images.githubusercontent.com/104063079/230986302-e336bbe0-e91b-43b5-915e-2e3a8f53f7ae.png)

Ruangan :
![Ruangan](https://user-images.githubusercontent.com/104063079/230986889-d011ffcc-2d0e-4971-ba17-a283a9663d40.png)

Tamu :
[!Tamu](https://user-images.githubusercontent.com/104063079/230987026-b1708777-e350-413b-a902-39249b457b63.png)

UndanganRapat:
![Undangan Rapat](https://user-images.githubusercontent.com/104063079/230987064-5942e7b9-5391-4268-b8bf-dd3e0b66c172.png)

Universitas:
![Universitas](https://user-images.githubusercontent.com/104063079/230987104-10c105f3-e418-4ef2-b7ea-591e64d1008e.png)

Screenshoot Postman:

- UndanganRapat
  ![Undangan Rapat Postman](https://user-images.githubusercontent.com/104063079/230987313-4907358d-1a24-4193-9e9e-82dc32f8a92a.png))

- Jam Rapat
  ![Jam Rapat Postman](https://user-images.githubusercontent.com/104063079/230987350-0b11645a-2887-4f7a-a1b2-d428bb884b5c.png)

- Lokasi
  ![Lokasi Postman](https://user-images.githubusercontent.com/104063079/230987372-9e8f5d21-01ed-4756-b034-043a527267b1.png)

- Ruangan
  ![Ruangan Postman](https://user-images.githubusercontent.com/104063079/230987521-3941a219-8cd0-4d82-b57d-96e99c97a284.png)

- Tamu
  ![Tamu Postman](https://user-images.githubusercontent.com/104063079/230987556-af4be3c3-89ba-4c64-a0e2-c5cd57bac248.png)

- Universitas
  ![Universitas Postman](https://user-images.githubusercontent.com/104063079/230987573-ca485db1-4d1e-4d46-9552-3f4869c45797.png)

Screenshoot FE

- Undangan Rapat
  ![Undangan Rapat FE](https://user-images.githubusercontent.com/104063079/230987874-a991ef8e-00e7-4dca-b0f4-0eca422b8f48.png)

- Jam Rapat
  ![Jam Rapat FE](https://user-images.githubusercontent.com/104063079/230987895-4e7387c2-8779-4c62-8d34-56e978695290.png)

- Tamu
  ![Tamu FE](https://user-images.githubusercontent.com/104063079/230987920-c0fc3bf2-8acb-4c18-bb1e-80590770aa3b.png)

- Ruangan
  ![Ruangan FE](https://user-images.githubusercontent.com/104063079/230988061-62b3b62b-be4f-4030-8f00-d8d73cc4bb1e.png)

- Lokasi
  ![LokasiFE](https://user-images.githubusercontent.com/104063079/230988075-b6468599-7d7b-44e6-ab94-84baee64c88e.png)

- Universitas
  ![Universitas FE](https://user-images.githubusercontent.com/104063079/230988093-32418993-12bb-41a5-a597-f2c5ad79756d.png)
