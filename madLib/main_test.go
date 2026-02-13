package main

import "testing"

type TestInput struct {
	noun      string
	verb      string
	adjective string
	adverb    string
}

func TestBuildingStrings(t *testing.T) {
	testCase := []struct {
		name         string
		input        TestInput
		stringPicker int
		got          string
		want         string
	}{
		{
			name: "TestCorrectInput",
			input: TestInput{
				noun:      "dog",
				verb:      "walk",
				adjective: "blue",
				adverb:    "quickly",
			},
			stringPicker: 0,
			got:          "",
			want:         "Do you walk your blue dog quickly? That's hilarious!",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := BuildingStrings(tc.input.noun, tc.input.verb, tc.input.adjective, tc.input.adverb, tc.stringPicker)
			if result != tc.want {
				t.Errorf("want %v, got %v instead", tc.want, result)
			}
		})
	}
}
