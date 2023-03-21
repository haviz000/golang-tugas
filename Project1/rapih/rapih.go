package main

import (
	"fmt"
	"sync"
)

func printData(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex, name string) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(data, name)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printData(data1, &wg, &mutex, "1")
		go printData(data2, &wg, &mutex, "2")
	}

	wg.Wait()
}
