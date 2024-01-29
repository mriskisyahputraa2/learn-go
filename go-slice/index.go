/* CATATAN

# SLICE DI GOLANG
 - MIRIP DENGAN ARRAY, SLICE ADALAH KUMPULAN NILAIDARI TIPE DATA YANG SAMA,
   NAMUN SLICE LEBIH DINAMIS DAN FLEKSIBEL
 - TIDAK SEPERTI ARRAY, SLICE TIDAK MEMILIKI UKURAN TETAP, SLICE BISA BERTAMBAH
   DAN MENYUSUT SESAUI KEINGINAN KITA
 - DALAM PRAKTIKNYA DI GOLANG, SLICIE LEBIH SERING DIGUNAKAN DIBANDING ARRAY
 - UNTUK MEMBUAT SLICE ADA BEBERAPA CARA:
	* MEMBUAT LANGSUNG SEPERTI ARRAY
	* MEMBUAT SLICE DARI ARRAY
	* MENGGUNAKAN FUNGSI MAKE()

1. MEMBUAT SLICE LANGSUNG
	* var nama_slice = [] tipedata{values}
	* nama_slice := [] tipedata{values}
2. MEMBUAT SLICE DARI ARRAY
	* var iniArray = [length]tipedata{values}
	* iniSlice := iniArray[low:high]

 a. arrayp[low : high] = Membuat slice dari array di mulai dari `index low sampai index sebelum high`
 b. array[low:] = Membuat slice dari array dimulai dari `index low sampai index sebelum high`
 c. array[:high] = Membuat slice dari array dimulai dari `index 0 sampai index sebelum high`
 d. array[:] = Membuat slice dari array dimuali dari `index 0 sampai index akhir di array`



3. MEMBUAT SLICE DENGAN FUNGSI MAKE()
	* nama_slice := make([]tipedata,length, capacity)

# Fungsi len() - untuk mengetahui jumlah elemen dari slice
# Fungsi cap() - mengembalikan kapasitas irisan (jumlah elemen yang dapat bertambah atau berkurang)




*/

package main

import "fmt"

func main() {

	// Membuat slice secara langsung
	mySlice := [7]int{1, 2, 3, 4, 5}
	myName := [4]string{"Muhammad", "Rizki", "Syahputra"}

	// Membuat slice dari array
	// contoh a (low:height):
	iniArrayA := [10]int{1, 3, 4, 6, 67, 73, 63, 53, 21, 45} // ini array nya
	sliceA := iniArrayA[4:6]                                 // low=67, high=63 - 1 = 73

	iniArray2 := [12]int{1, 2, 4, 21, 6, 2, 5, 46, 2, 3, 25, 6} // ini array
	slice2 := iniArray2[3:7]                                    // low = 3, high=7

	// contoh b (low:)
	iniArray3 := [10]int{1, 3, 4, 6, 67, 73, 63, 53, 21, 45}
	sliceB := iniArray3[4:]

	// contoh c (high:)
	iniArray4 := [10]int{1, 3, 4, 6, 67, 73, 63, 53, 21, 45}
	sliceC := iniArray4[:6]

	// contoh d (:). artinya kita membuat slice dari semua array
	iniArray5 := [10]int{1, 3, 4, 6, 67, 73, 63, 53, 21, 45}
	sliceD := iniArray5[:]

	// Ouput slice secara langsung
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(mySlice)
	fmt.Println(cap(myName))
	fmt.Println(len(myName))
	fmt.Println(myName)

	// Ouput slice dari array
	// Ouput contoh a (low:height)
	fmt.Printf("ini adalah sliceA = %v\n", sliceA) // [67 73]. 67 didapat dari low=4 dan dihitung dari index 0, 73 didapat dari hight 6 dan dikurang 1 jadi 63-1 = 73
	fmt.Printf("height = %d\n", len(sliceA))       // 2. ini adalah panjang slice yaitu [67 dan 73]
	fmt.Printf("capacity = %d\n", cap(sliceA))     // 6. ini adalah capacity slice yaitu [67,73,63, 53, 21, 45]. capacity akan dihitung sampai akhir elemen array

	fmt.Printf("ini adalah sliceA = %v\n", slice2) // [21, 6, 2, 5]
	fmt.Printf("height = %d\n", len(slice2))       // height = 4
	fmt.Printf("capacity = %d\n", cap(slice2))     // capacity = 9

	// Output contoh b (low:)
	fmt.Printf("ini adalah sliceB = %v\n", sliceB) // [67, 73, 63, 53, 21, 45]
	fmt.Printf("height = %d\n", len(sliceB))       // height = 6
	fmt.Printf("capacity = %d\n", cap(sliceB))     // capacity = 6

	// Output contoh c (:height)
	fmt.Printf("ini adalah sliceC = %v\n", sliceC) // [1, 3, 4, 6, 67, 73], sebenarnya sampai 63 tapi dikurang 1 maka hasilnya 73
	fmt.Printf("height = %d\n", len(sliceC))       // height = 6
	fmt.Printf("capacity = %d\n", cap(sliceC))     // capacity = 10. kenapa? karena capacity dihitung dari 0

	// Output contoh d (:). ouputnya akan membuatslice dari semua array
	fmt.Printf("ini adalah sliceD = %v\n", sliceD)
	fmt.Printf("height = %d\n", len(sliceD))
	fmt.Printf("capacity = %d\n", cap(sliceD))

	// Contoh pass by reference (hati2 jika ada data array yang terubah)
	Array := [8]int{1, 2, 4, 21, 6, 2, 5, 46}
	slice := Array[3:5]

	// Ouput pass by reference (hati2 jika ada data array yang terubah)
	Array[0] = 10
	Array[4] = 300
	fmt.Printf("ini contoh pass by reference dari: slice = %v\n", slice)
	fmt.Printf("ini contoh pass by reference dari: array = %v\n", Array)

	// MEMBUAT SLICE DENGAN FUNGSI make()
	sliceMake := make([]int, 5, 5)
	sliceMake[0] = 10
	sliceMake[1] = 11

	fmt.Printf("Slice dari make adalah = %v\n", sliceMake)
	fmt.Printf("lenght = %d\n", len(sliceMake))
	fmt.Printf("capacity = %d\n", cap(sliceMake))

}
