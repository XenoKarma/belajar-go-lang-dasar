package main

import "fmt"

func main() {
	var color string
	fmt.Print("Masukkan Warna Lampu lalu lintas: ")
	fmt.Scanln(&color)

	switch color {
	case "red":
		fmt.Println("Lampu Merah: Berhenti")
	case "yellow":
		fmt.Println("Lampu Kuning: Hati-hati")
	case "green":
		fmt.Println("Lampu Hijau: Jalan")
	default:
		fmt.Println("Warna lampu tidak valid")
	}
	
}