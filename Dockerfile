# Menggunakan image golang sebagai base image
FROM golang:1.19-alpine

# Menentukan direktori kerja di dalam container
WORKDIR /app

# Menyalin go.mod dan go.sum untuk mendownload dependencies
COPY go.mod go.sum ./

# Mendownload dependencies
RUN go mod download

# Menyalin semua file ke direktori kerja di dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Menentukan port yang akan digunakan oleh aplikasi
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./main"]
