/* CATATAN

# TYPE ASSERTIONS
	- Type Assertion di golang adalah cara untuk merubah tipe data tertentu
	menjadi tipe data yang kita inginkan
	- Type Assertion dapat digunakan dengan nilai yang mengimplementasikan
	interface kosong


*/

package main

import "fmt"

func main() {
	var data interface{} = "Hello World"

	dataString := data.(string)
	fmt.Println(dataString)
	fmt.Printf("Value: %v dan tipe data: %T\n", dataString, dataString)
}

// push ke github type-declaration, sama error interface, dan ini
