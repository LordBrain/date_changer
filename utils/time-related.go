package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/gravwell/gravwell/v3/timegrinder"
)

// ConvertTime returns the time for the new time layout
func ConvertTime(timestamp, originalFormat, newFormat string) (string, error) {
	// Get the time format for the original date
	originalFormatLayout, _, err := GetLayoutFromString(originalFormat)
	if err != nil {
		return "", err
	}
	// Get the time format for the new date
	newFormatLayout, _, err := GetLayoutFromString(newFormat)
	if err != nil {
		return "", err
	}
	// Parse the original date
	timeObj, _ := time.Parse(originalFormatLayout, timestamp)
	// Convert the date to the new format
	newTime := timeObj.Format(newFormatLayout)
	// Return new date
	return newTime, nil

}

// GetLayoutFromString returns the time layout and the regex to match that time layout
func GetLayoutFromString(layoutName string) (string, string, error) {
	// Set it to upper, just in case.
	switch strings.ToUpper(layoutName) {
	case "ANSIC":
		return time.ANSIC, timegrinder.AnsiCRegex, nil
	case "UNIXDATE":
		return time.UnixDate, timegrinder.UnixRegex, nil
	case "RUBYDATE":
		return time.RubyDate, timegrinder.RubyRegex, nil
	case "RFC822":
		return time.RFC822, timegrinder.RFC822Regex, nil
	case "RFC822Z":
		return time.RFC822Z, timegrinder.RFC822ZRegex, nil
	case "RFC850":
		return time.RFC850, timegrinder.RFC850Regex, nil
	case "RFC1123":
		return time.RFC1123, timegrinder.RFC1123Regex, nil
	case "RFC1123Z":
		return time.RFC1123Z, timegrinder.RFC1123ZRegex, nil
	case "RFC3339":
		return time.RFC3339, timegrinder.RFC3339Regex, nil
	case "RFC3339NANO":
		return time.RFC3339Nano, timegrinder.RFC3339NanoRegex, nil
	case "DATETIME":
		return time.DateTime, timegrinder.UnpaddedDateTimeRegex, nil
	default:
		return "", "", errors.New("no matching date format")
	}

}
