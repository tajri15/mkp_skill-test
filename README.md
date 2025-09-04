# Proyek Skill Test - Platform Tiket Bioskop Online

**Nama Lengkap:** Tajri Ramal Muzi

---

## Deskripsi Singkat

Ini adalah sebuah API backend sederhana yang dibuat menggunakan Golang untuk mensimulasikan sistem manajemen jadwal tayang bioskop. API ini mencakup fitur autentikasi (Login) dan otorisasi (khusus Admin) menggunakan JWT, serta operasi CRUD untuk data jadwal tayang.

---

## Struktur Folder

* `/system-design`: Berisi file gambar untuk Flowchart dan Topologi Sistem.
* `/database-design`: Berisi file gambar ERD dan skrip `schema.sql` untuk PostgreSQL.
* `/`: Berisi kode sumber utama untuk API Golang, yang dibagi lagi menjadi folder `controllers`, `models`, dan `middleware`.

---

## Prasyarat (Prerequisites)

Pastikan perangkat Anda sudah terinstal:
* [Go](https://golang.org/dl/) (versi 1.21 atau lebih baru)
* [PostgreSQL](https://www.postgresql.org/download/)
* [Git](https://git-scm.com/downloads/)
* [Postman](https://www.postman.com/downloads/) (Opsional, untuk testing)

---

## Cara Menjalankan Proyek

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

## Cara Testing dengan Postman

1.  Buka Postman dan impor file `MKP Test API.postman_collection.json` yang ada di repositori ini.
2.  Jalankan request **"User Login"** terlebih dahulu untuk mendapatkan token JWT.
3.  Untuk request yang membutuhkan otorisasi Admin (seperti Create, Update, Delete Schedule), salin token tersebut dan masukkan ke tab **Authorization** -> **Bearer Token**.