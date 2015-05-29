// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package process

// LaunchDetails holds information about an existing process as provided by
// a workload process plugin.
type LaunchDetails struct {
	// UniqueID is provided by the plugin as a guaranteed way
	// to identify the process to the plugin.
	UniqueID string

	// Status is the status of the process as reported by the plugin.
	Status string
}

// ParseDetails parses the input string in to a LaunchDetails struct.
func ParseDetails(input string) (LaunchDetails, error) {
	// TODO(ericsnow) Finish!
	return LaunchDetails{}, nil
}
