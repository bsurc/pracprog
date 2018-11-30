// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ebird

import (
	"os"
	"testing"
)

func TestTaxa(t *testing.T) {
	fin, err := os.Open("taxa.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer fin.Close()
	taxa, err := LoadTaxa(fin)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := taxa["amekes"]; !ok {
		t.Error("no kestrel")
	}
}

func TestTaxaDownload(t *testing.T) {
	taxa, err := DownloadTaxa()
	if err != nil {
		t.Skip(err)
	}
	t.Log(len(taxa))
	if _, ok := taxa["amekes"]; !ok {
		t.Error("no kestrel")
	}
}
