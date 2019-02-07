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
	GlobalUniqueIdentifier    string
	LastEditedDate            string
	TaxonomicOrder            string
	Category                  string
	CommonName                string
	ScientificName            string
	SubspeciesCommonName      string
	SubspeciesScientificName  string
	ObservationCount          int
	BreedingBirdAtlasCode     string
	BreedingBirdAtlasCategory string
	AgeSex                    string
	Country                   string
	CountryCode               string
	State                     string
	StateCode                 string
	County                    string
	CountyCode                string
	IBACode                   string
	BCRCode                   string
	USFWSCode                 string
	AtlasBlock                string
	Locality                  string
	LocalityID                string
	LocalityType              string
	Latitude                  float64
	Longitude                 float64
	ObservationDate           string
	TimeObservationsStarted   string
	ObserverID                string
	SamplingEventIdentifier   string
	ProtocolType              string
	ProtocolCode              string
	ProjectCode               string
	DurationMinutes           int
	EffortDistanceKM          float64
	EffortAreaHA              float64
	NumberObservers           int
	AllSpeciesReported        bool
	GroupIdentifier           string
	HasMedia                  bool
	Approved                  bool
	Reviewed                  bool
	Reason                    string
	TripComments              string
	SpeciesComments           string
}

func decodeObs(vals []string) (Obs, error) {
	var birds Obs
	var err error
	if len(vals) != 47 {
		return birds, fmt.Errorf("bad values, got:%d, want: 46", len(vals))
	}
	birds.GlobalUniqueIdentifier = vals[0]
	birds.LastEditedDate = vals[1]
	birds.TaxonomicOrder = vals[2]
	birds.Category = vals[3]
	birds.CommonName = vals[4]
	birds.ScientificName = vals[5]
	birds.SubspeciesCommonName = vals[6]
	birds.SubspeciesScientificName = vals[7]
	if vals[8] == "X" {
		birds.ObservationCount = -1
	} else {
		birds.ObservationCount, err = strconv.Atoi(vals[8])
		if err != nil {
			return birds, err
		}
	}
	birds.BreedingBirdAtlasCode = vals[9]
	birds.BreedingBirdAtlasCategory = vals[10]
	birds.AgeSex = vals[11]
	birds.Country = vals[12]
	birds.CountryCode = vals[13]
	birds.State = vals[14]
	birds.StateCode = vals[15]
	birds.County = vals[16]
	birds.CountyCode = vals[17]
	birds.IBACode = vals[18]
	birds.BCRCode = vals[19]
	birds.USFWSCode = vals[20]
	birds.AtlasBlock = vals[21]
	birds.Locality = vals[22]
	birds.LocalityID = vals[23]
	birds.LocalityType = vals[24]
	if vals[25] != "" {
		birds.Latitude, err = strconv.ParseFloat(vals[25], 64)
		if err != nil {
			return birds, err
		}
	}
	if vals[26] != "" {
		birds.Longitude, err = strconv.ParseFloat(vals[26], 64)
		if err != nil {
			fmt.Println(vals[21:27])
			return birds, err
		}
	}
	birds.ObservationDate = vals[27]
	birds.TimeObservationsStarted = vals[28]
	birds.ObserverID = vals[29]
	birds.SamplingEventIdentifier = vals[30]
	birds.ProtocolType = vals[31]
	birds.ProtocolCode = vals[32]
	birds.ProjectCode = vals[33]
	if vals[34] != "" {
		birds.DurationMinutes, err = strconv.Atoi(vals[34])
		if err != nil {
			return birds, err
		}
	}

	if vals[35] != "" {
		birds.EffortDistanceKM, err = strconv.ParseFloat(vals[35], 64)
		if err != nil {
			return birds, err
		}
	}
	if vals[36] != "" {
		birds.EffortAreaHA, err = strconv.ParseFloat(vals[36], 64)
		if err != nil {
			return birds, err
		}
	}
	if vals[37] != "" {
		birds.NumberObservers, err = strconv.Atoi(vals[37])
		if err != nil {
			return birds, err
		}
	}
	birds.AllSpeciesReported = vals[38] == "1"
	birds.GroupIdentifier = vals[39]
	birds.HasMedia = vals[40] == "1"
	birds.Approved = vals[41] == "1"
	birds.Reviewed = vals[42] == "1"
	birds.Reason = vals[43]
	birds.TripComments = vals[44]
	birds.SpeciesComments = vals[45]
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
	db, err = sql.Open("sqlite3", os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE ebird (
	guid TEXT, last_edited_date TEXT, taxonomic_order TEXT, category TEXT,
	common_name TEXT, scientific_name TEXT, subspecies_common_name TEXT,
	subspecies_scientific_name TEXT, observation_count INTEGER,
	breeding_bird_atlas_code TEXT, breeding_bird_atlas_category TEXT, age_sex TEXT,
	country TEXT, country_code TEXT, state TEXT, state_code TEXT, county TEXT,
	county_code TEXT, iba_code TEXT, bcr_code TEXT, usfws_code TEXT, atlas_block TEXT,
	locality TEXT, locality_id TEXT, locality_type TEXT, latitude float64,
	longitude float64, observation_date TEXT, time_observations_started TEXT,
	observer_id TEXT, sampling_event_identifier TEXT, protocol_type TEXT,
	protocol_code TEXT, project_code TEXT, duration_minutes INTEGER,
	effort_distance_km float64, effort_area_ha float64, number_observers INTEGER,
	all_species_reported INTEGER, group_identifier TEXT, has_media INTEGER,
	approved INTEGER, reviewed INTEGER, reason TEXT, trip_comments TEXT,
	species_comments TEXT);`)

	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(`INSERT INTO ebird VALUES(?,?,?,?,?,?,?,?, ?,?,?,?,
		?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
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
		_, err = stmt.Exec(values[0], values[1], values[2], values[3], values[4],
			values[5], values[6], values[7], values[8], values[9], values[10], values[11],
			values[12], values[13], values[14], values[15], values[16], values[17], values[18],
			values[19], values[20], values[21], values[22], values[23], values[24], values[25],
			values[26], values[27], values[28], values[29], values[30], values[31], values[32],
			values[33], values[34], values[35], values[36], values[37], values[38], values[39],
			values[40], values[41], values[42], values[43], values[44], values[45],
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
