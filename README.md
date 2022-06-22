### Go Command List

> Daftar dibawah ini adalah beberapa perintah CLI Bahasa Go

```
 go mod init
 go run
 go test
 go build
 go get
 go mod tidy
 go mod vendor
```

> command **go mod init**
>
> Digunakan untuk inisialisasi project pada Go (Menggunakan Go Modules). Untuk nama project bisa menggunakan nama apapun, tapi umum nya adalah disamakan dengan nama directori

> command **go run**
>
> digunakan untuk eksekusi file program (file ber-ektensi _.go_). Cara penggunaan nya dengan menuliskan command tersebut diikuti argumen nama file, contoh:
>
> > _go run main.go_
>
> command **go run** hanya bisa digunakan pada file yang didalam nya terdapat _package main_

> command **go test\***
>
> Go menyediakan package testing, berguna untuk keperluan _unit test._ File yang akan di test harus memiliki akhiran _\_test.go_. berikut contoh:
>
> > _go test main_test.go_

> command **go build**
>
> Command ini digunakan untuk mengkompilasi file program.
>
> Sebenarnya ketika eksekusi program menggunakan _go run_, terjadi proses kompilasi juga. File hasil kompilasi akan disimpan pada folder temporary untuk selanjutnya langsung dieksekusi.
>
> Berbeda dengan _go build_, command ini menghasilkan file executable atau binary pada folder yang sedang aktif.
>
> Untuk nama executable sendiri bisa diubah menggunakan flag -o. Contoh:
>
> > _go build -o program.exe_

> command **go get**
>
> Command _go get_ digunakan untuk men-download package. Sebagai contoh ingin men-download package Kafka driver untuk Go pada projek.
>
> > _go get github.com/segmentio/kafka-go_
>
> [github.com/segmentio/kafka-go](github.com/segmentio/kafka-go) adalah URL package kafka-go. Package yang sudah terunduh tersimpan dalam temporary folder yang ter-link dengan project folder di mana command go get dieksekusi, menjadikan projek tersebut bisa meng-import package terunduh. Untuk mengunduh dependensi versi terbaru, gunakan _flag -u_ pada command _go get_, misalnya:
>
> > _go get -u github.com/segmentio/kafka-go_
>
> Command go get harus dijalankan dalam folder project.
>
> Command **go mod tidy**

> command **go mod tidy**
> Command _go mod tidy_ digunakan untuk memvalidasi dependensi. Jika ada dependensi yang belum ter-download, maka akan otomatis di-download.

> command **go mpd vendor**
>
> Command ini digunakan untuk vendoring.
