/* CATATAN
CONST, ADALAH SEBUAH VARIABEL YANG MEMPUNYAI NILAI TETAP (FIXED). ARTINYA NILAI NYA TIDAK BISA DIUBAH-UBAH

# ATURAN PENULIASAN CONST
 1. PENAMAAN VARIABEL CONST HARUS SESUAI DENGAN PERATAURAN MEN-DEKLARASI VARIABEL
 2. PENAMANA VARIABEL CONST LEBIH DISARANKAN DI AWALI DENGAN HURUF BESAR (OPTIONAL)
 3. VARIABEL CONST BISA DI TULIS DARI LUAR FUNCTION DAN DARI DALAM FUNCTION (BEBAS)



*/

package main

import "fmt"

// penulisan constanta
const nama = "riski"
const num = 12.6663

// penulisan variabel const seperti ini tidak bisa, ini akan error
// const num int;
// num = 10;

func main() {

	// Multiple Constanta Declaration
	const (
		A = 1
		B = "ahmad"
		C = 3.14
	)

	fmt.Println(nama)
	fmt.Println(num)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
