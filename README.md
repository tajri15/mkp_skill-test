# Proyek Skill Test - Platform Tiket Bioskop Online

**Nama Lengkap:** Tajri Ramal Muzi

---

## Deskripsi Singkat

Ini adalah sebuah API backend sederhana yang dibuat menggunakan Golang untuk mensimulasikan sistem manajemen jadwal tayang bioskop. API ini mencakup fitur autentikasi (Login) dan otorisasi (khusus Admin) menggunakan JWT, serta operasi CRUD untuk data jadwal tayang.

---

## 1. System Design

Bagian ini menjelaskan arsitektur dan alur kerja dari sistem tiket bioskop online.

### Flowchart

Diagram berikut menggambarkan alur proses dari sisi pengguna, mulai dari login hingga mendapatkan tiket, serta penanganan jika terjadi kegagalan.

![Flowchart Sistem](https://raw.githubusercontent.com/tajri15/mkp_skill-test/main/System-Design/flowchart.png)

### Topologi Arsitektur

Diagram berikut menunjukkan komponen-komponen teknis yang membangun sistem ini dan bagaimana mereka saling berinteraksi.

![Topologi Sistem](https://raw.githubusercontent.com/tajri15/mkp_skill-test/main/System-Design/topology.png)

### Penjelasan Solusi Sistem

#### a. Solusi Sistem Pemilihan Tempat Duduk (Anti-Bentrok)

Untuk menangani banyak pengguna yang mencoba memesan kursi yang sama secara bersamaan (*race condition*), sistem ini menggunakan pendekatan **kunci sementara (temporary locking)** dengan **Redis** yang dikombinasikan dengan **WebSocket** untuk pembaruan *real-time*.

* **Alur Kerja**: Saat pengguna memilih kursi, sistem akan membuat sebuah *key* di Redis dengan masa berlaku terbatas (misalnya 10 menit). Hal ini jauh lebih cepat daripada mengunci baris di database utama (PostgreSQL). WebSocket akan secara instan memberitahu semua pengguna lain yang melihat denah yang sama bahwa kursi tersebut sedang tidak tersedia. Booking permanen hanya akan dicatat di PostgreSQL setelah pembayaran berhasil.
* **Keuntungan**: Pendekatan ini sangat **berperforma tinggi**, mencegah *double-booking*, dan memberikan pengalaman pengguna yang responsif.

#### b. Solusi Sistem Restock Tiket

"Restock" adalah proses membuat kursi tersedia kembali. Ini ditangani secara otomatis dalam dua skenario:

1.  **Transaksi Gagal/Waktu Habis**: Jika pengguna tidak menyelesaikan pembayaran dalam batas waktu, *key* di Redis akan **kadaluarsa secara otomatis**. Sistem akan menganggap kursi tersebut tersedia kembali tanpa perlu proses manual.
2.  **Pembatalan oleh Bioskop**: Jika jadwal dibatalkan, semua tiket terkait akan otomatis di-restock.

#### c. Solusi Refund/Pembatalan dari Pihak Bioskop

Untuk menangani pembatalan massal secara andal dan tanpa membuat server *hang*, sistem menggunakan **Message Queue (seperti RabbitMQ)**.

* **Alur Kerja**: Saat admin membatalkan jadwal, API hanya bertugas mengirimkan "pesan tugas refund" untuk setiap booking ke dalam antrean. Sebuah layanan *worker* terpisah akan mengambil tugas ini satu per satu, memproses refund melalui *payment gateway*, mengubah status di database, dan mengirim notifikasi ke pengguna.
* **Keuntungan**: Proses ini **asinkron** dan **andal (reliable)**. Jika terjadi kegagalan (misalnya *payment gateway* error), tugas tidak akan hilang dan bisa dicoba kembali, memastikan tidak ada pelanggan yang terlewat.

---

## 2. Struktur Folder

* `/system-design`: Berisi file gambar untuk Flowchart dan Topologi Sistem.
* `/database-design`: Berisi file gambar ERD dan skrip `schema.sql` untuk PostgreSQL.
* `/`: Berisi kode sumber utama untuk API Golang, yang dibagi lagi menjadi folder `controllers`, `models`, dan `middleware`.

---

## 3. Prasyarat (Prerequisites)

Pastikan perangkat Anda sudah terinstal:
* [Go](https://golang.org/dl/)
* [PostgreSQL](https://www.postgresql.org/download/)
* [Git](https://git-scm.com/downloads/)
* [Postman](https://www.postman.com/downloads/)

---

## 4. Cara Menjalankan Proyek

1.  **Clone Repositori**
    ```bash
    git clone [https://github.com/tajri15/mkp_skill-test.git](https://github.com/tajri15/mkp_skill-test.git)
    cd mkp_skill-test
    ```

2.  **Setup Database**
    * Pastikan server PostgreSQL Anda berjalan.
    * Buat database baru dengan nama `bioskop_db`.
    * Jalankan skrip SQL yang ada di `database-design/schema.sql` untuk membuat semua tabel dan data awal.

3.  **Konfigurasi Environment**
    * Salin file `.env.example` menjadi file baru bernama `.env`.
    * Buka file `.env` dan sesuaikan kredensial database (`DB_USER`, `DB_PASSWORD`, dll.) dengan konfigurasi lokal Anda.

4.  **Install Dependensi & Jalankan**
    ```bash
    # Install semua library yang dibutuhkan
    go mod tidy

    # Jalankan server API
    go run main.go
    ```
    Server akan berjalan di `http://localhost:8080`.

---

## 5. Cara Testing dengan Postman

1.  Buka Postman dan impor file `MKP Test API.postman_collection.json` yang ada di repositori ini.
2.  Jalankan request **"Admin Login"** terlebih dahulu untuk mendapatkan token JWT. Disarankan untuk menggunakan fitur variabel Postman untuk menyimpan token secara otomatis.
3.  Untuk request yang membutuhkan otorisasi Admin (seperti Create, Update, Delete Schedule), pastikan Bearer Token sudah terpasang.