## Tahapan deployment fe presensi mahasiswa menggunakan template baru

Dengan mengikuti tahapan pada latihan Week6(https://github.com/indrariksa/WS/blob/main/Week6/README.md), dan melanjutkan tugas pada Week6(https://github.com/Fatwaff/Tugas_Pemrograman3_Fatwa-Fatahillah-Fatah_1214038/tree/main/Week6/Tugas/1214038), tahapan membuat tampilan baru untuk aplikasi fe presensi mahasiswa adalah sebagai berikut:

1. Menambahkan file index2.html pada folder template

![ss uts](https://user-images.githubusercontent.com/96001238/230802557-0ee8f2ce-1614-40a6-9382-9f14e30ce521.PNG)

2. Mencari template yang ingin digunakan, kemudian salin source code-nya

![ss uts](https://user-images.githubusercontent.com/96001238/230806187-11d9fe15-683a-4a38-8aa1-deeffed7dc34.png)

3. Paste source code template pada index2.html, tidak lupa membuat tag html terlebih dahulu, dan menambahkan tag script di dalam section head untuk memanggil file js/main.js dan js/fetch.js

![ss uts](https://user-images.githubusercontent.com/96001238/230806645-cac73ff2-91d0-4a87-a99a-a088b5d821c0.png)

4. Mengedit dan menghapus tampilan yang tidak diperlukan, kemudian menambahkan tag tabel yang nanti akan digunakan untuk menampilkan data yang diambil dari API yang sudah dibuat

![ss uts](https://user-images.githubusercontent.com/96001238/230807666-51bec65f-8f2d-4ab2-9daa-d42676d1955c.png)

5. Menambahkan variabel pada table.js yang digunakan sebagai isian dari data tabel yang ingin ditampilkan, disini saya ingin menampilkan data untuk 7 tabel

![ss uts](https://user-images.githubusercontent.com/96001238/230808327-b4c02737-9752-46df-adf2-13772dc687da.png)

6. Menambahkan function pada get.js untuk mengisi tabel dengan data dari API, function yang dibuat berjumlah sebanyak tabel yang ingin ditampilkan

![ss uts](https://user-images.githubusercontent.com/96001238/230809794-86edde6f-1f45-43a1-a48c-bad0d029ceb4.png)

7. Menambah lagi function yang digunakan untuk looping data yang diambil dari masing-masing function diatas

![ss uts](https://user-images.githubusercontent.com/96001238/230812131-0c2675ef-b2ce-49b1-b765-396fde38487e.png)

8. Pada url.js, menambahkan link API data yang ingin ditampilkan

![ss uts](https://user-images.githubusercontent.com/96001238/230852034-1e7a3579-1a47-44a9-b0bb-78a191358ad5.png)

9. Setelah itu import API dan function looping yang sudah dibuat ke dalam fetch.js

![ss uts](https://user-images.githubusercontent.com/96001238/230873326-d6143614-829b-4503-b6f6-5a8b91ea1e21.png)

10. Kemudian pada index2.html, menambahkan tag tabel sebanyak data tabel yang ingin ditampilkan yaitu 7 tag tabel, namun 6 dari tag tabel diberi display none agar pada saat membuka halaman hanya 1 tabel yang terlihat

![ss uts](https://user-images.githubusercontent.com/96001238/231061625-2496265c-739f-4da7-a839-b48acbd55398.png)

11. Membuat file activeTable.js yang di dalamnya terdapat function-function yang digunakan untuk mengatur display dari tabel-tabel tersebut

![ss uts](https://user-images.githubusercontent.com/96001238/231064708-661f8d3e-1ad0-4565-b891-deb9159fdfd5.png)

12. Setelah itu memanggil function-function dari activeTable.js tersebut pada index2.html dan dimasukkan kedalam card list yang sudah dibuat. Di sini masing-masing function yang dipanggil menyesuaikan dengan nama dari masing-masing card list. Untuk memanggilnya function-nya menggunakan atribute onclick.

![ss uts](https://user-images.githubusercontent.com/96001238/231069723-3e6f1e53-e083-43b9-ba91-1e1b8ee7d227.png)

## URL Heroku

- https://ws-fatwa.herokuapp.com/semua-presensi
- https://ws-fatwa.herokuapp.com/semua-mahasiswa
- https://ws-fatwa.herokuapp.com/semua-kelas
- https://ws-fatwa.herokuapp.com/semua-prodi
- https://ws-fatwa.herokuapp.com/semua-matkul
- https://ws-fatwa.herokuapp.com/semua-dosen
- https://ws-fatwa.herokuapp.com/semua-ruangan

## URL github pages frontend

https://fatwaff.github.io/fe_presensi-mhs/template/index2.html

## Skrinsut collection mongoDB

- collection presensi

![ss uts)](https://user-images.githubusercontent.com/96001238/231209096-6ae34637-8ecf-4bca-bf6e-65a6542457dd.png)

- collection mahasiswa

![ss uts)](https://user-images.githubusercontent.com/96001238/231210028-b8543b33-3006-472f-a1b5-1f2cbd8c5bc4.png)

- collection kelas

![ss uts)](https://user-images.githubusercontent.com/96001238/231209981-5b430f3a-a21a-49f5-843a-7dc402bf661c.png)

- collection prodi

![ss uts)](https://user-images.githubusercontent.com/96001238/231209993-44044e89-ffb2-4b2c-9e44-5a2f1dd8cf33.png)

- collection matkul

![ss uts)](https://user-images.githubusercontent.com/96001238/231210003-16bb9350-2c2c-4b97-809f-9c2010b62dd0.png)

- collection dosen

![ss uts)](https://user-images.githubusercontent.com/96001238/231210015-25d118c7-b828-46c8-9804-4926997f529f.png)

- collection ruang

![ss uts)](https://user-images.githubusercontent.com/96001238/231210021-2ef5a75a-814e-4f54-8639-61434e8c9f5e.png)

## Skrinsut postman

- GET semua data presensi

![ss uts)](https://user-images.githubusercontent.com/96001238/231211365-c675a9dc-199c-463e-b431-daf8911dcc78.png)

- GET semua data mahasiswa

![ss uts)](https://user-images.githubusercontent.com/96001238/231211372-4dae97b1-8194-4b5c-80f2-56bf26dce4b6.png)

- GET semua data kelas

![ss uts)](https://user-images.githubusercontent.com/96001238/231211380-5b08fa28-6512-47b2-90e0-40eadd0ebc46.png)

- GET semua data prodi

![ss uts)](https://user-images.githubusercontent.com/96001238/231211322-7f216843-c382-40f9-9292-8ff0d0426cd2.png)

- GET semua data matkul

![ss uts)](https://user-images.githubusercontent.com/96001238/231211340-cc632cb4-f693-404b-9257-c30fed1c6335.png)

- GET semua data dosen

![ss uts)](https://user-images.githubusercontent.com/96001238/231211346-1149bc1e-6e58-4e4b-ab6a-2050eb56ba4f.png)

- GET semua data ruangan

![ss uts)](https://user-images.githubusercontent.com/96001238/231211355-3cecb732-003e-4ab1-9446-122780d8a98a.png)

## Skrinsut Frontend

- Tabel Presensi

![ss uts)](https://user-images.githubusercontent.com/96001238/231212859-49612712-42aa-45cc-9f9e-97cc2c347042.png)

- Tabel Mahasiswa

![ss uts)](https://user-images.githubusercontent.com/96001238/231212864-a1b4a2ab-9e5a-4472-8f93-fa0b928a56a5.png)

- Tabel Kelas

![ss uts)](https://user-images.githubusercontent.com/96001238/231212869-1b51505f-9ae9-41d3-b14a-bef9750a0df1.png)

- Tabel Prodi

![ss uts)](https://user-images.githubusercontent.com/96001238/231212848-2f601383-a466-4708-b74e-2e4ca5ba8e9a.png)

- Tabel Matkul

![ss uts)](https://user-images.githubusercontent.com/96001238/231212874-924248fe-3261-48db-a8aa-db895a5d90f8.png)

- Tabel Dosen

![ss uts)](https://user-images.githubusercontent.com/96001238/231212834-6f1a57f0-f8ee-4f7d-a82b-2e37bf0cc6c1.png)

- Tabel Ruangan

![ss uts)](https://user-images.githubusercontent.com/96001238/231212851-93381b3d-6c43-4600-accd-89832d775d67.png)
