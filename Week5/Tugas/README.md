## Tugas

## Kerjakan

* Pull Request Semua Logic Function(di folder module) dan Type Struct (di folder model) di pull request ke repo https://github.com/indrariksa/Week5/Tugas/NPM
* Buat repo masing-masing untuk menyimpan backend sesuai tema masing-masing (contoh backend : https://github.com/indrariksa/be_presensi), dan jangan lupa listing package repo backend ke https://pkg.go.dev/.
```sh
git tag
git tag v0.0.1
git push origin --tags
go list -m github.com/{Repo Mahasiswa}/{Nama Repo Backend}@v0.0.1
```
* Mencoba deploy boiler plate https://github.com/bukulapak/iteung ke heroku
* Boiler plate sudah diisi dengan pemanggilan package dari repo https://github.com/{Repo Mahasiswa}/{Nama Repo Backend} *poin ke-2
* Pull Request dengan Judul 5-KELAS-NPM-NAMA, menyertakan file README.md dalam folder Week5/Tugas/{NPM}/ yang berisi :
  * Daftar URL Heroku app yang sudah memanggil function di package masing-masing mahasiswa yang sudah publish di https://pkg.go.dev/. [contoh](https://ws-ulbi.herokuapp.com/presensi) beserta skinsutannya (nilai 25 per URL)
  * Skrinsutan isi mongodb dari mongocompass yang sudah berisi data dummy kasus (nilai 15 per collection atau tabel)
  * URL Package yang sudah publish di pkg.go.dev dari https://github.com/{Repo Mahasiswa}/{Nama Repo Backend} beserta skinsutannya (nilai 20)
