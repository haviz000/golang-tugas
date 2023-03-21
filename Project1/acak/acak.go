package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func printData(data interface{}, prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Println(data, i+1)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	data1 := []string{"bisa1", "bisa2", "bisa3"}
	data2 := []string{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printData(data1, "Data 1", &wg)
	}()

	go func() {
		printData(data2, "Data 2", &wg)
	}()

	wg.Wait()
}
