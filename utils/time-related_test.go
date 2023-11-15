package utils

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/gravwell/gravwell/v3/timegrinder"
)

func TestConvertTime(t *testing.T) {
	tests := []struct {
		timestamp      string
		originalFormat string
		newFormat      string
		expectedResult string
		expectedError  error
	}{
		{
			timestamp:      "2023-11-14T13:45:00Z",
			originalFormat: "RFC3339",
			newFormat:      "UnixDate",
			expectedResult: "Tue Nov 14 13:45:00 UTC 2023",
			expectedError:  nil,
		},
		{
			timestamp:      "Mon, 02 Jan 2006 15:04:05 MST",
			originalFormat: "RFC1123",
			newFormat:      "RubyDate",
			expectedResult: "Mon Jan 02 15:04:05 +0000 2006",
			expectedError:  nil,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result, err := ConvertTime(test.timestamp, test.originalFormat, test.newFormat)

		if !reflect.DeepEqual(err, test.expectedError) {
			t.Errorf("For (%s,%s,%s), Expected error: %v, got: %v", test.timestamp, test.originalFormat, test.newFormat, test.expectedError, err)
		}

		if result != test.expectedResult {
			t.Errorf("For (%s,%s,%s), Expected: %s, got: %s", test.timestamp, test.originalFormat, test.newFormat, test.expectedResult, result)
		}
	}
}

func TestGetLayoutFromString(t *testing.T) {
	tests := []struct {
		layoutName     string
		expectedLayout string
		expectedRegex  string
		expectedError  error
	}{
		{
			layoutName:     "ANSIC",
			expectedLayout: time.ANSIC,
			expectedRegex:  timegrinder.AnsiCRegex,
			expectedError:  nil,
		},
		{
			layoutName:     "RFC822Z",
			expectedLayout: time.RFC822Z,
			expectedRegex:  timegrinder.RFC822ZRegex,
			expectedError:  nil,
		},
		{
			layoutName:     "Unknown",
			expectedLayout: "",
			expectedRegex:  "",
			expectedError:  errors.New("no matching date format"),
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		layout, regex, err := GetLayoutFromString(test.layoutName)

		if layout != test.expectedLayout || regex != test.expectedRegex || !reflect.DeepEqual(err, test.expectedError) {
			t.Errorf("For %s, Expected layout: %s, regex: %s, error: %v; Got layout: %s, regex: %s, error: %v",
				test.layoutName, test.expectedLayout, test.expectedRegex, test.expectedError, layout, regex, err)
		}
	}
}
