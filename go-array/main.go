/* CATATAN

# INISIALISAI HANYA SATU SPESIFIC ELEMENT
CONTOH CODE DIBAWAH:
1. DIMANA TERDAPAT ARRAY DENGAN PANJANG 5 ELEMENT
2. 1:10 berarti, menetapkan 10 ke indeks 1
3. 2:40 berarti, menetapkan 40 ke indeks 2, dsb,

# ARRAY UNTUK MENDAPATKAN PANJANG PADA SEBUAH ARRAY
MENGGUNAKAN FUNGSI:
len() : fungsi nya digunakan untuk mengetahui panjang dari sebuah array

*/

package main

import "fmt"

func main() {

	// array dengan panjang yang sudah di definisikan
	var number = [5]int{1, 2, 3, 4, 5}
	text := [3]string{"Muhammad", "Rizki", "Syahputra"}

	// array dengan panjang yang disimpulkan atau dengan panjang yang menyesuaikan
	var a = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := [...]string{"a", "b", "c", "d", "e"}

	// mengubah elemen dari array
	var c = [3]int{10, 20, 30}
	c[1] = 40

	// array jika belum di deklarasi yang tipe int maka defaultnya akan 0, jika string defaultnya akan kosong
	var arr1 = [3]string{}     // not initialized
	var arr2 = [3]int{1, 2}    // partially initialized
	var arr3 = [3]int{1, 2, 3} // fully initialized

	// array yang hanya deklarasi element sesuai index yang ditentukan disini index 1 dan 2
	var ar1 = [5]int{1: 20, 2: 40}
	ar2 := [10]string{0: "Muhammad", 7: "Rizki", 9: "Syahputra"}

	// array untuk mendapatkan panjang dari sebuah array
	name := [5]string{"Rizki", "Yanto", "Ahmad", "Bambang", "Rendy"}
	age := [...]int{18, 22, 10, 45, 33, 55, 32, 13, 53}

	// ouput array sudah ada panjang nya
	fmt.Println("data =", number)
	fmt.Println("data =", text[2])

	// output array dengan panjnag yang menyesuaikan
	fmt.Println("data =", a)
	fmt.Println("data =", b)

	// output mengubah elemen dari array
	fmt.Println("data =", c)

	// output array yang belum sepenuhnya dideklarasi
	fmt.Println("data =", arr1)
	fmt.Println("data =", arr2)
	fmt.Println("data =", arr3)

	// output array hanya inisialisasi sesuai index array
	fmt.Println("data =", ar1)
	fmt.Println("data =", ar2)

	// output array untuk mendapatkan panjang dari sebuah array
	fmt.Println("data =", len(name), "lenght")
	fmt.Println("data =", len(age), "lenght")
}
