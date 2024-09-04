package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"soufian/funcs"
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

	fmt.Println("Average:", int(math.Round(funcs.Average(data))))
	fmt.Println("Median:", int(math.Round(funcs.Median(data))))
	fmt.Println("Variance:", int(math.Round(funcs.Variance(data))))
	fmt.Println("Standard Deviation:", int(math.Round(funcs.StandardDeviation(funcs.Variance(data)))))
}
