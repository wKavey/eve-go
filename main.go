package main

import (
	"eve-go/universe"
	"fmt"
	"os"
)

func main() {

	eveUniverse, err := universe.MakeUniverse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s, err := eveUniverse.GetSystem(30000142)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(s)

}
