/* CATATAN

# Function As Value
	- Di golang function juga merupakan sebuah tipe data dan bisa kita anggap sebagai sebuah value
	- Sehingga kita dapat menyimpan function ke dalam sebuah variabel


*/

package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	total := sum // Baris ini mengatribusikan fungsi sum ke variabel total. Ini berarti total sekarang menyimpan referensi ke fungsi sum.
	fmt.Print(total(1, 3))

}
