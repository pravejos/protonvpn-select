package main

import (
	"flag"
	"fmt"
	"os"
)

var verbose = flag.Bool("v", false, "verbose")

func main() {
	flag.Parse()

	debug("Application started")

	if !isInitialized() {
		debug("Application not initialized. Initializing.")
		initialize()
	}

	logicals, err := getLogicals()
	if err != nil {
		fatal("Unable to fetch server list. Error - %s", err)
		os.Exit(2)
	}

	fmt.Println(filter(logicals))

	debug("Application stopped")
}
