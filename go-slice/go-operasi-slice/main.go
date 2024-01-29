/* CATATAN

# OPERASI SLICE YANG AKAN DI PELAJARI:
	1. MENGAKSES ELEMEN DARI SEBUAH SLICE
	2. MENGUBAH ELEMEN DARI SEBUAH SLICE
	3. MENAMBAHKAN ELEMEN KE SLICE
		* fungsi append membuat slice baru dengan menambahkan data ke `posisi terakhir slice`
	4. MENGUBAH PANJANG SLICE
	5. MENYALIN ELEMEN DARI SEBUAH SLICE
		* Ketika menggunakan irisan, Go memuat semua elemen yang mendasarinya ke dalam memori.
		* Jika arranya besar dan Anda hanya membutuhkan beberapa elemen, lebih baik menyalin elemen-elemen tersebut menggunakan fungsi copy().
		* Fungsi copy() membuat array baru dengan elemen yang dibutuhkan untuk slice. Hal ini akan mengurangi memori yang digunakan untuk program.

*/

package main

import "fmt"

func main() {

	// MENGAKSES ELEMEN SLICE
	num := []int{1, 2, 3, 4, 5, 6}

	// OUTPUT MENGAKSES ELEMEN SLICE
	fmt.Println("MENGAKSES ELEMEN")
	fmt.Println(num[0])
	fmt.Println(num[4])
	fmt.Println("\n")

	// MENGUBAH ELEMEN SLICE
	num2 := []string{"Rizki", "Kamal", "Ahmad"}
	num2[1] = "Rendy"

	// OUTPUT MENGUBAH ELEMEN SLICE
	fmt.Println("MENGUBAH ELEMEN")
	fmt.Println(num2[1])
	fmt.Println(num2)
	fmt.Println("\n")

	// MENAMBAHKAN ELEMEN SLICE MENGGUNAKAN append
	num3 := []int{10, 20, 30, 40, 50}

	// OUTPUT MENAMBAKAN ELEMEN SLICE (BEFORE)
	fmt.Println("MENAMBAHKAN ELEMEN SLICE SEBELUM")
	fmt.Printf("Slice3 adalah = %v\n", num3)
	fmt.Printf("leght = %d\n", len(num3))
	fmt.Printf("capacity = %d\n", cap(num3))
	fmt.Println("\n")

	// OUTPUT MENAMBAKAN ELEMEN SLICE (AFTER)
	fmt.Println("MENAMBAHKAN ELEMEN SLICE SESUDAH")
	num3 = append(num3, 60, 70, 90) // MENAMBAHKAN ELEMENT SLICE
	fmt.Printf("Slice3 adalah = %v\n", num3)
	fmt.Printf("leght = %d\n", len(num3))
	fmt.Printf("capacity = %d\n", cap(num3))
	fmt.Println("\n")

	// MENGUBAH PANJANG SLICE
	num4 := [6]int{15, 25, 35, 45, 55, 65} // ini array
	mySlice1 := num4[1:5]                  // slice array

	// OUTPUT MENGUBAH PANJANG SLICE (BEFORE)
	fmt.Println("MENGUBAH PANJANG ELEMENT SLICE SEBELUM 1:5")
	fmt.Println("ini slice array sebelum diubah =", mySlice1)
	fmt.Println("leght =", len(mySlice1))
	fmt.Println("capacity =", cap(mySlice1))
	fmt.Println("array asli =", num4)
	fmt.Println("\n")

	// OUTPUT MENGUBAH PANJANG SLICE (AFTER)
	mySlice1 = num4[1:3] // INI MENGUBAH LENGHT DENGAN RE-SLICE SEBUAH ARRAY
	fmt.Println("MENGUBAH PANJANG ELEMENT SLICE SESUDAH 1:3")
	fmt.Println("ini slice array sesudah diubah =", mySlice1) // awalnya [25 35]
	fmt.Println("leght =", len(mySlice1))
	fmt.Println("capacity =", cap(mySlice1))
	fmt.Println("array asli =", num4)
	fmt.Println("\n")

	// OUTPUT MENAMBAHKAN PANJANG ELEMENT MENGGUNAKAN APPEND
	fmt.Println("MENAMBAHKAN PANJANG ELEMEN DENGAN SLICE ")
	mySlice1 = append(mySlice1, 75, 85, 95, 100)              // MENGUBAH PANJANG SLICE DENGAN MENAMBAHLAN ELEMENT
	fmt.Println("menambahkan slice sebuah array =", mySlice1) // setelah ditambah dengan append jadi [25 35 75 85 95]
	fmt.Println("leght =", len(mySlice1))
	fmt.Println("capacity =", cap(mySlice1))
	fmt.Println("array asli =", num4)
	fmt.Println("\n")

	// MENYALIN(COPY) ELEMEN DARI SEBUAH SLICE
	angka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println("MENYALIN/COPY PANJANG ELEMEN DENGAN SLICE BEFORE ")
	fmt.Println("original angka =", angka)
	fmt.Println("leght =", len(angka))
	fmt.Println("capacity =", cap(angka))
	fmt.Println("\n")

}
