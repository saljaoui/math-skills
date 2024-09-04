package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: go run . [data.txt]")
		return
	}

	fileName := os.Args[1]
	file, err := os.ReadFile("./stat-bin-dockerized/stat-bin/" + fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	content := strings.Split(string(file), "\n")
	var data []int
	for _, s := range content {
		if s != "" {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("invalid integer in file: %v", err)
			}
			data = append(data, i)
		}
	}

	fmt.Println("Average:", int(math.Round(Average(data))))
	fmt.Println("Median:", int(math.Round(Median(data))))
	fmt.Println("Variance:", int(math.Round(Variance(data))))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDeviation(Variance(data)))))
}

func Average(data []int) float64 {
	var sum int
	for _, s := range data {
		sum += s
	}

	// Convert sum to float64 before dividing
	return float64(sum) / float64(len(data))
}

func Median(data []int) float64 {
	// Sort the slice
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	// Find the median
	length := len(data)
	if length%2 == 0 {
		return float64(data[length/2-1]+data[length/2]) / 2
	}

	return float64(data[length/2])
}

func Variance(data []int) float64 {
	avrg := Average(data)
	var sum float64
	for _, s := range data {
		diff := avrg - float64(s)
		sum += diff * diff
	}
	return sum / float64(len(data))
}

func StandardDeviation(data float64) float64 {
	return math.Sqrt(data)
}
