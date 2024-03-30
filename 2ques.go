package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var numbers []float64
var windowSize = 10

func main() {
	http.HandleFunc("/numbers/", handleNumbers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleNumbers(w http.ResponseWriter, r *http.Request) {
	numberID := strings.TrimPrefix(r.URL.Path, "/numbers/")
	switch numberID {
	case "p", "f", "e", "r":
		
		numbers = append(numbers, 1.0, 2.0, 3.0, 4.0, 5.0)

		numbers = uniqueAndLimit(numbers, windowSize)

		avg := calculateAverage(numbers)

		fmt.Fprintf(w, "WindowPrevState: []\nWindowCurState: %v\nNumbers: %v\nAvg: %.2f\n", numbers, numbers, avg)
	default:
		http.Error(w, "Invalid number ID", http.StatusBadRequest)
	}
}

func uniqueAndLimit(numbers []float64, maxSize int) []float64 {
	uniqueNumbers := make(map[float64]bool)
	var result []float64
	for _, num := range numbers {
		if !uniqueNumbers[num] {
			uniqueNumbers[num] = true
			result = append(result, num)
		}
		if len(result) >= maxSize {
			break
		}
	}
	return result
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}
