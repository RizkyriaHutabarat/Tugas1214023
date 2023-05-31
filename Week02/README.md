# TOOLS

1. POSTMAN
   * Install postman. Pilih salah satu dari sini :
   ```sh
   https://github.com/public-apis/public-apis
   https://github.com/farizdotid/DAFTAR-API-LOKAL-INDONESIA
   ```
   * lakukan request ke public api dengan method GET dengan postman contoh : https://alexwohlbruck.github.io/cat-facts/docs/
   ![image](https://user-images.githubusercontent.com/26703717/220868366-271e81d8-707e-4611-9aa7-74f8b31cd008.png)
   
2. TAILWIND
   * Untuk membuat antarmuka menggunakan komponen tailwind contoh :https://tailwindcomponents.com/component/multi-line-table
   
3. Plugin LiveServer (VScode)
   ![image](https://user-images.githubusercontent.com/11188109/218396548-483f109a-c88c-4bc6-96d0-5d784a447556.png)

# Menghubungkan situs dengan Public API

1. Buat folder NPM masing (contoh :123456) di dalam folder Site
2. Buat template tailwindcss simpan sebagai index.html
3. buat file js dengan nama fetch.js panggil di bawah dengan script, sebelum tag penutup body
   ```html
   <script src="./fetch.js"></script>
   ```
4. Pada aplikasi Postman, pilih menu </> atau code pilih javascript fetch ,kemudian paste pada file fetch.js
   ![image](https://user-images.githubusercontent.com/26703717/220873867-a5685abb-3e1b-43ce-a1dc-57a5b3b38065.png)
5. Kemudian ubah dari script fetch.js contoh :
    ```js
    var myHeaders = new Headers();
    myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

    var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow'
    };

    hasil=""

    fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
      .then(response => response.text())
      .then(result => tampilkan(result))
      .catch(error => console.log('error', error));

      function tampilkan(result){
        console.log(result);
        hasil=JSON.parse(result);
      }
    ```
6. Buka dengan live server, inspect lihat di console.

    ![image](https://user-images.githubusercontent.com/26703717/220874890-4ba431d2-a221-4a75-bf3f-cc914ef0976e.png)
    
    akan terlihat variabel hasil di console log
    
    ![image](https://user-images.githubusercontent.com/11188109/218419862-983ddbba-a445-41a5-a703-f2d56829888a.png)
    
    untuk akses variabel global hasil kita pakai console log
    
    ![image](https://user-images.githubusercontent.com/11188109/218420311-3aa5700f-f6ec-4d5a-82d3-2000fdb24df6.png)
    
    kemudian tambahkan ramuan looping untuk mengeluarkan semuanya
    
    ![image](https://user-images.githubusercontent.com/26703717/220876915-e12233d6-0b17-4a08-a737-03d0dfc9c6f4.png)
    
# Setting Repo untuk Github Pages
  * Lakukan pada repo masing-masing mahasiswa
  * Setting pages pilih deployment from branch pilih main. Maka form kita bisa diakses dari menjadi github pages. Alamat github pages biasanya sub domain dari github.io. Contoh : {username github}.github.io --> https://indrariksa.github.io/WS/Week2/Site/NPM
  ![image](https://user-images.githubusercontent.com/26703717/220871197-19aa1b56-fa71-49b7-bcc7-91cad9b783ed.png)
  
## Kerjakan (Minggu depan sudah dikumpul semua = 3 Maret 2023)
1. Buat folder NPM masing-masing (contoh : 123456) di dalam folder Week2-->Tugas.
2. Pilih Public API di bawah atau bisa cari sendiri :
  * https://github.com/public-apis/public-apis
  * https://github.com/farizdotid/DAFTAR-API-LOKAL-INDONESIA
3. Kirim link API yang dipakai di grup WA
4. Ada dua file yang dibuat yaitu dengan nama index.html dan fetch.js. Boleh menambahkan file css
5. Buat html bisa membaca semua data dari API dan disimpan dalam tabel (Minimal data yang ditampilkan yaitu 4 data), pastikan setiap orang berbeda, tidak boleh sama API nya
  ![image](https://user-images.githubusercontent.com/26703717/221107173-eb43145c-0ac8-4880-8af7-6d185292246f.png)
6. Nilai tambah jika API yang ditampilkan selain text (misal : menampilkan gambar, embed video, dll)
7. Sertakan link Github Pages Sudah jalan di repo masing-masing (misal : https://indrariksa.github.io/WS/Week2/Tugas/NPM/). Note : buat file README.md di dalam folder NPM masing-masing kemudian isikan linknya disitu
8. Pull Request Dengan Subjek : 2-KELAS-NPM-NAMA, 
9. Pastikan akses API hanya menggunakan protokol https

