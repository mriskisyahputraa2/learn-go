package main

import "fmt"

func main() {

	// if statement
	num := 10
	if num == 10 {
		fmt.Println("Number adalah 10")
	}

	// if else statement
	angka := 5

	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	// else if statement
	nilai := 100
	absen := 100

	if nilai > 90 && absen > 95 {
		fmt.Println("Sangat Baik")
	} else if nilai > 80 && absen > 85 {
		fmt.Println("Baik")
	} else if nilai > 79 && absen > 89 {
		fmt.Println("nilai kamu kurang baik tapi absen kamu bagus")
	} else {
		fmt.Println("Kurang Baiik")
	}

	// nested if
	number := 70

	if number >= 80 {
		fmt.Println("Nilai lebih dari 80")
		if number > 90 {
			fmt.Println("nilai lebih dari 90")
		}
	} else {
		fmt.Println("nilai kurang dari 80")
	}

	username := "riskiqa"
	password := 123

	if username == "riski" {
		fmt.Println("Username Benar")
		if password == 123 {
			fmt.Println("Password benar")
		} else {
			fmt.Println("Password salah")
		}
	} else {
		fmt.Println("username & password salah, gagal Login")
	}
}
