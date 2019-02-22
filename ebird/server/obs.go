package main

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
