package main

import (
	"fmt"
	"os"
	"github.com/severinraez/cotenoer/hud"
	// "github.com/severinraez/cotenoer/inventory"
)

type arguments struct {
	Command string
}

func main() {
	arguments := parseArguments(os.Args[1:])

	if arguments.Command == "hud" {
		hud.Start()
	} else {
		fmt.Printf("Don't know what to do. Orders: %+v\n", arguments)
		os.Exit(0)
	}
}

func parseArguments(argumentArray []string) arguments {
	if len(argumentArray) == 0 {
		return arguments{}
	}

	return arguments{
		Command: argumentArray[0]}
}
