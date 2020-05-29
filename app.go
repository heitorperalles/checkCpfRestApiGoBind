package main

import (
  "log"
  "io/ioutil"
)

// Logging
const (
		// Logs on/off
    VERBOSE bool = false
)

// Main function
func main() {

	if VERBOSE == false {
		log.SetOutput(ioutil.Discard) // Turning logs OFF
	}

	startRouting()
}
