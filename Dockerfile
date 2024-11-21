# Gunakan base image untuk Go
FROM golang:1.20-alpine

# Set working directory dalam container
WORKDIR /app

# Copy semua file dari project ke dalam container
COPY . .

# Install dependency dan build aplikasi
RUN go mod tidy
RUN go build -o main .

# Expose port yang akan digunakan
EXPOSE 9122

# Jalankan aplikasi
CMD ["./main"]
