package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/theskch/binary-tree-path-calculation/path"
)

func main() {
	tree := &treeFlag{}
	flag.Var(tree, "tree", "Tree flag used as an input")
	flag.Parse()

	path, err := path.Calculate(*tree)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(path)
}

// A helper struct used to parse input flag with tree definition.
type treeFlag []int

// Implement the flag.Value interface for the custom type.
func (tf *treeFlag) String() string {
	return fmt.Sprintf("%v", *tf)
}

// Implement the flag.Value interface for the custom type.
func (tf *treeFlag) Set(value string) error {
	stringTree := strings.Split(value, ",")
	for _, str := range stringTree {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*tf = append(*tf, intValue)
	}

	return nil
}
