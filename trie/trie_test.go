package trie

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	tree := New()
	tree.Add("hello")
	fmt.Println(tree.Exits("hellw"))
}
