/* CATATAN

ANONYMOUS FUNCTION
	- DI MATERI-MATERI SEBELUMNYA KETIKA MEMBUAT FUNCTION KITA SELALU MEMBERIKAN
	NAMA FUNCTION PADA FUNCTION TERSERBUT
	- SEBENARNYA KITA BISA MEMBUAT FUNGSI TANP NAMA, ATAU YANG BIASA DISEBUT
	ANONYMOUS FUNCTION
	- KITA AKAN MENETAPKAN ANONYMOUS FUNCTION KE SEBUAH VARIABEL DAN
	MENGGUNAKAN NAMA VARIABEL TERSEBUT UNTUK MEMANGGIL ANONYMOUS FUNCTION
*/

package main

import "fmt"

func main() {

	// contoh program anonymous function
	var getHello = func() {
		fmt.Println("Hello anonymous function")
	}
	getHello()

	viewNumber := func() {
		fmt.Println(1, 2, 3, 4, 5)
	}
	viewNumber()

	// program menggunakan parameter
	person := func(fist string, age int) {
		fmt.Println("Halo nama saya", fist, "umur saya", age, "tahun")
	}
	person("Riski", 18)

	// program menggunakan return valur
	addNumber := func(a int, b int) int { // int kosong ini digunakan untuk return value
		result := a + b
		return result
	}
	total := addNumber(18, 10)
	fmt.Println(total)

}
