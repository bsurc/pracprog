package main

import "testing"

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

func TestMax(t *testing.T) {
}

func TestMode(t *testing.T) {
}

func TestMedian(t *testing.T) {
}

func BenchmarkMean(b *testing.B) {
}

func BenchmarkMax(b *testing.B) {
}

func BenchmarkMin(b *testing.B) {
}

func BenchmarkMode(b *testing.B) {
}

func BenchmarkMedian(b *testing.B) {
}
