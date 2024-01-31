/* CATATAN

# Swicth statement di golang mirip seperti bahasa pemorgraman lainnya
perbedaannya di golang `tidak memerlukan statement break`

*/

package main

import "fmt"

func main() {

	// swicth statement
	warnaFavorit := "merah"

	switch warnaFavorit {
	case "biru":
		fmt.Println("Warna favorit kamu adalah biru")
	case "merah":
		fmt.Println("Warna favorit kamu adalah merah")
	case "hitam":
		fmt.Println("Warna favorit kamu adalah hitam")
	case "kuning":
		fmt.Println("Warna favorit kamu adalah kuning")
	case "hijau":
		fmt.Println("Warna favorit kamu adalah hijau")

	default:
		fmt.Println("kamu aneh")
	}

	// multi case switch statement\

	hari := "Minggu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jum'at":
		fmt.Println("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Hari tidak valid")
	}

}
