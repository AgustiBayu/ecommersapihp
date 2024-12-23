# Ecommers API Hp
Rancang bangun ecommers sederhana untuk penjualan handphone adalah backend API yang dirancang untuk mengelola data penjualan handphone, termasuk menu, pesanan, detail pesanan, keranjang dan produk. Aplikasi ini mendukung operasi CRUD.

![GitHub Logo](https://cdn.prod.website-files.com/6100d0111a4ed76bc1b9fd54/62217e885f52b860da9f00cc_Apa%20Itu%20Golang%3F%20Apa%20Saja%20Fungsi%20Dan%20Keunggulannya%20-%20Binar%20Academy.jpeg)

## Fitur Utama
- **Auth Basic:** Implementasi Auth Basic ketika signin dan memberikan otorisasi pada setiap API   
- **Produk:** Terdapat proses implementasi CRUD pada produk dengan otorisasi dari basic auth.
- **Manajemen Pesanan:** Terdapat proses implementasi CRUD pada pesanan dengan otorisasi dari basic auth.
- **Manajemen Detail Pesanan:** Terdapat proses implementasi CRUD pada detail pesanan dengan otorisasi dari basic auth.
- **Keranjang:** Terdapat proses implementasi CRUD pada keranjang dengan otorisasi dari basic auth.

## Teknologi
- **Bahasa:** Golang
- **Golang Httprouter:**
    ```bach
    github.com/julienschmidt/httprouter
- **Golang Validate:**
    ```bach
    github.com/go-playground/validator/v10
- **Database:** MySql
    ```bach
    github.com/go-sql-driver/mysql
## Instalasi
1. Clone repository:
   ```bash
   git clone https://github.com/AgustiBayu/ecommersapihp.git
2. cd ecommersapihp
3. go mod tidy
4. Atur konfigurasi database di file app.
5. Jalankan aplikasi:
    ```bash
    go run main.go