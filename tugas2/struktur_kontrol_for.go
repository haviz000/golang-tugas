package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai 1 =", i)
	}

	for j := 0; j <= 10; j++ {
		if j < 5 {
			fmt.Println("Nilai j =", j)
		} else {
			break
		}
	}

	str := "CAШABРBO"

	for i := 0; i < len(str); i++ {

		if str[i] == 0xd0 && str[i+1] == 0xa8 {
			fmt.Println("character U+0428 'Ш' starts at byte position", i)
		}

		if str[i] == 0x41 {
			fmt.Println("character U+0410 'A' starts at byte position", i)
		}

		if str[i] == 0x42 {
			fmt.Println("character U+0412 'B' starts at byte position", i)
		}

		if str[i] == 0x4f {
			fmt.Println("character U+041E 'O' starts at byte position", i)
		}
	}

	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j =", j)
	}
}
