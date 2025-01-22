package gotfy

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageMarshalJSON(mainTest *testing.T) {
	testCases := []struct {
		name        string
		arg         Message
		expected    string
		expectedErr error
	}{
		{
			name:     "base case",
			expected: `{"topic":""}`,
		},
		{
			name:     "topic",
			arg:      Message{Topic: "topic"},
			expected: `{"topic":"topic"}`,
		},
		{
			name:     "Message",
			arg:      Message{Message: "Message"},
			expected: `{"topic":"","message":"Message"}`,
		},
		{
			name:     "Title",
			arg:      Message{Title: "Title"},
			expected: `{"topic":"","title":"Title"}`,
		},
		{
			name:     "Tags",
			arg:      Message{Tags: []string{"tag1", "tag2"}},
			expected: `{"topic":"","tags":["tag1","tag2"]}`,
		},
		{
			name:     "Priority negative",
			arg:      Message{Priority: -1},
			expected: `{"topic":""}`,
		},
		{
			name:     "Priority greater than 0",
			arg:      Message{Priority: 1},
			expected: `{"topic":"","priority":1}`,
		},
		{
			name: "Actions",
			arg: Message{Actions: []ActionButton{&ViewAction{
				Label: "action",
				Link:  "http://host.com",
				Clear: true,
			}}},
			expected: `{"topic":"","actions":[{"action":"view","label":"action","url":"http://host.com","clear":true}]}`,
		},
		{
			name:     "ClickURL",
			arg:      Message{ClickURL: "h://t.com"},
			expected: `{"topic":"","click":"h://t.com"}`,
		},
		{
			name:     "IconURL",
			arg:      Message{IconURL: "h://t.com"},
			expected: `{"topic":"","icon":"h://t.com"}`,
		},
		{
			name:     "Delay as int",
			arg:      Message{Delay: 1},
			expected: `{"topic":"","delay":1}`,
		},
		{
			name:     "Delay as string",
			arg:      Message{Delay: "1ns"},
			expected: `{"topic":"","delay":"1ns"}`,
		},
		{
			name:     "Email",
			arg:      Message{Email: "Email"},
			expected: `{"topic":"","email":"Email"}`,
		},
		{
			name:     "Call",
			arg:      Message{Call: "Call"},
			expected: `{"topic":"","call":"Call"}`,
		},
		{
			name:     "AttachURLFilename",
			arg:      Message{AttachURLFilename: "AttachURLFilename"},
			expected: `{"topic":"","filename":"AttachURLFilename"}`,
		},
		{
			name:     "AttachURL",
			arg:      Message{AttachURL: "h://t.com"},
			expected: `{"topic":"","attachurl":"h://t.com"}`,
		},
		{
			name: "everything",
			arg: Message{
				Topic:    "Topic",
				Message:  "Message",
				Title:    "Title",
				Tags:     []string{"tag1", "tag2"},
				Priority: High,
				Actions: []ActionButton{&ViewAction{
					Label: "ajisdiopa",
					Link:  "h://t.com",
					Clear: true,
				}},
				ClickURL:          "h://t.com",
				IconURL:           "h://t.com",
				Delay:             "10m",
				Email:             "Email",
				Call:              "Call",
				AttachURLFilename: "AttachURLFilename",
				AttachURL:         "h://t.com",
			},
			expected: `{"topic":"Topic","message":"Message","title":"Title","tags":["tag1","tag2"],"priority":4,"actions":[{"action":"view","label":"ajisdiopa","url":"h://t.com","clear":true}],"click":"h://t.com","icon":"h://t.com","delay":"10m","email":"Email","call":"Call","filename":"AttachURLFilename","attachurl":"h://t.com"}`,
		},
		{
			name: "test case failure 1/28/2024",
			arg: Message{
				Topic:             "9mm Luger brass 115 grain: $225.00/round, 1000 rounds",
				Title:             "",
				Tags:              []string{Nine},
				Priority:          0,
				Actions:           []ActionButton{},
				ClickURL:          "",
				IconURL:           "",
				Delay:             0,
				Email:             "",
				Call:              "",
				AttachURLFilename: "",
				AttachURL:         "",
				Message:           "ZSR: 9mm - ZSR Buffalo Cartridge 115 Grain Full Metal Jacket - 1000 Rounds 8683262441013 - FREE SHIPPING\n5.0/5 stars, 204 ratings",
			},
			expected: `{"topic":"9mm Luger brass 115 grain: $225.00/round, 1000 rounds","message":"ZSR: 9mm - ZSR Buffalo Cartridge 115 Grain Full Metal Jacket - 1000 Rounds 8683262441013 - FREE SHIPPING\n5.0/5 stars, 204 ratings","tags":["nine"]}`,
		},
	}

	t := assert.New(mainTest)
	for _, tc := range testCases {
		actual, actualErr := json.Marshal(&tc.arg)

		if t.Nil(actualErr, fmt.Sprintf("should not return %s", actualErr)) {
			t.Equal([]byte(tc.expected), actual, tc.name)
		}
	}
}
