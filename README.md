# TokoOnline API

TokoOnline API adalah aplikasi RESTful API berbasis Golang yang menyediakan manajemen produk, pesanan, dan autentikasi pengguna. API ini menggunakan JWT untuk keamanan endpoint tertentu.

## Fitur
- **Autentikasi**: Login, register, profil, reset password, dan logout
- **Manajemen Produk**: Tambah, lihat, edit, dan hapus produk
- **Manajemen Pesanan**: Lihat daftar pesanan, tandai pesanan sebagai diproses atau selesai
- **Keamanan**: JWT untuk akses terbatas di endpoint profil, produk, dan pesanan

## Prasyarat

- **Golang** v1.18 atau lebih baru
- **MySQL** untuk basis data
- **Git**

## Instalasi

1. **Clone Repository**
   ```bash
   git clone https://github.com/danigilang17/tokoonline/
   cd tokoonline-api
   
2. **Instal Dependensi Jalankan perintah di bawah ini untuk mengunduh semua paket yang diperlukan.**
   ```bash
    go mod tidy
  
3. **Atur Database MySQL**
   - Buat database baru di MySQL, contoh: hr_ptzenmultimediaindonesia.
   - Buka file database.go dan sesuaikan konfigurasi koneksi database:

4. **Migrasi Database** Gunakan skrip SQL berikut untuk membuat tabel yang diperlukan dalam database:
     ```bash
     CREATE TABLE users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       email VARCHAR(255) NOT NULL UNIQUE,
       password VARCHAR(255) NOT NULL
     );
     
     CREATE TABLE products (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       price DECIMAL(10, 2) NOT NULL,
       description TEXT NOT NULL
     );
     
     CREATE TABLE orders (
       id INT AUTO_INCREMENT PRIMARY KEY,
       user_id INT NOT NULL,
       product_id INT NOT NULL,
       qty INT NOT NULL,
       status ENUM('pending', 'processed', 'completed') DEFAULT 'pending',
       FOREIGN KEY (user_id) REFERENCES users(id),
       FOREIGN KEY (product_id) REFERENCES products(id)
     );

5. **Jalankan Aplikasi** Mulai server menggunakan perintah berikut:
   ```bash
      go run main.go
   
6. Server akan berjalan pada http://localhost:8000.

## Endpoint Utama

### Autentikasi
   - Register: POST /register
   - Login: POST /login
   - Profile: GET /profile (dengan JWT)
   - Reset Password: POST /reset-password
   - Logout: POST /logout
### Manajemen Produk
   - Tambah Produk: POST /products (dengan JWT)
   - Lihat Produk: GET /products
   - Update Produk: PUT /products/{id} (dengan JWT)
   - Hapus Produk: DELETE /products/{id} (dengan JWT)
### Manajemen Pesanan
   - Tambah Pesanan: POST /orders (dengan JWT)
   - Lihat Pesanan: GET /orders (dengan JWT)
   - Update Status Pesanan: PUT /orders/{id} (dengan JWT)

## Pengujian
Gunakan Postman atau alat serupa untuk menguji endpoint. Pastikan Anda menambahkan header Authorization: Bearer {jwt_token} untuk endpoint yang memerlukan autentikasi.

## Kontribusi
Jika Anda ingin berkontribusi, silakan fork repository ini dan kirim pull request. Semua masukan sangat dihargai.

## Lisensi
Proyek ini menggunakan lisensi MIT. Silakan lihat file LICENSE untuk informasi lebih lanjut.

## Note
Proyek ini hasil rekruitmen di posisi Backend di PT Zen Multimedia Indonesia
