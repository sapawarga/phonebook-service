[![Go Report Card](https://goreportcard.com/badge/github.com/sapawarga/phonebook-service)](https://goreportcard.com/report/github.com/sapawarga/phonebook-service)
[![Maintainability](https://api.codeclimate.com/v1/badges/d620fba429567c496754/maintainability)](https://codeclimate.com/github/sapawarga/phonebook-service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d620fba429567c496754/test_coverage)](https://codeclimate.com/github/sapawarga/phonebook-service/test_coverage)
# phonebook-service
Sapawarga service for "Nomor Penting" feature.

## Quick Setup

1. Download dahulu [bloomRPC](https://appimage.github.io/BloomRPC/) karena aplikasi ini menggunakan protokol GRPC jadi untuk mengetesnya menggunakan bloomRPC. BloomRPC ini seperti Postman. Ikuti saja langkah untuk instalasinya.
2. Siapkan databasenya, sebaiknya minta file `sql.dump` dari database sapawarga stagging yang existing karena masih menggunakan struktur data yang sama
3. Jalankan `go mod tidy` atau `go mod download`
4. Untuk melakukan test dan melihat coveragenya ketikkan command `make test`
5. Buat file `config.json` dan copy dari `config.example.json`
6. Jalankan command `make build`
7. Jalankan command `make run`
8. Clone repository [proto-file](https://github.com/sapawarga/proto-file)
9. Buka bloomRPC dan pilih file protonya dari repository sapawarga/proto-file

## Stack Libraries
1. [Gomock](https://github.com/golang/mock)
2. [Gokit](https://github.com/go-kit/kit)
3. [GRPC](https://grpc.io/docs/languages/go/basics/)


## Package Structure
Pada bahasa pemrograman Golang, satu folder atau *directory* itu dikatakan satu *package* sehingga berikut ini adalah strukur *package* / *directory* dengan menerapkan konsep *clean architecture*

```sh
cmd 
    - database
    - grpc
endpoint
repository
    - mysql
    - postgres
transport
    - grpc
    - http
usecase
    - phonebook
mocks
    - testcases
    - mock_repository
config
model
helper
```
### 1. CMD
Folder `cmd` ini berfungsi sebagai `infrastructure` dari keseluruhan service. Di dalam folder inilah segala jenis inisialisasi baik itu untuk database, aplikasi *thidrdparty*, *routing*, *modul* diinisialisasi di dalam folder ini. Packagenya menggunakan nama `main` yang artinya ini adalah package utama yang akan diakses oleh client saat menggunakan service ini. 

### 2. Endpoint
Folder `endpoint` berfungsi sebagai tempat untuk mengeksekusi terkait request dari client dan response untuk client. Segala bentuk validasi dari client terkait pengisian body request dapat dilakukan di sini sehingga ketika sebelum masuk usecase, body json yang digunakan adlaah body request yang sudah bersih

### 3. Transport
Package ini berfungsi untuk mengatur bagaimana caranya aplikasi/service yang dibuat ini dalam menerima request maupun memberikan response ke client. Di dalam sini dapat diatur apakah requestnya yang diterima dapat dari akses `http`, `grpc`, `tcp` atau protokol lainnya.

### 4. Usecase
Package ini merupakan bagian dimana segala bisnis logic dari service yang dibuat tersimpan. Seharusnya di dalam package ini tidak terdapat logika dari repository maupun pengecekan validasi dari inputan client, hanya ada bisnis logik saja. Usecase ini dapat terdiri dari kondisi konkret maupun unit testnya. Di package ini juga unit test ditaruh untuk memastikan logik yang dijalankan sudah sesuai dengan ekspektasi.

### 5. Mocks
Package ini terdiri atas dua bagian, yaitu package `mock` yang merupakan hasil generate dari `mockgen` dan package `testcases` yang merupakan skenario untuk unit test. Untuk package `mock` ini adalah hasil generate dari interface di repository, untuk mensimulasikan response dari repository aslinya, sedangkan `testcase` itu tempat menampung segala skenario unit test baik itu untuk bagian sukses maupun failed. Pada kasus ini untuk skenarionya saya buat menjadi satu list sehingga ketika menjalankan unit test, tidak perlu mendeklarasikan satu per satu skenarionya, tinggal melooping saja dari skenario yang sudah dibuat.

### 6, Config
Package ini berfungsi sebagai konfigurasi untuk service yang digunakan. Di sini dapat mendeklarasikan variabel-variabl yang diperlukan dalam menjalankan servicenya semisal terkait settingan database yang digunakan, settingan port dari service dan juga settingan untuk akses thirdparty. Selain itu dapat juga digunakan untuk menyimpan `constant` maupun `enum` di sini.

### 7. Model
Package ini berfungsi untuk menyimpan semua permodelan data yang digunakan di dalam keseluruhan package

### 8. Helper
Package ini berfungsi menampung semua fungsi yang bersifat utilitas dan dapat digunakan di seluruh package tanpa harus melakukan inisialisasi terlebih dahulu.



