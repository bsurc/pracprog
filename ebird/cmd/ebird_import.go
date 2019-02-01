// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println(`
ebird_import ebird_data.csv output.db

	ebird_import reads ebird data from the supplied csv file
	with ebird data, and populates a SQLite3 database.  It does
	not create any indices or any other side effects.  Only
	CREATE and INSERT SQL commands are issued.  The table is named
	'ebird'.`)
}

func main() {
	if len(os.Args) != 3 {
		help()
		os.Exit(0)
	}
	var fin *os.File
	var err error
	fin, err = os.Open(os.Args[1])
	if err != nil {
		fmt.Println("failed to open file", err)
		help()
		os.Exit(1)
	}
	var scn *bufio.Scanner
	scn = bufio.NewScanner(fin)
	var hasRow bool
	var values []string
	hasRow = scn.Scan()
	for hasRow == true {
		hasRow = scn.Scan()
		if hasRow == false {
			break
		}
		values = strings.Split(scn.Text(), "\t")
		_ = values
	}

}
