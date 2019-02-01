// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "testing"

func TestDecodeObsLength(t *testing.T) {
	var badVals [12]string
	var err error
	_, err = decodeObs(badVals[:])
	if err == nil {
		t.Error("bad values passed")
	}
}

func TestDecodeObs(t *testing.T) {
	var got Obs
	var want Obs
	want.GlobalUniqueIdentifier = "URN:CornellLabOfOrnithology:EBIRD:OBS67812684"
	var vals []string
	vals = []string{
		"URN:CornellLabOfOrnithology:EBIRD:OBS67812684", // 0
		"2013-05-16 16:17:25.0",
		"1513",
		"species",
		"Gray Partridge",
		"Perdix perdix", // 5
		"",
		"",
		"X",
		"",
		"", // 10
		"",
		"United States",
		"US",
		"Idaho",
		"US-ID", // 15
		"Ada",
		"US-ID-001",
		"",
		"9",
		"", // 20
		"",
		"Kuna",
		"L191852",
		"T",
		"43.49179", // 25
		"-116.41996",
		"1969-02-01",
		"",
		"obsr180108",
		"S4840007", // 30
		"Incidental",
		"P20",
		"EBIRD",
		"",
		"",
		"",
		"",
		"0",
		"", "0", "1", "0",
	}

	var err error
	got, err = decodeObs(vals)
	if err != nil {
		t.Error(err)
	}
	if got.GlobalUniqueIdentifier != want.GlobalUniqueIdentifier {
		t.Errorf("decoding failed, got: %s, want: %s", got.GlobalUniqueIdentifier, want.GlobalUniqueIdentifier)
	}
}
