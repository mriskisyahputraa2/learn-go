/* CATATAN

#Type Declarations
	- Type declaration adalah bentuk "penulisan ulang dari tipe data yang sudah
	ada menjadi nama baru"
	- Dengan type declaration kita bisa memberikan nama alias dari tipedata
	yang sudah ada, sehingga lebih mudah diingat dan disebut
	(membuat type declaration dengan nama keinginan kita)
*/

package main

import "fmt"

func main() {

	// contoh program type declaration
	type MyString string // membuat type declaration sesuai keiginan sendiri
	type MyInt int
	type Boolean bool

	var txt MyString = "Rizki"
	var num MyInt = 17
	var std Boolean = true

	fmt.Println("Nama aku:", txt)
	fmt.Println("Umur aku:", num)
	fmt.Println("Mahasiswa:", std)

}
