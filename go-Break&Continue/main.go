/* CATATAN

# BREAK & CONTINUE
	- BREAK DAN CONTINUE ADALAH KEYWORD YANG BISA DIGUNAKAN DALAM PERULANGASN
	- BREAK DIGUNAKAN UNTUK MENGHENTIKAN PERULANGAN
	- CONTINUE DIGUNAKAN UNTUK MELEWATKAN SATU ATAU LEBIH ITERASI, KEMUDIAN
	MELANJUTKAN ITERASI SELANJUTNYA DALAM LOOP
	- BREAK DAN CONTINUE BIASNYA DIGUNAKAN DENGAN PENGKONDISIAN



*/

package main

import "fmt"

func main() {

	// Continue
	for i := 10; i <= 100; i += 10 {
		if i == 10 || i == 100 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Selesai")
	fmt.Println("\n")

	for j := 1; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("Selesai")
	fmt.Println("\n")

	// Break
	for j := 1; j < 10; j++ {
		if j == 7 {
			break
		}
		fmt.Println(j)
	}
	fmt.Println("Selesai")

}
