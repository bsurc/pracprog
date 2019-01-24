// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Println(`
ebird_import ebird_data.csv output.db

	ebird_import reads ebird data from the supplied csv file
	with ebird data, and populates a SQLite3 database.  It does
	not create any indices or any other side effects.  Only
	CREATE and INSERT SQL commands are issued.  The table is named
	'ebird'.
`)
}

func main() {
	if len(os.Args) != 3 {
		help()
		os.Exit(0)
	}
}
