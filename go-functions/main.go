/* CATATAN

# MULTIPLE RETURN VALUES
	- FUNCTION TIDAK HANYA DAPAT MENGEMBALIKAN SATU NILAI, TAPI JUGA BISA
	MENGEMBALIKAN BEBERAPA NILAI
	- JIKA KITA INGIN FUNGSI MENGEMBALIKAN BEBERAPA NILAI, MAKA KITA PERLU MENULIS
	SEMUA TIPE DATA, DAN JUGA BISA MEMBERI NAMA DARI NILAI YANG DIKEMBALIKAN
*/

package main

import "fmt"

// create function
func myMessage() {
	fmt.Println("Hello, Functions")
}

// function parameter & multiple function
func fullName(full string, age int) {
	fmt.Println("Hello", full, "Syahputra", age, "Tahun")
}

// function return ada banyak cara penulisan
// 1. dengan return memanggil result
func person(fist string, last string) (result string) {
	result = fist + " " + last

	return result
}

// 2. dengan hanya paramter result aja yang di panggil
func tambahAngka() (hasil int) { // untuk return dipisahkan / dibedakan dengan kurung untuk parameter
	a := 2
	b := 53
	hasil = a + b
	return hasil

}

// 3. dengan hanya menuliskan retrun saja tidak usah memanggil result nya
func kaliKali(x int, y int) (result int) {
	result = x * y
	return
}

// multiple function retrun value
func getCombine(i int, j string) (sum int, txt string) {
	sum = i + i
	txt = j + "Syahputra"

	return
}

// atau penulisan gini yang mana int tidak ditampilkan atau sebaliknya liat di output /
func getCombine2(i int, j string) (sum int, txt string) {
	sum = i + i
	txt = j + "Syahputra"

	return
}

func main() {
	myMessage()
	fmt.Println("\n")

	// output function parameters
	fullName("Rizki", 18)
	fullName("Rendy", 19)
	fmt.Println("\n")

	// output function return
	fmt.Println(tambahAngka())
	fmt.Println("\n")

	fmt.Println(kaliKali(100, 7)) // bisa gini pemanggilan fungsinya atau
	total := kaliKali(10, 10)     // bisa juga gini, awalnya disimpan dalam variabel
	fmt.Println(total)

	fmt.Println("\n")

	fmt.Println(person("Rizki", "Syahputra"))
	fmt.Println("\n")

	// output multiple return function
	a, b := getCombine(5, "Rizki")
	fmt.Println(a, b)
	fmt.Println("\n")

	// output nilai integernya tidak ditampilkan hanya string saja
	_, y := getCombine2(10, "Rizki")
	fmt.Println(y)
	fmt.Println("\n")

	// output nilai strinf tidak ditampilkan hanya integernya saja
	z, _ := getCombine2(20, "Rizki")
	fmt.Println(z)
	fmt.Println("\n")

}
