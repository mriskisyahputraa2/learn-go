/* CATATAN

# STRUCT
	- "Struct" adalah kumpulan data field yang dideklarasikan dengan tipe data
	- "Struct" digunakan untuk "menyimpan beberapa nilai dari tipe data yang
	berbeda kedalam satu variabel"
	- "Struct" adalah sebuah template, "setelah struct dibuat tidak bisa langsung
	digunakan, kita bisa membuat data struct yang sudah kita buat"

# STRUCT LITERALS
	- Sebelumnya kita sudah membuat data dari struct, tapi sebenarnya ada
	beberapa cara yang bisa kita gunakan untuk membuat data dari struct:
	seperti penulisan array, slice, dan map
	- Kita bisa membuat struct dan mengisi datanya secara langsung dengan
	menyebutkan nama field atau tanpa menyebutkan nama field

*/

package main

import "fmt"

// contoh program struct
type User struct {
	Name, Email string
	Age         int
}

func main() {
	// menyimpan data Struct kedalam var-user, ini disebut mengambali field dari data struct
	var user User

	// setalah disimpan dalam v-user, maka tentukan value dari field-nya
	user.Name = "Rizki"
	user.Email = "riskideveloper2@gmail.com"
	user.Age = 18

	// contoh program struct literal, cuma ini yang terdapat perbedaan, selebihnya sama semua
	var user1 = User{
		Name:  "Syahputra",
		Email: "syahputra@gmail.com",
		Age:   17,
	}

	// contoh struct literlas(Tanpa Menyebutkan Nama Field)
	// ini terdapat kelemahan, jika datanya tidak sesuai dengan field-nya atau menambahkan field-nya maka akan terjadi error. Not Recommended
	user2 := User{
		"Muhammad",
		"muhammad@gmail.com",
		19,
	}

	// ouput struct
	fmt.Println("Data Full dari User Struct =", user)
	fmt.Println("Nama:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Umur:", user.Age, "Tahun")
	fmt.Println("\n")

	// ouput struct literals
	fmt.Println("Data Full dari User1 Struct =", user1)
	fmt.Println("Nama:", user1.Name)
	fmt.Println("Email:", user1.Email)
	fmt.Println("Umur:", user1.Age, "Tahun")
	fmt.Println("\n")

	// ouput struct literals(Tanpa Menyebutkan Nama Field)
	fmt.Println("Data Full dari User1 Struct =", user2)
	fmt.Println("Nama:", user2.Name)
	fmt.Println("Email:", user2.Email)
	fmt.Println("Umur:", user2.Age, "Tahun")
	fmt.Println("\n")
}
