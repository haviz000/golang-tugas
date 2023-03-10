package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")

	j := true
	fmt.Printf("%v\n", j)

	russianChar := 'Я'
	fmt.Printf("%c\n", russianChar)

	decimalNum := 21
	fmt.Printf("%d\n", decimalNum)

	octalNum := 025
	fmt.Printf("%o\n", octalNum)

	hexNumLower := 0xf
	fmt.Printf("%x\n", hexNumLower)

	hexNumUpper := 0xF
	fmt.Printf(" %X\n", hexNumUpper)

	unicodeChar := 'Я'
	fmt.Printf("%U\n", unicodeChar)

	k := 123.456
	fmt.Printf("%f\n", k)
	fmt.Printf("%e\n", k)
}
