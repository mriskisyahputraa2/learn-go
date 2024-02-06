/* CATATAN

# DEFER
	- Defer adalah sebuah "statement yang akan menunda pengeksekusian
	sebuah function hingga function terderkatnya selesai dieksekusi" atau
	mengembalikan nilai
	- Function yang dipanggil menggunakan "Defer" statement akan selalu
	"dijalankan meskipun terjadi error"
	- Penting: Pemanggilan function defer harus ditulis dipaling atas, tidak boleh di dibawah
	- dan satu lagi ketika di bawah function defer ada kode yang sedang dieksekusi dan terdapat error maka defer akan menampilkan error dengan bertulis "Panic: <pesan error>".

# PANIC
	- "Panic" adalah sebuah function yang digunakan untuk menghentikan program
	- "Panic" function biasanya akan terpanggil saat terjadi error pada program kita
	- "Panic" function juga "bisa kita panggil secara manual", ini berguna ketika
	misalnya, kita melkukan pengecekan terhadap suatu kondisi
	- Ketika "panic" function dipanggil, function yang menggunakan statement defer akan tetap
	dijalankan meskipun program dihentikan

# RECOVER
	- "Recover" adalah function yang berfungsi untuk menangkap data dari panic function
	- Dengan "Recover" proses panic akan terhenti, sehingga program akan tetap berjalan
	- "Recover" function fungsinya mirip seperti "catch" pada bahasa pemrograman lain
	- Program tetap bejalan jika terjadi error, tapi akan terdapat pesan error
*/

package main

import "fmt"

// contoh program defer
func txtNumber(txt string) string {
	return txt
}

// contoh program "Panic" sambung dengan "Recover"
func division(num1, num2 int) { // create function division and passing parameters num1 and num2 type int

	// menangkan defer dari function handlePanic
	defer handlePanic()

	if num2 == 0 { // if num2 == 0, return message error
		panic("Tidak dapat membagi angka dengan nol") // view message panic
	} else { // if not error
		result := num1 / num2           // num1 / num2 result 1 or etc
		fmt.Println("Result =", result) // return view message
	}
}

// Contoh Program "Recover" sambungan dari program "Panic"
func handlePanic() { // create function handlePanic, for debugging function division if terjadi error

	err := recover() // create keyword recover() and disimpan in v-err, baca diatas definisi recover

	if err != nil { // if err != nil == jika v-err tidak sama dengan nill(nilai nol)
		fmt.Println("RECOVER", err) // print messange RECOVER dengan pesan error dari panic di function division
	} else {
		fmt.Println("Program tidak terjadi error dari saya yang menggunakan fungsi recover")
	}
}

func main() {

	// ouput dari defer, fungsi defer ini akan ditampilkan setelah "Println Dua". baca definisinya diatas
	defer fmt.Println("Kode ini dipanggil dengan fungsi defer:", txtNumber("Tiga"))
	// defer txtNumber()
	// ini jika terjadi error di pertengahan kode, akan muncul "panic"
	// num := 0
	// result := 10 / num
	// fmt.Println(result)
	fmt.Println("Satu")
	fmt.Println("Dua")
	fmt.Println("\n")

	// output program "Panic" dan "Recover"
	// division(5, 0) // 5 is num1 and 0 is num2, this code return message panic = 0
	// division(5, 5) // this code return message not error / success = 1

	// output program "Recover"
	division(3, 3) // contoh program error
	fmt.Println("\n")

}
