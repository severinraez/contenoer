package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/severinraez/cotenoer/hud"
	"github.com/severinraez/cotenoer/inventory"
)

type arguments struct {
	Command string
	Parameters []string
	SessionFile string
}

const sessionEnvVar = "COTENOER_SESSION"

func main() {
	arguments := parseArguments(os.Args[1:], os.Getenv(sessionEnvVar))

	if arguments.Command == "hud" {
		hud.Start()
	} else if arguments.Command == "add" {
		session := loadSession(arguments.SessionFile)

		usage("add NAME DOCKERFILE", 2, arguments.Parameters)

		name, path := arguments.Parameters[0], arguments.Parameters[1]

		session, err := inventory.Add(session, name, path)
		exitOnError(err)

		saveSession(arguments.SessionFile, session)
	} else {
		fmt.Printf("Don't know what to do. Orders: %+v\n", arguments)
		os.Exit(1)
	}
}

func usage(description string, parameterCount int, parameters []string) {
	if len(parameters) != parameterCount {
		fmt.Printf("Usage: contenör %s", description)
		os.Exit(1)
	}
}

func exitOnError(err error) {
	if err == nil { return }

	fmt.Println("Error, aborting: %v", err)
	os.Exit(2)
}

func parseArguments(argumentArray []string, sessionFile string) arguments {
	if len(argumentArray) == 0 {
		return arguments{
			SessionFile: sessionFile}
	}

	return arguments{
		Command: argumentArray[0],
		Parameters: argumentArray[1:],
		SessionFile: sessionFile}
}

func loadSession(sessionFile string) inventory.Inventory {
	_, err := os.Stat(sessionFile)

	if err == nil {
		contents, err := ioutil.ReadFile(sessionFile)
		exitOnError(err)

		session, err := inventory.Deserialize(contents)
		exitOnError(err)

		return session
	}

	fmt.Println("Initializing new Session at %s", sessionFile)

	return inventory.New()
}

func saveSession(sessionFile string, session inventory.Inventory) {
	content, err := inventory.Serialize(session)
	exitOnError(err)

	err = ioutil.WriteFile(sessionFile, content, 0644)
	exitOnError(err)
}
