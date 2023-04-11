# 1. Penjelasan Source Code

Kita sudah membuat backend pada di week 4 dengan tema profile yang dimana profile ini terdapat 5 collection yaitu user, pendidikan, pengalaman ,skill dan profile yang dimana 5 collection ini saling berhubungan satu dengan yang lain. jumlah data yaitu sebnayak 20 data.

Selanjutnya pada week 5 kita membuat boilerplate yang terhubung dengan heroku dan menampilkan data-data kita ke json backend pada halaman browser kita.

Pada week 6 kita diminta untuk membuat frontend untuk menampilkan semua data yang ada pada week 4

Selanjutnya Pada UTS ini kita akan mengganti template dan dibuatkan frontend dari yang sebelumnya

Berikut ini alur untuk membuat template dan frontend

1. Langkah pertama yang kita lakukan adalah dengan go get terlebih dahulu dengan memanggil pakage yang sudah kerjakan pada week sebelumnya di vcode masing-masing.

```bash
go get github.com/rizkyriahutabarat/be_profile
```

jika sudah berhasil, kemudian kita mengimport pakage pada `controller/coba.go` dibawah ini

```bash
inimodel "github.com/rizkyriahutabarat/be_profile/model"
inimodul "github.com/rizkyriahutabarat/be_profile/module"
```

setelah tahap ini selesai, kemudia import fungsi berikut ini dengan memanggil fungsi GetAllPresensi yang sudah masuk ke pakage

```bash
func GetAllProfile(c *fiber.Ctx) error {
	nl := inimodul.GetAllProfile(config.Ulbimongoconn, "profile")
	return c.JSON(nl)
}
```

Kemudian rapikan dengan cara mengetikan perintah

```bash
go mod tidy
go run main.go
```

Kemudian akan ada terminal sebagai berikut
![Screenshot (3561)](https://user-images.githubusercontent.com/98501177/230940386-3873b99a-4208-4786-9c3e-3babb0ae5bac.png)
kemudian kita klik http://127.0.0.1.8080
kemudian tambahkan `/allprofile` maka tampil seperti pada gambar berikut ini
![Screenshot (3562)](https://user-images.githubusercontent.com/98501177/230940409-ea5efa03-797f-4550-a29c-5285390ec7fd.png)
Jika data sudah tampil kan kita lakukan ke tahap selanjutnya yaitu push ke heroku. Git status ini berfungsi untuk mengcek terlebih dahulu yang akan kita push sudah sesuai.

```bash
git status
```

Jika sudah lakukan lagi dengan perintah git add

```bash
git add.
git status
```

jika semua berwarna hijau langkah selanjutnya commit dan push

```bash
git commit -m "BEBAS"
```

Jika tahap ini sudah selesai tahap berikutnya membuat repo baru dengan nama fe_profile. Public, add a README File dan MIT license
kemudian copy code HTTPS
![Screenshot (3563)](https://user-images.githubusercontent.com/98501177/230943469-df48648b-271e-4a23-9815-c806d80304a3.png)

kemudian lakukan git clone

```bash
git clone {link https}
```

Pada repo kita terdapat folder `js`, `template`, `LICENSE` dan `README`. Pada bagian folder `js` terdapat folder `config`,`controller` dan `temp` yang dimana setiap folder memiliki fungsi yang berbeda-beda.

1.  Folder `config/url.js`

```javascript
export let urlAPI = "https://rizkyria.herokuapp.com/allprofile";
```

2. Folder `controller/get,js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
  results.forEach(isiRow);
}
function isiRow(value) {
  let content = isiTabel
    .replace("#USERNAME#", value.nama_user)
    .replace("#TAHUNMULAI#", value.data_pendidikan.tahunmulai)
    // .replace("#EMAIL#", value.biodata.email)
    .replace("#SEKOLAH#", value.data_pendidikan.sekolah)
    .replace("#JABATAN#", value.data_pengalaman.jabatan)
    .replace("#LULUSAN#", value.data_pendidikan.lulusan)
    // .replace("#TAHUNSELESAI#", value.biodata.pengalaman)
    .replace("#SKILL#", value.skills.nama)
    .replace("#LEVEL#", value.skills.level)
    .replace("#WARNA#", getRandomColor())
    .replace(/#WARNALOGO#/g, getRandomColorName());
  addInner("iniTabel", content);
}
```

3. Folder `temp/table.js`

```javascript
export let isiTabel = `
<tr class="h-18 border-b border-coolGray-100">
   <th class="whitespace-nowrap px-4 bg-white text-left">
       <div class="flex items-center -m-2">
           <div class="w-auto p-2">
               <input class="w-4 h-4 bg-white rounded" type="checkbox">
           </div>
           <div class="w-auto p-2">
               <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">PRF</div>
           </div>
           <div class="w-auto p-2">
           <p class="text-xs font-semibold text-coolGray-800">#USERNAME#</p>
           <p class="text-xs font-medium text-coolGray-500">#TAHUNMULAI#</p>
       </div>
       </div>
   </th>
   <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#SEKOLAH#</th>
   <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#JABATAN#</th>
   <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#LULUSAN#</th>
   <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#SKILL#</th>
   <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#LEVEL#</th>
   <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
       <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
           <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
       </svg>
   </th>
</tr>
`;
```

4. `fetch`

```javascript
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```

setelah tahap ini selesai makan hasil browser seperti ini :
![Screenshot (3564)](https://user-images.githubusercontent.com/98501177/231102144-23e95d8f-5756-4ef4-9104-64a39e50f1eb.png)

# 2. Link URL Heroku

https://rizkyria.herokuapp.com/allprofile

# 3. Link URL Frontend

https://rizkyriahutabarat.github.io/fe_profile/template/

# 4. SS MongoDB

1. user
   ![Screenshot (3556)](https://user-images.githubusercontent.com/98501177/231102664-21756abf-93ce-4214-8675-21e4ba95abed.png)
   ![Screenshot (3557)](https://user-images.githubusercontent.com/98501177/231102679-4c02e010-0074-45cd-8112-4440391c68e3.png)
   ![Screenshot (3558)](https://user-images.githubusercontent.com/98501177/231102693-0b472b9e-97e4-41ef-a66a-1128c8f4a21b.png)
   ![Screenshot (3559)](https://user-images.githubusercontent.com/98501177/231102708-e641504a-85e9-49f6-a0f0-eabb62d81bdf.png)
   ![Screenshot (3560)](https://user-images.githubusercontent.com/98501177/231102723-52e9e40b-366d-46e4-93b7-be73eacdbca5.png)
2. pendidikan
   ![Screenshot (3538)](https://user-images.githubusercontent.com/98501177/231103083-9c5678e0-680e-4b7f-95dc-fa338e45a680.png)
   ![Screenshot (3539)](https://user-images.githubusercontent.com/98501177/231103088-ab902dc8-04a2-4c3f-9fdb-cdd5e37407d7.png)
   ![Screenshot (3540)](https://user-images.githubusercontent.com/98501177/231103092-13bf8f50-34e0-4aa5-9df4-aa28dea2dac1.png)
   ![Screenshot (3541)](https://user-images.githubusercontent.com/98501177/231103093-89c0db2c-9a48-4f07-8350-ba27baa9fcef.png)
   ![Screenshot (3542)](https://user-images.githubusercontent.com/98501177/231103098-3683dab3-8654-4c98-85d3-23e29e02077c.png)
   ![Screenshot (3543)](https://user-images.githubusercontent.com/98501177/231103110-d3d8106e-3339-4077-b570-c743a5df3b53.png)
   ![Screenshot (3544)](https://user-images.githubusercontent.com/98501177/231103117-98352840-eacf-4a19-88bb-537004ff5026.png)
3. Pengalaman
   ![Screenshot (3545)](https://user-images.githubusercontent.com/98501177/231103472-6234713d-076a-4f10-8999-69570057402a.png)
   ![Screenshot (3546)](https://user-images.githubusercontent.com/98501177/231103478-f92645a4-b6e9-4622-a6c5-b6bc63d017db.png)
   ![Screenshot (3547)](https://user-images.githubusercontent.com/98501177/231103482-4d6ed0b6-7481-4472-b9df-a8e3b83f13a6.png)
   ![Uploading Screenshot (3548).png…]()
   ![Screenshot (3549)](https://user-images.githubusercontent.com/98501177/231103504-65ef37e6-2ffd-47fd-bfa9-cb28865e6a56.png)
   ![Screenshot (3550)](https://user-images.githubusercontent.com/98501177/231103516-79bce46a-b345-4d3d-8101-f94406c39240.png)
   ![Screenshot (3551)](https://user-images.githubusercontent.com/98501177/231103531-e4aa3c47-1dfc-4ac6-b12a-4b83357a922b.png)
4. Skill
   ![Screenshot (3552)](https://user-images.githubusercontent.com/98501177/231103775-0bdfac0d-29ee-4a94-83fa-1c0edf1efbbd.png)
   ![Screenshot (3553)](https://user-images.githubusercontent.com/98501177/231103780-ca32338e-258a-4ed1-a842-813d6ca6189e.png)
   ![Screenshot (3554)](https://user-images.githubusercontent.com/98501177/231103784-794cbcd7-524d-4f16-8cf0-f63af6fb1b0e.png)
   ![Screenshot (3555)](https://user-images.githubusercontent.com/98501177/231103796-0524d1f9-cf95-4361-9694-6addf5985451.png)
5. Profile
   ![Screenshot (3522)](https://user-images.githubusercontent.com/98501177/231104043-0e64a63f-9596-4cbc-bb11-30a56d95fa96.png)
   ![Screenshot (3523)](https://user-images.githubusercontent.com/98501177/231104063-0a169263-187a-4358-9ea1-ef5022881645.png)
   ![Screenshot (3524)](https://user-images.githubusercontent.com/98501177/231104069-f52e1bdd-3a03-4501-8fe7-e5174e18cacd.png)
   ![Screenshot (3525)](https://user-images.githubusercontent.com/98501177/231104076-8a86ec7c-8f7c-43df-94af-84c03f679c62.png)
   ![Screenshot (3526)](https://user-images.githubusercontent.com/98501177/231104085-9c903c22-cb9e-4cf3-a5f0-56e7472a24ec.png)
   ![Screenshot (3527)](https://user-images.githubusercontent.com/98501177/231104094-d0fb6e85-c6b6-4b9b-97af-190f0ce335c6.png)

# 4. SS Postman

![Screenshot (3521)](https://user-images.githubusercontent.com/98501177/231104508-06830b3b-bfee-4186-a080-8639deba8885.png)

# 5. SS Frontend

![Screenshot (3565)](https://user-images.githubusercontent.com/98501177/231104926-a2da9c92-1c45-4ea0-8698-143bc6406824.png)
![Screenshot (3566)](https://user-images.githubusercontent.com/98501177/231104927-c924649c-6636-4b19-9e8e-1052be070a63.png)
![Screenshot (3567)](https://user-images.githubusercontent.com/98501177/231104931-be97cb7d-e475-4844-9f71-65ab4be98663.png)
