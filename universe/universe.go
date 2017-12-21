package universe

import (
	"eve-go/esi"
	"fmt"
)

const universeAPIEndpoint = "universe/systems/"

// Universe is composed of star systems
// It contains a map from SystemID to System
type Universe struct {
	apiInfo esi.APIInfo
	systems map[int]*System
	// moons map[int]*Moons
	// planets map[int]*Planets
	// regions map[int]*Regions
}

// universe is the singleton instance
var entireUniverse *Universe

// GetUniverse implements a singleton pattern
// It returns entireUniverse, and loads it from the API
// if necessary
func GetUniverse() (*Universe, error) {

	var err error

	if entireUniverse == nil {
		entireUniverse = &Universe{apiInfo: esi.MakeAPIInfo(universeAPIEndpoint), systems: make(map[int]*System)}
		err = entireUniverse.UpdateUniverse()
	}

	return entireUniverse, err

}

// UpdateUniverse refreshes the universe's system list from the API
func (u *Universe) UpdateUniverse() error {

	var systemIDs []int // Create empty integer struct to fill with JSON
	args := ""

	err := esi.Get(u, &systemIDs, args)
	if err != nil {
		return err
	}

	newSystemMap := make(map[int]*System)

	// For every system,
	for _, systemID := range systemIDs {
		val, ok := u.systems[systemID]
		if ok && val != nil { // If the system is mapped and loaded, update it and insert in new map
			val.UpdateSystem()
			newSystemMap[systemID] = val
		} else { // If the system does not exist in the map, insert nil in new map
			newSystemMap[systemID] = nil
		}
	}

	u.systems = newSystemMap

	return nil
}

// GetSystem returns a loaded system if it is in the universe
// If systemID is not in the universe, it returns an error
func (u *Universe) GetSystem(systemID int) (*System, error) {

	var err error

	s, ok := u.systems[systemID]

	if !ok {
		err = fmt.Errorf("Universe does not contain system with [systemID: %v]", systemID)
		return nil, err
	}

	if s == nil {
		if s, err = MakeSystem(systemID); err == nil {
			u.systems[systemID] = s
		} else {
			return nil, err
		}
	}
	return s, nil
}

// APIInfo returns the apiInfo object
func (u *Universe) APIInfo() *esi.APIInfo {
	return &u.apiInfo
}
