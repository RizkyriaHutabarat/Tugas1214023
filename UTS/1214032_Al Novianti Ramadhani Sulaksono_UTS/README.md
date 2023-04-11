# Penjelasan Tahapan dan Sourcecodenya

Hal pertama yang perlu dilakukan adalah menyelesaikan file boiler plate, backend dan frontend yang ada pada tugas minggu sebelumnya. Pada UTS kita diminta untuk menampilkan data backend ke template frontend yang baru. Hal-hal yang perlu di perhatikan seperti berikut :

- Pada file controller.go/backend ditambahkan perintah berikut pada setiap collection yang dipanggil, contohnya pada perwalian:

```go
	func GetAllPerwalian(db *mongo.Database, col string) (data []model.Perwalian) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
```

- Pada file dev_test.go/backend ditambahkan perintah berikut, contohnya pada perwalian:

```go
func TestGetAll(t *testing.T) {
	data :=module.GetAllPerwalian(module.MongoConn, "perwalian")
	fmt.Println(data)
}
```

Pada file url.go/boiler plate ditambahkan perintah berikut :

```go
	page.Get("/perwalian", controller.GetAllPerwalian)
	page.Get("/allmahasiswa", controller.GetAllMahasiswa)
	page.Get("/alldosen", controller.GetAllDosen)
	page.Get("/alllocation", controller.GetAllLocation)
	page.Get("/allwaktu", controller.GetAllWaktu)
	page.Get("/allruangan", controller.GetAllRuangan)
}
```

Selanjutnya lakukan push ke heroku, dengan perintah berikut :

```
git push heroku main
```

Maka akan muncul link yang apabila dibuka seperti ini :
![d](https://user-images.githubusercontent.com/125644091/231106328-99ebf326-ed5f-4fdc-90a5-c118c5fdea8c.png)

Lakukan pengecekan apakah semua function berhasil

Selanjutnya membuat new repository dengan judul "uts_perwalian", lalu masukkan folder template yang sudah dicari ke dalamnya
![a](https://user-images.githubusercontent.com/125644091/231030205-adb05f2a-fbb3-4555-b1a2-34b3f9a3ed46.png)

Lakukan git clone terlebih dahulu

```go
git clone https://github.com/alnoviantirs/uts_perwalian.git
```

Menambahkan folder index.html, mahasiswa.html, dosen.html, location.html, ruangan.html, waktu.html, lalu edit bagian tabelnya sesuai dengan isi dari MongoDB
![b](https://user-images.githubusercontent.com/125644091/231030942-ba84c61c-50c1-40b7-8813-fbab7122e433.png)

Menambahkan id pada table sesuai isinya

```javascript
<table class="w-full whitespace-nowrap" id="iniTabel"> Pada index.html
<table class="w-full whitespace-nowrap" id="iniTabelMahasiswa"> Pada mahasiswa.html
<table class="w-full whitespace-nowrap" id="iniTabelDosen"> Pada dosen.html
<table class="w-full whitespace-nowrap" id="iniTabelLocation"> Pada location.html
<table class="w-full whitespace-nowrap" id="iniTabelRuangan"> Pada ruangan.html
<table class="w-full whitespace-nowrap" id="iniTabelWaktu"> Pada waktu.html
```

Selanjutnya masuk ke url.js/config dan masukkan Link Heroku dengan judul yang berbeda-beda

```javascript
export let urlAPI = "https://alnovianti.herokuapp.com/perwalian";
export let urlAPImahasiswa = "https://alnovianti.herokuapp.com/allmahasiswa";
export let urlAPIdosen = "https://alnovianti.herokuapp.com/alldosen";
export let urlAPIlocation = "https://alnovianti.herokuapp.com/alllocation";
export let urlAPIwaktu = "https://alnovianti.herokuapp.com/allwaktu";
export let urlAPIruangan = "https://alnovianti.herokuapp.com/allruangan";
```

Masuk ke table.js/temp buat nama untuk memanggil isi dari MongoDB seperti berikut ini

```javascript
export let isiTabel = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JURUSAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#NAMAWALDOS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#TANGGAL#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#HARI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#JAM#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#LOKASI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
</tr>
`;
export let isiTabelMahasiswa = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#NOHP#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#JURUSAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
</tr>
`;
export let isiTabelDosen = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#JABATAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
</tr>
`;
export let isiTabelLocation = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMALOKASI#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#ALAMAT#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
</tr>
`;
export let isiTabelWaktu = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#JAM#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#HARI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#TANGGAL#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
    </th>
</tr>
`;
export let isiTabelRuangan = `
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="text-center flex items-center -m-2">
            <div class="text-center w-auto p-2">
                <p class="text-center text-xs font-semibold text-coolGray-800">#LOKASIRUANGAN#</p>
            </div>
        </div>
    </th>
</tr>
`;
```

Menambahkan get.js/controller isi sesuai nama yang sudah dibuat pada table.js

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
import {
  getRandomColor,
  getRandomColorName,
} from "https://bukulapak.github.io/image/process.js";
import {
  isiTabel,
  isiTabelMahasiswa,
  isiTabelDosen,
  isiTabelLocation,
  isiTabelWaktu,
  isiTabelRuangan,
} from "../temp/table.js";
export function isiTablePresensi(results) {
  results.forEach(isiRow);
}
export function isiTable_Mahasiswa(results) {
  results.forEach(isiRowMahasiswa);
}
export function isiTabel_Dosen(results) {
  results.forEach(isiRowDosen);
}
export function isiTabel_Location(results) {
  results.forEach(isiRowLocation);
}
export function isiTabel_Waktu(results) {
  results.forEach(isiRowWaktu);
}
export function isiTabel_Ruangan(results) {
  results.forEach(isiRowRuangan);
}
function isiRow(value) {
  let content = isiTabel
    .replace("#NAMA#", value.biodata.nama)
    .replace("#NOHP#", value.biodata.phone_number)
    .replace("#JURUSAN#", value.biodata.jurusan)
    .replace("#NAMAWALDOS#", value.walidosen.nama)
    .replace("#TANGGAL#", value.time.tanggal)
    .replace("#HARI#", value.time.hari)
    .replace("#JAM#", value.time.jam)
    .replace("#LOKASI#", value.lokasi)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}
function isiRowMahasiswa(value) {
  let content = isiTabelMahasiswa
    .replace("#NAMA#", value.nama)
    .replace("#NOHP#", value.phone_number)
    .replace("#JURUSAN#", value.jurusan)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelMahasiswa", content);
}
function isiRowDosen(value) {
  let content = isiTabelDosen
    .replace("#NAMA#", value.nama)
    .replace("#JABATAN#", value.jabatan)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelDosen", content);
}
function isiRowLocation(value) {
  let content = isiTabelLocation
    .replace("#NAMALOKASI#", value.nama_lokasi)
    .replace("#ALAMAT#", value.alamat)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelLocation", content);
}
function isiRowWaktu(value) {
  let content = isiTabelWaktu
    .replace("#JAM#", value.jam)
    .replace("#HARI#", value.hari)
    .replace("#TANGGAL#", value.tanggal)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelWaktu", content);
}
function isiRowRuangan(value) {
  let content = isiTabelRuangan
    .replace("#LOKASIRUANGAN#", value.lokasi_ruangan)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabelRuangan", content);
}
```

Selanjutnya pada fetch.js dimasukkan code berikut sesuai apa yang telah dibuat sebelumnya

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import {
  isiTablePresensi,
  isiTable_Mahasiswa,
  isiTabel_Dosen,
  isiTabel_Location,
  isiTabel_Waktu,
  isiTabel_Ruangan,
} from "./controller/get.js";
import {
  urlAPI,
  urlAPImahasiswa,
  urlAPIdosen,
  urlAPIlocation,
  urlAPIwaktu,
  urlAPIruangan,
} from "./config/url.js";
get(urlAPI, isiTablePresensi);
get(urlAPImahasiswa, isiTable_Mahasiswa);
get(urlAPIdosen, isiTabel_Dosen);
get(urlAPIlocation, isiTabel_Location);
get(urlAPIwaktu, isiTabel_Waktu);
get(urlAPIruangan, isiTabel_Ruangan);
```

Setelah itu dilakukan git push dan mencoba membuka link pada github pages
![c](https://user-images.githubusercontent.com/125644091/231033324-4494d952-9c05-4adc-90cc-9aac4007bacd.png)

Maka akan muncul Tabel - Tabel yang sudah dibuat
![Frontend1](https://user-images.githubusercontent.com/125644091/230999457-b75431c3-f9b6-4cb0-a7d7-9ae419272a0e.png)

# URL HEROKU

- https://alnovianti.herokuapp.com/perwalian
- https://alnovianti.herokuapp.com/allmahasiswa
- https://alnovianti.herokuapp.com/alldosen
- https://alnovianti.herokuapp.com/alllocation
- https://alnovianti.herokuapp.com/allwaktu
- https://alnovianti.herokuapp.com/allruangan

# URL Github Pages Frontend

- https://alnoviantirs.github.io/uts_perwalian/template/

# Screenshot Setiap Collection dan Isi Pada MongoDB

- MongoDB Dosen
  ![Mongo1](https://user-images.githubusercontent.com/125644091/230994878-e48f5102-c3ad-48ee-91e2-acf77b5ec030.png)
- MongoDB Location
  ![Mongo2](https://user-images.githubusercontent.com/125644091/230994898-9d958c1d-7e77-4cb9-9f5e-a06dd02a1e5a.png)
- MongoDB Mahasiswa
  ![Mongo3](https://user-images.githubusercontent.com/125644091/230994911-8d740a32-e300-4fa4-ab0e-da2d8c13f1b3.png)
- MongoDB Perwalian
  ![Mongo4](https://user-images.githubusercontent.com/125644091/230994924-67846412-d6cb-43da-9532-2f9ae7ac34df.png)
- MongoDB Ruangan
  ![Mongo5](https://user-images.githubusercontent.com/125644091/230994934-a524a145-ced4-48ea-ae2c-28f6386a7284.png)
- MongoDB Waktu
  ![Mongo6](https://user-images.githubusercontent.com/125644091/230994952-944a6a7c-7207-45ed-8eca-4873ac47d6c1.png)

# Screenshot Pada Postman

- Postman Perwalian
  ![Postman1](https://user-images.githubusercontent.com/125644091/230995402-a5f9fe8b-09c0-499b-a7cc-9391b3705af9.png)
- Postman Mahasiswa
  ![Postman2](https://user-images.githubusercontent.com/125644091/230995418-220275e4-8cf5-406b-a17c-870fcc3538f7.png)
- Postman Dosen
  ![Postman3](https://user-images.githubusercontent.com/125644091/230995431-08f453da-42bd-4dd2-a6d9-64922cbd5b98.png)
- Postman Location
  ![Postman4](https://user-images.githubusercontent.com/125644091/230995442-ebaee5cf-78c2-4d0a-bc68-93a6088a018e.png)
- Postman Waktu
  ![Postman5](https://user-images.githubusercontent.com/125644091/230995468-1bda9009-2025-46ff-8a15-7ceb267543de.png)
- Postman Ruangan
  ![Postman6](https://user-images.githubusercontent.com/125644091/230995489-35432679-e284-4180-ae14-f5c074b1650e.png)

# Screenshot Hasil Frontend

- Tabel Perwalian
  ![Frontend1](https://user-images.githubusercontent.com/125644091/230999457-b75431c3-f9b6-4cb0-a7d7-9ae419272a0e.png)
- Tabel Mahasiswa
  ![Frontend2](https://user-images.githubusercontent.com/125644091/230999486-1b4fae05-fde2-4562-a693-de13eeb89aad.png)
- Tabel Dosen
  ![Frontend3](https://user-images.githubusercontent.com/125644091/230999497-0a884732-2d5f-4e0c-8ad8-fd5df1ff0ab1.png)
- Tabel Waktu
  ![Frontend4](https://user-images.githubusercontent.com/125644091/230999523-45d785c0-a601-4995-babb-2ec8bac1b997.png)
- Tabel Ruangan
  ![Frontend5](https://user-images.githubusercontent.com/125644091/230999543-928b7b3a-b9ec-416d-a6a0-45b0dcc155f5.png)
- Tabel Location
  ![Frontend6](https://user-images.githubusercontent.com/125644091/230999555-1859c334-1d5f-4781-bac4-197db6e59768.png)
