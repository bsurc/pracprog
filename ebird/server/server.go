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

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
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
	mux.HandleFunc("/", htmlHandler)
	mux.HandleFunc("/obs", obsHandler)
	mux.HandleFunc("/species", speciesHandler)
	mux.HandleFunc("/codes", codeHandler)
	go func() {
		fmt.Println("open your browser to http://127.0.0.1" + addr)
	}()
	log.Fatal(http.ListenAndServe(addr, mux))
}
