package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"testing"
)

// cfsbData loads the Chocolate Frosted Sugar Bomb data into a Stats struct
func cfsbData(tb testing.TB) *Stats {
	fin, err := os.Open("CFSB.csv") //reading in data set
	if err != nil {
		tb.Fatal(err)
	}
	defer fin.Close()
	var st Stats
	s := bufio.NewScanner(fin)
	for s.Scan() {
		x, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			tb.Fatal(err)
		}
		st.Add(x)
	}
	return &st
}

func TestMean(t *testing.T) {
	var st Stats
	st.Add(4, 8)
	var got float64
	var want float64
	got = st.Mean()
	want = 6.0
	if got != want {
		t.Errorf("got: %f, want: %f", got, want)
	}
}

func TestMax(t *testing.T) {
	var st Stats
	st.Add(2, 3, 33, 5, 67, 8, 10, 12, 15)
	var got float64
	var want float64
	got = st.Max()
	want = 67
	if got != want {
		t.Errorf("got: %f, want: %f", got, want)
	}
}

func TestMin(t *testing.T) {
	var st Stats
	st.Add(1, 4, 5, 7, 8, 3) //this gives a test dataset, using st.Min accesses this function inside of the test, leading to an error
	var got float64
	var want float64
	got = st.Min()
	want = 1.0
	if got != want {
		t.Errorf("got: %f, want: %f", got, want)
	}
}

func TestStan(t *testing.T) {
	var st Stats
	st.Add(4, 8, 10)
	var got float64
	var want float64
	got = st.Stan()
	want = 2.494438258
	if math.Abs(got-want) > 0.0001 {
		t.Errorf("got: %f, want: %f", got, want)
	}
}

func TestSort(t *testing.T) {
}

func TestMode(t *testing.T) {
}

func TestMedian(t *testing.T) {
}

var sink float64

func BenchmarkMean(b *testing.B) {
	st := cfsbData(b)
	for i := 0; i < b.N; i++ {
		sink = st.Mean()
	}
}

func BenchmarkMax(b *testing.B) {
}

func BenchmarkMin(b *testing.B) {
}

func BenchmarkMode(b *testing.B) {
}

func BenchmarkMedian(b *testing.B) {
}

func BenchmarkSort(b *testing.B) {
}
