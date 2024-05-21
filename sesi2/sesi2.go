package main

import "fmt"

func main() {
	var panjang int = 5 // Change this variable to get a different size of pyramid

	var initiate int = 1
	result := []int{}

	for i := 0; i < panjang; i++ {
		if len(result) == 0 {
			result = append(result, initiate)
		} else {
			initiate += 2
			result = append(result, initiate)
		}
	}

	lastVal := result[len(result)-1]

	for i := 0; i < panjang; i++ {

		var startIndex int
		var lastIndex int

		if lastVal == result[i] {
			startIndex = 0
			lastIndex = lastVal - 1
		} else {
			startIndex = (lastVal - result[i]) / 2
			if result[i] == 1 {
				lastIndex = startIndex
			} else {
				lastIndex = (startIndex + result[i]) - 1
			}
		}

		for counter := 0 ; ; {

			if(counter >= startIndex && counter <= lastIndex)  {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

			if(counter == lastVal) {
				fmt.Println()
				break
			}else {
				counter++
			}

		}
	}
}