package universe

import (
	"eve-go/common"
	"eve-go/esi"
	"strconv"
)

const systemAPIEndpoint = "universe/systems/"

type rawSystem struct {
	StarID   int    `json:"star_id"`
	SystemID int    `json:"system_id"`
	Name     string `json:"name"`
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"position"`
	SecurityStatus  float64 `json:"security_status"`
	ConstellationID int     `json:"constellation_id"`
	Planets         []struct {
		PlanetID int   `json:"planet_id"`
		Moons    []int `json:"moons"`
	} `json:"planets"`
	SecurityClass string `json:"security_class"`
	Stargates     []int  `json:"stargates"`
}

// System holds meaningful information regarding a system
type System struct {
	apiInfo        esi.APIInfo
	starID         int
	SystemID       int
	Name           string
	Position       common.Position
	SecurityStatus float64
	//Constellation  ConstellationData
	//Planets        []PlanetData
	SecurityClass string
	Stargates     []int
}

// MakeSystem creates a system and sets its SystemID and apiInfo.initialized = false
func MakeSystem(systemID int) (*System, error) {
	s := System{SystemID: systemID, apiInfo: esi.MakeAPIInfo(systemAPIEndpoint)}
	err := s.UpdateSystem()
	return &s, err
}

// UpdateSystem refreshes the system's information from the API
func (s *System) UpdateSystem() error {

	var err error

	rawData := rawSystem{}           // Create empty struct to fill with JSON
	args := strconv.Itoa(s.SystemID) // system endpoint takes SystemID

	err = esi.Get(s, &rawData, args)
	if err != nil {
		return err
	}

	s.starID = rawData.StarID
	s.Name = rawData.Name
	s.Position = common.Position{
		X: rawData.Position.X,
		Y: rawData.Position.Y,
		Z: rawData.Position.Z,
	}
	s.SecurityClass = rawData.SecurityClass
	s.Stargates = rawData.Stargates

	return err
}

// APIInfo returns the apiInfo object
func (s *System) APIInfo() *esi.APIInfo {
	return &s.apiInfo
}
