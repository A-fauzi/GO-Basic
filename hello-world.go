/*
	Setiap file program harus memiliki package, Setiap project harus ada minimal satu file dengan nama package main.
	File yang berpackage main, akan di eksekusi pertama kali ketika program pertama kali dijalankan
*/
package main

/*
	Keyword import digunakan untuk meng-import atau memasukan package lain kedalam file program, agar isi dari package yang di import bisa di manfaatkan.
*/
// Package fmt merupakan salah satu package bawaan yang di sediakan oleh Go, isinya banyak fungsi untuk keperluan I/O yang berhubungan dengan text
import "fmt"

// Dalam sebuah proyek harus ada file program yang di dalamnya berisi sebuah fungsi bernama main().
// Fungsi tersebut harus berada di file yang package-nya bernama main.
// Fungsi main() adalah yang dipanggil pertama kali pada saat eksekusi program.
func main() {
	// Fungsi fmt.Println() digunakan untuk memunculkan text ke layar (pada konteks ini, terminal atau CMD). 
	// Di program pertama yang telah kita buat, fungsi ini memunculkan tulisan Hello world.
	fmt.Println("Hello World")
}