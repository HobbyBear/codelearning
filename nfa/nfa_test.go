package nfa

import (
	"fmt"
	"testing"
)

func TestNFA_Match(t *testing.T) {
	nfa := New("a*")
	fmt.Println(nfa.Match("aa"))
}
