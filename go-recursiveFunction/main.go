/* CATATAN

# RECURSIVE FUNCTIONS
	- RECURSIVE FUNCTIONS ADALAH FUNCTION YANG MEMANGGIL FUNCTION DIRINYA SENDIRI
	- SUATU FUNGSI BERSIGAT RECURSIVE JIKA MEMANGGIL DIRINYA SENDIRI DAN MENCAPAI
	KONDISI BERHENTI
	- HATI-HATI DALAM MENULISA RECURSIVE FUNCTION KARNA RAWAN INFINIT LOOP, TAPI INI SANGAT BAGUS

*/

package main

import "fmt"

// recursive function
func countTest(x int) int { // type int kedua, artinya mengembalikan return type int
	// if ini adalah kondisi berhenti dimana induk pokoknya ada di return
	if x == 11 { // jika nilai x sama dengan 11
		return 0 // maka kembalikan nilai 0
	}

	fmt.Println(x) // cetak nilai x
	// return ini yang dinamakan recursive function yang artinya function memanggil dirinya sendiri
	return countTest(x + 1) // kembalikan fungsinya sendiri dengan x + 1. (x(1) == 1, x(2) == 11 etc). sampai perulangannya x(11) == 11 maka kembalikan nilai 0
}

// recursive function factorial. ini contoh langsung dpt hasil menggunakan if
func factorialRecursive(z int) int {

	// membuat logika berhenti faktorialnya
	if z == 0 { // jika z == 0. contohnya 5 == 0? no 4 == 0?no sampai 0 == 0
		return 1 // kembalikan nilai 1, untuk memberhatikan nilai hati-hati dengan ini
	}

	// function yang mengembalikan fungsi nya sendiri.
	return z * factorialRecursive(z-1) //kembalikan nilai z * nilai factorialnya yaitu (5) dan dikurangi 1. contoh nya 5-1 = 4 sampai nilainya 0. jadi nanti akan gini 5 x 4 x 3 x 2 x 1 = 120
}

// contoh mencari faktorial menggunakan for
func factorialLoops(n int) int { // nilai (n) adalah 10, bisa berubah-ubah sesaui pembuat
	result := 1 // deklarasi variabel result dengan value 1

	// lakukan perulangan dengan v-i dgn nilai n yaitu (10).
	for i := n; i > 0; i-- { // cek apakah n(10) > 0 jika ya, maka loop dari bawah ke atas karna menggunakan decrement(--)
		result *= i // kembalikan nilai result yaitu 1 dan kali dengan nilai i yaitu 10
	}

	return result // kembalikan hasil dari result
}

func main() {
	// output recursive function
	countTest(1) // dengan fungsinya diawal dari nilai 1 atau bisa juga 2,3,4,5, trsh kamu
	fmt.Println("\n")

	// ouput recursion function factorial menggunakan if
	fmt.Println("hasil dari faktorial 5 adalah =", factorialRecursive(5)) // nilai z nya adalah 5
	fmt.Println("\n")

	// output menggunakan for
	loop := factorialLoops(10) // nilai n adalah 10
	fmt.Println("hasil dari faktorial 10 adalah =", loop)

	// ini saya gunakan untuk memahami kode recursive function dengan for
	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}

}
