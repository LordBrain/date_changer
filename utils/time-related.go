package utils

import (
	"errors"
	"fmt"
	"os"
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
	timeObj, err := time.Parse(originalFormatLayout, timestamp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
		ansicFixedRegex := `[A-Za-z]{3}\s+[(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Sept|Oct|Nov|Dec)]+\s+\d{1,2}\s+\d\d:\d\d:\d\d\s+\d{4}`
		return time.ANSIC, ansicFixedRegex, nil
	case "UNIXDATE":
		unixdateFixedRegex := `[A-Za-z]{3}\s+[(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Sept|Oct|Nov|Dec)]+\s+\d{1,2}\s+\d\d:\d\d:\d\d\s+[A-Z]{3}\s+\d{4}`
		return time.UnixDate, unixdateFixedRegex, nil
	case "RUBYDATE":
		rubydateFixedRegex := `[A-Za-z]{3}\s+[(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Sept|Oct|Nov|Dec)]+\s+\d{1,2}\s+\d\d:\d\d:\d\d\s+[\-|\+]\d{4}\s+\d{4}`
		return time.RubyDate, rubydateFixedRegex, nil
	case "RFC822":
		return time.RFC822, timegrinder.RFC822Regex, nil
	case "RFC822Z":
		return time.RFC822Z, timegrinder.RFC822ZRegex, nil
	case "RFC850":
		rfc850FixedRegex := `[Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday]+\,+\s\d{2}\-[JFMASOND][anebriyunlgpctov]+\-\d{2}\s\d\d:\d\d:\d\d\s[A-Z]{3}`
		return time.RFC850, rfc850FixedRegex, nil
	case "RFC1123":
		rfc1123FixedRegex := `[Mon|Tue|Wed|Thu|Fri|Sat|Sun]+\,+\s\d{2} [JFMASOND][anebriyunlgpctov]+ \d{4}\s\d\d:\d\d:\d\d\s[A-Z]{3}`
		return time.RFC1123, rfc1123FixedRegex, nil
	case "RFC1123Z":
		rfc1123zFixedRegex := `[Mon|Tue|Wed|Thu|Fri|Sat|Sun]+\,+\s\d{2} [JFMASOND][anebriyunlgpctov]+ \d{4}\s\d\d:\d\d:\d\d\s[\-|\+]\d{4}`
		return time.RFC1123Z, rfc1123zFixedRegex, nil
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
