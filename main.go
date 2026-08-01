package main

import "fmt"

func data() string {
	nama := "Andika"
	umur := 20
	tinggi := 170.5
	lulus := true

	return fmt.Sprintf("Halo, nama saya %s, umur saya %d tahun, dengan tinggi badan %.1f cm, dan saya %t dari covid", nama, umur, tinggi, lulus)
}

func main() {
	fmt.Println(data())
}