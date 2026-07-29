package main

import "fmt"

// cara nyimpen varabel didalam golang ada 2 cara
// 1 var
// var nama string = "Dimas"
// var usia uint = 27
// var isCompleted bool = true

func main() {
	// short declaration :=
	nama := "Budi"
	usia := 30
	isCompleted := false
	tinggiBadan := 169.6
	// rune
	// merupakan tipe data yang digunakan untuk menyimpan karakter
	emoji := '😎'


	fmt.Println("Nama saya : ", nama)
	fmt.Println("Usia saya : ", usia)
	fmt.Println("Apakah tugas selesai? : ", isCompleted)
	fmt.Println("Tinggi badan saya : ", tinggiBadan)
	fmt.Println("Emoji saya : ", emoji)
}