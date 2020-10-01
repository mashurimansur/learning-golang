# Belajar Golang Dasar
[Slide Materi by Programmer Zaman Now](https://docs.google.com/presentation/d/1QNFV9kjV4TfN-FVFLT6-8Urq2MmadAmgc1puk-YE5Fs/edit)

## GOPATH
* gopath adalah environment variable yang berisikan lokasi tempat kita menyimpan project dan library golang
* gopath wajib dibuat ketika kita mulai membuat aplikasi lebih dari satu file atau butuh menggunakan library dari luar

## Package & Import
1. Pakcage
   * Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
   * dengan menggunakan package, kita bisa merapikan kode program yang kita buat
   * package sendiri sebenarnya hanya direktori folder di sistem operasi kita
2. Import
   * secara standart, file golang hanya bisa mengakses file golang lainnya yang berada dalam package yang sama
   * jika kita ingin mengakses file golang yang berada diluar package, maka kita bisa menggunakan import 

## Access Modifier
* di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadapa suatu function atau variable
* di golang, untuk menentukan access modifier, cukup dengan nama function atau variable
* jika namanya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai degan huruf kecil, artinya tidak bisa diakses dari package lain

## Package Initialization
* saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
* ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kitamembuat funciton inisialisasi untuk mebuka koneksi ke database
* untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

## Blank Identifier
* kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
* secara default, golang akan komplen ketika ada package yang diimport namun tidak digunakan
* untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import

## Pakcage Bawan Golang
1. Pakcage os
    * golang telah menyediakan banyak sekali pakcage bawaan, salah satunya adalah package os
    * package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
    * https://golang.org/pkg/os/
2. Package flag
    * packge flag berisikan fungsionalitas untuk meparsing command line argument
    * https://golang.org/pkg/flag/
3. Package strings
    * package strings adalah package yang berisikan function-function untuk memanipulasi tipe data string
    * ada banyak sekali function yang bisa kita gunakan
    * https://golang.org/pkg/strings/
    * beberapa function di package strings

        | Function                               | Description                                      |
        | -----------                            | -----------                                      |
        | strings.Trim(string, cutset)           | memotong cutset di awal dan akhir string         |
        | strings.ToLower(string)                | membuat semua karakter string menjadi lowercase  |
        | strings.ToUpper(string)                | membuat semua karakter string menjadi uppercase  |
        | strings.Split(string, operator)        | membuat string berdasarkan separator             |
        | strings.Contains(string, search)       | mengecek apakah string mengandung string lain    |
        | strings.ReplaceAll(string, from, to)   | mengubah semua string dari from ke to            |
4. Package strconv
    * sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 to int64
    * bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? misal int ke string, atau sebaliknya
    * hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
    * https://golang.org/pkg/strconv/
    * beberapa function di package strconv
    
        | Function                          | Description               |
        | -----------                       | -----------               |
        | strconv.ParseBool(string)         | mengubah string ke bool   |
        | strconv.ParseFloat(string)        | mengubah string ke float  |
        | strconv.ParsInt(string)           | mengubah string ke int64  |
        | strconv.FormatBool(bool)          | mengubah bool ke string   |
        | strconv.FormatFloat(float, ...)   | mengubah float ke string  |
        | strconv.FormatInt(int, ...)       | mengubah int64 ke string  |
5. Package math
    * package math merupakan package yang berisikan constant dan fungsi matematika
    * https://golang.org/pkg/math/
    * beberapa function di package math
    
        | Function                   | Description                                                              |
        | -----------                | -----------                                                              |
        | math.Round(float64)        | membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat |
        | math.Floor(float64)        | mebulatkan float64 kebawah                                               |
        | math.Ceil(float64)         | mebulatkan float64 keatas                                                |
        | math.Max(float64, float64) | mengembalikan nilai float64 paling besar                                 |
        | math.Min(float64, float64) | mengembalikan nilai float64 paling kecil                                 |
6. Package time
    * package time adalah package yang berisikan fungsionalitas untuk management waktu di golang
    * https://golang.org/pkg/time/
    * beberapa function di package time
    
        | Function                   | Description                      |
        | -----------                | -----------                      |
        | time.Now()                 | untuk mendapatkan waktu saat ini |
        | time.Date(...)             | untuk membuat waktu              |
        | time.Parse(layout, string) | untuk memparsing waktu ke string |
7. Package regexp
    * package regexp adalah utilitas di golang untuk melakukan pencarian regular expression
    * regular expression di golang menggunakan library C  yang dibuat Google bernama RE2
    * https://github.com/google/re2/wiki/Syntax
    * http://golang.org/pkg/regexp/
    
        | Function                          | Description                                           |
        | -----------                       | -----------                                           |  
        | regexp.MustCompile(string)        | membuat regexp                                        |
        | regexp.MatchString(string) bool   | mengecek apakah regexp match dengan string            |
        | regexp.FindAllString(string, max) | mencari string yang match degan maximum jumlah hasil  | 
8. Package container/list
    * package container/list adalah implementasi struktur data double linked list di golang
    * https://golang.org/pkg/container/list/
9. Package container/ring
    * package container/ring adalah implementasi struktur data circular list
    * circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
    * https://golang.org/pkg/container/ring/
10. Package sort
    * package sort adalah package yang berisikan utilitas untuk proses pengurutan
    * agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
    * https://golang.org/pkg/sort/
11. Package replect