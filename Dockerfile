# Menggunakan versi Go yang sesuai
FROM golang:1.23 AS builder

# Set direktori kerja
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependencies
RUN go mod download

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
