package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
)

func main() {
	// Define the flags.
	versionFlag := flag.String("v", "", "Version number (required)")
	methodFlag := flag.String("m", "", "Update method: major|minor|patch (required)")

	// Parse the flags.
	flag.Parse()

	// Check if both flags are provided.
	if *versionFlag == "" || *methodFlag == "" {
		fmt.Fprintln(os.Stderr, "Both -v (version) and -m (method) flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	currentVersion, err := semver.NewVersion(*versionFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid version: %s\n", *versionFlag)
		os.Exit(2)
	}

	// Ensure the methodFlag is either 'major', 'minor', or 'patch'.
	switch *methodFlag {

	case "major", "minor", "patch":
		var nextVersion semver.Version
		switch *methodFlag {
		case "major":
			nextVersion = currentVersion.IncMajor() // Assign the incremented version to nextVersion.
		case "minor":
			nextVersion = currentVersion.IncMinor() // Assign the incremented version to nextVersion.
		case "patch":
			nextVersion = currentVersion.IncPatch() // Assign the incremented version to nextVersion.
		}

		fmt.Printf("version.current=%s\n", currentVersion.String())
		fmt.Printf("version.next=%s\n", nextVersion.String())
		fmt.Printf("method=%s\n", *methodFlag)
	default:
		fmt.Fprintf(os.Stderr, "invalid method: %s\n", *methodFlag)
		fmt.Fprintln(os.Stderr, "Please specify one of the following for -m flag: major, minor, or patch.")
		os.Exit(2)
	}
}
