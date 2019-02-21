.timer on
CREATE TABLE ebird (
	global_unique_identifier TEXT,
	last_edited_date TEXT,
	taxonomic_order TEXT,
	category TEXT,
	common_name TEXT,
	scientific_name TEXT,
	subspecies_common_name TEXT,
	subspecies_scientific_name TEXT,
	observation_count INTEGER,
	breeding_bird_atlas_code TEXT,
	breeding_bird_atlas_category TEXT,
	age_sex TEXT,
	country TEXT,
	country_code TEXT,
	state TEXT,
	state_code TEXT,
	county TEXT,
	county_code TEXT,
	iba_code TEXT,
	bcr_code TEXT,
	usfws_code TEXT,
	atlas_block TEXT,
	locality TEXT,
	locality_id TEXT,
	locality_type TEXT,
	latitude FLOAT,
	longitude FLOAT,
	observation_date TEXT,
	time_observations_started TEXT,
	observer_id TEXT,
	sampling_event_identifier TEXT,
	protocol_type TEXT,
	protocol_code TEXT,
	project_code TEXT,
	duration_minutes INTEGER,
	effort_distance_km FLOAT,
	effort_area_ha FLOAT,
	number_observers INTEGER,
	all_species_reported TEXT,
	group_identifier TEXT,
	has_media INTEGER,
	approved INTEGER,
	reviewed INTEGER,
	reason TEXT,
	trip_comments TEXT,
	species_comments TEXT
);

/* Taxa for lookups of species codes */
CREATE TABLE taxa (
	taxon_order INTEGER,
	category TEXT,
	species_code TEXT,
	primary_common_name TEXT,
	scientific_name TEXT,
	species_order TEXT,
	family TEXT,
	species_group TEXT,
	report_as TEXT
);

.mode csv
.header on
.import taxa.csv taxa
.sep \t

BEGIN;
.import ebird.txt tmp_ebird
END;

INSERT INTO ebird SELECT
	"GLOBAL UNIQUE IDENTIFIER",
	"LAST EDITED DATE",
	"TAXONOMIC ORDER",
	"CATEGORY",
	"COMMON NAME",
	"SCIENTIFIC NAME",
	"SUBSPECIES COMMON NAME",
	"SUBSPECIES SCIENTIFIC NAME",
	"OBSERVATION COUNT",
	"BREEDING BIRD ATLAS CODE",
	"BREEDING BIRD ATLAS CATEGORY",
	"AGE/SEX",
	"COUNTRY",
	"COUNTRY CODE",
	"STATE",
	"STATE CODE",
	"COUNTY",
	"COUNTY CODE",
	"IBA CODE",
	"BCR CODE",
	"USFWS CODE",
	"ATLAS BLOCK",
	"LOCALITY",
	"LOCALITY ID",
	"LOCALITY TYPE",
	"LATITUDE",
	"LONGITUDE",
	"OBSERVATION DATE",
	"TIME OBSERVATIONS STARTED",
	"OBSERVER ID",
	"SAMPLING EVENT IDENTIFIER",
	"PROTOCOL TYPE",
	"PROTOCOL CODE",
	"PROJECT CODE",
	"DURATION MINUTES",
	"EFFORT DISTANCE KM",
	"EFFORT AREA HA",
	"NUMBER OBSERVERS",
	"ALL SPECIES REPORTED",
	"GROUP IDENTIFIER",
	"HAS MEDIA",
	"APPROVED",
	"REVIEWED",
	"REASON",
	"TRIP COMMENTS",
	"SPECIES COMMENTS" FROM tmp_ebird;

DROP TABLE tmp_ebird;

UPDATE ebird SET observation_count=-1 WHERE observation_count='X';

CREATE INDEX idx_spp_date ON ebird(common_name, observation_date);
CREATE INDEX idx_spp_code ON taxa(species_code);
