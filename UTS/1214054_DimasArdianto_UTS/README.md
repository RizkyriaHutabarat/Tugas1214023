# Penjelasan Langkah Langkah Pengerjaan

Disini akan dijelaskan langkah langkah yang dilakukan untuk membuat frontend dan menampilkan data Kemahasiswaan yang sudah dibuat sebelumnya pada Week 4.

Pertama tama yaitu kita cek terlebih dahulu data yang ingin diambil menggunakan link heroku boiler plate yang sudah dibuat sebelumnya pada Week 5. Disini akan dicoba untuk memanggil data kemahasiswaan dengan link heroku [https://dimasardnt6-ulbi.herokuapp.com/all-kemahasiswaan](https://dimasardnt6-ulbi.herokuapp.com/all-kemahasiswaan)

![heroku-kemahasiswaan](https://user-images.githubusercontent.com/94734096/230945374-fc451697-20a7-4234-a372-bce38b1e31c7.png)

Data sudah bisa tampil dan nantinya bisa diambil dan ditampilkan pada frontend yang akan dibuat

Sekarang kita akan membuat Frontend dengan template yang berbeda, dan akan menampilkan data yang sudah dibuat

Pertama tama buat repositori yang isinya yaitu folder `js` , `kemahasiswaan`, dan file `README` , `LICENSE`.
Lalu buat Folder `config` , `controller` , `temp` , dan file `fetch.js` pada Folder `js`
Gambar struktur folder nya seperti ini :

![Struktur Folder](https://user-images.githubusercontent.com/94734096/230948620-27abb145-fb50-4340-ad7d-a767efcd0ef1.png)

Jika sudah, jangan lupa untuk setting cors pada file `config/cors.go` terlebih dahulu pada boilerplate sebelumnya

```go
var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```

Lalu pada `var origins = []string` tambahkan link github pages frontend kita

```go
"https://dimasardnt6.github.io"

```

Isi Folder kemahasiswaan dengan template baru yang akan kita pakai. Kemudian pada folder `build/pages` buatlah halaman table untuk menampung data yang nantinya akan ditampilkan. file yang dibuat yaitu : `tabel-kemahasiswaan.html`,`tabel-mahasiswa.html`,`tabel-keuangan.html`,`tabel-nilai.html`

Lalu buatlah file `index.html` untuk untuk menampung halaman tabel tabel diatas. Jangan lupa untuk menambahkan tag di dalam section head untuk memanggil file `js/fetch.js` dan juga `js/main.js`

```html
<script src="../../../js/main.js"></script>
<script type="module" src="../../../js/fetch.js"></script>
```

Edit bagian table pada setiap halaman tabel yang sudah dibuat,tambahkan id `<table class="items-center w-full mb-0 align-top border-gray-200 text-slate-500" id="iniTabel2">` agar terisi dengan template html dari `table.js` yang isinya seperti #NAMA# yang akan mereplace dengan data dari API.

![index](https://user-images.githubusercontent.com/94734096/230958431-d70dec2d-4965-493a-b376-d51116aab34f.png)

Selanjutnya yaitu mengisi file `fetch.js` dengan code :

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import {
  isiTable,
  isiTableKemahasiswaan,
  isiTableKeuangan,
  isiTableMahasiswa,
  isiTableNilai,
} from "./controller/get.js";
import {
  urlAPI,
  urlAPIKemahasiswaan,
  urlAPIKeuangan,
  urlAPIMahasiswa,
  urlAPINilai,
} from "./config/url.js";
get(urlAPI, isiTable);

get(urlAPIKemahasiswaan, isiTableKemahasiswaan);
get(urlAPIKeuangan, isiTableKeuangan);
get(urlAPIMahasiswa, isiTableMahasiswa);
get(urlAPINilai, isiTableNilai);
```

Code diatas digunakan untuk mengambil data dari beberapa API. Data yang diambil dari API tersebut kemudian diproses menggunakan beberapa fungsi yang disediakan di dalam file `controller/get.js`, yaitu `isiTable`, `isiTableKemahasiswaan`, `isiTableKeuangan`, `isiTableMahasiswa`, dan `isiTableNilai`.

Setiap pemanggilan fungsi get akan mengambil data dari API yang berbeda dengan URL yang diambil dari file
`config/url.js`, yaitu `urlAPI`, `urlAPIKemahasiswaan`, `urlAPIKeuangan`, `urlAPIMahasiswa`, dan `urlAPINilai`

Selanjutnya pada folder `config` buatlah file `url.js` yang berisi code :

```javascript
export let urlAPI = "https://dimasardnt6-ulbi.herokuapp.com/all-kemahasiswaan";

export let urlAPIKemahasiswaan =
  "https://dimasardnt6-ulbi.herokuapp.com/all-kemahasiswaan";
export let urlAPIKeuangan =
  "https://dimasardnt6-ulbi.herokuapp.com/all-keuangan";
export let urlAPIMahasiswa =
  "https://dimasardnt6-ulbi.herokuapp.com/all-mahasiswa";
export let urlAPINilai = "https://dimasardnt6-ulbi.herokuapp.com/all-nilai";
```

Code di atas adalah code yang menentukan URL dari beberapa API yang akan digunakan dalam aplikasi. Setiap variabel merupakan sebuah string yang berisi URL dari masing-masing API.
Dengan menentukan URL dari setiap API, aplikasi dapat mengambil data yang diperlukan dan memprosesnya sesuai dengan kebutuhan.

Pada folder `controller` buatlah file `get.js` yang berisi code :

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
import {
  getRandomColor,
  getRandomColorName,
} from "https://bukulapak.github.io/image/process.js";
import {
  isiTabel,
  isiTabel2,
  isiTabel3,
  isiTabel4,
  isiTabel5,
} from "../temp/table.js";

export function isiTableKemahasiswaan(results) {
  results.forEach(isiRow2);
}
function isiRow2(value) {
  let content = isiTabel2
    .replace("#NPM#", value.identitas.npm)
    .replace("#NAMA#", value.identitas.nama)
    .replace("#NOHP#", value.identitas.no_hp)
    .replace("#NAMAPRODI#", value.identitas.prodi.nama)
    .replace("#JURUSAN#", value.identitas.jurusan)
    .replace("#PEMBAYARAN#", value.status_keuangan.total_pembayaran)
    .replace("#MATKUL#", value.nilai_mhs.matakuliah.nama_matkul)
    .replace("#DOSEN#", value.nilai_mhs.matakuliah.dosen)
    .replace("#ANGKA#", value.nilai_mhs.nilai_angka)
    .replace("#HURUF#", value.nilai_mhs.nilai_huruf)
    .replace("#KELAS#", value.identitas.kelas);
  addInner("iniTabel2", content);
}
```

Code di atas digunakan untuk mengisi sebuah tabel HTML dengan data yang berasal dari API
Fungsi `isiRow2` akan menerima value, yang berisi data. Fungsi ini akan mengganti nilai dari beberapa placeholder dalam sebuah template tabel `isiTabel2` dengan nilai yang sesuai dari data yang diterima.
Dalam template tabel `isiTabel2`, ada beberapa placeholder yang diganti dengan nilai yang sesuai dari data, seperti #NPM#, #NAMA#, #NOHP#, #NAMAPRODI#, #JURUSAN#, #PEMBAYARAN#, #MATKUL#, #DOSEN#, #ANGKA#, #HURUF#, dan #KELAS#. Setiap placeholder akan diganti dengan nilai yang sesuai untuk membuat baris tabel yang berisi data.

Selanjutnya pada folder `temp` buatlah file `table.js` yang berisi code :

```javascript
export let isiTabel2 = `
<tr>
    <td class="p-2 align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <div class="flex px-2 py-1">
        <div>
            <img src="../assets/img/user.png" class="inline-flex items-center justify-center mr-4 text-white transition-all duration-200 ease-soft-in-out text-sm h-9 w-9 rounded-xl" alt="user1" />
        </div>
            <div class="flex flex-col justify-center">
                <h6 class="mb-0 leading-normal text-sm">#NAMA#</h6>
                <p class="mb-0 leading-tight text-xs text-slate-400">#NPM#</p>
                <p class="mb-0 leading-tight text-xs text-slate-400">#NOHP#</p>
            </div>
        </div>
    </td>
    <td class="p-2 align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <p class="mb-0 font-semibold leading-tight text-xs">#NAMAPRODI#</p>
        <p class="mb-0 font-semibold leading-tight text-xs">#JURUSAN#</p>
    </td>
    <td class="p-2 text-center align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <p class="mb-0 font-semibold leading-tight text-xs">#KELAS#</p>
    </td>
    <td class="p-2 text-center align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <p class="mb-0 font-semibold leading-tight text-xs">#PEMBAYARAN#</p>
    </td>
    <td class="p-2 text-center align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <p class="mb-0 font-semibold leading-tight text-xs">#MATKUL#</p>
        <p class="mb-0 font-semibold leading-tight text-xs">#DOSEN#</p>
    </td>
    <td class="p-2 text-center align-middle bg-transparent border-b whitespace-nowrap shadow-transparent">
        <p class="mb-0 font-semibold leading-tight text-xs">#ANGKA#</p>
        <p class="mb-0 font-semibold leading-tight text-xs">#HURUF#</p>
    </td>
</tr>
`;
```

Kode di atas merupakan variabel yang berisi string HTML yang akan digunakan sebagai template untuk mengisi data pada tabel aplikasi.
Setiap bagian yang akan diisi dengan data memiliki penanda yang diganti dengan data yang sesuai melalui method `replace()`. Penanda tersebut diawali dengan tanda pagar `(#)` dan diakhiri dengan tanda pagar juga. Misalnya, `#NAMA#` adalah penanda untuk data nama yang nantinya akan diganti dengan data yang sudah dipanggil pada file `controller/get.js`

Jadilah Tabel Kemahasiswaan yang sudah terisi data dari API yang sudah dibuat

![github pages-kemahasiswaan](https://user-images.githubusercontent.com/94734096/230960458-e6d8a636-325b-4641-b2b1-0db7fec969ce.png)

# URL HEROKU

https://dimasardnt6-ulbi.herokuapp.com/all-kemahasiswaan

https://dimasardnt6-ulbi.herokuapp.com/all-mahasiswa

https://dimasardnt6-ulbi.herokuapp.com/all-keuangan

https://dimasardnt6-ulbi.herokuapp.com/all-nilai

# URL GITHUB PAGES

https://dimasardnt6.github.io/fe_kemahasiswaan/kemahasiswaan/build/

# SS MongoDb

- kemahasiswaan

![kemahasiswaan](https://user-images.githubusercontent.com/94734096/230961545-a6220734-21f2-4b7b-9d13-53cd90fa23aa.png)

- data_mahasiswa

![data-mahasiswa](https://user-images.githubusercontent.com/94734096/230961657-18ff0206-7b16-4b94-a22c-8a9ff907ced1.png)

- keuangan_mahasiswa

![keuangan](https://user-images.githubusercontent.com/94734096/230961753-f9b25900-9442-4972-83d0-e2ddad038c7e.png)

- nilai_mahasiswa

![nilai](https://user-images.githubusercontent.com/94734096/230961865-164b299d-b64a-455b-b184-5ebc3d61df3b.png)

# SS Postman

- get kemahasiswaan

![get-kemahasiswaan](https://user-images.githubusercontent.com/94734096/230962594-96e816cb-0547-443f-b81f-7c394d5c5891.png)

- get mahasiswa

![get-mahasiswa](https://user-images.githubusercontent.com/94734096/230962630-dbaaa805-b94b-4028-baa0-7bbd0a69587f.png)

- get keuangan

![get-keuangan](https://user-images.githubusercontent.com/94734096/230962657-ece4fb7d-9657-4e79-8dd4-5216604daf19.png)

- get nilai

![get-nilai](https://user-images.githubusercontent.com/94734096/230962683-b53a4636-33be-4758-a115-c81fc9a2aa40.png)

# SS Frontend

- Tabel Kemahasiswaan

![tabel-kemahasiswaan](https://user-images.githubusercontent.com/94734096/230965808-9fc71f54-26c3-4073-9164-15ee8be74e0f.png)

- Tabel Mahasiswa

![tabel-mahasiswa](https://user-images.githubusercontent.com/94734096/230965819-8c7e9764-094e-4ba9-b8c0-14420a7d2e94.png)

- Tabel Keuangan

![tabel-keuangan](https://user-images.githubusercontent.com/94734096/230965813-b8d31300-14de-4beb-b492-b6e2e9368d96.png)

- Tabel Nilai

![tabel-nilai](https://user-images.githubusercontent.com/94734096/230965823-8925248f-d828-4905-aed2-fde51c1a578a.png)

---
