## Tugas Besar
Pembuatan API Server dan Frontend untuk Mengambil Data

Deskripsi Tugas:
Anda diminta untuk membuat sebuah API server menggunakan Golang dengan framework Fiber. API server ini akan menyediakan beberapa endpoint untuk mengambil dan mengelola data. Selain itu, Anda juga diminta untuk membuat sebuah tampilan frontend menggunakan HTML dan JavaScript yang akan mengambil data dari API server tersebut.

### Spesifikasi Tugas:

Pembuatan API Server:
Buatlah API server menggunakan Golang dan framework Fiber.

API server harus memiliki endpoint untuk melakukan operasi CRUD (Create, Read, Update, Delete).

Endpoints yang harus disediakan:

a. GET /users - Mengambil daftar semua pengguna.

b. GET /users/{id} - Mengambil informasi pengguna berdasarkan ID.

c. POST /users - Membuat pengguna baru.

d. PUT /users/{id} - Memperbarui informasi pengguna berdasarkan ID.

e. DELETE /users/{id} - Menghapus pengguna berdasarkan ID.

Tanggapi setiap permintaan dengan status HTTP yang sesuai dan format response JSON yang konsisten.

Tambahan:
- Anda dapat menggunakan Postman atau aplikasi serupa untuk menguji endpoint API yang telah dibuat.
- Dokumentasikan API yang telah dibuat dengan menggunakan Swagger untuk menjelaskan cara menggunakan setiap endpoint dan format data yang diharapkan.

---

Pembuatan Frontend:
- Buatlah tampilan frontend menggunakan HTML dan JavaScript.
- Tampilan frontend harus memiliki fitur untuk menampilkan semua data, menampilkan detail data berdasarkan ID, menambah data melalui form input, mengubah data  melalui form input, dan menghapus data.
- Gunakan AJAX atau Fetch API untuk mengambil data dari API server menggunakan metode GET, POST, PUT, dan DELETE.
- Tampilkan hasil data yang diperoleh dari API server ke dalam tampilan frontend.

Catatan: 
Tugas besar ini dirancang untuk melatih kemampuan mahasiswa dalam membuat frontend sederhana menggunakan JavaScript dan HTML untuk mengambil dan menampilkan data dari API server yang telah dibuat menggunakan Golang dengan framework Fiber.


### Ketentuan TB :
- Membuat Project API (Backend dan Frontend)
- Backend diperbolehkan menggunakan project saat UTS
- Bahasa pemrograman Backend (server) menggunakan GO Lang
- Bahasa pemrograman Frontend bisa menggunakan JavaScript + HTML seperti yang sudah dipraktekkan sebelumnya, atau juga bisa menggunakan Framework JavaScript
- Database diperbolehkan menggunakan selain MongoDB (Dengan catatan tidak boleh berada di localhost / local)
- Dokumentasi API menggunakan Swagger
- Implementasikan validasi input (misal : kolom nama wajib diisi, format kolom email yang harus sesuai, dll) pada data di sisi server (API server) atau pada sisi frontend untuk memastikan data yang diterima sesuai dengan format yang diharapkan.
- Terdapat alert message ketika melakukan insert, update, delete pada frontend.
- Nilai tambah jika terdapat fitur atau apapun yang belum pernah dipelajari di kelas (Bisa berupa implementasi register dan login pada frontend ,autentikasi (contoh: JWT), otorisasi, enkripsi data, dll). Dapat kalian jelaskan hal ini ketika presentasi.
- Buat PPT terkait pendahuluan, alur aplikasi yang dibuat dan database. (Tidak perlu memasukkan koding pada PPT)
- Membuat dokumen dalam bentuk pdf (diagram erd database, printscreen semua output untuk masing-masing endpoint baik pada server API (Go Lang) maupun client/frontend) serta berikan penjelasan setiap langkah yang kalian lakukan (setiap penjelasan alur proses tidak boleh mirip dengan kelompok lain, jika terdapat kemiripan maka mahasiswa yang bersangkutan mendapat nilai 0). Note: Disarankan menggunakan daftar isi
- Presentasi di pertemuan UAS (Asesmen)
- Nama kelompok dan judul dikumpulkan pada Folder WS/TugasBesar/KELOMPOK.md
- Menyertakan link atau url repo backend, boilerplate, frontend pada Folder WS/TugasBesar/KELOMPOK.md. (Bisa menyusul)

---
Selamat mengerjakan tugas besar!!
