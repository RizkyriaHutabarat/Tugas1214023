## Tugas

## Kerjakan
* Disarankan menggunakan Text Editor GoLand (jika laptop tidak kuat skip aja ges langsung ke poin selanjutnya), untuk GoLand bisa ambil educational dahulu disini menggunakan email ulbi https://www.jetbrains.com/shop/eform/students, kemudian download https://www.jetbrains.com/go/download/, jika sudah aktivasi saja
* Masih menggunakan tema yang sama pada Week4
* Buat repo masing-masing untuk menyimpan backend sesuai tema masing-masing (contoh backend : https://github.com/indrariksa/be_presensi), struktur folder harus sama seperti https://github.com/indrariksa/be_presensi Logic Function(di folder module) dan Type Struct (di folder model)
* Jangan lupa listing package repo backend ke https://pkg.go.dev/.
```sh
git tag
git tag v0.0.1
git push origin --tags
go list -m github.com/{Repo Mahasiswa}/{Nama Repo Backend}@v0.0.1
```
* Mencoba deploy boiler plate https://github.com/bukulapak/iteung ke heroku, sebelum melakukan push ke heroku alangkah baiknya running via local dulu dengan cara
```sh
go run main.go
```
![image](https://user-images.githubusercontent.com/26703717/225911044-28bd321d-8502-499d-b818-c69f06782aaa.png)
Buka link tanda merah di atas di browser untuk testing
Kemudian push ke repo github dan ke heroku

* Boiler plate sudah diisi dengan pemanggilan package dari repo https://github.com/{Repo Mahasiswa}/{Nama Repo Backend} *poin ke-2
  Untuk pemanggilan package contohnya menggunakan go get {nama package} :
  ```sh
  go get github.com/indrariksa/be_presensi
  ```
  Di atas merupakan cara memanggil package github.com/indrariksa/be_presensi yang sudah listing di https://pkg.go.dev/github.com/indrariksa/be_presensi
* Pull Request dengan Judul 5-KELAS-NPM-NAMA, menyertakan : 
  * Pull Request ke repo https://github.com/indrariksa/Week5/Tugas/NPM
  * file README.md dalam folder Week5/Tugas/{NPM}/ yang berisi :
    * Daftar URL Heroku app yang sudah memanggil function di package masing-masing mahasiswa yang sudah publish di https://pkg.go.dev/. [contoh](https://ws-ulbi.herokuapp.com/presensi) beserta skinsutannya (nilai 25 per URL) jika terdapat 3 url atau 3 endpoint maka sertakan 3 URL
    * URL Package yang sudah publish di pkg.go.dev dari https://github.com/{Repo Mahasiswa}/{Nama Repo Backend} beserta skinsutannya (nilai 20)
  * Skrinsutan isi mongodb dari mongocompass yang sudah berisi data dummy kasus (nilai 15 per collection atau tabel)

  
