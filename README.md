mytools adalah Command Line Tools(CLI) yang berfungsi untuk mengonversi file .log yang ada di folder linux /var/log menjadi plaintext atau file JSON.

# Overview

mytools ini dibuat menggunakan bahasa pemrograman Golang dan menggunakan [Cobra library](https://github.com/spf13/cobra) sebagai library untuk membuat Aplikasi CLI.

## Installation

Untuk menggunakan mytools, silahkan ikuti langkah-langkah berikut:

1. Lakukan clone pada repository ini, gunakan perintah `git clone https://github.com/AldiZulfikar/mytools.git`
2. Run `go mod tidy` dan pastikan kamu sudah menginstall go
3. Install aplikasi mytools dengan perintah `go install mytools`
4. Gunakan `mytools -h` sebagai petunjuk dalam menggunakan mytools

### 1. Mengonversi file log menjadi PlainText atau JSON

Jalankan perintah ini untuk mengonversi file .log menjadi Plaintext atau JSON: `mytools [path of log files] [flag of type] [type convert]`

1. Misal untuk mengonversi file .log menjadi Plaintext: `mytools /var/log/nginx/error.log -t text`
2. Misal untuk mengonversi file .log menjadi Json: `mytools /var/log/nginx/error.log -t json`
3. Misal jika tidak memasukan [flags of type], maka secara default akan mengonversi menjadi Plaintext: `mytools /var/log/nginx/error.log`

### 2. Menyimpan file .log yang telah dikonversi

kamu juga bisa memilih dimana akan menyimpan file output dari konversi tersebut
Jalankan perintah ini untuk menyimpan file hasil konversi ke dalam direktori yang kamu inginkan: `mytools [path of log files] [flag of type] [type convert] [flag of output] [output destination]`

1. Misal untuk mengonversi file .log menjadi Plaintext: `mytools /var/log/nginx/error.log -t text -o /User/johnmayer/Desktop/nginxlog.txt`
2. Misal untuk mengonversi file .log menjadi Json: `mytools /var/log/nginx/error.log -t json -o /User/johnmayer/Desktop/nginxlog.json`
3. Misal jika tidak memasukan [flags of type], maka secara default akan mengonversi menjadi Plaintext: `mytools /var/log/nginx/error.log -o /User/johnmayer/Desktop/nginxlog.txt`
