package main

import (
	"fmt"
	"strings"
)

func main() {
	nama := "andika"
	animals := "kucing, anjing, burung"
	filter := strings.Split(animals, ",")
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