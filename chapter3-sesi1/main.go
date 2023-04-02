package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Elemen struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		waterStatus := getStatus(water, "water")
		windStatus := getStatus(wind, "wind")

		data := Elemen{water, wind}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(string(jsonData)))
		if err != nil {
			fmt.Println("Error:", err, "response", resp)
			return
		}

		io.Copy(ioutil.Discard, resp.Body)
		fmt.Println("Status Air:", waterStatus)
		fmt.Println("Status Angin:", windStatus)

		time.Sleep(15 * time.Second)
	}
}

func getStatus(value int, elemen string) string {
	switch elemen {
	case "water":
		if value < 5 {
			return "aman"
		} else if value >= 5 && value <= 8 {
			return "siaga"
		} else {
			return "bahaya"
		}
	case "wind":
		if value < 6 {
			return "aman"
		} else if value >= 6 && value <= 15 {
			return "siaga"
		} else {
			return "bahaya"
		}
	default:
		return ""
	}
}
