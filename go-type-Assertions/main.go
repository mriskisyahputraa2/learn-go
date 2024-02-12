/* CATATAN

# TYPE ASSERTIONS
	- Type Assertion di golang adalah cara untuk merubah tipe data tertentu
	menjadi tipe data yang kita inginkan
	- Type Assertion dapat digunakan dengan nilai yang mengimplementasikan
	interface kosong

# Menghadle Panic Pad Type Assertion
	- "Panic" pada type assertion terjadi saat tipe data yang kita ubah tidak
	sesuai
	- Jika terjadi panic dan tidak "ter-recover" maka program kita akan terhenti
	- Agar program kita tidak panic atau terhenti kita bisa menggunakan recover
	atau "menggunakan switch expression"

# Fungsi recover() dalam bahasa Go digunakan untuk menangkap panic yang terjadi dalam program.

*/

package main

import "fmt"

func main() {
	var data interface{} = "Hello World" // membuat v-data dengan tipe data interface kosong dan memasukkan nilai string

	// ini adalah format melakukan type assertion
	dataString := data.(string) // .(string) harus sesuai dengan isi data dari v-data
	fmt.Println(dataString)
	fmt.Printf("Value: %v dan tipe data: %T\n", dataString, dataString)
	fmt.Println("\n")

	// contoh program menghandle panic type assertion
	var MyData interface{} = 123 // mendefinisikan MyData dengan tipe interface

	// melakukan type assertion
	num, ok := MyData.(int) // num = untuk menyimpan nilai data MyData
	// ok = untuk menentukan nilai true dan false

	// jika ok tidak ada atau tidak sesuai
	if !ok {
		recoverMyData := recover()                               // menangkap panic yang terjadi, ini adalah fungsi recover()
		fmt.Println("Tipe data MyData bukan int", recoverMyData) // tampilkan pesan
	} else {
		// jika berhasil, tampilkan pesan ini beserta nilai num nya
		fmt.Println("Tipe data MyData adalah int", num)
	}
	fmt.Println("\n")

	// contoh program type assertions menggunakan swicth
	var data3 interface{} = 32

	switch data3.(type) {
	case string:
		fmt.Println("data3 adalah sebuah string", data3)
	case int:
		fmt.Println("data3 adalah sebuah int", data3)
	default:
		fmt.Println("data3 adalah sebuah tipe data lainnnya", data3)

	}

}
