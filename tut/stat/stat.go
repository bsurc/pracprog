package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Stats struct {
	data []float64
}

func (s *Stats) Add(x ...float64) {
	s.data = append(s.data, x...)
}

func (s *Stats) Mean() float64 {
	var x float64
	x = 0
	var i int
	for i = 0; i < len(s.data); i++ {
		x = x + s.data[i]
	}
	x = x / float64(len(s.data))
	return x
}

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

	var mean float64
	for _, x := range data {
		mean = mean + x
	}
	mean = mean / float64(len(data))
	fmt.Printf("Mean: %.3f\n", mean)

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
		hist[x]++
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

	sorted := make([]float64, len(data))
	copy(sorted, data)
	// Insertion sort from https://en.wikipedia.org/wiki/Insertion_sort
	i := 1
	for i < len(sorted) {
		j := i
		for j > 0 && sorted[j-1] > sorted[j] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
		i++
	}
	fmt.Printf("Median: %.3f\n", sorted[len(sorted)/2])

	var stdev float64
	for _, x := range data {
		stdev += (x - mean) * (x - mean)
	}
	stdev /= float64(len(data) - 1)
	stdev = math.Sqrt(stdev)
	fmt.Printf("Std Dev: %.3f\n", stdev)
}
