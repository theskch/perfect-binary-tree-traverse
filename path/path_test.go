package path

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	for _, testCase := range []struct {
		name          string
		binaryTree    []int
		expectedPath  string
		expectedError error
	}{
		{
			name:          "empty tree",
			binaryTree:    []int{},
			expectedPath:  "",
			expectedError: fmt.Errorf("tree must have at least two elements"),
		},
		{
			name:          "single element tree",
			binaryTree:    []int{0, 4},
			expectedPath:  "",
			expectedError: fmt.Errorf("tree only has root element"),
		},
		{
			name:          "tree is not perfectly balanced",
			binaryTree:    []int{0, 4, 5, 6, 2, 3},
			expectedPath:  "",
			expectedError: fmt.Errorf("tree is not perfectly balanced"),
		},
		{
			name:          "expected path LLL",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, -1, 3, 6, 5, 3, 2, 1, 2},
			expectedPath:  "LLL",
			expectedError: nil,
		},
		{
			name:          "expected path LLR",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, -3, 6, 5, 3, 2, 1, 2},
			expectedPath:  "LLR",
			expectedError: nil,
		},
		{
			name:          "expected path LRL",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, -6, 5, 3, 2, 1, 2},
			expectedPath:  "LRL",
			expectedError: nil,
		},
		{
			name:          "expected path LRR",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, 6, -5, 3, 2, 1, 2},
			expectedPath:  "LRR",
			expectedError: nil,
		},
		{
			name:          "expected path RLL",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, 6, 5, -3, 2, 1, 2},
			expectedPath:  "RLL",
			expectedError: nil,
		},
		{
			name:          "expected path RLR",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, 6, 5, 3, -2, 1, 2},
			expectedPath:  "RLR",
			expectedError: nil,
		},
		{
			name:          "expected path RRL",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, 6, 5, 3, 2, -1, 2},
			expectedPath:  "RRL",
			expectedError: nil,
		},
		{
			name:          "expected path RRR",
			binaryTree:    []int{0, 4, 2, 5, 4, -8, -2, 6, 1, 3, 6, 5, 3, 2, 1, -2},
			expectedPath:  "RRR",
			expectedError: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			path, err := Calculate(testCase.binaryTree)
			assert.Equal(t, err, testCase.expectedError)
			assert.Equal(t, testCase.expectedPath, path)
		})

	}
}
