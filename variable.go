package main

import "fmt"

func main() {
	// Keyword var digunakan untuk deklarasi variable

	// Nilai variable di isi langsung ketika deklarasi
	var firstname string = "john"

	var lastname string
	// nilai variable yang di isi setelah baris kode deklarasi
	lastname = "wick"

	fmt.Printf("Hello %s %s! \n", firstname, lastname)

	main1()
	main2()
	main3()
	main4()
	main5()
}

func main1() {
	// Deklarasi variable menggunakan keyword var
	var lastname string
	// Nilai variable bisa di isi langsung pada saat deklarasi variable
	var firstname string = "jhon"

	// Penggunaan fungsi fmt.Printf()
	/*
		perhatikan bagian "Hello %s %s! \n",
		karakter %s disitu akan diganti dengan data string yang berbeda di parameter ke 2, dan ke 3 dan seterusnya

	*/
	fmt.Printf("Hello %s %s! \n", firstname, lastname)

	lastname = "wick"
	// Ketiga baris kode berikut ini akan menghasilkan output yang sama, meskipun cara penulisan nya berbeda.
	fmt.Printf("Hello jhon wick! \n")
	fmt.Printf("Hello %s %s! \n", firstname, lastname)
	fmt.Println("Hello", firstname, lastname+"!")
	// Tanda plus (+) jika ditempatkan di antara string, fungsinya adalah untuk penggabungan string atau string concatenation.
	/*
		Fungsi fmt.Printf() tidak menghasilkan baris baru di akhir text, oleh karena itu digunakanlah literal newline yaitu \n, untuk memunculkan baris baru di akhir. Hal ini sangat berbeda
		jika dibandingkan dengan fungsi fmt.Println() yang secara otomatis menghasilkan new line (baris baru) di akhir
	*/

}

func main2() {
	// Deklarasi variable tanpa tipe data
	/*
		Selain manifest typing,
		Go juga mengadopsi konsep type inference, yaitu metode deklarasi variabel yang tipe data-nya ditentukan oleh tipe data nilainya, cara kontradiktif jika dibandingkan dengan cara pertama.
		Dengan metode jenis ini, keyword var dan tipe data tidak perlu ditulis.
	*/
	var firstname string = "jhon"

	// Variable lastname di deklarasikan dengan menggunakan metode type interface
	// tipe data tidak di tulis
	// menggunakan operand (:=)
	// keyword var di hilangkan
	lastname := "wick"

	fmt.Printf("Hello %s %s! \n", firstname, lastname)

	// Menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstnameNotDataType = "jhon"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastnameVariableTypeInterface := "wick"

	// Tanda := hanya digunakan sekali di awal pada saat deklarasi.
	// Untuk assignment nilai selanjutnya harus menggunakan tanda =, contoh:
	lastnameVariableTypeInterface = "Ethan"
	lastnameVariableTypeInterface = "bourne"

	fmt.Printf("Hello %s %s! \n", firstnameNotDataType, lastnameVariableTypeInterface)
}

func main3() {
	// Deklarasi multi variable

	/*
		Go mendukung metode deklarasi banyak variabel secara bersamaan, caranya dengan menuliskan variabel-variabel-nya dengan pembatas tanda koma (,).
		Untuk pengisian nilainya-pun diperbolehkan secara bersamaan.
	*/
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Pengisian nilai juga bisa dilakukan bersamaan pada saat deklarasi.
	var fourth, fifth, sixth string = "Empat", "Lima", "Enam"
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth)

	// Lebih ringkas menggunakan type interface
	seventh, eighth, ninth := "Tujuh", "Delapan", "Sembilan"
	fmt.Println(seventh)
	fmt.Println(eighth)
	fmt.Println(ninth)

	// Dengan menggunakan teknik type inference,
	// deklarasi multi variabel bisa dilakukan untuk variabel-variabel yang tipe data satu sama lainnya berbeda.
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "Hello"
	fmt.Println(one)
	fmt.Println(isFriday)
	fmt.Println(twoPointTwo)
	fmt.Println(say)
}

func main4() {
	// Variable undescore (_)
	// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
	// Bisa dibilang variabel ini merupakan keranjang sampah

	_ = "Belajar Golang"
	_ = "Golang is easy"
	name, _ := "jhon", "wick"
	fmt.Println(name)
}

func main5() {
	// Deklarasi variable menggunakan keyword new
	// Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu.
	// Nilai data default-nya akan menyesuaikan tipe datanya.

	name := new(string)
	fmt.Println(name)  // 0xc0000483b0
	fmt.Println(*name) // ""
}

func main6() {
	// Deklarasi variable menggunakan keyword make

	/*
		Keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:

		channel
		slice
		map
	*/
}
