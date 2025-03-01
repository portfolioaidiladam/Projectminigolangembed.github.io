// Package main adalah package utama yang mengeksekusi program
package main

import (
	"embed" // Package untuk menyematkan file statis ke dalam binary
	"fmt"   // Package untuk fungsi I/O formatting
	"io/fs" // Package untuk operasi filesystem
	"os"    // Mengganti io/ioutil dengan os
)

// version menyimpan isi dari file version.txt
//go:embed version.txt
var version string

// logo menyimpan data binary dari file logo.png
//go:embed logo.png
var logo []byte

// path menyimpan seluruh file .txt yang ada di dalam folder files
//go:embed files/*.txt
var path embed.FS

// main adalah fungsi utama yang akan dieksekusi saat program dijalankan
func main() {
	// Menampilkan isi dari file version.txt
	fmt.Println(version)

	// Menulis ulang file logo ke logo_new.png menggunakan os.WriteFile
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err) // Menghentikan program jika terjadi error
	}

	// Membaca semua entri dalam direktori "files"
	dirEntries, _ := path.ReadDir("files")
	// Iterasi setiap entri dalam direktori
	for _, entry := range dirEntries {
		// Memeriksa apakah entri bukan direktori
		if !entry.IsDir() {
			// Menampilkan nama file
			fmt.Println(entry.Name())
			// Membaca dan menampilkan isi file
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
