/* CATATAN
# NIL
	- "nil" bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya "nil" berarti memiliki "s".
	- Di bahasa pemrograman lain "suatu object yang belum di inisialisasi biasanya
	secara otomatis bernilai null atu nil"
	- Golang memiliki perilaku yang berbeda, saat kita membuat variabel dengan
	tipe data tertentu, variabel tersebut tidak diisi null atau nill melainkan
	berisi default value dari tipe datanya
	- Di golang ada "nil, yaitu sebuah data yang kosong(tidak bernilai)"
	- Nil hanya digunakan di beberapa tipe data seperti "map, slice, function,
	interface, pointer dan channel
"


*/

package main

import (
	"fmt"
)

// contoh program ke-6 // menggunakan function
func newMap(name string) map[string]string {
	if name == "" { // jika name sama dengan(==) kosong("")
		return nil // kembalikan nil
	}

	//jika tidak kosong, kembalikan nilai name dalam bentuk map
	return map[string]string{"name": name}
}

func main() {

	// contoh program ke-1 // slice
	var s []int
	var slice = []int{1, 2, 3, 4}
	fmt.Println("Slice", s)
	fmt.Println("Slice == nil?", s == nil) // true, karna nilai x kosong
	fmt.Println("Slice", slice)
	fmt.Println("Slice == nil?", slice == nil) // false, karna ada nilainya
	fmt.Println("\n")

	// contoh program ke-2 // map
	var m map[string]string
	var Map = map[string]string{
		"name":     "Rizki",
		"activity": "sebagai mahasiswa",
		"hobby":    "coding",
	}
	fmt.Println("map", m)
	fmt.Println("map == nil?", m == nil) // true, karna nilai m kosong
	fmt.Println("map", Map)
	fmt.Println("map == nil?", Map == nil) // false, karna ada nilainya
	fmt.Println("\n")

	// contoh program ke-3 // interface kosong
	var i interface{}
	fmt.Println("interface", i)

	// contoh program ke-4 // pointer
	var ptr *int
	fmt.Println("pointer", ptr)

	// contoh program ke-5 // channel
	var ch chan int
	fmt.Println("channel", ch)

	// output dari program ke-6 // menggunakan function
	user := newMap("Rizki") // pemanggilan parameter "name" dari function newMap

	if user == nil { // jika user sama dengan nil
		fmt.Println("Kosong") // tampilkan kosong
	} else { // jika tidak kosong
		fmt.Println(user) // tampilkan paramater dari name
	}

}
