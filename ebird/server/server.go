package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Sync with cmd/ebird_import.go
type Obs struct {
	GlobalUniqueIdentifier    string  `json:"global_unique_identifier,omitempty"`
	LastEditedDate            string  `json:"last_edited_date,omitempty"`
	TaxonomicOrder            string  `json:"taxonomic_order,omitempty"`
	Category                  string  `json:"category,omitempty"`
	CommonName                string  `json:"common_name,omitempty"`
	ScientificName            string  `json:"scientific_name,omitempty"`
	SubspeciesCommonName      string  `json:"subspecies_common_name,omitempty"`
	SubspeciesScientificName  string  `json:"subspecies_scientific_name,omitempty"`
	ObservationCount          int     `json:"observation_count,omitempty"`
	BreedingBirdAtlasCode     string  `json:"breeding_bird_atlas_code,omitempty"`
	BreedingBirdAtlasCategory string  `json:"breeding_bird_atlas_category,omitempty"`
	AgeSex                    string  `json:"age_sex,omitempty"`
	Country                   string  `json:"country,omitempty"`
	CountryCode               string  `json:"country_code,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateCode                 string  `json:"state_code,omitempty"`
	County                    string  `json:"county,omitempty"`
	CountyCode                string  `json:"county_code,omitempty"`
	IBACode                   string  `json:"iba_code,omitempty"`
	BCRCode                   string  `json:"bcr_code,omitempty"`
	USFWSCode                 string  `json:"usfws_code,omitempty"`
	AtlasBlock                string  `json:"atlas_block,omitempty"`
	Locality                  string  `json:"locality"`
	LocalityID                string  `json:"locality_id,omitempty"`
	LocalityType              string  `json:"locality_type,omitempty"`
	Latitude                  float64 `json:"latitude,omitempty"`
	Longitude                 float64 `json:"longitude,omitempty"`
	ObservationDate           string  `json:"observation_date,omitempty"`
	TimeObservationsStarted   string  `json:"time_observations_started,omitempty"`
	ObserverID                string  `json:"observer_id,omitempty"`
	SamplingEventIdentifier   string  `json:"sampling_event_identifier,omitempty"`
	ProtocolType              string  `json:"protocol_type,omitempty"`
	ProtocolCode              string  `json:"protocol_code,omitempty"`
	ProjectCode               string  `json:"project_code,omitempty"`
	DurationMinutes           int     `json:"duration_minutes,omitempty"`
	EffortDistanceKM          float64 `json:"effort_distance_km,omitempty"`
	EffortAreaHA              float64 `json:"effort_area_ha,omitempty"`
	NumberObservers           int     `json:"number_observers,omitempty"`
	AllSpeciesReported        bool    `json:"all_species_reported,omitempty"`
	GroupIdentifier           string  `json:"group_identifier,omitempty"`
	HasMedia                  bool    `json:"has_media,omitempty"`
	Approved                  bool    `json:"approved,omitempty"`
	Reviewed                  bool    `json:"reviewed,omitempty"`
	Reason                    string  `json:"reason,omitempty"`
	TripComments              string  `json:"trip_comments"`
	SpeciesComments           string  `json:"species_comments"`
}

func obsHandler(w http.ResponseWriter, r *http.Request) {
	row := db.QueryRow(`
	SELECT common_name, age_sex, observation_count, locality, longitude, latitude,
		observation_date, species_comments
		FROM ebird
		WHERE county_code='US-ID-037' AND species_comments!=''
		ORDER BY RANDOM() LIMIT 1`)
	var obs Obs
	err := row.Scan(&obs.CommonName, &obs.AgeSex, &obs.ObservationCount,
		&obs.Locality, &obs.Longitude, &obs.Latitude, &obs.ObservationDate,
		&obs.SpeciesComments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(obs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func speciesHandler(w http.ResponseWriter, r *http.Request) {
	spp := r.FormValue("spp")
	if spp == "" {
		http.Error(w, "request URI must have a spp query parameter", http.StatusBadRequest)
		return
	}
	start := r.FormValue("start")
	_, err := time.Parse("2006-01-02", start)
	if err != nil || start == "" {
		http.Error(w, fmt.Sprintf("invalid start date: '%s'", start), http.StatusBadRequest)
		return
	}
	end := r.FormValue("end")
	_, err = time.Parse("2006-01-02", end)
	if err != nil || end == "" {
		http.Error(w, fmt.Sprintf("invalid end date: '%s'", end), http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`
	SELECT common_name, age_sex, observation_count, locality, longitude, latitude,
		observation_date, time_observations_started, species_comments
		FROM ebird JOIN taxa ON common_name=primary_common_name
		WHERE species_code=? AND observation_date>? AND observation_date<?
		GROUP BY group_identifier
		ORDER BY observation_date`, spp, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var (
		obs   Obs
		obss  []Obs
		count string
	)
	for rows.Next() {
		err := rows.Scan(&obs.CommonName, &obs.AgeSex, &count, &obs.Locality,
			&obs.Longitude, &obs.Latitude, &obs.ObservationDate,
			&obs.TimeObservationsStarted, &obs.SpeciesComments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count == "X" || count == "" {
			obs.ObservationCount = -1
		} else {
			obs.ObservationCount, _ = strconv.Atoi(count)
		}
		obss = append(obss, obs)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(obss)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	tw := &tabwriter.Writer{}
	tw.Init(w, 0, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "category", "code", "common", "sci")
	rows, err := db.Query(`SELECT category, species_code, primary_common_name, scientific_name
		FROM taxa ORDER BY species_code`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var (
		cat  string
		code string
		comm string
		sci  string
	)
	for rows.Next() {
		err = rows.Scan(&cat, &code, &comm, &sci)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", cat, code, comm, sci)
	}
	tw.Flush()
}

var (
	db *sql.DB
)

const addr = ":8888"

func main() {
	var err error
	db, err = sql.Open("sqlite3", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	mux := &http.ServeMux{}
	mux.HandleFunc("/", obsHandler)
	mux.HandleFunc("/species", speciesHandler)
	mux.HandleFunc("/codes", codeHandler)
	go func() {
		fmt.Println("open your browser to http://127.0.0.1" + addr)
	}()
	log.Fatal(http.ListenAndServe(addr, mux))
}
