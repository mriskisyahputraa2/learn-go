/* CATATAN

# ERROR INTERFACE
	- Di Golang mempunyai interface yang digunakan sebagai kontrak untuk membuat
	error, interface itu bernama `error`
	- Error Interface memiliki sebuah method bernama "Error" dengan return value
	berupa string
	- untuk membuat error di golang kita bisa membuat struct yang memiliki
	function "Error" sehingga memenuhi kontrak error interface
	- Tapi daripada kita membuat struct secara manual, kita bisa memanfaatkan
	helper yang sudah disediakan oleh golang, "yaitu package 'error'"
	- penulisan syntax
		type error interface {
			Error() string
		}



*/

package main

import (
	"errors"
	"fmt"
)

// membuat function untuk melakukan pembagian angka
func divide(x, y int) (int, error) { // menerima params x dan y bertipe int, dan return-nya nilai int dan error
	if y == 0 { // cek jika, params y sama dengan 0. y disini mksudnya jika nilai y nya adalah 0, tidak akan bisa melukakan pembagian
		// maka kembalikan nilai 0 dengan fungsi error dari golang dan menampilkan pesan
		return 0, errors.New("Tidak bisa melakukan pembagian dengan nol")
	} else { // jika nilai y bukan 0
		result := x / y    // maka nilai x dan y dibagi
		return result, nil // kembalikan hasil bagi dan nil nya. mksudnya jika nilainya nil maka valid, jika tidak nil maka error
	}

}

func main() {

	// contoh kode program error yang standar
	var contohError error = errors.New("Ini adalah contoh error, menggunakan fungsi error") // error itu adalah sebuah interface juga
	fmt.Println(contohError)

	// mendeklarasikan nilai x dan y dan disimpan dalam v-result dan err
	result, err := divide(0, 5) // kenapa harus 2 variabel? karena func nya mengembalikan 2 fungsi. fungsi ke-1 adalah x dan y, fungsi ke-2 adalah error

	if err != nil { // jika err tidak sama dengan nil
		fmt.Println(err.Error()) // maka tampilkan persan error
	} else { // jika sama dengan nil
		fmt.Println(result) // tampilkan hasil dari pembagian
	}
}
