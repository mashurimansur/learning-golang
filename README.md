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