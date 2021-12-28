package main

import (
	"fmt"
)

var (
	// The git commit that was compiled. These will be filled in by the compiler.
	GitCommit string

	// The main version number that is being run at the moment.
	Version string

	// BuildDate contains the date and time of build process.
	BuildDate string
)

// GetHumanVersion composes the parts of the version in a way that's suitable
// for displaying to humans.
func GetHumanVersion() {
	fmt.Printf("secrets-init %s, commit %s, built at %s", Version, GitCommit, BuildDate)
}
