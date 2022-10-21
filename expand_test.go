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
	var file *token.File
	fileSet.Iterate(func(f *token.File) bool { file = f; return false })
	if file.Name() != "testdata/example.go" {
		t.Fatal(file.Name())
	}

	cases := []struct {
		input    Selection
		expected Selection
	}{
		{
			input:    Selection{167, 167},
			expected: Selection{165, 170},
		},
		{
			input:    Selection{40, 40},
			expected: Selection{40, 41},
		},
		{
			input:    Selection{40, 41},
			expected: Selection{37, 45},
		},
		{
			input:    Selection{37, 45},
			expected: Selection{36, 59},
		},
		{
			input:    Selection{140, 142},
			expected: Selection{139, 144},
		},
		{
			input:    Selection{139, 144},
			expected: Selection{127, 144},
		},
		{
			input:    Selection{127, 144},
			expected: Selection{123, 144},
		},
		{
			input:    Selection{123, 144},
			expected: Selection{119, 147},
		},
		{
			input:    Selection{119, 147},
			expected: Selection{98, 147},
		},
	}

	for _, tc := range cases {
		actual := expandSelection(file, a, tc.input)
		if actual != tc.expected {
			t.Errorf("expandSelection(%v) = %v, expected %v", tc.input, actual, tc.expected)
		}

	}

}
