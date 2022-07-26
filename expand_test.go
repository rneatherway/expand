package expand

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestExpand(t *testing.T) {
	fileSet := token.FileSet{}
	a, err := parser.ParseFile(
		&fileSet,
		"testdata/example.go",
		nil,
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		input    Selection
		expected Selection
	}{
		{
			input:    Selection{14, 11, 14, 12},
			expected: Selection{14, 9, 14, 14},
		},
		{
			input:    Selection{5, 13, 5, 13},
			expected: Selection{5, 13, 5, 14},
		},
		{
			input:    Selection{5, 13, 5, 14},
			expected: Selection{5, 10, 5, 18},
		},
		{
			input:    Selection{5, 10, 5, 18},
			expected: Selection{5, 9, 5, 32},
		},
		{
			input:    Selection{10, 22, 10, 23},
			expected: Selection{10, 19, 10, 24},
		},
		{
			input:    Selection{10, 19, 10, 24},
			expected: Selection{10, 7, 10, 24},
		},
		{
			input:    Selection{10, 7, 10, 24},
			expected: Selection{10, 3, 10, 24},
		},
		{
			input:    Selection{10, 3, 10, 24},
			expected: Selection{9, 23, 11, 3},
		},
		{
			input:    Selection{9, 23, 11, 3},
			expected: Selection{9, 2, 11, 3},
		},
	}

	for _, tc := range cases {
		actual := expandSelection(&fileSet, a, tc.input)
		if actual != tc.expected {
			t.Errorf("expandSelection(%v) = %v, expected %v", tc.input, actual, tc.expected)
		}

	}

}
