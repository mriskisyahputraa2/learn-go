/* CATATAN

# Inteface Kosong
	-Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any, merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing.
	- Interface kosong adalah sebuah interface "yang tidak memiliki method
	apapun"
	- Untuk mendefinisikan interface kosong "kita bisa tulis dengan iterface
	dulu diikuti dengan curly bracket kosong: interface{}"
	- Secara sederhana interface kosong bisa kita anggap "sebagai tipe data
	apapun"
	- Interface kosong bisa juga memasukkan data argument
	"apapun itu jenis tipe datanya"

# Type Alias Any
	- "Tipe any" merupakan alias dari "interface{}", keduanya adalah "sama".

# Casting Variabel Interface Kosong Ke Objek Pointer
	- Variabel interface{} bisa menyimpan data apa saja, termasuk data objek, pointer, ataupun gabungan keduanya.

# Kombinasi Slice, map, dan interface{}
	- Tipe []map[string]interface{} adalah salah satu tipe yang paling sering digunakan (menurut saya), karena tipe data tersebut bisa menjadi alternatif tipe slice struct.

*/

package main

import "fmt"

// contoh penulisan ke-1 // menggunakan function
func printType(i interface{}) {
	fmt.Printf("%v, %T\n", i, i) // %v mencetak value, %T mencetak tipe data
}

func main() {

	// output ke-1
	printType("Riski")
	printType(18)
	printType(true)
	fmt.Println("\n")

	// contoh penulisan ke-2 // menggunakan variabel
	var x interface{}

	x = "Muhammad Rizki Syahputra"
	fmt.Println(x)
	x = 18
	fmt.Println(x)
	x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = false
	fmt.Println(x)
	fmt.Println("\n")

	// contoh program ke-3 // menggunakan swicth
	var data interface{}

	data = []string{"halo", "a,"}

	switch data.(type) {
	case string:
		fmt.Println("Data adalah string: ", data)
	case int:
		fmt.Println("Data adalah int: ", data)
	case bool:
		fmt.Println("Data adalah bool: ", data)
	default:
		fmt.Println("Data adalah tipe lain: ", data)
	}
	fmt.Println("\n")

	// contoh program ke-3 // menggunakan map
	var data2 map[string]interface{}

	data2 = map[string]interface{}{
		"name":      "riski",
		"age":       18,
		"breakfast": []string{"rise", "fish", "vegetable"},
	}
	fmt.Println("Nama saya:", data2["name"])
	fmt.Println("Umur saya:", data2["age"])
	fmt.Println("Sarapan pagi saya:", data2["breakfast"])
	fmt.Println("SELESAI DI TYPE INTEFACE ,\n")

	fmt.Println("MASUK DI TYPE ANY")
	// masuk ke any
	// contoh program ke-4 // menggunakan map
	var data3 map[string]any

	data3 = map[string]any{
		"name":  "riski",
		"age":   18,
		"fruit": []string{"apple", "manggo", "banana"},
	}
	fmt.Println("Nama saya:", data3["name"])
	fmt.Println("Umur saya:", data3["age"])
	fmt.Println("Buah kesukaan saya:", data3["fruit"])
	fmt.Println("SELESAI DI TYPE ANY ,\n")

	fmt.Println("masuk ke Casting Variabel Interface Kosong Ke Objek Pointer")
	type person struct { // membuat struct dengan nama person
		name   string
		age    int
		status bool
	}
	// Variabel secret dideklarasikan bertipe interface{} menampung referensi objek cetakan struct person
	var secret interface{} = &person{name: "Rizki", age: 18, status: true}
	var name = secret.(*person).name //Cara casting dari interface{} ke struct pointer adalah dengan menuliskan nama struct-nya dan ditambahkan tanda asterisk (*) di awal, contohnya seperti secret.(*person)
	fmt.Println(name)                // tampilkan output dari name
	fmt.Println("SELESAI Casting Variabel Interface Kosong Ke Objek Pointer\n")

	fmt.Println("masuk ke Kombinasi Slice, map, dan interface{}")

	// [] adalah slice
	var person2 = []map[string]interface{}{ // kombinasi slice dan map
		{"name": "Rizki", "age": 18, "job": "Software Engineer"},
		{"name": "Faris", "age": 19, "job": "Frontend Engineer"},
		{"name": "Rian", "age": 20, "job": "Backend Engineer"},
	}

	// melakukan perulangan dengan index dan value
	//each: Variabel ini menyimpan setiap elemen map selama setiap iterasi.
	//range person: Ini menentukan bahwa loop mengulangi elemen array person.
	for _, each := range person2 { // disini index tidak ditampilkan, hanya valuenya
		fmt.Println(each["name"], each["age"], each["job"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	// ini jika menggunakan index array / map / slice
	for i, each := range fruits {
		fmt.Println(i, each)
	}
}
