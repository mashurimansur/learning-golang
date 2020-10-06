package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/mnt/c/Users/mashu/Documents/WebDeveloper/GoLinux/src/github.com/mashurimansur/learning-golang/text.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}

func createFile() {
	//deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	//buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile() {
	//buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//tulis data ke file
	_, err = file.WriteString("halo \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang \n")
	if isError(err) {
		return
	}

	//simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				return
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dihapus")
}

func main() {
	//createFile()
	//writeFile()
	//readFile()
	deleteFile()
}
