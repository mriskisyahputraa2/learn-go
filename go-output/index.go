/* CATATAN
ADA 3 CARA DALAM MENAMPILKAN OUTPUT PADA PEMROGRAMAN GOLANG, MASING-MASING MEMILIKI FUNGSI YANG BERBEDA, YAITU:

1. Print(): INI ADALAH DEFAULT DARI FORMAT OUTPUT GO, DIMANA OUTPUT YANG DI TAMPILKAN TIDAK ADA SPASI DAN ENTER YANG TERCETAK JADI HANYA SATU BARIS

2. Println() : SEDANGKAN INI OUTPUT YANG DI TAMPILKAN AKAN MENAMBAH SPASI DAN ENTER (MEMBUAT BARIS BARU)

3. Printf(): fungsi Printf memformat argument berdasarkan format 'verb' yang diberikan, lalu kemudian mencetaknnya. terdapat empat format yaitu:
1. %v : mencetak `value` argumen default // yang umum digunakan
2. %T : mencetak `tipe data` value // yang umum digunakan
3. %#v : mencetak value dalam format Go-Syntax
4. %% : mencetak tanda %
5. %d : mencetak nilai int

# \n: MEMBUAT BARIS BARU
# " ": MENAMBAH SPASI

*/

package main

import "fmt"

// Printf (format %v & %T)
var i = "Hello"
var j = 12

// Printf (format %v%% & %#v)
var number = 15.211
var text = "Hello World!"

func main() {
	// Print()
	var firstName, LastName = "Muhammad Rizki", "Syahputra"

	// Println()
	var name, age = "rizki", 18

	// Output Print
	fmt.Print(firstName, " ")
	fmt.Print(LastName, "\n")

	//  Output Println
	fmt.Println(name)
	fmt.Println(age)

	// Output Printf dengan (format %v & %T)
	fmt.Printf("Value-nya adalah: '%v' dan tipe datanya: '%T'\n", i, i)
	fmt.Printf("Value-nya adalah: '%v' dan tipe datanya: '%T'\n", j, j)

	// Output Printf dengan (format %v%% & %#v)
	fmt.Printf("%T%%\n", number)
	fmt.Printf("%#v\n", text)
}
