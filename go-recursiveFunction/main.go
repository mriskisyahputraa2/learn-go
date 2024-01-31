/* CATATAN

# RECURSIVE FUNCTIONS
	- RECURSIVE FUNCTIONS ADALAH FUNCTION YANG MEMANGGIL FUNCTION DIRINYA SENDIRI
	- SUATU FUNGSI BERSIGAT RECURSIVE JIKA MEMANGGIL DIRINYA SENDIRI DAN MENCAPAI
	KONDISI BERHENTI
	- HATI-HATI DALAM MENULISA RECURSIVE FUNCTION KARNA RAWAN INFINIT LOOP, TAPI INI SANGAT BAGUS

*/

package main

import "fmt"

func countTest(x int) int {
	// if ini adalah kondisi berhenti dimana induk pokoknya ada di return
	if x == 11 { // jika nilai x sama dengan 11
		return 0 // maka kembalikan nilai 0
	}

	fmt.Println(x) // cetak nilai x
	// return ini yang dinamakan recursive function yang artinya function memanggil dirinya sendiri
	return countTest(x + 1) // kembalikan fungsinya sendiri dengan x + 1. (x(1) == 1, x(2) == 11 etc). sampai perulangannya x(11) == 11 maka kembalikan nilai 0
}

func main() {
	countTest(1) // dengan fungsinya diawal dari nilai 1 atau bisa juga 2,3,4,5, trsh kamu
}
