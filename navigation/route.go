package navigation

import (
	"eve-go/universe"
	"fmt"
)

type routeNode struct {
	*universe.System
	entrance *routeNode
	system   *universe.System
	exits    map[int]*routeNode
}

// Route holds information representing a set of directions through space
type Route struct {
	Origin      *routeNode
	Destination *routeNode
	Directions  []*routeNode
}

// MakeRoute takes a list of routeNodes in order and creates a Route
func MakeRoute(directions ...*routeNode) (*Route, error) {

	if len(directions) == 0 {
		return nil, fmt.Errorf("Error using MakeRoute: must provide at least 1 routeNodes as arguments")
	}

	r := Route{
		Origin:      directions[0],
		Destination: directions[len(directions)-1],
		Directions:  directions[1 : len(directions)-1],
	}

	return &r, nil
}

func aStarAlgorithm(u *universe.Universe, waypoints ...*universe.System) (*Route, error) {
	// Initialize open and closed lists
	openSystems := []universe.System{}
	visitedSystems := map[int]*routeNode{}

	// Put the starting node on the open list
	openSystems = append(openSystems, waypoints[0])

	// While the open list is not empty
	for len(openSystems) > 0 {

		// Find the node with the least f on the open list
	}

	return &Route{}, nil
}

// FindStargateRoute uses a user-specified algorithm to find a path between two systems
func FindStargateRoute(waypoints ...*universe.System) (*Route, error) {

	if len(waypoints) < 2 {
		return nil, fmt.Errorf("Error using FindStargateRoute: must provide at least 2 systems as arguments")
	}

	var r Route
	var err error

	return &r, err
}
