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
