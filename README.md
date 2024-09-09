<h1 align="center" >Hello 👋, I'm Alfredo</h1>

<h1 align="center" >Welcome To CRUD GraphQL Golang</h1>

> Program ini digunakan sebagai API untuk manajemen users menggunakan GraphQL dan Go.

## Requirement

- [Go](https://go.dev/) v1.22.6+

## How To Run ?🤔

```bash
# Clone this repositories
git clone https://github.com/dwiptrdo/CRUD_GraphGO.git

# go into the directory
cd CRUD_GraphGO

```

Buat file .env di root proyek untuk konfigurasi database. Contoh isinya:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=1234
```

Untuk menjalankan Programnya kamu hanya perlu menjalankan dengan command

```bash
go run main.go
```

## 🚀Structure

```
📦CRUD_GraphGO
 ┣ 📂controller   # Berisi logika untuk menangani request dan response
 ┃ ┗ 📜controller.go
 ┣ 📂database     # Koneksi ke database dan konfigurasi terkait
 ┃ ┗ 📜database.go
 ┣ 📂model        # Definisi struktur data atau model aplikasi
 ┃ ┗ 📜model.go
 ┣ 📂schema       # Definisi skema GraphQL
 ┃ ┣ 📜schema.go
 ┃ ┗ 📜user.go
 ┣ 📂utils        # Fungsi-fungsi utilitas
 ┃ ┗ 📜util.go
 ┣ 📜go.mod       # File module Go
 ┣ 📜go.sum       # File dependencies Go
 ┗ 📜main.go      # Entry point aplikasi

```

## Author

👤 **dwiptrdo**


- Github: [@dwiptrdo](https://github.com/dwiptrdo)
- Instagram: [@_dwiptrdo](https://www.instagram.com/_dwiptrdo/)
