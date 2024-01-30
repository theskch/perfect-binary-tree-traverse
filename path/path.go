package path

import (
	"fmt"
	"math"
)

func print(path []int) string {
	stringPath := ""
	for i := len(path) - 1; i >= 0; i-- {
		if path[i]%2 == 1 {
			stringPath += "R"
		} else {
			stringPath += "L"
		}
	}

	return stringPath
}

func Calculate(tree []int) (string, error) {
	if len(tree) < 2 {
		return "", fmt.Errorf("tree must have at least two elements")
	}

	// Skip the first level [0, {root}]
	// First level should have two elements.
	countPerLevel := 2
	leftover := tree[countPerLevel:]

	if len(leftover) == 0 {
		return "", fmt.Errorf("tree only has root element")
	}

	// Iterate through the array until the leaf level is reached.
	for len(leftover) > 2*countPerLevel {
		// Remove the current level.
		leftover = leftover[countPerLevel:]
		// Next level should always have double the number of prevous level
		// in a perfectly balanced tree.
		countPerLevel *= 2
	}

	// If the leaf level doesn't have the expected number of elements
	// it means it's not perfectly balanced.
	if len(leftover) != countPerLevel {
		return "", fmt.Errorf("tree is not perfectly balanced")
	}

	// Find the index of the minimal element in the leaf level.
	minItem := math.MaxInt
	index := 0
	for i, item := range leftover {
		if item < minItem {
			minItem = item
			index = i
		}
	}

	// Save the index of the leaf with the lowest value.
	indexes := []int{index}
	// Next level should always have half the number of prevous level.
	currentLevelCount := len(leftover) / 2

	// Find the index of the parent node in the lower level.
	for currentLevelCount >= 2 {
		// Current level is half the size of the previous level.
		currentLevelCount /= 2
		// The index of the parent node is the rounded division by two of the
		// index of the child node in the current level.
		index = index / 2
		indexes = append(indexes, index)
	}

	return print(indexes), nil
}
