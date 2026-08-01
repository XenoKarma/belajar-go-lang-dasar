package main //package utama. Program Go selalu dimulai dengan package main. Package ini digunakan untuk menandai bahwa file ini adalah file utama yang akan dijalankan.

import (
	"fmt"     //package untuk input dan output
	"strings" //package untuk manipulasi string
)

func main() {
	//deklarasi variabel
	nama := "andika"

	//manipulasi string
	animals := "kucing, anjing, burung"

	//memisahkan string menjadi slice
	filter := strings.Split(animals, ",")

	//menghapus spasi di awal dan akhir string
	cihuy := "          hello world           "
	filterCihuy := strings.TrimSpace(cihuy)

	fmt.Println(strings.ToUpper(nama))
	fmt.Println(strings.ToLower(nama))
	fmt.Println(strings.Contains(animals, "kucing"))
	fmt.Println(animals)
	fmt.Println(strings.ReplaceAll(animals, "kucing", "harimau"))
	fmt.Println(filter)
	fmt.Println(filterCihuy)
}