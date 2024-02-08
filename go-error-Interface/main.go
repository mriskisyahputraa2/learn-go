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

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak bisa melakukan pembagian dengan nol")
	} else {
		result := x / y
		return result, nil
	}

}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
