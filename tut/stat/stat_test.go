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
}

func TestMin(t *testing.T) {
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
