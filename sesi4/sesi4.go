package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	slc := []int{1,2,3,4,5,6,7,8,9,0}

	wg.Add(len((slc)))

	for _, val := range slc {
		go IIFE(&wg, val)
	}

	wg.Wait()
}

func IIFE (wg *sync.WaitGroup, val int) {
	defer wg.Done()
	fmt.Println(val)
}