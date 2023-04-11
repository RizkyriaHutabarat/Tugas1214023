# Penjelasan Langkah-langkah Pengerjaan
Langkah Awal yaitu mencari template yang akan digunakan pada frontend dan masukan pada directory fe-p1 yang telah dibuat pada week 6

menambahkan tabel pada `index.html` 
```html
<section class="bg-coolGray-50 py-4">
            <div class="container px-4 mx-auto">
                <div class="flex flex-wrap items-center justify-between -m-2 mb-4">
                    <div class="w-full md:w-1/2 p-2">
                        <p class="font-semibold text-xl text-coolGray-800">Data Proyek 1</p>
                        <p class="font-medium text-sm text-coolGray-500">45 Data</p>
                    </div>
                    <div class="w-full md:w-1/2 p-2">
                        <div class="relative md:max-w-max md:ml-auto">
                            <svg class="absolute left-3 transform top-1/2 -translate-y-1/2" width="16" height="16"
                                viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M14.0467 11.22L12.6667 9.80667C12.3699 9.5245 11.9955 9.33754 11.5916 9.26983C11.1876 9.20211 10.7727 9.25673 10.4001 9.42667L9.80008 8.82667C10.5071 7.88194 10.83 6.70445 10.7038 5.53122C10.5775 4.358 10.0115 3.27615 9.1197 2.50347C8.22787 1.73078 7.07643 1.32464 5.89718 1.36679C4.71794 1.40894 3.59844 1.89626 2.76405 2.73065C1.92967 3.56503 1.44235 4.68453 1.4002 5.86378C1.35805 7.04302 1.76419 8.19446 2.53687 9.08629C3.30956 9.97812 4.3914 10.5441 5.56463 10.6704C6.73786 10.7966 7.91535 10.4737 8.86007 9.76667L9.45341 10.36C9.26347 10.7331 9.1954 11.1564 9.25879 11.5702C9.32218 11.984 9.51383 12.3675 9.80674 12.6667L11.2201 14.08C11.5951 14.4545 12.1034 14.6649 12.6334 14.6649C13.1634 14.6649 13.6717 14.4545 14.0467 14.08C14.2372 13.8937 14.3886 13.6713 14.4919 13.4257C14.5953 13.1802 14.6485 12.9164 14.6485 12.65C14.6485 12.3836 14.5953 12.1198 14.4919 11.8743C14.3886 11.6287 14.2372 11.4063 14.0467 11.22ZM8.39341 8.39333C7.9269 8.85866 7.33294 9.1753 6.68657 9.30323C6.0402 9.43117 5.37041 9.36466 4.76181 9.11212C4.15321 8.85958 3.63312 8.43234 3.26722 7.88436C2.90132 7.33638 2.70603 6.69224 2.70603 6.03333C2.70603 5.37442 2.90132 4.73029 3.26722 4.18231C3.63312 3.63433 4.15321 3.20708 4.76181 2.95454C5.37041 2.702 6.0402 2.6355 6.68657 2.76343C7.33294 2.89137 7.9269 3.208 8.39341 3.67333C8.70383 3.98297 8.95012 4.35081 9.11816 4.75577C9.2862 5.16074 9.3727 5.59488 9.3727 6.03333C9.3727 6.47178 9.2862 6.90592 9.11816 7.31089C8.95012 7.71586 8.70383 8.08369 8.39341 8.39333ZM13.1067 13.1067C13.0448 13.1692 12.971 13.2187 12.8898 13.2526C12.8086 13.2864 12.7214 13.3039 12.6334 13.3039C12.5454 13.3039 12.4583 13.2864 12.377 13.2526C12.2958 13.2187 12.2221 13.1692 12.1601 13.1067L10.7467 11.6933C10.6843 11.6314 10.6347 11.5576 10.6008 11.4764C10.567 11.3951 10.5495 11.308 10.5495 11.22C10.5495 11.132 10.567 11.0449 10.6008 10.9636C10.6347 10.8824 10.6843 10.8086 10.7467 10.7467C10.8087 10.6842 10.8825 10.6346 10.9637 10.6007C11.0449 10.5669 11.1321 10.5495 11.2201 10.5495C11.3081 10.5495 11.3952 10.5669 11.4765 10.6007C11.5577 10.6346 11.6314 10.6842 11.6934 10.7467L13.1067 12.16C13.1692 12.222 13.2188 12.2957 13.2527 12.3769C13.2865 12.4582 13.3039 12.5453 13.3039 12.6333C13.3039 12.7213 13.2865 12.8085 13.2527 12.8897C13.2188 12.971 13.1692 13.0447 13.1067 13.1067Z"
                                    fill="#BBC3CF"></path>
                            </svg>
                            <input
                                class="w-full md:w-64 pl-8 pr-4 py-2 text-sm text-coolGray-400 font-medium outline-none focus:border-green-500 border border-coolGray-200 rounded-lg shadow-input"
                                type="text" placeholder="Pencarian" />
                        </div>
                    </div>
                </div>
                <div class="mb-6 border border-coolGray-100"></div>
                <div class="overflow-hidden border border-coolGray-100 rounded-md shadow-dashboard">
                    <div class="overflow-x-auto">
                        <table class="w-full" id="iniTabel">
                            <tr
                                class="whitespace-normal h-11 bg-coolGray-50 bg-opacity-80 border-b border-coolGray-100">
                                <th class="px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    <div class="flex items-left">
                                        <p>Nama & NPM</p>
                                    </div>
                                </th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    Jurusan</th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    PRODI</th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    Dosen Pembimbing</th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    Dosen Penguji</th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    Judul</th>
                                <th
                                    class="whitespace-normal px-4 font-semibold text-xs text-coolGray-500 uppercase text-left">
                                    Tanggal Sidang</th>
                                <th
                                    class="whitespace-normal font-semibold text-xs text-coolGray-500 uppercase text-left">
                                </th>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </section>
```
Sesuaikan table dengan format yang diinginkan

jangan lupa menambahkan tag untuk memanggil `main.js`, `table.js`, dan `fetch.js`

```html
    <script src="../../js/main.js"></script>
    <script src="../../js/temp/table.js"></script>
    <script type="module" src="../../js/fetch.js"></script>
```
copy `index.html` untuk membuat tampilan lainnya (`datamhs.html` dan `datadsn.html`)
sesuaikan tampilan sesuai keinginan 

setelah tampilan selesai selanjutnya bisa dicek dengan Go live

pada `url.js` tambahkan url API lainnya yang dibutuhkan

```javascript
export let urlAPI = "https://aufa-ulbi.herokuapp.com/allproyek1";
export let urlAPImhs = "https://aufa-ulbi.herokuapp.com/allmahasiswa";
export let urlAPIdsn = "https://aufa-ulbi.herokuapp.com/alldosen";
```

selanjutnya menambahkan beberapa fungsi untuk mengisi tabel pada `get.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel, isiTabelmhs, isiTabeldsn } from "../temp/table.js";


export function isiTable(results) {
  results.forEach(isiRow);
}
function isiRow(value) {
  let content =
    isiTabel.replace("#NAMA#", value.biodata_mahasiswa.nama)
      .replace("#NPM#", value.biodata_mahasiswa.npm)
      .replace("#JURUSAN#", value.biodata_mahasiswa.jurusan.nama)
      .replace("#PRODI#", value.biodata_mahasiswa.prodi.nama)
      .replace("#DOSENPEMBIMBING#", value.dosen_pembimbing.nama)
      .replace("#DOSENPENGUJI#", value.dosen_penguji.nama)
      .replace("#JUDUL#", value.judul)
      .replace("#TANGGALSIDANG#", value.tanggal_sidang)
      .replace("#WARNA#", getRandomColor())
      .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}

export function isiTabelmahasiswa(results) {
  results.forEach(isiRowmhs);
}
function isiRowmhs(value) {
  let content =
    isiTabelmhs.replace("#NAMA#", value.nama)
      .replace("#NPM#", value.npm)
      .replace("#JURUSAN#", value.jurusan.nama)
      .replace("#PRODI#", value.prodi.nama)
      .replace("#KELAS#", value.kelas)
      .replace("#WARNA#", getRandomColor())
      .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelmhs", content);
}

export function isiTabeldosen(results) {
  results.forEach(isiRowdsn);
}
function isiRowdsn(value) {
  let content =
    isiTabeldsn.replace("#NAMA#", value.nama)
      .replace("#NID#", value.nid)
      .replace("#PRODI#", value.prodi.nama)
      .replace("#WARNA#", getRandomColor())
      .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabeldsn", content);
}
```

selanjutnya melakukan tambahan pada `fetch.js` sebagai berikut

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTable, isiTabelmahasiswa, isiTabeldosen } from "./controller/get.js";
import { urlAPI, urlAPIdsn, urlAPImhs } from "./config/url.js";
get(urlAPI, isiTable);
get(urlAPImhs, isiTabelmahasiswa);
get(urlAPIdsn, isiTabeldosen);
```

selanjutnya mengubah `table.js` sesuai dengan tabel yang telah dibuat pada `index.html`, `datamhs.html`, dan `datadsn.html`

lalu save dan commit 

setelah data pada github terupdate buka link pages 



# URL Heroku
https://aufa-ulbi.herokuapp.com/allproyek1
https://aufa-ulbi.herokuapp.com/allmahasiswa
https://aufa-ulbi.herokuapp.com/alldosen

# URL Github pages frontend
https://aulianafahrian.github.io/fe-p1/template_baru/public/index.html
# Screenshot Collection
## Collection Proyek 1
![image.png]( {https://raw.githubusercontent.com/aulianafahrian/fe-p1/main/screenshot/Screenshot%20dbproyek1.png
} )
## Collection Mahasiswa
![image.png]( {
https://raw.githubusercontent.com/aulianafahrian/fe-p1/main/screenshot/Screenshot%20dbmhs.png} )
## Collection Dosen
![image.png]( {https://github.com/aulianafahrian/fe-p1/blob/main/screenshot/Screenshot%20dbdsn.png} )
# Screenshot Postman
## Get Proyek 1
![image.png]( https://github.com/aulianafahrian/fe-p1/blob/main/screenshot/Screenshot%20postman%20dbproyek1.png )
## Get Mahasiswa
![image.png]( https://github.com/aulianafahrian/fe-p1/blob/main/screenshot/Screenshot%20postman%20dbmhs.png )
## Get Dosen
![image.png]( https://raw.githubusercontent.com/aulianafahrian/fe-p1/main/screenshot/Screenshot%20dbdsn.png )
# Screenshot Frontend
## Proyek 1
![image.png]( {https://raw.githubusercontent.com/aulianafahrian/fe-p1/main/screenshot/Screenshot%20fe%20pyk1.png} )

## Data Mahasiswa
![image.png]( {https://github.com/aulianafahrian/fe-p1/blob/main/screenshot/Screenshot%20fe%20mhs.png} )
## Data Dosen
![image.png]( {https://github.com/aulianafahrian/fe-p1/blob/main/screenshot/Screenshot%20fe%20dsn.png} )