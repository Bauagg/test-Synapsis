# test-Synapsist

# About
Aplikasi ini adalah backend RESTful API yang dikembangkan sebagai bagian dari test Synapsis. Aplikasi ini dirancang untuk mengelola sistem perpustakaan yang memungkinkan pengguna mengelola buku, peminjaman, dan pengembalian. Dibangun menggunakan bahasa Golang dengan framework Gin untuk server HTTP, Gorm sebagai ORM, dan PostgreSQL sebagai database.

# Features
- Manajemen Buku: Tambah, hapus, dan perbarui informasi buku dalam perpustakaan.
- Manajemen Pengguna: Daftar pengguna untuk meminjam dan mengembalikan buku.
- Manajemen Peminjaman dan Pengembalian: Pengguna dapat meminjam dan mengembalikan buku.
- Autentikasi: Implementasi autentikasi untuk mengamankan akses.

# Tech Stack
- Golang: Bahasa pemrograman utama.
- Gin: Framework HTTP untuk routing.
- Gorm: ORM untuk manipulasi database.
- PostgreSQL: Database relasional untuk penyimpanan data.
  
# Installation
Prerequisites
- Golang
- PostgreSQL

# Setup
1. Clone repository:
      - git clone https://github.com/Bauagg/test-Synapsis.git
      - cd test-Synapsis
2. Install dependencies:
      - go mod tidy
3. Buat file .env di direktori root proyek dan tambahkan konfigurasi berikut:
      - DB_HOST=127.0.0.1
      - DB_PORT=5433
      - DB_NAME=book-tes-Synapsis
      - DB_USER=postgres
      - DB_PASSWORD=root
      - APP_PORT=:8080
      - DB_TIMEZONE=Asia/Jakarta
      - SECRETKEY_TOKEN=kufruiuerjhioisduytuklp
      - URL_HOST=http://localhost:8080
4. Buat database baru di PostgreSQL sesuai dengan DB_NAME di file .env.
5. Jalankan aplikasi:
      - go run main.go

![image](https://github.com/user-attachments/assets/4d9dba49-2508-4383-ab4c-c9c791a31ce7)

# Contributing
1. Fork repository ini.
2. Buat branch baru (git checkout -b feature-branch).
3. Commit perubahan Anda (git commit -m 'Add some feature').
4. Push ke branch (git push origin feature-branch).
5. Buka Pull Request.

# License
This project is open-source and available under the MIT License.




