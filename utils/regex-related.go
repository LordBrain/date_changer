package utils

import (
	"errors"
	"regexp"
)

// GetMatchingTimeStrings returns a list of timestamps that match the regex
func GetMatchingTimeStrings(regex string, text []byte) ([]string, error) {
	var listOfTimes []string
	// Setup regex to be used
	regexFilter, err := regexp.Compile(regex)
	if err != nil {
		// Error with creating the regex. Return a empty list and the error
		return listOfTimes, errors.New("issue with regex")
	}
	// Use regex to match all the date formats
	listOfTimes = regexFilter.FindAllString(string(text), -1)
	// Return the list of matching dates
	return listOfTimes, nil
}
