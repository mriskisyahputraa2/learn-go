/* CATATAN:

uint8 dalam golang adalah tipe data bilangan cacah (unsigned integer) dengan lebar 8 bit, atau 1 byte. Nilai minimumnya adalah 0 dan nilai maksimumnya adalah 255.

Dalam Golang, tipe data uint8 sering digunakan untuk menyimpan nilai-nilai seperti:

    Kode warna
    Kode ASCII
    Kode biner
    Nilai bit



*/

package main

import "fmt"

func main() {
	var a uint8 = 10

	var (
		b = 3
		c = 3
		d = 15
		e = 13
		f = 11
		g = 3
	)

	b += 1
	c -= 10
	d /= 3
	e *= 2
	f %= 2
	g &= 3

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
}
