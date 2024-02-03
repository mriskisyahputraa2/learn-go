/* CATATAN

# VARIADIC FUNCTION
	- VARIADIC FUNCTION ADALAH SEBUAH FUNCTION YANG DAPAT MENERIMA SEJUMLAH VARIABEL
	SEBAGAI ARGUMEN, ARGUMEN INI DISEBUT DENGAN "varargs"
	- UNTUK MENDEKLARASI VARAIDIC FUNCTION, TIPE DATA PADA PARAMETER TERAKHIR DIDAHULUI
	DENGAN ...(titik tiga)
	- JENIS FUNGSI INI BERGUNA KETIKA KITA TIDAK MENGETAHUI JUMLAH ARGUMEN YANG KITA
	BERIKAN KE FUNGSI TERSEBUT
	- SALAH SATU CONTOH VARIADIC FUNCTION YAITU FUNGSI "Println", BAWAAN DARI PAKET "fmt"
	- ... (titik tiga) BERFUNGSI UNTUK MEMBUAT BANYAK NILAI/VALUE ARGUMENT

# NOTE PENTING PENULISAN VARIADIC FUNCTION:
	- PENULISAN VARIADIC FUNCTION TIDAK BOLEH DIAWAL DAN DITENGAH JIKA ANDA MAU MEMBUAT BANYAK PARAMETER.
	- JIKA MEMBUAT 2 ATAU LEBIH PARAMETER MAKA YANG PARAMETER VARIADIC DITEMPATKAN SEBELAH KIRI
	- JIKA HANYA MEMBUAT 2 PARAMATER SAJA VARIADIC FUNCTION TIDAK BOLEH DITEMPATKAN DI SEBELAH KANAN
	- VARIADIC FUNCTION HANYA BOLEH DITEMPATKAN PALING UJUNG SEBELAH KANAN

# CONTOH SYNTAX YANG TIDAK BOLEH:
	1. membuat paramter lebih dari 2
		A. func sum (param1 int, nums ...int, param2 int) {

		}
		B. func sum (nums ...int, param1 int, param2 int) {

		}
	2. func sum (nums ...int, param1 int) {

		}

# SLICE SEBAGAI PARAMETER
	- PADA CONTOH SEBELUMNYA KITA MENULISAKAN ANGKA-ANGKA LANGSUNG SAAT MEMENGGIL
	FUNCTION
	- SEBENARNYA DEFAULTNYA KITA BISA MEMASUKKAN SLICE SEBAGAI ARGUMENT PADA VARIADIC FUNCTION
	- KARENA Varargs PADA VARIADIC FUNCTION AKAN DIUBAH MENJADI "SLICE", OLEH KARNA ITU
	KITA BISA MELAKUKAN PARULANGAN PADA Varrags UNTUK MELAKUKAN PENJUMLAHAN.

#  Fungsi range:
	- Mengulang setiap elemen slice, array, map, atau string.
	- Mendapatkan nilai dan index (opsional) dari setiap elemen.

# (Underscore) _ : digunakan untuk tidak menggunakan index dari slice, jadi indexnya tidak ditampilkan dan digunakan. karna default dari slice mengembalikan index dan value


*/

package main

import "fmt"

// ...(titik tiga) ini digunakan untuk bisa memasukkan banyak nilai dari argument disebut dengan slice
func sum(nums ...int) int {

	// v-total ini digunakan untuk menampung nilai parameter nums diawali dengan 0
	total := 0

	// melakukan perulangan paramter nums
	// _(undersroce) digunakan untuk tidak perlu menggunakan index pada slice, karna default dari slice mengembalikan index dan value, sedangkan disini kita tidak perlu dengan indexnya hanya mengambil valuenya yaitu "num"
	for _, num := range nums { // range digunakan untuk mengulang setiap element slice
		total += num // nilai dari total ditambah dengan num. jadi awalnya 1+1=2+1=3+3=6+4=10+5=15
	}
	return total // lalu kembalikan nilai dari total

}

func main() {
	// contohnya ini kita memasukkan banyak nilai dengan satu argument dipisahkan dengan koma
	sumAll := sum(1, 2, 3, 4, 5) // ini nilai paramaternya sudah masuk ke perulangan num, ini data slice
	fmt.Println(sumAll)          // 15
}
