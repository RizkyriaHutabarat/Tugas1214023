# HTTP Header and Body Capture

## Setting Endpoint di Pipedream

Pipedream merupakan endpoint testing API, untuk membaca header dan body(raw message) dari http client yang kita request. Http client request bisa dari postman maupun kode javascript kita. Pada langkah ini kita akan melakukan:
1. Membuat Akun pipedream
2. Testing Endpoint menggunakan Postman
3. Membangun HTML Form dengan Tailwind CSS Component
4. Membuat file js untuk mengirimkan data form ke endpoint pipedream
5. Melihat hasilnya di dashboard pipedream
6. Production mode menggunakan github pages

Disini kita akan melakukan request dengan menggunakan javascript fetch.

### Membuat Akun

* Kunjungi https://pipedream.com/ , terus klik Sign Up Free! 

  ![image](https://user-images.githubusercontent.com/11188109/220200037-5f556ae6-3bd6-4aa3-9869-00c1d119bb51.png)
  
* Klik sign up with github

  ![image](https://user-images.githubusercontent.com/26703717/222384917-1c39dada-86f2-401e-a394-bdb561eaeac3.png) 
* Authorize

  ![image](https://user-images.githubusercontent.com/26703717/222384917-1c39dada-86f2-401e-a394-bdb561eaeac3.png) 
  
* Pilih Nama Workspace yang tersedia, ditandai dengan tanda centang hijau, kemudian klik Continue

  ![image](https://user-images.githubusercontent.com/26703717/222385437-4206ac91-7c76-46b5-9ed1-807f59f47136.png)
* Opsional membagikan link workspace seperti ini, https://pipedream.com/@chapter03wsulbi/invite?token=d271f288ef995089871c5bdfbec91fd4
* Lanjutkan dengan menekan tombol Skip.
  
  ![image](https://user-images.githubusercontent.com/26703717/222385986-dbf17b1e-2d40-45d9-b6f9-6ad5ce111a0d.png)
* Klik New HTTP / Webhook REquest
  ![image](https://user-images.githubusercontent.com/26703717/222385768-722b1c42-a0a7-4d67-8508-e5645308e928.png)
* Pilih Event Data : Raw REquest, HTTP Response : 200 OK . Kemudian klik Save and Continue
  ![image](https://user-images.githubusercontent.com/26703717/222386318-d91cdc75-9699-4a9d-b893-d6292b9a3d2a.png)
* Akan keluar unique URL untuk endpoint : https://eogm9v38644wi49.m.pipedream.net . Kemudian kita coba dengan postman dahulu
  ![image](https://user-images.githubusercontent.com/26703717/222386457-5d42b873-4aa1-4825-ba2c-737b9605a959.png)

### Melakukan Testing Endpoint

Disini kita akan membuka Postman untuk melakukan testing endpoint dahulu, dengan contoh :
* Method POST Headers kita isi dengan Key : Login , Value : Bebas. 
  ![image](https://user-images.githubusercontent.com/26703717/222386790-3f76db30-b0c8-4abf-8b16-a61667c616a9.png)
  
* Pada bagian body isi dengan data json dengan cara klik Body-->raw-->JSON. Kemudian klik Send
  ![image](https://user-images.githubusercontent.com/26703717/222387563-194e886d-21bf-44ac-96a5-592259fa2343.png)

* Setelah klik Send maka akan menampilkan response seperti berikut
  ![image](https://user-images.githubusercontent.com/26703717/222387763-12060805-9276-4585-ad62-7d17771ae23d.png)

* Dashboard Pipedream akan muncul 1 New Event, kita buka event tersebut.
  ![image](https://user-images.githubusercontent.com/26703717/222388021-55f0b019-90ab-4875-a60b-aa50946ee0c9.png)

* Disana akan terlihat pada bagian headers ada Login yang kita masukkan dan pada bagian body ada json yang kita masukkan dari postman. Artinya endpoint dan http request bekerja dengan baik untuk menangkap header dan body yang dikirimkan.
  ![image](https://user-images.githubusercontent.com/26703717/222388910-72abefa1-8cb4-4a80-aca8-ce137617b5be.png)

* Simpan kode javascript yang dibuat oleh postman dengan menekan tanda </> dibagian pojok kanan atas, kemudian pilih javascipt - Fetch
  ![image](https://user-images.githubusercontent.com/26703717/222389196-429c9015-38e2-4333-aa99-fc914b7d58a0.png)
  ```javascript
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
  "npm": "123456789",
  "nama": "Virgil De Nguyen",
  "no_hp": 628122221814,
  "alamat": "bandung"
  });

  var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
  };

  fetch("https://eos4gcaqdr2ultq.m.pipedream.net?Login=chapter03", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
  ```


### Membuat Form Pendaftaran

* Kita cari dulu di google dengan keyword sign up form tailwind component free. Contoh : https://tailwindcomponents.com/component/simple-registersign-up-form/landing
  ![image](https://user-images.githubusercontent.com/11188109/220203961-34229f29-9cce-4352-b158-bcf7ad55b6bc.png)
* Kita unduh kodenya untuk kita pakai ditaruh di repo kita dan coba kita live server dari vscode.
  ![image](https://user-images.githubusercontent.com/26703717/222399786-36156e16-adc7-4327-8bdd-92f3f2e9de2c.png)
  ![image](https://user-images.githubusercontent.com/26703717/222391921-f577c304-96ec-4701-b674-52d154cbe022.png)
* Cari element dari input dan buttonnya kemudian kita beri id
  ![image](https://user-images.githubusercontent.com/26703717/222392252-ead9dc6b-4bb2-4e02-b211-fc44fbedea0c.png)
  ![image](https://user-images.githubusercontent.com/11188109/220205510-205c1381-073b-451f-b29c-ede13072c333.png)
* Buat file js kita masukkan dahulu javascript yang didapatkan dari postman ke js tersebut, kemudian panggil pada bagian bawah sebelum tag penutup body

  ![image](https://user-images.githubusercontent.com/26703717/222392397-900e8868-f866-46e4-b939-abaad852c444.png)
  ```html
  <script src="ws.js"></script>
  ```


### Membuat Fungsi Javascript

Pada bagian ini kita akan membuat fungsi-fungsi di javascript untuk mengirimkan data dari html menuju endpoint pipedream yang kita buat. Kenapa harus dibuat fungsi? agar kode program javascript tetap rapih dan mudah terbaca.
1. Membuat fungsi PostSignUp() yang berfungsi untuk melakukan Post Form Data Sign Up. Fungsi ini diambil dari postman, yang kita modif data inputan dari id form html.
   ```javascript
    function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "email": email,
      "password": password
    });

    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };

    fetch("https://eos4gcaqdr2ultq.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
    }
   ```
3. Membuat fungsi PushButton() untuk melakukan aksi setelah menekan tombol, pada bagian html button tambahkan atribut onclick.

   ![image](https://user-images.githubusercontent.com/11188109/220208507-0d5cb2cc-4979-410c-a3e0-a804caa732c4.png)
   ```javascript
    function PushButton(){
          namadepan=document.getElementById("namadepan").value;
          namabelakang=document.getElementById("namabelakang").value;
          email=document.getElementById("email").value;
          password=document.getElementById("password").value;
          PostSignUp(namadepan,namabelakang,email,password);
    }
   ```
4. Kita test dengan klik kanan Open with live server, kita isi form nya sambil inspect console lalu klik button yang kita buat. Terlihat dari console data success dikirim. kemudian kita lihat pada dashboard pipedream data sudah diterima dengan baik oleh endpoint baik itu header maupun body.
   ![image](https://user-images.githubusercontent.com/26703717/222400423-db80721c-07df-412f-bee6-cd5928b76e98.png)
   ![image](https://user-images.githubusercontent.com/26703717/222394243-80ff61f4-0b54-4d37-94f0-90a6ff15e41c.png)
   ![image](https://user-images.githubusercontent.com/26703717/222394382-d64a13b6-6bff-419c-9fcb-53059e72bfef.png)

### Tambahan Estetika UX

Pasti merasa aneh bukan setelah menekan tombol, tapi tampilan tidak berubah sama sekali seolah tidak terjadi apa-apa. Disini kita akan coba untuk mengubah tampilan jika tombol di klik, maka form akan disembunyikan dan menampilkan data yang diterima dari endpoint pipedream. Oke kita cukup menambahkan satu fungsi lagi dan memodifikasi fungsi PostSignUp() khususnya dibagian result then fetch nya. Langkahnya sebagai berikut :

* Pertama kita cari dulu element yang akan kita hidden, bisa menggunakan inspect elemetns untuk identifikasinya, kemudian kita kasih id
  ![image](https://user-images.githubusercontent.com/26703717/222395087-e80e49dc-e46e-4900-a052-1a995ae20e0c.png)
  ![image](https://user-images.githubusercontent.com/11188109/220210342-a81493b4-453b-4cd1-a526-3ca1092f6ebc.png)
* Kita coba script style display pada bagian console dengan menggunakan id yang sudah kita buat, style display mana yang bisa menghilangkan formsignup apakah block atau none. Terlihat di gambar none bisa mengilangkan element formsignup, kita **mungkin** akan pakai script ini di fungsi kita yang akan datang. Dari sini kita paham bagaimana menyembunyikan element dengan js.
  ![image](https://user-images.githubusercontent.com/26703717/222395574-f10f7e72-7090-45c2-a912-b94f51d4f1da.png)
  ```javascript
    document.getElementById("formsignup").style.display = 'none';
    document.getElementById("formsignup").style.display = 'block';
  ```
* Tambahkan fungsi GetResponse() di file js kita yang sudah dibuat sebelumnya, dan melakukan modifikasi dari fungsi PostSignUp pada bagian then result.
  ```javascript
    function GetResponse(result){
          document.getElementById("formsignup").innerHTML = result;
    }
  ```
  ![image](https://user-images.githubusercontent.com/26703717/222396080-4d8a0cd5-2b4d-4ee0-9d9d-0d77d00d991e.png)
* Kita ujicoba dengan mengisi form dan klik tombol submit, maka form kita sudah berhasil.
  ![image](https://user-images.githubusercontent.com/26703717/222396657-b8a826d3-6a80-4d4b-85f2-65603951a6b5.png)
  ![image](https://user-images.githubusercontent.com/26703717/222396763-fb243c41-aebd-49a2-a59f-a40d3789dacd.png)
* Cek pada web pipedream, seharusnya terdapat event baru seperti gambar di bawah
  ![image](https://user-images.githubusercontent.com/26703717/222401220-a090b549-e759-40d3-93e8-fba22bbdea55.png)
* Push pada repo github masing-masing Week3/Site/{NPM masing-masing}, kemudian lakukan pull request, dan coba jalankan hasil yang sudah dibuat pada Github Pages. Alamat github pages biasanya sub domain dari github.io. Contoh : https://indrariksa.github.io/WS/Week3/Site/NPM

## Kerjakan (Tugas Take Home)

* Buatlah form sign up yang melakukan POST ke pipedream.com ketika klik button.
* Form inputan yang dibuat minimal berisi 5 form.
* Buat Folder NPM masing-masing didalam Week3/Tugas/{NPM masing-masing} yang berisi 2 file minimal js dan html, dengan nama index.html, ws.js. Boleh menambahkan file css.
* file ws.js minimal berisi 3 fungsi, tidak boleh ada kode js diluar dari fungsi, semua harus masuk ke dalam fungsi js.
* Form dibangun dengan menggunakan CSS tailwind, desain html tidak boleh sama, harus berbeda satu sama lain.
* Pull Request dengan nama 3-Kelas-NPM-NAMA di folder Week3/Site/{NPM masing-masing}, dengan deskripsi disertakan di bawah ini.
* Sertakan screenshoot dari live server aplikasi dan pipedream.com
* Sertakan link Github Pages Sudah jalan di repo masing-masing yang dimasukkan pada file README.md
* Sertakan screenshoot juga sertifikat dari sertifikasi berikut (Minimal 1 sertifikat): 
  * https://www.mygreatlearning.com/academy/learn-for-free/courses/go-programming-language 
  * https://www.mindluster.com/certificate/3394
  * https://www.codecademy.com/learn/learn-go
  * https://www.coursera.org/specializations/google-golang


