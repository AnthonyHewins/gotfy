package gotfy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewMarshalJSON(mainTest *testing.T) {
	uri, err := url.Parse("https://docs.ntfy.sh/publish/#icons")
	if err != nil {
		mainTest.Fatalf("failed setting up test URL: %v", err)
	}

	testCases := []struct {
		name        string
		arg         ViewAction
		expected    string
		expectedErr error
	}{
		{
			name:     "base case",
			expected: `{"action":"view","label":""}`,
		},
		{
			name:     "only label",
			arg:      ViewAction{Label: "label"},
			expected: `{"action":"view","label":"label"}`,
		},
		{
			name:     "url",
			arg:      ViewAction{Link: uri},
			expected: `{"action":"view","label":"","url":"https://docs.ntfy.sh/publish/#icons"}`,
		},
		{
			name: "clear",
			arg: ViewAction{
				Clear: true,
			},
			expected: `{"action":"view","label":"","clear":true}`,
		},
		{
			name: "everything",
			arg: ViewAction{
				Label: "label",
				Link:  uri,
				Clear: true,
			},
			expected: `{"action":"view","label":"label","url":"https://docs.ntfy.sh/publish/#icons","clear":true}`,
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := tc.arg.MarshalJSON()
		t.Equal([]byte(tc.expected), actual, tc.name)
		t.Equal(tc.expectedErr, actualErr, tc.name)
	}
}
