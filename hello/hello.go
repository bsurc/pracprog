// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
)

// hello returns the famous 'Hello, World!' string replacing world with
// name.  If name is an empty string, and empty string is returned.
func hello(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hello %s!", name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
