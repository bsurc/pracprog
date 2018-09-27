package main

import (
	"fmt"
)

// STRUCT_START_OMIT
type Bird struct {
	Species string `json:"spp" xml:"spp"`
	Sex     string `json:"sex" xml:"sex"`
}

// STRUCT_END_OMIT

func main() {
	// USE_START_OMIT
	var b Bird
	b.Species = "AMKE"
	b.Sex = "F"
	fmt.Printf("I saw a %s, it was a %s\n", b.Species, b.Sex)
	// USE_END_OMIT
}
