package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var friends = []Friend{
	Friend{1, "Haviz", "Jalan Haji kodir", "Frontend engineer", "meningkatkan kemampuan backend"},
	Friend{2, "Tasmara", "Jalan Haji Musa", "backend engineer", "meningkatkan kemampuan analisis"},
	Friend{3, "Jaka", "Jalan Haji Muchtar", "UI/UX Designer", "meningkatkan kemampuan programming"},
}

func getFriend(absen int) Friend {
	for _, f := range friends {
		if f.Absen == absen {
			return f
		}
	}
	return Friend{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage go run biodata.go [absen]")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	friend := getFriend(absen)

	if friend.Absen == 0 {
		fmt.Println("Absen", absen, "Tidak ditemukan")
		return
	}

	fmt.Println("Nama :", friend.Nama)
	fmt.Println("Alamat :", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan :", friend.Alasan)
}
