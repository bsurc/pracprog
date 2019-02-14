package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fin, err := os.Open("CFSB.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	s := bufio.NewScanner(fin)
	var data []float64
	for s.Scan() {
		x, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, x)
	}

	var avg float64
	for _, x := range data {
		avg = avg + x
	}
	avg = avg / float64(len(data))
	fmt.Printf("Average: %.3f\n", avg)

	var max float64
	for _, x := range data {
		if x > max {
			max = x
		}
	}
	fmt.Printf("Max: %.3f\n", max)

	var min float64
	min = max + 1
	for _, x := range data {
		if x < min {
			min = x
		}
	}
	fmt.Printf("Min: %.3f\n", min)

	var hist = map[float64]int{}
	for _, x := range data {
		hist[x] += 1
	}
	var key float64
	var val int
	for k, v := range hist {
		if v > val {
			val = v
			key = k
		}
	}
	fmt.Printf("Mode: %.3f (%d)\n", key, val)

	// Median
}
