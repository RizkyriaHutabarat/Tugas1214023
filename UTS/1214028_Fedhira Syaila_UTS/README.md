# Tahapan dalam memunculkan data backend di frontend:

1. Kita harus membuat databasenya terlebih dahulu pada mongoDB (terdapat pada week 4)
2. Dalam repo backend kita harus mendefinisikan terlebih dahulu struct yang akan kita gunakan, misalnya disini tema backend saya yaitu TAGIHAN. Jadi, saya membuat struct nasabah, penagih, tagihan dan bank
   ![struct](https://user-images.githubusercontent.com/93858982/230938929-198b51c2-df83-4176-ba01-05b0d4b700c3.png)

3. Membuat folder dan file sesuai dengan gambar dibawah ini

   ![folder_be](https://user-images.githubusercontent.com/93858982/230939308-af28ceb9-2cc8-4d39-94d9-2331581e850d.png)

4. Lalu, buatlah function untuk menambahkan data dan mengambil data di mongoDB. Dalam repo backend ini saya membuat function di folder controller dalam file controller.go. Contohnya seperti gambar dibawah ini.
   <img width="519" alt="controller_go" src="https://user-images.githubusercontent.com/93858982/230939550-7f25cec9-2993-47ae-9818-f6d0883f8eec.png">

5. Lalu, saya membuat file dev_test.go untuk file testing insert data dan get data. Contohnya seperti gambar dibawah ini
   ![dev_test](https://user-images.githubusercontent.com/93858982/230939630-5204990d-d2ce-4ae8-9906-fb6c66c5ed87.png)

6. Jangan lupa untuk selalu go mod tidy di terminal jika ada perubahan
7. Untuk menjalankan testing, kitab isa menggunakan keyword go test di terminal, jika berhasil akan keluar data yang di insert seperti gambar dibawah ini
   ![go_test](https://user-images.githubusercontent.com/93858982/230939861-ec3c5eab-0de5-4823-9715-e912f1aaffa4.png)

8. Hasilnya bisa dilihat pada mongoDB masing-masing, seperi pada gambar dibawah ini
   ![mongo](https://user-images.githubusercontent.com/93858982/230940014-be11f300-94da-469e-b5fb-e3b36a016d6b.png)

9. Jika repo backend sudah benar dan tidak ada error ataupun masalah lainnya, kita akan melakukan publish package

```
git tag
git tag v0.0.1
git push origin --tags
go list -m github.com/(username github)/(nama repo)@v0.0.1

```

10. Untuk mengetahui apakah package kita sudsah di publish atau belum bisa kunjungi link https://pkg.go.dev/github.com/(username github)/(nama repo)

11. Langkah selanjutnya, pada week 5 kita sudah membuat repo boiler plate, serperti pada gambar dibawah ini
    ![boiler](https://user-images.githubusercontent.com/93858982/230940171-e948f658-6d52-4c85-9790-e6ce89e27bb4.png)

12. Pada repo boiler plate ini, kita harus melakukan go get repo backend yang sudah kita buat sebelumnya dengan keyword

```
go get github.com/(username github)/(repo backend)
```

13. Lalu, pada folder controller di file coba.go kita mengimport package yang diperlukan, seperti gambar dibawah ini. Saya melakukan import package repo backend untuk menghubungkan backend dengan repo boiler plate. Dengan keyword

```
inimodel "github.com/Fedhira/be_tagihan/model"
inimodul "github.com/Fedhira/be_tagihan/module"

```

<img width="279" alt="import_controller" src="https://user-images.githubusercontent.com/93858982/230940268-58668b7d-175d-477c-b583-ab65ec87614c.png">

14. Setelah itu, tambahkan fungsi GetAll data tanpa filter untuk menampilkan semua data yang ada di mongoDB seperti pada gambar dibawah ini
    <img width="161" alt="url_getall" src="https://user-images.githubusercontent.com/93858982/230940891-40ed8ce3-f98d-4899-b6d0-8c55550090a1.png">

15. Jangan lupa untuk selalu lakukan go mod tidy di terminal, lalu ketik go run main.go di terminal untuk melihat apakah datanya sudah bisa tampil atau belum
16. Jika sudah nanti akan tampil link seperti gambar dibawah ini

    ![go_run](https://user-images.githubusercontent.com/93858982/230941002-6f6df9b4-d72d-4160-becf-6d62641c9459.png)

17. Cek dengan url yang sudah kita buat, misalnya saya url /all untuk menampilkan semua data bank, jika berhasil akan tampil semua data seperti gambar dibawah ini
    ![hasil_gorun](https://user-images.githubusercontent.com/93858982/230941072-3f6adc3c-c457-45d9-af72-3ed08938a7f1.png)

18. Jika tidak ada masalah ataupun error, langkah selanjutnya yaitu melakukan commit and push pada repo github dan Heroku. Untuk melakukan push pada Heroku dengan keyword

```
git status
git add.
git status
git commit -m "hhh"
git push
git push heroku main

```

Jika berhasil akan muncul link Heroku seperti gambar dibawah ini
![link_heroku](https://user-images.githubusercontent.com/93858982/230941253-3a75c415-f78d-4bbd-88cc-6f303e730307.png)

20. Langkah selanjutnya, kita akan membuat repo baru untuk frontend di github
    ![new_repo](https://user-images.githubusercontent.com/93858982/230941316-e96a5e40-1ff4-423c-985a-974b7a84cfd2.png)

21. Kemudian, jangan lupa untuk copy link repo untuk di git clone
    ![copy_linkrepo](https://user-images.githubusercontent.com/93858982/230941444-5fa30889-d11d-4c14-b0fc-5cbbb7296cb2.png)

22. Lalu, kita akan melakukan git clone dengan git bash here di folder mana pun sesuai kemauan. Git clone berfungsi untuk mengunduh code yang ada pada repository.
    ![git_clone](https://user-images.githubusercontent.com/93858982/230941543-4868725d-252f-497a-a8f5-d1ca0481fb1c.png)

23. Selanjutnya, masukkan template yang ingin dipakai ke dalam repo frontend. Disini saya menggunakan template dari windmill
    ![windmill_template](https://user-images.githubusercontent.com/93858982/230941677-007a56a6-73d1-4ef4-9bbc-1976b56a0b6b.png)

24. Kemudian, membuat folder js terdiri dari 3 folder yaitu temp, controller dan config dan folder template adalah template yang akan kita gunakan

    <img width="81" alt="folder_fe" src="https://user-images.githubusercontent.com/93858982/230941761-63fa1c58-209d-4165-8c15-e8d48ce45aaa.png">

25. Lalu, tambahkan code dibawah ini pada file index.html

```
<script type="module" src="../js/fetch.js"></script>
<script src="../js/main.js"></script> pada file index.html
```

![index_html](https://user-images.githubusercontent.com/93858982/230941805-e9c0767a-4a1e-43a8-b41c-4c4a4ebbf722.png)

26. Kemudian, lakukan setting repo github di bagian pages
    ![setting_githubpages](https://user-images.githubusercontent.com/93858982/230941917-097c2052-e828-443e-b828-18eebaa4a701.png)

27. Dalam file fetch.js berisi

```
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```

28. Lalu, lakukan commit and push. Kemudian, tambahkan code serta url github kita seperti dibawah ini dalam file cors.go di repo boiler plate
    ![cors_go](https://user-images.githubusercontent.com/93858982/230941992-953f5713-a660-449b-b754-c35eee5d232f.png)

29. Langkah selanjutnya, tambahkan id table seperti gambar dibawah ini
    <img width="736" alt="id_table" src="https://user-images.githubusercontent.com/93858982/230942046-e654e581-34ea-40a5-b0c6-d7f23ec125d9.png">

30. Dalam table.js berisi code

```
export let isiTabel = `
<tbody
                    class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800"
                  >
<tr class="h-18 border-b border-coolGray-100 text-gray-700 dark:text-gray-400">
    <th class="whitespace-nowrap px-4 text-left px-4 py-3">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
            </div>
            <div class="w-auto p-2">
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ABC</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA_BANK#</p>
                <p class="text-xs text-gray-600 dark:text-gray-400">#LOKASI_BANK#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 text-sm font-medium text-coolGray-800 text-left px-4 py-3">
    <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#NAMA_NASABAH#</p>
        <p class="text-xs text-gray-600 dark:text-gray-400">#PHONE_NUMBER#</p>
        <p class="text-xs text-gray-600 dark:text-gray-400">#EMAILN#</p>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 text-sm font-medium text-coolGray-500 text-left px-4 py-3">
    <div class="w-auto p-2">
    <p class="text-xs font-semibold text-coolGray-800">#TAGIHAN#</p>
    <p class="text-xs text-gray-600 dark:text-gray-400">#TANGGAL#</p>
    </div>
    </th>
    <th class="whitespace-nowrap px-4 text-sm font-medium text-#col#-500 text-left px-4 py-3">#DESKRIPSI#</th>
    <th class="whitespace-nowrap px-4 text-sm font-medium text-#col#-500 text-left px-4 py-3">#STATUS#</th>
    <th class="whitespace-nowrap px-4 text-sm font-medium text-#col#-500 text-left px-4 py-3">
    <div class="w-auto p-2">
    <p class="text-xs font-semibold text-coolGray-800">#NAMA_PENAGIH#</p>
    <p class="text-xs text-gray-600 dark:text-gray-400">#NO_HP#</p>
    <p class="text-xs text-gray-600 dark:text-gray-400">#EMAILP#</p>
    </div>
    </th>
    <th class="whitespace-nowrap pr-4 text-sm font-medium text-coolGray-800 px-4 py-3">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
</tbody>
`;
```

31. Lalu, kita coba melihat hasilnya apakah tabel data backendnya sudah muncul atau belum
    ![githubpages](https://user-images.githubusercontent.com/93858982/230942292-e8063086-7eb3-4692-8d1f-e480a34b6764.png)
32. Jika sudah berhasil makan data akan tampil seperti ini
    ![frontend](https://user-images.githubusercontent.com/93858982/230942186-67383091-0570-4e39-8fc7-885eb6dd5a48.png)

# URL HEROKU:

https://fedhirasyaila.herokuapp.com/all

# URL FRONTEND:

https://fedhira.github.io/fe_UTS/template/index.html

# Screenshot MongoDB

- bank
  ![bank](https://user-images.githubusercontent.com/93858982/230943664-efa76f89-af10-487a-af35-d013a0ae8029.png)

- nasabah
  ![nasabah1](https://user-images.githubusercontent.com/93858982/230943499-af31cbdb-9287-4634-9d71-54055e41a49d.png)

- penagih
  ![penagih1](https://user-images.githubusercontent.com/93858982/230943594-61bdcbb2-2708-4c77-b082-b4da391e727f.png)

- tagihan
  ![tagihan1](https://user-images.githubusercontent.com/93858982/230943387-137cd144-095d-4995-8992-cc7557a15c81.png)

# Screenshot Postman

![postman_tagihan](https://user-images.githubusercontent.com/93858982/230943853-2410e204-3d43-400f-b83f-7c110debc1c6.png)

# Screenshot Frontend

![frontend](https://user-images.githubusercontent.com/93858982/230944110-ffc61552-2422-4b20-a0e8-57c932d7ef01.png)

![frontend1](https://user-images.githubusercontent.com/93858982/230944103-32c8c5af-7176-4865-82c0-12bb5c8be1e8.png)

![frontend2](https://user-images.githubusercontent.com/93858982/230944099-58f73c33-9de3-4f51-895b-939b81934e44.png)

![frontend3](https://user-images.githubusercontent.com/93858982/230944090-d83f030d-334d-4618-a235-46822cc31cb9.png)

![frontend4](https://user-images.githubusercontent.com/93858982/230944129-dc206460-af7f-4254-af26-93f63b0313d8.png)
