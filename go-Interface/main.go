/* CATATAN

# INTERFACE
	- Interface di golang adalah sebuah tipe data yang tidak memiliki
	implementasi secara lengsung(abstrak) ini bukan interface UI
	- Sebuah interface berisi definisi-definisi function / method
	- Interface biasanya digunakan sebagai "kontrak" contoh nya seperti,
	struct
*/

package main

import "fmt"

// membuat interface Gretting yang berfungsi sebagai kontrak atau blueprint untuk tipe data yang memiliki metode Greet().
type Gretting interface {
	Greet() string // Greet() adalah "Method" yang harus mengembalikan nilai string.
}

// membuat struct
type Person struct {
	Name string // membuat field Name bertipe string

}

// membuat struct method dengan v-p menerima params Person dengan method Greet() dan mengembalikan nilai string
func (p Person) Greet() string {
	return p.Name // kembalikan value v-p dengan field Name
}

// membuat function "sayHello" dengan v-g dan menerima params Greeting
func sayHello(g Gretting) {
	fmt.Println("Hello my name is", g.Greet()) // tampilkan "Hello my name is"  dari v-g yang menerima method Greet()
}

// contoh program ke-2
// Interface GetStudent untuk menentukan kontrak metode GetStudent()
type GetStudent interface {
	GetStudent() string
}

// Struct Students untuk menampung data student
type Students struct {
	Name  string
	Nim   int
	Prodi string
}

// Metode GetStudent() pada struct Students untuk memenuhi kontrak interface GetStudent
func (s Students) GetStudent() string {
	return fmt.Sprintf("Hello nama saya %s, Nim saya %d, Prodi saya %s", s.Name, s.Nim, s.Prodi)
}

// Fungsi viewStudent untuk mencetak informasi student
func viewStudent(a GetStudent) {
	fmt.Println(a.GetStudent())
}
func main() {
	person := Person{Name: "Rizki"} // tambahkan value Name yang disimpan dalam v-person
	sayHello(person)
	fmt.Println("\n")

	std := Students{
		Name:  "Syahputra",
		Nim:   2023903430029,
		Prodi: "TRKJ",
	}
	viewStudent(std)

}
