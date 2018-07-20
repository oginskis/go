// sumfloats project main.go
package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]
	result := calculateSum(numbers)
	Println(result)
}

func calculateSum(numbers []string) float64 {
	sum := 0.0
	for _, x := range numbers {
		parsed, err := strconv.ParseFloat(x, 64)
		if err == nil {
			sum += parsed
		}
	}
	return sum
}
