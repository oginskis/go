// bubblesort project main.go
package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	input := os.Args[1:]
	decimals := func() []float64 {
		result := make([]float64, 0)
		for _, value := range input {
			parsed, err := strconv.ParseFloat(value, 64)
			if err == nil {
				result = append(result, parsed)
			}
		}
		return result
	}
	result := sort(decimals(),0)
	fmt.Println(result)
}

func sort(numbers []float64, iteration int) []float64{
	if iteration == len(numbers) - 1 {
		return numbers
	}
	for x, y := range numbers[:len(numbers)-iteration-1] {
		if y > numbers[x+1] {
			tmp := numbers[x+1]
			numbers[x+1] = y
			numbers[x] = tmp
		}
	}
	return sort(numbers,iteration+1)
}
