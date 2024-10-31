# GoRoutines
- Parallel and Concurency Programming

## Cara runing unit test
- `go test -v -run=NamaFunctionUnitTest`

## Process
- Adalah sebuah eksekusi dari program
- Mengkonsumsi memory besar
- Saling terisolasi antar process lain
- Lama untuk dijalankan dan dihentikan

## Thread
- Adalah segmen atau bagian dari process
- Menggunakan memory kecil
- Bisa saling terhubung antar thread
- Cepat untuk dijalankan dan dihentikan

## CPU-Bound
- Hanya membutuhkan CPU untuk menjalankannya, sangat bergantung pada kecepatan CPU
- Algoritma seperti ini tidak cocok menggunakan Concurency Programming, namun bisa dibantu dengan implementasi Parallel Programming
- Untuk Algoritma CPU-Bound kurang cocok untuk GoLang, lebih cocok di Java

## I/O Bound
- Sangat bergantung pada kecepatan input output device
- Contohnya aplikasi seperti membaca data dari file, database. Umumnya aplikasi web dan backend menggunakan algoritma ini
- Aplikasi jenis IO-Bound lebih baik menggunakan Concurrency Programing, meskipun bisa implementasi Parallel Programing

## Parallel Programming
- Definisi: Parallel programming adalah teknik menjalankan beberapa tugas secara bersamaan pada berbagai core CPU, memungkinkan eksekusi tugas-tugas tersebut benar-benar berjalan di waktu yang sama (simultan).
Tujuan: Mempercepat eksekusi tugas dengan membaginya ke beberapa bagian yang dapat dijalankan secara paralel.
- Contoh: Prosesor dengan banyak core yang menjalankan tugas besar dengan membagi bagian-bagian tugas ke setiap core.
- Penggunaan: Sangat berguna dalam aplikasi yang memerlukan komputasi berat seperti pemrosesan data dalam jumlah besar, simulasi ilmiah, atau aplikasi machine learning.
Contoh Parallel Programming: Misalkan Anda memiliki 1 juta angka yang ingin dijumlahkan. Dengan parallel programming, Anda bisa membagi data ini menjadi beberapa bagian, dan tiap bagian dijumlahkan secara paralel oleh core yang berbeda, lalu hasilnya digabungkan.

## Concurrency Programming
- Definisi: Concurrency programming adalah teknik yang memungkinkan beberapa tugas tampak berjalan secara bersamaan dengan mengatur pengaksesan dan pengalihan di antara tugas-tugas tersebut, meskipun tidak harus berjalan di waktu yang sama.
- Tujuan: Mengelola banyak tugas secara efisien dengan membuat mereka beralih secara cepat dan meminimalkan waktu tunggu.
Contoh: Sistem operasi yang mengelola beberapa aplikasi berjalan "bersamaan" dengan bergantian mengalokasikan waktu prosesor ke tiap aplikasi.
- Penggunaan: Concurrency lebih umum pada aplikasi yang menunggu masukan/keluaran (I/O-bound), seperti server web yang harus menangani banyak permintaan atau aplikasi GUI yang merespons input pengguna.
Contoh Concurrency Programming: Misalkan ada dua tugas: membaca file besar dan menghitung sesuatu. Dengan concurrency, alih-alih menunggu seluruh file terbaca, program bisa memulai perhitungan di sela-sela membaca bagian-bagian file. Tugas-tugas ini mungkin terlihat berjalan bersamaan, tetapi sebenarnya mereka hanya bergantian.

### Perbedaan Utama
- Aspek	|| Parallel Programming	|| Concurrency Programming
- Tujuan || Meningkatkan kecepatan dengan pemrosesan simultan || Mengelola beberapa tugas dengan efisien agar tetap responsif
- Eksekusi || Tugas benar-benar berjalan secara simultan di beberapa core CPU || Tugas bergantian eksekusi (tidak harus simultan)
- Jenis Tugas || CPU-bound (komputasi intensif) || I/O-bound atau multitasking
- Contoh Penerapan || Komputasi ilmiah, Machine Learning || Web server, aplikasi GUI

## Goroutine
- Goroutine adalah thread ringan yang dikelola oleh Go Runtime
- Ukuran Goroutine sangat kecil (sekitar 2kb), dibandingkan Thread (bisa mencapai 1mb)
- Goroutine berjalan secara concurency, tidak seperti Thread yang berjalan parallel
- Namun Goroutine berjalan di dalam Thread dan dijalankan oleh Go Scheduler, dimana jumlah threadnya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)

## Terminologi GoScheduler
- G -> Goroutine
- M -> Thread (Machine)
- P -> Processor

## Membuat Goroutine
- Untuk membuat goroutine cukup menambahkan perintah `go` sebelum memanggil function
- function tersebut akan berjalan secara asynchronous, tidak ditunggu sampai function tersebut selesai
- Goroutine kurang cocok untuk function yang mengembalikan value

## Channel
- Tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
- Di Channel terdapat pengirim dan penerima, yang biasanya pengirim dan penerima adalah goroutine yang berbeda
- Channel cocok sebagai alternatif seperti mekanisme async await, atau seperti Future di Java
- Ketika mengirim data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut
- Channel secara default hanya bisa menerima satu data, ketika data sudah diterima maka Channel akan kembali kosong
- Channel harus diclose ketika tidak digunakan
- Ketika Channel digunakan sebagai parameter pada sebuah function, maka otomatis akan pass by refference, tidak perlu menggunakan pointer untuk channel