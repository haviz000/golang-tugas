package main

import "fmt"

func main() {
	input := "selamat malam"
	mapping := make(map[string]int)

	for _, huruf := range input {
		fmt.Println(string(huruf))
		mapping[string(huruf)]++
	}

	fmt.Println(mapping)
}
