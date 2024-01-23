/* CATATAN:

DI DALAM BAHASA GOLANG VARIABEL HANYA TERDAPAT DUA YAITU:
1. VAR
2. :=

# PERBEDAAN VARIABEL VAR & :=
1. VAR : - BISA DI GUNAKAN DIDALAM DAN DI LUAR FUNGSI(FUNCTION)
		 - DEKLARASI VARIABEL DAN PENETAPAN VALUE DAPAT DILAKUKAN SECARA TERPISAH TAPI TIDAK BISA DILUAR FUNCTION

2. := : - HANYA DAPAT DIGUNAKAN DI DALAM FUNGSI(FUNCTION)
		- DEKLARASI VARIABEL DAN PENETAPAN VALUE TIDAK DAPAT DILAKUKAN SECARA TERPISAH(HARUS DILAKUKAN PADA BARIS YANG SAMA)


*/

package main

import "fmt"

// menggunakan variabel var bisa dideklarasi di luar fungsi
var x = 4
var y = 4

var hasil = x + y

// ini contoh kasus error jika menggunakan variabel :=, karena tidak boleh diluar fungsi
// y := 4;

func main() {
	var MyName string
	MyName = "Muhammad Rizki Syahputra"
	var nickName = "Rizki"
	myAge := 18

	// jika variabel tidak ada value, maka nilai value nya adalah:
	var a string // output: kosong / tidak ada apa-apa
	var b int    // output: 0
	var c bool   // output: false

	// deklarasi multiple variabel
	var d, e, f, g, text = 3, 4, 5, 6, "semua ini bertipe integer dan string"

	var fist, last, fullName = "Muhammad", "Rizki", "Muhammad Rizki Syahputra"

	fmt.Println(MyName)
	fmt.Println(nickName)
	fmt.Println(myAge, "tahun")

	// output deklarasi variabel di luar fungsi (function)
	fmt.Println("hasil penambahan", x, "dan", y, "adalah =", hasil)
	// fmt.Println(y)

	// output deklarasi tanpa value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// output deklarasi multiple variabel
	fmt.Println(d)    // 3
	fmt.Println(e)    // 4
	fmt.Println(f)    // 5
	fmt.Println(g)    // 6
	fmt.Println(text) // semua ini bertipe integer dan string
	fmt.Println(fist)
	fmt.Println(last)
	fmt.Println(fullName)
}
