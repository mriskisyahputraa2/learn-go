/* CATATAN

# PERULANGAN FOR LOOP ADALAH SALAH-SATUNYA PERULANGAN YANG ADA DI GOLANG
JADI TIDAK ADA TUH PERULANGAN DO WHILE, WHILE


1. for loops
2. nested for loops
3. for range loops
	- For range digunakan untuk mempermudah dalam mengulang array, slice atau map
	- For range akan mengembalikan index dan nilainya atau key dan nilainya
	- syntax for range:
		for index, value := range array|slice|map {
			// kode yang akan dieksekusi untuk setiap iterasi
		}

*/

package main

import "fmt"

func main() {

	// for loops
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}
	fmt.Println("\n")

	// nested loops / for didalam for
	a := [2]string{"buah", "jus"} // array
	b := [3]string{"jambu", "mangga", "alpukat"}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}
	fmt.Println("\n")

	// For Range
	cars := [3]string{"toyota", "supra", "civic"} // array

	// contoh yang mengembalikan index dan nilainya
	for index, cars := range cars {
		fmt.Println(index, cars)
	}
	fmt.Println("\n")

	// for ini jika tidak mau menampilkan indexnya menggunakan _ [ignore] / sebaliknya hanya ingin menampilkan index
	for _, cars := range cars {
		fmt.Println(cars)
	}
	fmt.Println("\n")

	var person = map[string]string{ // map
		"name":  "Rizki",
		"age":   "18 tahun",
		"hobby": "coding",
	}
	// contoh yang mengembalikan key dan nilainya
	for key, person := range person {
		fmt.Println(key, person)
	}
	fmt.Println("\n")

}
