package main

import (
	"math"
)

type Stats struct {
	data []float64
}

func (s *Stats) Add(x ...float64) { //append
	s.data = append(s.data, x...)
}

func (s *Stats) Mean() float64 {
	mean := 0.0
	for _, x := range s.data {
		mean += x
	}
	return mean / float64(len(s.data)) //has to divide float by float, not float by integer
}

func (s *Stats) Min() float64 {
	min := math.MaxFloat64
	for _, x := range s.data {
		if x < min {
			min = x
		}
	}
	return min
}

func (s *Stats) Max() float64 {
	max := -math.MaxFloat64
	for _, x := range s.data {
		if x > max {
			max = x
		}
	}
	return max
}

func (s *Stats) Stan() float64 {
	stdev := 0.0
	for _, x := range s.data {
		stdev += (x - s.Mean()) * (x - s.Mean())
	}
	stdev = stdev / float64(len(s.data))
	stdev = math.Sqrt(stdev)
	return stdev
}

func (s *Stats) Sort() []float64 {
	sorted := make([]float64, len(s.data))
	copy(sorted, s.data)
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
	return sorted
}

func main() {
}
