# Stage 1: Build
FROM golang:1.20-alpine AS builder

# Install dependencies tambahan jika diperlukan
RUN apk add --no-cache git

# Set working directory dalam container
WORKDIR /app

# Copy semua file dari project ke dalam container
COPY . .

# Install dependency dan build aplikasi
RUN go mod tidy && go build -o main .

# Stage 2: Runtime
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary dari stage build
COPY --from=builder /app/main .

# Expose port yang akan digunakan
EXPOSE 9122

# Jalankan aplikasi
CMD ["./main"]
