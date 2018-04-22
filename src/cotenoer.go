package main

import (
	"fmt"
	"os"
	"github.com/severinraez/cotenoer/hud"
	// "github.com/severinraez/cotenoer/inventory"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "hud" {
		hud.Start()
	} else {
		fmt.Println(argsWithoutProg)
		os.Exit(0)
	}
}
