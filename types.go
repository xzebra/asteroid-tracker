package main

type estimatedDiameter struct {
	Min float64 `json:"estimated_diameter_min"`
	Max float64 `json:"estimated_diameter_max"`
}

func (d estimatedDiameter) Med() float64 {
	return (d.Min + d.Max) / 2
}

type diameter struct {
	Kilometers estimatedDiameter `json:"kilometers"`
	Meters     estimatedDiameter `json:"meters"`
}

type relativeVelocity struct {
	KilometersPerSecond float64 `json:"kilometers_per_second,string"`
	KilometersPerHour   float64 `json:"kilometers_per_hour,string"`
}

type approachData struct {
	Date              string           `json:"close_approach_date_full"`
	DateShort         string           `json:"close_approach_date"`
	EpochDateApproach int64            `json:"epoch_date_close_approach"`
	RelativeVelocity  relativeVelocity `json:"relative_velocity"`
	//MissDistance      `json:miss_distance`
	OrbitingBody string `json:"orbiting_body"`
}

type orbitClass struct {
	OrbitClassType string `json:"orbit_class_type"`
	OrbitClassDescription string `json:"orbit_class_description"`
	OrbitClassRange string `json:"orbit_class_range"`
}

type orbitalData struct {
	ID                   string `json:"orbit_id"`
	DeterminationDate    string `json:"orbit_determination_date"`
	FirstObservation     string `json:"first_observation_date"`
	LastObservation      string `json:"last_observation_date"`
	ArcInDays            int    `json:"data_arc_in_days"`
	Observations         int    `json:"observations_used"`
	MinOrbitIntersection string `json:"minimum_orbit_intersection"`
	Eccentricity         string `json:"eccentricity"`
	SemiMajorAxis        string `json:"semi_major_axis"`
	Inclination			 string `json:"inclination"`
	OrbitalPeriod		 string `json:"orbital_period"`
	PerihelionDistance 	string `json:"perihelion_distance"`
	PerihelionArgument string `json:"perihelion_argument"`
	AphelionDistance string	`json:"aphelion_distance"`
	PerihelionTime string `json:"perihelion_time"`
	OrbitClass orbitClass `json:"orbit_class"`
}

type asteroidData struct {
	IsHazardous  bool           `json:"is_potentially_hazardous_asteroid"`
	IsSentry     bool           `json:"is_sentry_object"`
	ID           string         `json:"neo_reference_id"`
	Name         string         `json:"name"`
	LimName      string         `json:"name_limited"`
	Designation  string         `json:"designation"`
	JPLUrl       string         `json:"nasa_jpl_url"`
	Magnitude    float64        `json:"absolute_magnitude_h"`
	Diameter     diameter       `json:"estimated_diameter"`
	ApproachData []approachData `json:"close_approach_data"`
	OrbitalData  orbitalData    `json:"orbital_data"`
	SentryData   string         `json:"sentry_data"`
}

type responseData struct {
	Objects []asteroidData
}
