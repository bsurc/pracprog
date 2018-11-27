// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	var s = "world!"
	fmt.Println(hello(s))
}
