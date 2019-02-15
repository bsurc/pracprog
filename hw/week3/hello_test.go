package main

import (
	"fmt"
	"strings"
	"testing"
)

type Point struct {
	X float64
	Y float64
}

type Rect struct {
	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
}

func PointInRect(p Point, r Rect) bool {
	if p.X <= r.MaxX && p.X >= r.MinX &&
		p.Y <= r.MaxY && p.Y >= r.MinY {
		return true
	}
	return false
}

var sink bool

func BenchmarkPointInRect(b *testing.B) {
	var r Rect
	var p Point
	r.MinX = 0
	r.MaxX = 10
	r.MinY = 0
	r.MaxY = 10
	p.X = 5
	p.Y = 5
	for i := 0; i < b.N; i++ {
		sink = PointInRect(p, r)
	}
}

func TestPointRect(t *testing.T) {
	var r Rect
	var p Point
	r.MinX = 0
	r.MaxX = 10
	r.MinY = 0
	r.MaxY = 10
	p.X = 5
	p.Y = 5
	var got bool
	got = PointInRect(p, r)
	if got != true {
		t.Errorf("got: %t, want: %t, %+v, %+v", got, true, p, r)
	}
	p.X = 15
	p.Y = 15
	got = PointInRect(p, r)
	if got != false {
		t.Errorf("got: %t, want: %t, %+v, %+v", got, false, p, r)
	}
	p.X = 10
	p.Y = 10
	got = PointInRect(p, r)
	if got != true {
		t.Errorf("got: %t, want: %t, %+v, %+v", got, true, p, r)
	}
}

func HelloKyle() string {
	return "Hello, Kyle!"
}

// HelloWorld returns the famous 'Hello, World!' string, substituting 'World'
// with the name provided by the caller.  If name is an empty string, or a
// string that contains only white space, an empty string is returned.
func HelloWorld(name string) string {
	if strings.TrimSpace(name) == "" {
		return ""
	}
	s := fmt.Sprintf("Hello, %s!", name)
	return s
}

// Write tests for the function above.  Identify special cases, and make sure
// you test for them.  At a minumum, two cases should be tested, a third may be
// more clear, but not necessary.  Make sure to read the comment to understand
// what the function *should* do.
func TestHelloWorld(t *testing.T) {
	type test struct {
		in   string
		want string
	}
	var tests = []test{
		test{in: "Kyle", want: "Hello, Kyle!"},
		test{in: "", want: ""},
		test{in: "   ", want: ""},
	}
	var got string
	for i, helloTest := range tests {
		got = HelloWorld(helloTest.in)
		if got != helloTest.want {
			t.Errorf("at %d, got: %s, want: %s", i, got, helloTest.want)
		}
	}
}

func TestHelloKyle(t *testing.T) {
	var got, want string
	want = "Hello, Kyle!"
	got = HelloKyle()
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
