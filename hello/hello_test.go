// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if got, want := hello("kyle"), "Hello, kyle!"; got != want {
		t.Errorf("hello failed, got: '%s', want: '%s'", got, want)
	}
}
