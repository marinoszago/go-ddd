package geolocation

import (
	"regexp"
)

const (
	// IPRegex contains the regex for a valid IP address.
	IPRegex = "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
)

// ContainsValidFieldForRegex checks a field based on a regex.
func ContainsValidFieldForRegex(regex, field string) (bool, error) {
	match, err := regexp.MatchString(regex, field)
	if err != nil {
		return false, err
	}

	return match, nil
}
