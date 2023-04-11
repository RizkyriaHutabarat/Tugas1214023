# UJIAN TENGAH SEMESTER PEMROG 3

|                |                                     |
| -------------- | ----------------------------------- |
| NPM            | 1214050                             |
| NAMA           | DANI FERDINAN                       |
| PRODI/KELAS    | D4TI/2B                             |
| MATAKULIAH     | PEMOGRAMAN III                      |
| DOSEN PENGAMPU | INDRA RIKSA HERLAMBANG M.KOM., SFPC |
| HARI & TANGGAL | RABU, 12 APRIL 2023                 |
|                |                                     |

# PENJELASAN ALUR PROSES

## SOURCE CODE DAN PENJELASAN TAHAPAN
Pertama kita cek data yang telah dibuat sebelumnya. Berikut merupakan URL web service yang telah dibuat sebelumnya.

- https://ws-dani.herokuapp.com/dhs-all

- https://ws-dani.herokuapp.com/mahasiswa-all

- https://ws-dani.herokuapp.com/dosen-all

- https://ws-dani.herokuapp.com/matkul-all

kita cek cukup dengan browser

- DHS 
- Mata Kuliah
- Mahasiswa
- Dosen

Dapat dilihat data yang akan ditampilkan di frontend terdapat semua.

Selanjutnya buat template sesuai pertemuan 6, yang berisi folder `js` dan `template`.

![image](https://user-images.githubusercontent.com/55969069/231115709-5da8b3c4-cdd1-4927-8dd3-eba96a29f1da.png)

Jangan Lupa pada boiler plate atau projek web service nya diatur terlebih dahulu. pada file `config/cors.go` tambahkan domain pages kita di variabel `origins` 
```go
"https://daniferdinandall.github.io"
```
dan atur variabel 'Cors' menjadi seperti ini
```go
var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```

Kembali lagi pada Projek Frontend

pada folder template isikan template html untuk file `index.html`, `mahasiswa.html`, `dosen.html` dan `matkul.html`.

![image](https://user-images.githubusercontent.com/55969069/231115470-870c2699-e645-4530-ac46-e62423f358e6.png)

- Isi `index.html`

```javascript
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#0ed3cf">
    <meta name="msapplication-TileColor" content="#0ed3cf">
    <meta name="theme-color" content="#0ed3cf">

    <meta property="og:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />
    <meta property="og:image:width" content="1280" />
    <meta property="og:image:height" content="640" />
    <meta property="og:image:type" content="image/png" />

    <meta property="og:url" content="https://tailwindcomponents.com/component/simple-registersign-up-form/landing" />
    <meta property="og:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta property="og:description" content="Just a simple responsive sign up form with icons" />

    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:site" content="@TwComponents" />
    <meta name="twitter:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta name="twitter:description" content="Just a simple responsive sign up form with icons" />
    <meta name="twitter:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />

    <title>Simple Register/Sign Up Form by Scott Windon. </title>

    <script src="https://cdn.tailwindcss.com"></script>

    <script type="module" src="../js/fetch.js"></script>
</head>

<body>
    <aside
        class="ml-[-100%] fixed z-10 top-0 pb-3 px-6 w-full flex flex-col justify-between h-screen border-r bg-white transition duration-300 md:w-4/12 lg:ml-0 lg:w-[25%] xl:w-[20%] 2xl:w-[15%]">
        <div>
            <div class="-mx-6 px-6 py-4">
                <a href="#" title="home">
                    <img src="./img/LOGO ULBI - WIDE DARK.png" class="w-32"
                        alt="tailus logo">
                </a>
            </div>

            <div class="mt-8 text-center">
                <img src="./img/DMP_1868 - MASTER BIRU.jpg" alt=""
                    class="w-10 h-10 m-auto rounded-full object-cover lg:w-28 lg:h-28">
                    <h5 class="hidden mt-4 text-xl font-semibold text-gray-600 lg:block">Dani Ferdinan</h5>
                    <span class="hidden text-gray-400 lg:block">D4 TI / 2B</span>
            </div>

            <ul class="space-y-2 tracking-wide mt-8">
                <li>
                    <a href="#" aria-label="dashboard"
                        class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400">
                        <svg class="-ml-1 h-6 w-6" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M6 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8ZM6 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-1Z"
                                class="fill-current text-cyan-400 dark:fill-slate-600"></path>
                            <path d="M13 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2V8Z"
                                class="fill-current text-cyan-200 group-hover:text-cyan-300"></path>
                            <path d="M13 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-1Z"
                                class="fill-current group-hover:text-sky-300"></path>
                        </svg>
                        <span class="-mr-1 font-medium">DHS</span>
                    </a>
                </li>
                <li>
                    <a href="mahasiswa.html"
                        class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300" fill-rule="evenodd"
                                d="M2 6a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1H8a3 3 0 00-3 3v1.5a1.5 1.5 0 01-3 0V6z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M6 12a2 2 0 012-2h8a2 2 0 012 2v2a2 2 0 01-2 2H2h2a2 2 0 002-2v-2z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mahasiswa</span>
                    </a>
                </li>
                <li>
                    <a href="matkul.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600" fill-rule="evenodd"
                                d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M15 7h1a2 2 0 012 2v5.5a1.5 1.5 0 01-3 0V7z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mata Kuliah</span>
                    </a>
                </li>
                <li>
                    <a href="dosen.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Dosen</span>
                    </a>
                </li>
            </ul>
        </div>

        <div class="px-6 -mx-6 pt-4 flex justify-between items-center border-t">
            <button class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span class="group-hover:text-gray-700">Logout</span>
            </button>
        </div>
    </aside>
    <div class="ml-auto mb-6 lg:w-[75%] xl:w-[80%] 2xl:w-[85%]">
        <div class="sticky z-10 top-0 h-16 border-b bg-white lg:py-2.5">
            <div class="px-6 flex items-center justify-between space-x-4 2xl:container">
                <h5 hidden class="text-2xl text-gray-600 font-medium lg:block">DATA DHS</h5>
                <button class="w-12 h-16 -mr-2 border-r lg:hidden">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 my-auto" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                </button>
                <div class="flex space-x-4">
                    <!--search bar -->
                    <div hidden class="md:block">
                        <div class="relative flex items-center text-gray-400 focus-within:text-cyan-400">
                            <span class="absolute left-4 h-6 flex items-center pr-3 border-r border-gray-300">
                                <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 fill-current"
                                    viewBox="0 0 35.997 36.004">
                                    <path id="Icon_awesome-search" data-name="search"
                                        d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                                    </path>
                                </svg>
                            </span>
                            <input type="search" name="leadingIcon" id="leadingIcon" placeholder="Search here"
                                class="w-full pl-14 pr-4 py-2.5 rounded-xl text-sm text-gray-600 outline-none border border-gray-300 focus:border-cyan-300 transition">
                        </div>
                    </div>
                    <!--/search bar -->
                    <button aria-label="search"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200 md:hidden">
                        <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 mx-auto fill-current text-gray-600"
                            viewBox="0 0 35.997 36.004">
                            <path id="Icon_awesome-search" data-name="search"
                                d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                            </path>
                        </svg>
                    </button>
                    <button aria-label="chat"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" fill="none"
                            viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                        </svg>
                    </button>
                    <button aria-label="notification"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" viewBox="0 0 20 20"
                            fill="currentColor">
                            <path
                                d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <div class="px-12 pt-12 2xl:container">
            <!-- <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3"> -->

            <div class="flex flex-col">
                <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
                        <div class="overflow-hidden">
                            <table id="iniTabel" class="min-w-full text-left text-sm font-medium">
                                <thead class="border-b font-medium dark:border-neutral-500">
                                    <tr>
                                        <th scope="col" class="px-6 py-4">Nama & NPM</th>
                                        <th scope="col" class="px-6 py-4">PROGRAM STUDI & FAKULTAS</th>
                                        <th scope="col" class="px-6 py-4">Matkul & Nilai 1</th>
                                        <th scope="col" class="px-6 py-4">Matkul & Nilai 2</th>
                                        <th scope="col" class="px-6 py-4">Matkul & Nilai 3</th>
                                        <th scope="col" class="px-6 py-4">Matkul & Nilai 4</th>
                                        <!-- <th scope="col" class="px-6 py-4">Handle</th> -->
                                    </tr>
                                </thead>
                                <tbody>
                                    <!-- START -->
                                    <!-- <tr class="border-b dark:border-neutral-500">
                                        <td class="whitespace-nowrap px-6 py-4 font-medium">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#NAMA#</p>
                                                <p class="text-xs font-medium text-coolGray-500">#NPM#</p>
                                            </div>
                                        </td>
                                        <td class="whitespace-nowrap px-6 py-4">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#PROGRAM_STUDI#</p>
                                                <p class="text-xs font-medium text-coolGray-500">#FAKULTAS#</p>
                                            </div>
                                        </td>
                                        <td class="whitespace-nowrap px-6 py-4">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
                                                <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
                                            </div>
                                        </td>
                                        <td class="whitespace-nowrap px-6 py-4">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
                                                <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
                                            </div>
                                        </td>
                                        <td class="whitespace-nowrap px-6 py-4">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
                                                <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
                                            </div>
                                        </td>
                                        <td class="whitespace-nowrap px-6 py-4">
                                            <div class="w-auto">
                                                <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
                                                <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
                                            </div>
                                        </td>
                                    </tr> -->
                                    <!-- END -->
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>

            <!-- </div> -->
        </div>
    </div>
</body>

</html>
```

- Isi `mahasiswa.html`

```javascript
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#0ed3cf">
    <meta name="msapplication-TileColor" content="#0ed3cf">
    <meta name="theme-color" content="#0ed3cf">

    <meta property="og:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />
    <meta property="og:image:width" content="1280" />
    <meta property="og:image:height" content="640" />
    <meta property="og:image:type" content="image/png" />

    <meta property="og:url" content="https://tailwindcomponents.com/component/simple-registersign-up-form/landing" />
    <meta property="og:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta property="og:description" content="Just a simple responsive sign up form with icons" />

    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:site" content="@TwComponents" />
    <meta name="twitter:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta name="twitter:description" content="Just a simple responsive sign up form with icons" />
    <meta name="twitter:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />

    <title>Simple Register/Sign Up Form by Scott Windon. </title>
    <script type="module" src="../js/fetchMahasiswa.js"></script>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
    <aside
        class="ml-[-100%] fixed z-10 top-0 pb-3 px-6 w-full flex flex-col justify-between h-screen border-r bg-white transition duration-300 md:w-4/12 lg:ml-0 lg:w-[25%] xl:w-[20%] 2xl:w-[15%]">
        <div>
            <div class="-mx-6 px-6 py-4">
                <a href="#" title="home">
                    <img src="./img/LOGO ULBI - WIDE DARK.png" class="w-32"
                        alt="tailus logo">
                </a>
            </div>

            <div class="mt-8 text-center">
                <img src="./img/DMP_1868 - MASTER BIRU.jpg" alt=""
                    class="w-10 h-10 m-auto rounded-full object-cover lg:w-28 lg:h-28">
                    <h5 class="hidden mt-4 text-xl font-semibold text-gray-600 lg:block">Dani Ferdinan</h5>
                    <span class="hidden text-gray-400 lg:block">D4 TI / 2B</span>
            </div>

            <ul class="space-y-2 tracking-wide mt-8">
                <li>
                    <a href="index.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <!-- <a href="#" aria-label="dashboard" class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400"> -->
                        <svg class="-ml-1 h-6 w-6" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M6 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8ZM6 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-1Z"
                                class="fill-current text-cyan-400 dark:fill-slate-600"></path>
                            <path d="M13 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2V8Z"
                                class="fill-current text-cyan-200 group-hover:text-cyan-300"></path>
                            <path d="M13 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-1Z"
                                class="fill-current group-hover:text-sky-300"></path>
                        </svg>
                        <span class="-mr-1 font-medium">DHS</span>
                    </a>
                </li>
                <li>
                    <a href="#"
                        class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400">
                        <!-- <a href="#" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group"> -->
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300" fill-rule="evenodd"
                                d="M2 6a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1H8a3 3 0 00-3 3v1.5a1.5 1.5 0 01-3 0V6z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M6 12a2 2 0 012-2h8a2 2 0 012 2v2a2 2 0 01-2 2H2h2a2 2 0 002-2v-2z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mahasiswa</span>
                    </a>
                </li>
                <li>
                    <a href="matkul.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600" fill-rule="evenodd"
                                d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M15 7h1a2 2 0 012 2v5.5a1.5 1.5 0 01-3 0V7z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mata Kuliah</span>
                    </a>
                </li>
                <li>
                    <a href="dosen.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Dosen</span>
                    </a>
                </li>
            </ul>
        </div>

        <div class="px-6 -mx-6 pt-4 flex justify-between items-center border-t">
            <button class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span class="group-hover:text-gray-700">Logout</span>
            </button>
        </div>
    </aside>
    <div class="ml-auto mb-6 lg:w-[75%] xl:w-[80%] 2xl:w-[85%]">
        <div class="sticky z-10 top-0 h-16 border-b bg-white lg:py-2.5">
            <div class="px-6 flex items-center justify-between space-x-4 2xl:container">
                <h5 hidden class="text-2xl text-gray-600 font-medium lg:block">DATA MAHASISWA</h5>
                <button class="w-12 h-16 -mr-2 border-r lg:hidden">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 my-auto" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                </button>
                <div class="flex space-x-4">
                    <!--search bar -->
                    <div hidden class="md:block">
                        <div class="relative flex items-center text-gray-400 focus-within:text-cyan-400">
                            <span class="absolute left-4 h-6 flex items-center pr-3 border-r border-gray-300">
                                <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 fill-current"
                                    viewBox="0 0 35.997 36.004">
                                    <path id="Icon_awesome-search" data-name="search"
                                        d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                                    </path>
                                </svg>
                            </span>
                            <input type="search" name="leadingIcon" id="leadingIcon" placeholder="Search here"
                                class="w-full pl-14 pr-4 py-2.5 rounded-xl text-sm text-gray-600 outline-none border border-gray-300 focus:border-cyan-300 transition">
                        </div>
                    </div>
                    <!--/search bar -->
                    <button aria-label="search"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200 md:hidden">
                        <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 mx-auto fill-current text-gray-600"
                            viewBox="0 0 35.997 36.004">
                            <path id="Icon_awesome-search" data-name="search"
                                d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                            </path>
                        </svg>
                    </button>
                    <button aria-label="chat"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" fill="none"
                            viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                        </svg>
                    </button>
                    <button aria-label="notification"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" viewBox="0 0 20 20"
                            fill="currentColor">
                            <path
                                d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <div class="px-12 pt-12 2xl:container">
            <!-- <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3"> -->
                
                <div class="flex flex-col">
                    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
                      <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
                        <div class="overflow-hidden">
                          <table id="iniTabel" class="min-w-full text-left text-sm font-medium">
                            <thead class="border-b font-medium dark:border-neutral-500">
                              <tr>
                                <th scope="col" class="px-6 py-4">NPM</th>
                                <th scope="col" class="px-6 py-4">Nama</th>
                                <th scope="col" class="px-6 py-4">Fakultas</th>
                                <th scope="col" class="px-6 py-4">Program Studi</th>
                                <th scope="col" class="px-6 py-4">Dosen Wali</th>
                              </tr>
                            </thead>
                            <tbody>
                              <!-- <tr class="border-b dark:border-neutral-500">
                                <td class="whitespace-nowrap px-6 py-4 font-medium">#NPM#</td>
                                <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
                                <td class="whitespace-nowrap px-6 py-4">#FAKULTAS#</td>
                                <td class="whitespace-nowrap px-6 py-4">#PROGRAM_STUDI#</td>
                                <td class="whitespace-nowrap px-6 py-4">#DOSEN_WALI#</td>
                              </tr> -->
                            </tbody>
                          </table>
                        </div>
                      </div>
                    </div>
                  </div>

            <!-- </div> -->
        </div>
    </div>
</body>

</html>
```

- Isi `dosen.html`

```javascript
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#0ed3cf">
    <meta name="msapplication-TileColor" content="#0ed3cf">
    <meta name="theme-color" content="#0ed3cf">

    <meta property="og:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />
    <meta property="og:image:width" content="1280" />
    <meta property="og:image:height" content="640" />
    <meta property="og:image:type" content="image/png" />

    <meta property="og:url" content="https://tailwindcomponents.com/component/simple-registersign-up-form/landing" />
    <meta property="og:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta property="og:description" content="Just a simple responsive sign up form with icons" />

    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:site" content="@TwComponents" />
    <meta name="twitter:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta name="twitter:description" content="Just a simple responsive sign up form with icons" />
    <meta name="twitter:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />

    <title>Simple Register/Sign Up Form by Scott Windon. </title>
    <script type="module" src="../js/fetchDosen.js"></script>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
    <aside
        class="ml-[-100%] fixed z-10 top-0 pb-3 px-6 w-full flex flex-col justify-between h-screen border-r bg-white transition duration-300 md:w-4/12 lg:ml-0 lg:w-[25%] xl:w-[20%] 2xl:w-[15%]">
        <div>
            <div class="-mx-6 px-6 py-4">
                <a href="#" title="home">
                    <img src="./img/LOGO ULBI - WIDE DARK.png" class="w-32"
                        alt="tailus logo">
                </a>
            </div>

            <div class="mt-8 text-center">
                <img src="./img/DMP_1868 - MASTER BIRU.jpg" alt=""
                    class="w-10 h-10 m-auto rounded-full object-cover lg:w-28 lg:h-28">
                <h5 class="hidden mt-4 text-xl font-semibold text-gray-600 lg:block">Dani Ferdinan</h5>
                <span class="hidden text-gray-400 lg:block">D4 TI / 2B</span>
            </div>

            <ul class="space-y-2 tracking-wide mt-8">
                <li>
                    <a href="index.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <!-- <a href="#" aria-label="dashboard" class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400"> -->
                        <svg class="-ml-1 h-6 w-6" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M6 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8ZM6 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-1Z"
                                class="fill-current text-cyan-400 dark:fill-slate-600"></path>
                            <path d="M13 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2V8Z"
                                class="fill-current text-cyan-200 group-hover:text-cyan-300"></path>
                            <path d="M13 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-1Z"
                                class="fill-current group-hover:text-sky-300"></path>
                        </svg>
                        <span class="-mr-1 font-medium">DHS</span>
                    </a>
                </li>
                <li>
                    <!-- <a href="#"
                        class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400"> -->
                        <a href="mahasiswa.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300" fill-rule="evenodd"
                                d="M2 6a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1H8a3 3 0 00-3 3v1.5a1.5 1.5 0 01-3 0V6z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M6 12a2 2 0 012-2h8a2 2 0 012 2v2a2 2 0 01-2 2H2h2a2 2 0 002-2v-2z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mahasiswa</span>
                    </a>
                </li>
                <li>
                    <a href="matkul.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600" fill-rule="evenodd"
                                d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M15 7h1a2 2 0 012 2v5.5a1.5 1.5 0 01-3 0V7z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mata Kuliah</span>
                    </a>
                </li>
                <li>
                    <a href="dosen.html"
                        class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400">
                    <!-- <a href="#" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group"> -->
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Dosen</span>
                    </a>
                </li>
            </ul>
        </div>

        <div class="px-6 -mx-6 pt-4 flex justify-between items-center border-t">
            <button class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span class="group-hover:text-gray-700">Logout</span>
            </button>
        </div>
    </aside>
    <div class="ml-auto mb-6 lg:w-[75%] xl:w-[80%] 2xl:w-[85%]">
        <div class="sticky z-10 top-0 h-16 border-b bg-white lg:py-2.5">
            <div class="px-6 flex items-center justify-between space-x-4 2xl:container">
                <h5 hidden class="text-2xl text-gray-600 font-medium lg:block">DATA DOSEN</h5>
                <button class="w-12 h-16 -mr-2 border-r lg:hidden">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 my-auto" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                </button>
                <div class="flex space-x-4">
                    <!--search bar -->
                    <div hidden class="md:block">
                        <div class="relative flex items-center text-gray-400 focus-within:text-cyan-400">
                            <span class="absolute left-4 h-6 flex items-center pr-3 border-r border-gray-300">
                                <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 fill-current"
                                    viewBox="0 0 35.997 36.004">
                                    <path id="Icon_awesome-search" data-name="search"
                                        d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                                    </path>
                                </svg>
                            </span>
                            <input type="search" name="leadingIcon" id="leadingIcon" placeholder="Search here"
                                class="w-full pl-14 pr-4 py-2.5 rounded-xl text-sm text-gray-600 outline-none border border-gray-300 focus:border-cyan-300 transition">
                        </div>
                    </div>
                    <!--/search bar -->
                    <button aria-label="search"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200 md:hidden">
                        <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 mx-auto fill-current text-gray-600"
                            viewBox="0 0 35.997 36.004">
                            <path id="Icon_awesome-search" data-name="search"
                                d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                            </path>
                        </svg>
                    </button>
                    <button aria-label="chat"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" fill="none"
                            viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                        </svg>
                    </button>
                    <button aria-label="notification"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" viewBox="0 0 20 20"
                            fill="currentColor">
                            <path
                                d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <div class="px-12 pt-12 2xl:container">
            <!-- <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3"> -->
                
                <div class="flex flex-col">
                    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
                      <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
                        <div class="overflow-hidden">
                          <table id="iniTabel" class="min-w-full text-left text-sm font-medium">
                            <thead class="border-b font-medium dark:border-neutral-500">
                              <tr>
                                <th scope="col" class="px-6 py-4">Kode Dosen</th>
                                <th scope="col" class="px-6 py-4">Nama</th>
                                <th scope="col" class="px-6 py-4">Phone Number</th>
                              </tr>
                            </thead>
                            <tbody>
                              <!-- <tr class="border-b dark:border-neutral-500">
                                <td class="whitespace-nowrap px-6 py-4">#KODE#</td>
                                <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
                                <td class="whitespace-nowrap px-6 py-4">#PHONE#</td>
                              </tr> -->
                            </tbody>
                          </table>
                        </div>
                      </div>
                    </div>
                  </div>

            <!-- </div> -->
        </div>
    </div>
</body>

</html>
```

- Isi `matkul.html`

```javascript
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#0ed3cf">
    <meta name="msapplication-TileColor" content="#0ed3cf">
    <meta name="theme-color" content="#0ed3cf">

    <meta property="og:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />
    <meta property="og:image:width" content="1280" />
    <meta property="og:image:height" content="640" />
    <meta property="og:image:type" content="image/png" />

    <meta property="og:url" content="https://tailwindcomponents.com/component/simple-registersign-up-form/landing" />
    <meta property="og:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta property="og:description" content="Just a simple responsive sign up form with icons" />

    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:site" content="@TwComponents" />
    <meta name="twitter:title" content="Simple Register/Sign Up Form by Scott Windon" />
    <meta name="twitter:description" content="Just a simple responsive sign up form with icons" />
    <meta name="twitter:image"
        content="http://tailwindcomponents.com/storage/2335/temp66711.png?v=2023-02-20 21:54:06" />

    <title>Simple Register/Sign Up Form by Scott Windon. </title>
    <script type="module" src="../js/fetchMatkul.js"></script>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
    <aside
        class="ml-[-100%] fixed z-10 top-0 pb-3 px-6 w-full flex flex-col justify-between h-screen border-r bg-white transition duration-300 md:w-4/12 lg:ml-0 lg:w-[25%] xl:w-[20%] 2xl:w-[15%]">
        <div>
            <div class="-mx-6 px-6 py-4">
                <a href="#" title="home">
                    <img src="./img/LOGO ULBI - WIDE DARK.png" class="w-32"
                        alt="tailus logo">
                </a>
            </div>

            <div class="mt-8 text-center">
                <img src="./img/DMP_1868 - MASTER BIRU.jpg" alt=""
                    class="w-10 h-10 m-auto rounded-full object-cover lg:w-28 lg:h-28">
                    <h5 class="hidden mt-4 text-xl font-semibold text-gray-600 lg:block">Dani Ferdinan</h5>
                    <span class="hidden text-gray-400 lg:block">D4 TI / 2B</span>
            </div>

            <ul class="space-y-2 tracking-wide mt-8">
                <li>
                    <a href="index.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <!-- <a href="#" aria-label="dashboard" class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400"> -->
                        <svg class="-ml-1 h-6 w-6" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M6 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8ZM6 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-1Z"
                                class="fill-current text-cyan-400 dark:fill-slate-600"></path>
                            <path d="M13 8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2V8Z"
                                class="fill-current text-cyan-200 group-hover:text-cyan-300"></path>
                            <path d="M13 15a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-1Z"
                                class="fill-current group-hover:text-sky-300"></path>
                        </svg>
                        <span class="-mr-1 font-medium">DHS</span>
                    </a>
                </li>
                <li>
                    <!-- <a href="#"
                        class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400"> -->
                        <a href="mahasiswa.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300" fill-rule="evenodd"
                                d="M2 6a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1H8a3 3 0 00-3 3v1.5a1.5 1.5 0 01-3 0V6z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M6 12a2 2 0 012-2h8a2 2 0 012 2v2a2 2 0 01-2 2H2h2a2 2 0 002-2v-2z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mahasiswa</span>
                    </a>
                </li>
                <li>
                    <a href="matkul.html" aria-label="dashboard" class="relative px-4 py-3 flex items-center space-x-4 rounded-xl text-white bg-gradient-to-r from-sky-600 to-cyan-400">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600" fill-rule="evenodd"
                                d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z"
                                clip-rule="evenodd" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M15 7h1a2 2 0 012 2v5.5a1.5 1.5 0 01-3 0V7z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Mata Kuliah</span>
                    </a>
                </li>
                <li>
                    <a href="dosen.html" class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path class="fill-current text-gray-600 group-hover:text-cyan-600"
                                d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" />
                            <path class="fill-current text-gray-300 group-hover:text-cyan-300"
                                d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z" />
                        </svg>
                        <span class="group-hover:text-gray-700">Dosen</span>
                    </a>
                </li>
            </ul>
        </div>

        <div class="px-6 -mx-6 pt-4 flex justify-between items-center border-t">
            <button class="px-4 py-3 flex items-center space-x-4 rounded-md text-gray-600 group">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span class="group-hover:text-gray-700">Logout</span>
            </button>
        </div>
    </aside>
    <div class="ml-auto mb-6 lg:w-[75%] xl:w-[80%] 2xl:w-[85%]">
        <div class="sticky z-10 top-0 h-16 border-b bg-white lg:py-2.5">
            <div class="px-6 flex items-center justify-between space-x-4 2xl:container">
                <h5 hidden class="text-2xl text-gray-600 font-medium lg:block">DATA MATAKULIAH</h5>
                <button class="w-12 h-16 -mr-2 border-r lg:hidden">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 my-auto" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                </button>
                <div class="flex space-x-4">
                    <!--search bar -->
                    <div hidden class="md:block">
                        <div class="relative flex items-center text-gray-400 focus-within:text-cyan-400">
                            <span class="absolute left-4 h-6 flex items-center pr-3 border-r border-gray-300">
                                <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 fill-current"
                                    viewBox="0 0 35.997 36.004">
                                    <path id="Icon_awesome-search" data-name="search"
                                        d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                                    </path>
                                </svg>
                            </span>
                            <input type="search" name="leadingIcon" id="leadingIcon" placeholder="Search here"
                                class="w-full pl-14 pr-4 py-2.5 rounded-xl text-sm text-gray-600 outline-none border border-gray-300 focus:border-cyan-300 transition">
                        </div>
                    </div>
                    <!--/search bar -->
                    <button aria-label="search"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200 md:hidden">
                        <svg xmlns="http://ww50w3.org/2000/svg" class="w-4 mx-auto fill-current text-gray-600"
                            viewBox="0 0 35.997 36.004">
                            <path id="Icon_awesome-search" data-name="search"
                                d="M35.508,31.127l-7.01-7.01a1.686,1.686,0,0,0-1.2-.492H26.156a14.618,14.618,0,1,0-2.531,2.531V27.3a1.686,1.686,0,0,0,.492,1.2l7.01,7.01a1.681,1.681,0,0,0,2.384,0l1.99-1.99a1.7,1.7,0,0,0,.007-2.391Zm-20.883-7.5a9,9,0,1,1,9-9A8.995,8.995,0,0,1,14.625,23.625Z">
                            </path>
                        </svg>
                    </button>
                    <button aria-label="chat"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" fill="none"
                            viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                        </svg>
                    </button>
                    <button aria-label="notification"
                        class="w-10 h-10 rounded-xl border bg-gray-100 focus:bg-gray-100 active:bg-gray-200">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 m-auto text-gray-600" viewBox="0 0 20 20"
                            fill="currentColor">
                            <path
                                d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <div class="px-12 pt-12 2xl:container">
            <!-- <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3"> -->
                
                <div class="flex flex-col">
                    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
                      <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
                        <div class="overflow-hidden">
                          <table id="iniTabel" class="min-w-full text-left text-sm font-medium">
                            <thead class="border-b font-medium dark:border-neutral-500">
                              <tr>
                                <th scope="col" class="px-6 py-4">Kode Matkul</th>
                                <th scope="col" class="px-6 py-4">Nama Matkul</th>
                                <th scope="col" class="px-6 py-4">Dosen Pengampu</th>
                                <th scope="col" class="px-6 py-4">SKS</th>
                              </tr>
                            </thead>
                            <tbody>
                              <!-- <tr class="border-b dark:border-neutral-500">
                                <td class="whitespace-nowrap px-6 py-4 font-medium">#KODE#</td>
                                <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
                                <td class="whitespace-nowrap px-6 py-4">#DOSEN#</td>
                                <td class="whitespace-nowrap px-6 py-4">#SKS#</td>
                              </tr> -->
                            </tbody>
                          </table>
                        </div>
                      </div>
                    </div>
                  </div>

            <!-- </div> -->
        </div>
    </div>
</body>

</html>
```

foldet `img` dalam folder `template` digunakan untuk menyimpan gambar yang digunakan pada template.

lalu pada tabel di setiap template html diberi id dengan nama iniTabel.

![image](https://user-images.githubusercontent.com/55969069/231116124-b3029bc5-5823-4ba0-979c-5ce5870dff27.png)

setelah itu didalam folder `template` buat 3 folder yaitu `config`, `controller` dan `temp`.

didalam folder `config` di isi `url.js` yang isi nya seperti ini

```javascript
export let urlAPIdhs = "https://ws-dani.herokuapp.com/dhs-all";
export let urlAPImatkul = "https://ws-dani.herokuapp.com/matkul-all";
export let urlAPIdosen = "https://ws-dani.herokuapp.com/dosen-all";
export let urlAPImahasiswa = "https://ws-dani.herokuapp.com/mahasiswa-all";
```
lalu pada folder `temp` buat folder dengan nama `table.js`.

```javasript
export let isiTabel =
    `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#NAMA#</p>
            <p class="text-xs font-medium text-coolGray-500">#NPM#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#PROGRAM_STUDI#</p>
            <p class="text-xs font-medium text-coolGray-500">#FAKULTAS#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI#</p>
            </div>
        </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL1#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI1#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL2#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI2#</p>
        </div>
    </td>
    <td class="whitespace-nowrap px-6 py-4">
        <div class="w-auto">
            <p class="text-s font-semibold text-coolGray-800">#MATKUL3#</p>
            <p class="text-s font-medium text-coolGray-800">Grade: #NILAI3#</p>
        </div>
    </td>
</tr>
`
export let isiTabelDosen =
    `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4">#KODE#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#PHONE#</td>
</tr>
`
export let isiTabelMatkul =
    `
    <tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">#KODE#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#DOSEN#</td>
    <td class="whitespace-nowrap px-6 py-4">#SKS#</td>
  </tr>
  `

export let isiTabelMahasiswa = `
<tr class="border-b dark:border-neutral-500">
    <td class="whitespace-nowrap px-6 py-4 font-medium">#NPM#</td>
    <td class="whitespace-nowrap px-6 py-4">#NAMA#</td>
    <td class="whitespace-nowrap px-6 py-4">#FAKULTAS#</td>
    <td class="whitespace-nowrap px-6 py-4">#PROGRAM_STUDI#</td>
    <td class="whitespace-nowrap px-6 py-4">#DOSEN_WALI#</td>
</tr>
`
```

Lalu untuk folder `controller` tambahkan 4 file yang berguna untuk memasukan data kedalam tabel di tamplate html. file yang dibuat adalah 

- `get.js` untuk data DHS

- `getMahasiswa.js` untuk data Mahsiswa

- `getDosen.js` untuk data Dosen

- `getMatkul.js` untuk data Mata Kuliah

isi dari `get.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
        isiTabel.replace("#NAMA#", value.mahasiswa.nama)
            .replace("#NPM#", value.mahasiswa.npm)
            .replace("#PROGRAM_STUDI#", value.mahasiswa.program_studi?value.mahasiswa.program_studi.nama:"#PROGRAM_STUDI#")
            .replace("#FAKULTAS#", value.mahasiswa.fakultas?value.mahasiswa.fakultas.nama:"#FAKULTAS#")
            .replace("#MATKUL#", value.mata_kuliah?value.mata_kuliah[0].nama:"#MATKUL#")
            .replace("#NILAI#", value.mata_kuliah?value.mata_kuliah[0].nilai:"#NILAI#")
            .replace("#MATKUL1#", value.mata_kuliah?value.mata_kuliah[1].nama:"#MATKUL1#")
            .replace("#NILAI1#", value.mata_kuliah?value.mata_kuliah[1].nilai:"#NILAI1#")
            .replace("#MATKUL2#", value.mata_kuliah?value.mata_kuliah[2].nama:"#MATKUL2#")
            .replace("#NILAI2#", value.mata_kuliah?value.mata_kuliah[2].nilai:"#NILAI2#")
            .replace("#MATKUL3#", value.mata_kuliah?value.mata_kuliah[3].nama:"#MATKUL3#")
            .replace("#NILAI3#", value.mata_kuliah?value.mata_kuliah[3].nilai:"#NILAI3#")
            // .replace("#WARNA#", getRandomColor())
            // .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
}
```

isi dari `getMahasiswa.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelMahasiswa } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
        isiTabelMahasiswa.replace("#NAMA#", value.nama)
            .replace("#NPM#", value.npm)
            .replace("#PROGRAM_STUDI#", value.program_studi?value.program_studi.nama:"#PROGRAM_STUDI#")
            .replace("#FAKULTAS#", value.fakultas?value.fakultas.nama:"#FAKULTAS#")
            .replace("#DOSEN_WALI#", value.dosen_wali?value.dosen_wali.nama:"#DOSEN_WALI#");
    addInner("iniTabel", content);
}
```

isi dari `getDosen.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelDosen } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah?value.mata_kuliah[1].nama:"1")
    let content =
    isiTabelDosen.replace("#KODE#", value.kode_dosen)
            .replace("#NAMA#", value.nama)
            .replace("#PHONE#", value.phone_number?value.phone_number:"#PHONE#")
    addInner("iniTabel", content);
}
```

isi dari `getMatkul.js`

```javascript
import { addInner } from "https://bukulapak.github.io/element/process.js";
// import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabelMatkul } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    console.log(value)
    console.log(value.mata_kuliah ? value.mata_kuliah[1].nama : "1")
    let content =
        isiTabelMatkul.replace("#KODE#", value.kode_matkul)
            .replace("#NAMA#", value.nama)
            .replace("#DOSEN#", value.dosen ? value.dosen.nama : "#DOSEN#")
            .replace("#SKS#", value.sks ? value.sks : "#SKS#")
    addInner("iniTabel", content);
}
```

Lalu pada folder `js` tambahkan 4 file untuk mendapat 

isi dari `getMahasiswa.js`

```javascript

```

isi dari `getDosen.js`

```javascript

```

isi dari `getMatkul.js`

```javascript

```

## URL HEROKU
https://ws-dani.herokuapp.com/dhs-all

https://ws-dani.herokuapp.com/mahasiswa-all

https://ws-dani.herokuapp.com/dosen-all

https://ws-dani.herokuapp.com/matkul-all

## URL GITHUB PAGES FRONTEND
pages: https://daniferdinandall.github.io/fe_uts/template

repo: https://github.com/daniferdinandall/fe_uts

## SCREENSHOOT SETIAP COLLECTION PADA MONGO
### DHS
![image](https://user-images.githubusercontent.com/55969069/231034237-3aa10da5-49ce-464c-8915-c369aaa419d6.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231034342-00c2eaf7-fb77-408b-ba3a-7bdc50ca3cd0.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231034293-fbed9c88-55dd-4163-a50d-5f0342feec3e.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231034419-406ef1cc-95b0-4b50-becc-af056f13f5f6.png)

## SCREENSHOOT APLIKASI FRONTEND UNTUK SETIAP DATA/COLLECTION
### DHS
![image](https://user-images.githubusercontent.com/55969069/231035230-a6f618cc-a4e5-4e0a-9c8d-f0776ecf3d8b.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231035275-0839db16-65c5-4c44-9152-2a94eb931664.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231035385-85683c12-994a-4341-adca-fb4cda4ad0a0.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231035340-c198898f-7d52-4904-b291-584802fe0ade.png)


## SCREENSHOOT POSTMAN UNTUK SETIAP DATA/COLLECTION
### DHS
![image](https://user-images.githubusercontent.com/55969069/231034658-5a73313d-0b49-4ffb-9469-4360fe081eff.png)

### MAHASISWA
![image](https://user-images.githubusercontent.com/55969069/231034695-77db7087-5281-407e-a4cf-e7f4034d448d.png)

### DOSEN
![image](https://user-images.githubusercontent.com/55969069/231034722-3e50f782-87fb-4b10-854a-b6fd75858609.png)

### MATKUL
![image](https://user-images.githubusercontent.com/55969069/231034763-f277099e-d3fa-4fa4-ac7e-cea70587d33f.png)
