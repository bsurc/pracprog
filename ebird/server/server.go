package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"golang.org/x/crypto/acme/autocert"
)

func obsHandler(w http.ResponseWriter, r *http.Request) {

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
	}
	var mbr [4]float64
	for i, x := range []string{"minx", "maxx", "miny", "maxy"} {
		v := r.FormValue(x)
		mbr[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid coordinate for %s: '%s'", x, v), http.StatusBadRequest)
			return
		}
	}
	c := pool.Get(context.Background())
	defer pool.Put(c)
	var stmt *sqlite.Stmt
	stmt := c.Prep(`SELECT
		common_name,
		age_sex,
		observation_count,
		locality,
		longitude,
		latitude,
		observation_date,
		species_comments
			FROM ebird JOIN taxa ON common_name=primary_common_name
		WHERE species_code=? AND observation_date>? AND observation_date<?
		AND longitude>? AND longitude<? AND latitude>? AND latitude<?`)
	stmt.BindText(1, spp)
	stmt.BindText(2, start)
	stmt.BindText(3, end)
	stmt.BindFloat(4, mbr[0])
	stmt.BindFloat(5, mbr[1])
	stmt.BindFloat(6, mbr[2])
	stmt.BindFloat(7, mbr[3])
	type obs struct {
		CommonName string  `json:"common_name"`
		AgeSex     string  `json:"age_sex"`
		Count      int64   `json:"count"`
		Locality   string  `json:"locality"`
		Longitude  float64 `json:"longitude"`
		Latitude   float64 `json:"latitude"`
		Date       string  `json:"date"`
		Comments   string  `json:"comments"`
	}
	var os []obs

	for {
		if hasRow, err := stmt.Step(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if !hasRow {
			break
		}
		os = append(os, obs{
			stmt.ColumnText(0),
			stmt.ColumnText(1),
			stmt.ColumnInt64(2),
			stmt.ColumnText(3),
			stmt.ColumnFloat(4),
			stmt.ColumnFloat(5),
			stmt.ColumnText(6),
			stmt.ColumnText(7),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(os)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func codeHandler(w http.ResponseWriter, r *http.Response) {
	c := pool.Get(context.Background())
	defer pool.Put(c)
	stmt := c.Prep(`SELECT
		common_name,
		age_sex,
		observation_count,
		locality,
		longitude,
		latitude,
		observation_date,
		species_comments
			FROM ebird JOIN taxa ON common_name=primary_common_name
		WHERE species_code=? AND observation_date>? AND observation_date<?`)
	_ = stmt
}

var (
	pool   *sqlitex.Pool
	uptime time.Time
)

func main() {
	flagAddr := flag.String("addr", ":8888", "HTTP service address")
	flag.Parse()

	var err error
	pool, err = sqlitex.Open("ebird.db", sqlite.SQLITE_OPEN_READONLY, 10)
	if err != nil {
		log.Fatal(err)
	}

	mux := &http.ServeMux{}
	mux.HandleFunc("/obs", obsHandler)
	mux.HandleFunc("/species", speciesHandler)
	uptime = time.Now()
	srv := &http.Server{
		Addr:         *flagAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	if *flagAddr == ":443" || *flagAddr == ":https" {
		m := &autocert.Manager{
			Cache:      autocert.DirCache("/opt/acme/"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("wright-lxt-01.boisestate.edu"),
		}
		go func() {
			log.Fatal(http.ListenAndServe(":http", m.HTTPHandler(nil)))
		}()
		srv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		log.Fatal(srv.ListenAndServeTLS(*flagAddr, ""))
	} else {
		go func() {
			fmt.Println("open your browser to http://127.0.0.1" + *flagAddr)
		}()
		log.Fatal(srv.ListenAndServe())
	}
}
