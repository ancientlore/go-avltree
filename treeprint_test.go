package avltree

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tree := NewOrdered[rune](0)
	for i := 0; i < 50; i++ {
		tree.Add('A' + rune(rand.Intn(26)))
		tree.Add('a' + rune(rand.Intn(26)))
		tree.Add('0' + rune(rand.Intn(10)))
	}

	var buf bytes.Buffer
	Print(tree, &buf, func(c rune) bool {
		fmt.Fprintf(&buf, "\"%c\"", c)
		return true
	}, 3)

	t.Log(buf.String())
	fmt.Println(buf.String())
}
