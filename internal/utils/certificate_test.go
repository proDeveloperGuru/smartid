package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Certificate_Extract(t *testing.T) {
	value := "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"

	tests := []struct {
		name     string
		value    string
		expected *Person
		error    bool
	}{
		{
			name:  "Success",
			value: value,
			expected: &Person{
				IdentityNumber: "PNOEE-30303039914",
				PersonalCode:   "30303039914",
				FirstName:      "TESTNUMBER",
				LastName:       "OK",
			},
			error: false,
		},
		{
			name:     "Error: invalid certificate format",
			value:    "c29tZSBpbmNvcnJlY3QgY2VydGlmaWNhdGU=",
			expected: nil,
			error:    true,
		},
		{
			name:     "Error: invalid base64",
			value:    "invalid-base64",
			expected: nil,
			error:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(tt.value)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
