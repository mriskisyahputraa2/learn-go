/* CATATAN

# CLOSURE FUNCTION
	- CLOSURE ADALAH SEBUAH FUNCTION YANG BERADA DIDALAM FUNCTION /
	(FUNCTION DALAM FUNCTION)
	- CLOSURE ADALAH KOMBINASI DARI FUNCTION YANG DI "BUDLE" BERSAMA DENGAN
	"SCOPE" DISEKITARNYA
	- DI DALAM CLOSURE KITA BISA MENGAKSES SCOPE YANG "BERADA DI LUAR FUNCTION"
	DARI DALAM FUNCTION ITU SENDIRI, KARENA MERUPAKAN SATU SCOPE YANG SAMA

# HATI-HATI SAAT MENGGUNAKAN CLOSURE
	- SAAT MENGGUNAKAN CLOSURE KAMU HARUS BERHATI-HATI AGAR TERHINDAR DARI
	KETIKASENGAJAAN "MERUBAH DATA DI LUAR FUNCTION"



*/

package main

import "fmt"

// contoh penulisan closure function, menulis function didalam function
func main() {
	// ini yang dinamakan "closure", mengakses variabel diluar function, lalu digunakan
	name := "Muhammad Rizki Syahputra"

	// function anonymouse
	displayName := func() {
		// Anggap saja tdk sengaja merubah isi variabel name
		// name = "Muhammad Rizki"
		fmt.Println(name)
	}

	displayName()
}
