// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "",
			out: "",
		},
		{
			in:  "World",
			out: "Hello, World!",
		},
		{
			in:  "Kyle",
			out: "Hello, Kyle!",
		},
		{
			in:  "Kyle Shannon",
			out: "Hello, Kyle Shannon!",
		},
	}

	for _, test := range tests {
		if got, want := hello(test.in), test.out; got != want {
			t.Errorf("hello failed, got: '%s', want: '%s'", got, want)
		}
	}
}
