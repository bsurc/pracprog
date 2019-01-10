// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ebird

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// Taxa represents a observation species or bird.  Derived from the eBird
// taxonomy, see http://www.birds.cornell.edu/clementschecklist/download/
//
// This probably won't get used in the pracprog project, but if we run into
// a case where we do need it, we don't want to deal with the line ending
// issues.
type Taxa struct {
	TaxonOrder        int
	Category          string
	SpeciesCode       string
	PrimaryCommonName string
	SciName           string
	Order             string
	Family            string
	SpeciesGroup      string
	ReportAs          string
}

const eBirdTaxaURL = "http://www.birds.cornell.edu/clementschecklist/wp-content/uploads/2018/08/eBird_Taxonomy_v2018_14Aug2018.csv"

// DownloadTaxa fetches taxa information from the eBirdTaxaURL, and returns the loaded taxa.
//
// TODO(ksshannon): The line endings from the URL are garbled, find a way to fix them.  Until
// then, use the local taxa.csv file, it has normal line endings.
func DownloadTaxa() (map[string]Taxa, error) {
	if true {
		return nil, errors.New("fix line ending issue")
	}
	resp, err := http.Get(eBirdTaxaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	t, err := LoadTaxa(resp.Body)
	return t, err
}

// LoadTaxa extracts data from an io.Reader and returns a map represented with SpeciesCode as a key
// and a Taxa object as a value.
func LoadTaxa(r io.Reader) (map[string]Taxa, error) {
	taxa := map[string]Taxa{}
	c := csv.NewReader(r)
	c.Comment = '#'
	row, err := c.Read()
	if err != nil {
		return nil, err
	}
	for {
		row, err = c.Read()
		if err != nil {
			break
		}
		to, err := strconv.Atoi(row[0])
		if err != nil {
			to = -1
		}
		taxa[row[2]] = Taxa{
			to,
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8],
		}
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return taxa, nil
}
