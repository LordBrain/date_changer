package utils

import (
	"reflect"
	"testing"
)

func TestGetMatchingTimeStrings(t *testing.T) {
	testCases := []struct {
		regex     string
		text      []byte
		expected  []string
		expectErr bool
	}{
		{
			regex:     `\d{2}-\d{2}-\d{4}`,
			text:      []byte("Date: 12-31-2022, Date2: 06-06-2020"),
			expected:  []string{"12-31-2022", "06-06-2020"},
			expectErr: false,
		},
		{
			regex:     `\d{2}:\d{2}`,
			text:      []byte("Time: 23:59, Date: 2023-01-01"),
			expected:  []string{"23:59"},
			expectErr: false,
		},
		{
			regex:     `\d{3}-\d{3}-\d{3}`,
			text:      []byte("No matching pattern here"),
			expected:  nil,
			expectErr: false,
		},
		{
			regex:     `(`, // Invalid regex
			text:      []byte("Some text"),
			expected:  nil,
			expectErr: true,
		},
	}

	for _, testCase := range testCases {
		result, err := GetMatchingTimeStrings(testCase.regex, testCase.text)

		if (err != nil) != testCase.expectErr {
			t.Errorf("Expected error: %v, but got: %v", testCase.expectErr, err != nil)
		}

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}
