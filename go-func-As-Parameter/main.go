/* CATATAN

# Function As Paramater
	- Sebelumnnya kita sudh bahas di golang function bisa kita simpan sebagai value dari sebuah variabel
	- Kita juga bisa gunakan function sebagai parameter dari function lainnya

# Tyoe Declaration Untuk Alias Function
	- Jika function yang kita buat sudah terlalu panajng, maka akan membuat
	kode kita terlihat kurang rapi
	- Untuk mengatasi hal itu kita bisa membuat type declaration unutk membuat
	alias suatu function
	- Type declaration adalah kemampuan untuk membuat ulang tipe data baru dari tipe data yang  ada, kita juga bisa gunakan untuk membuat suatu alias pada function




*/

package main

import (
	"fmt"
	"strings" // Package untuk memanipulasi string.
)

/* penjelasan kode

 */

// contoh program function type declaration / agar meringkas atau merapikan function untuk penggunaan function as paramter
type Formatter func(string) string

// membuat function formatCase, dengan menerima dua parameter
// 1. txt string: untuk meng-format yang akan diformat yaitu: "hello world" bertipe string
// 2. formatter: Fungsi lain yang akan digunakan untuk memformat string txt. Parameter ini adalah fungsi itu sendiri, menunjukkan kemampuan fungsi sebagai nilai (function as values).
func formatCase(txt string, formatter Formatter) {
	// Memanggil fungsi "formatter" yang diteruskan sebagai argumen dan menyimpan hasil pemformatan dalam variabel "formatted".
	formatted := formatter(txt)
	fmt.Println(formatted)
}

// func upperCase(txt string) string {: Mendefinasikan fungsi bernama upperCase yang mengubah string menjadi huruf besar.
// penulisan function ini disebut dengan: "Function As Paramater"
func upperCase(txt string) string {

	// return strings.ToUpper(txt): Menggunakan fungsi strings.ToUpper dari package strings untuk mengubah string menjadi huruf besar dan mengembalikan hasilnya.
	return strings.ToUpper(txt)
}

func lowerCase(txt string) string {
	return strings.ToLower(txt)
}

func main() {
	// formatCase("hello world", upperCase): Memanggil fungsi formatCase dengan dua argumen:
	//1. "hello world": String yang akan diformat.
	//2. upperCase: Fungsi upperCase yang akan digunakan untuk memformat string.
	formatCase("Hello", upperCase)

	// atau gini pun bisa menampilkannya, ini disebut "Function as value"
	ToUpper := upperCase
	fmt.Println(ToUpper("riski"))

	// Output Function Type Declarations
	formatCase("apa kabar?", upperCase)
	formatCase("BAIK-BAIK SAJA", lowerCase)

}
