package main

import (
	"math"
)

const (
	flagMin   = 1 << 0
	flagMax   = 1 << 1
	flagStan  = 1 << 2
	flagMean  = 1 << 3
	flagClean = flagMin | flagMax | flagStan | flagMean
)

type Stats struct {
	data                 []float64
	min, max, stan, mean float64
	clean                uint8
}

func (s *Stats) Add(x ...float64) { //append
	s.data = append(s.data, x...)
	s.clean = 0
}

func (s *Stats) Mean() float64 {
	if s.clean&flagMean > 0 {
		return s.mean
	}
	mean := 0.0
	for _, x := range s.data {
		mean += x
	}
	s.mean = mean / float64(len(s.data))
	s.clean = s.clean | flagMean
	return s.mean
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
	mean := s.Mean()
	for _, x := range s.data {
		stdev += (x - mean) * (x - mean)
	}
	stdev = stdev / float64(len(s.data))
	stdev = math.Sqrt(stdev)
	return stdev
}

func (s *Stats) Sort() []float64 {
	sorted := make([]float64, len(s.data))
	copy(sorted, s.data)
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
