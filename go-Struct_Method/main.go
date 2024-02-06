/* CATATAN

# STRUCT METHOD
	- Di golang kit abisa membuat sebbuah struct agar seolah-olah memiliki
	"method seperti class dalam pemrograman berorientasi object"
	- Untuk membuat struct method kita cukup menuliskan struct yang ingin kita tambahkan method
	didepan deklarasi function




*/

package main

import "fmt"

// contoh program struct method
type User struct {
	Name, Status, Email string
	Age                 int
}

// membuat struct method
func (data User) getName() string { // membuat function getName, dengan struct method "User"
	return data.Name // dan mengembalikan nilai string
}

func (data User) getEmail() string {
	return data.Email
}

func main() {

	var data User
	data.Name = "Muhammad Rizki Syahputra"
	data.Status = "Student"
	data.Email = "riskideveloper2@gmail.com"
	data.Age = 18

	fmt.Println(data.getName())
	fmt.Println(data.getEmail())
}
