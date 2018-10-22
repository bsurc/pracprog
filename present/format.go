package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"os"
)

type Bird struct {
	Species string `json:"spp" xml:"spp"`
	Sex     string `json:"sex" xml:"sex"`
}

func main() {
	var b Bird
	b.Species = "AMKE"
	b.Sex = "F"

	// START_OMIT
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"spp", "sex"})
	w.Write([]string{b.Species, b.Sex})
	w.Flush()
	json.NewEncoder(os.Stdout).Encode(b)
	xml.NewEncoder(os.Stdout).Encode(b)
	// END_OMIT
}
