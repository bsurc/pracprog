// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Obs struct {
	GlobalUniqueIdentifier       string
	Last_edited_date             string
	Taxonomic_order              string
	Category                     string
	Common_name                  string
	Scientific_name              string
	Subspecies_common_name       string
	Subspecies_scientific_name   string
	Observation_count            string
	Breeding_bird_atlas_code     string
	Breeding_bird_atlas_category string
	AgeSex                       string
	Country                      string
	Country_code                 string
	State                        string
	State_code                   string
	County                       string
	County_code                  string
	Iba_code                     string
	Bcr_code                     string
	Usfws_code                   string
	Atlas_block                  string
	Locality                     string
	Locality_id                  string
	Locality_type                string
	Latitude                     float64
	Longitude                    float64
	Observation_date             string
	Time_observations_started    string
	Observer_id                  string
	Sampling_event_identifier    string
	Protocol_type                string
	Protocol_code                string
	Project_code                 string
	Duration_minutes             int
	Effort_distance_km           float64
	Effort_area_ha               float64
	Number_observers             int
	All_species_reported         bool
	Group_identifier             string
	Has_media                    bool
	Approved                     bool
	Reviewed                     bool
	Reason                       string
	Trip_comments                string
	Species_comments             string
}

func decodeObs(vals []string) (Obs, error) {
	var birds Obs
	var err error
	if len(vals) != 43 {
		return birds, fmt.Errorf("bad values, got:%d, want: 33", len(vals))
	}
	birds.GlobalUniqueIdentifier = vals[0]
	birds.Last_edited_date = vals[1]
	birds.Taxonomic_order = vals[2]
	birds.Category = vals[3]
	birds.Common_name = vals[4]
	birds.Scientific_name = vals[5]
	birds.Subspecies_common_name = vals[6]
	birds.Subspecies_scientific_name = vals[7]
	birds.Observation_count = vals[8]
	birds.Breeding_bird_atlas_code = vals[9]
	birds.Breeding_bird_atlas_category = vals[10]
	birds.AgeSex = vals[11]
	birds.Country = vals[12]
	birds.Country_code = vals[13]
	birds.State = vals[14]
	birds.State_code = vals[15]
	birds.County = vals[16]
	birds.County_code = vals[17]
	birds.Iba_code = vals[18]
	birds.Bcr_code = vals[19]
	birds.Usfws_code = vals[20]
	birds.Atlas_block = vals[21]
	birds.Locality = vals[22]
	birds.Locality_id = vals[23]
	birds.Locality_type = vals[24]
	birds.Latitude, err = strconv.ParseFloat(vals[25], 64)
	if err != nil {
		return birds, err
	}
	/*
		birds.Longitude, err = strconv.ParseFloat(vals[24], 64)
		if err != nil {
			return birds, err
		}
		birds.Observation_date = vals[25]
		birds.Time_observations_started = vals[26]
		birds.Observer_id = vals[27]
		birds.Sampling_event_identifier = vals[28]
		birds.Protocol_type = vals[29]
		birds.Protocol_code = vals[30]
		birds.Project_code = vals[31]
		birds.Duration_minutes, err = strconv.Atoi(vals[32])
		if err != nil {
			return birds, err
		}
			Effort_distance_km           float64
			Effort_area_ha               float64
			Number_observers             int
			All_species_reported         bool
			Group_identifier             string
			Has_media                    bool
			Approved                     bool
			Reviewed                     bool
			Reason                       string
			Trip_comments                string
			Species_comments             string
	*/
	return birds, nil

}

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

	var db *sql.DB
	db, err = sql.Open("sqlite3", "ebird.db")
	if err != nil {
		log.Fatal(err)
	}

	var sql string
	sql = `CREATE TABLE IF NOT EXISTS test_table (
		first_name TEXT,
		last_name TEXT,
		age INTEGER,
		height FLOAT
		)`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	sql = `INSERT INTO test_table(first_name, last_name, age, height)
			 VALUES(?,?,?,?)`
	stmt, err := db.Prepare(sql)
	fmt.Printf("%T\n", stmt)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("kyle", "shannon", 39, 5.75)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	var scn *bufio.Scanner
	scn = bufio.NewScanner(fin)
	var hasRow bool
	var values []string
	hasRow = scn.Scan()
	//var nextObs Obs
	for hasRow == true {
		hasRow = scn.Scan()
		if hasRow == false {
			break
		}
		values = strings.Split(scn.Text(), "\t")
		_ = values

	}

}
