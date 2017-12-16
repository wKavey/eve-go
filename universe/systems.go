package systems

import "fmt"

// System - a solar system
type System struct {
	SystemID int    `json:"system_id"`
	Name     string `json:"name"`
	Position struct {
		X int64 `json:"x"`
		Y int64 `json:"y"`
		Z int64 `json:"z"`
	} `json:"position"`
	SecurityStatus  float64 `json:"security_status"`
	ConstellationID int     `json:"constellation_id"`
	Planets         []struct {
		PlanetID int   `json:"planet_id"`
		Moons    []int `json:"moons,omitempty"`
	} `json:"planets"`
	SecurityClass string `json:"security_class"`
	Stargates     []int  `json:"stargates"`
	Stations      []int  `json:"stations"`
}
