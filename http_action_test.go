package gotfy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPActionMarshal(mainTest *testing.T) {
	path := "https://github.com/AnthonyHewins/gotfy"

	testCases := []struct {
		name        string
		arg         HttpAction[string]
		expected    string
		expectedErr error
	}{
		{
			name:     "base case",
			expected: `{"action":"http","label":"","url":""}`,
		},
		{
			name:     "label",
			arg:      HttpAction[string]{Label: "label"},
			expected: `{"action":"http","label":"label","url":""}`,
		},
		{
			name:     "url",
			arg:      HttpAction[string]{URL: path},
			expected: `{"action":"http","label":"","url":"https://github.com/AnthonyHewins/gotfy"}`,
		},
		{
			name:     "method",
			arg:      HttpAction[string]{Method: "method"},
			expected: `{"action":"http","label":"","method":"method","url":""}`,
		},
		{
			name:     "headers",
			arg:      HttpAction[string]{Headers: map[string]string{"header": "val"}},
			expected: `{"action":"http","headers":{"header":"val"},"label":"","url":""}`,
		},
		{
			name: "body",
			arg: HttpAction[string]{
				Body: "body",
			},
			expected: `{"action":"http","body":"IlwiYm9keVwiIg==","label":"","url":""}`,
		},
		{
			name: "clear",
			arg: HttpAction[string]{
				Clear: true,
			},
			expected: `{"action":"http","clear":true,"label":"","url":""}`,
		},
		{
			name: "everything",
			arg: HttpAction[string]{
				Label:   "label",
				URL:     path,
				Method:  "method",
				Headers: map[string]string{"header": "val"},
				Body:    "body",
				Clear:   true,
			},
			expected: `{"action":"http","body":"IlwiYm9keVwiIg==","clear":true,"headers":{"header":"val"},"label":"label","method":"method","url":"https://github.com/AnthonyHewins/gotfy"}`,
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := tc.arg.MarshalJSON()
		t.Equal([]byte(tc.expected), actual, tc.name)
		t.Equal(tc.expectedErr, actualErr, tc.name)
	}
}
