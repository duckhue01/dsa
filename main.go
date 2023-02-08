package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	a := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 100000; i++ {
			a++
		}
		wg.Done()
	}()
	for i := 0; i < 100000; i++ {
		a++
	}
	wg.Wait()
	fmt.Println(a)
}
