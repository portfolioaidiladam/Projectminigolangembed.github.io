// Package main_test berisi unit test untuk demonstrasi penggunaan fitur embed di Go
package main_test

import (
	"embed"     // Package untuk mendukung embedding file statis
	_ "embed"   // Import blank untuk mengaktifkan directive go:embed
	"fmt"       // Package untuk formatting dan printing
	"io/fs"     // Package untuk operasi filesystem
	"io/ioutil" // Package untuk utilitas I/O
	"testing"   // Package untuk unit testing
)

// version menyimpan isi dari file version.txt
//go:embed version.txt
var version string

// version2 menyimpan isi yang sama dari file version.txt (contoh multiple embedding)
//go:embed version.txt
var version2 string

// TestString menguji embedding file teks sebagai string
func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

// logo menyimpan konten binary dari file logo.png
//go:embed logo.png
var logo []byte

// TestByte menguji embedding dan penulisan file binary
func TestByte(t *testing.T) {
	// Menulis konten logo ke file baru dengan permission penuh
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// files adalah virtual filesystem yang berisi multiple file
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

// TestMultipleFiles menguji pembacaan multiple file yang di-embed
func TestMultipleFiles(t *testing.T) {
	// Membaca dan menampilkan isi file a.txt
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	// Membaca dan menampilkan isi file b.txt
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	// Membaca dan menampilkan isi file c.txt
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

// path adalah virtual filesystem yang menggunakan pattern matching
//go:embed files/*.txt
var path embed.FS

// TestPathMatcher menguji pembacaan file menggunakan pattern matching
func TestPathMatcher(t *testing.T) {
	// Membaca semua entri dalam direktori files
	dirEntries, _ := path.ReadDir("files")
	// Iterasi setiap entri
	for _, entry := range dirEntries {
		// Cek apakah entri bukan direktori
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			// Membaca dan menampilkan isi file
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
