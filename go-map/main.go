/* CATATAN
# MAP
	* MAP DIGUNAKAN UNTUK MENYIMPAN NILAI DATA DALAM (KEY: VALUE)
	* SETIAP ELEMEN DALAM MAP ADALAH PASANGAN (KEY: VALUE)
	* KEY PADA MAP (BERSIFAT UNIK TIDAK BOLEH SAMA, JIKA ADA KEY YANG SAMA MAKA
	SECARA OTOMATIS DATA SEBELUMNYA AKAN DIGANTI DENGAN DATA BARU)
	* (PANJANG MAP ADALAHH JUMLAH ELEMENNYA), KITA BISA MENDAPATKAN PANAJANG MAP DENGAN
	MENGGUNAKAN FUNGSI DALAM `len()`
	* DEFAULT VALUE DARI MAP ADALAH (NIL) / NILAI KOSONG

# SYNTAX MEMBUAT MAP
	1. DENGAN KATA KUNCI VAR DAN TANDA :=
		* var map_name = map[KeyType]ValueType{key1: value1, key2: value2,....}
		* map_name = map[KeyType]ValueType{key1: value1, key2: value2,....}

	2. DENGAN FUNCTION make()
		* var map_name = make(map[KeyType]ValueType)
		* map_name = make(map[KeyType]ValueType)
*/

package main

import "fmt"

func main() {

	// MEMBUAT MAP DENGAN KEYWORD VAR DAN :=
	// tipe data string pertama itu untuk key dan tipe data string kedua untuk value
	var person = map[string]string{
		"name":     "Rizki",
		"activity": "sebagai mahasiswa",
		"hobby":    "coding",
	}

	// tipe data string pertama itu untuk key dan tipe data integer kedua untuk value
	detailPerson := map[string]int{
		"tgl_lahir": 8,
		"bln_lahir": 7,
		"thn_lahir": 2005,
		"umur":      18,
	}

	fmt.Println("Nama saya", person["name"])
	fmt.Println("Aktivitas saya", person["activity"])
	fmt.Println("Hobby saya", person["hobby"])
	fmt.Println(person)
	fmt.Println("\n")

	fmt.Println("Tanggal Lahir", detailPerson["tgl_lahir"])
	fmt.Println("Bulan Lahir", detailPerson["bln_lahir"], "July")
	fmt.Println("Tahun Lahir", detailPerson["thn_lahir"])
	fmt.Println("Umur saya", detailPerson["umur"], "Tahun")
	fmt.Println(detailPerson)
	fmt.Println("\n")

	// MEMBUAT MAP DENGAN FUNGSI MAKE()
	var laptop = make(map[string]string)
	laptop["merek"] = "Lenovo"
	laptop["model"] = "Ideaped Gaming 3"

	computer := make(map[string]int)
	computer["stok"] = 1000
	computer["harga"] = 1200000

	fmt.Println("Merek Laptop saya", laptop["merek"])
	fmt.Println("Model Laptop saya", laptop["model"])
	fmt.Println(laptop)
	fmt.Println("\n")

	fmt.Println("Stok Komputer", computer["stok"])
	fmt.Println("Harga Komputer", computer["harga"])
	fmt.Println(computer)
	fmt.Println("\n")

	// OUTPUT MENDAPATKAN PANJANG MAP DENGAN FUNGSI LEN()
	fmt.Println("Panjang Map dari person adalah =", len(person))
	fmt.Println("Panjang Map dari Detailperson adalah =", len(detailPerson))
	fmt.Println("\n")

}
