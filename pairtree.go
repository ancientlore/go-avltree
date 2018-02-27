package avltree

import (
	"io"
)

// PairTree is a specialization of Tree that hides the wrapping of Elements around Pair structures.
type PairTree struct {
	ObjectTree
}

// Pair combines a key and value.
type Pair struct {
	Key   string
	Value interface{}
}

// Compare is the compare function for Pairs, based on Key.
func (a Pair) Compare(b Interface) int {
	if a.Key < b.(Pair).Key {
		return -1
	} else if a.Key > b.(Pair).Key {
		return 1
	}
	return 0
}

// PairIterateFunc is the type of function used for iterating across Pairs.
type PairIterateFunc func(v Pair) bool

// Init will initialize or reset a PairTree.
func (t *PairTree) Init(flags byte) *PairTree {
	t.ObjectTree.Init(flags)
	return t
}

// NewPairTree returns an initialized PairTree.
func NewPairTree(flags byte) *PairTree { return new(PairTree).Init(flags) }

// At returns the value at the given index.
func (t *PairTree) At(index int) *Pair {
	v := t.ObjectTree.At(index)
	if v != nil {
		x := v.(Pair)
		return &x
	}
	return nil
}

// Find returns the element where the comparison function matches
// the node's value and the given key value.
func (t *PairTree) Find(key string) *Pair {
	v := t.ObjectTree.Find(Pair{key, nil})
	if v != nil {
		x := v.(Pair)
		return &x
	}
	return nil
}

// Do calls function f for each element of the tree, in order.
// The function should not change the structure of the tree underfoot.
func (t *PairTree) Do(f PairIterateFunc) {
	t.ObjectTree.Do(func(v interface{}) bool { return f(v.(Pair)) })
}

// chanIterate should be used as a goroutine to produce all the values
// in the tree.
func (t *PairTree) chanIterate(c chan<- Pair) {
	t.Do(func(v Pair) bool { c <- v; return true })
	close(c)
}

// Iter returns a channel you can read through to fetch all the items.
func (t *PairTree) Iter() <-chan Pair {
	c := make(chan Pair)
	go t.chanIterate(c)
	return c
}

// Data returns all the elements as a slice.
func (t *PairTree) Data() []Pair {
	arr := make([]Pair, t.Len())
	var i int
	i = 0
	t.Do(func(v Pair) bool {
		arr[i] = v
		i++
		return true
	})
	return arr
}

// Add adds an item to the tree, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found. A duplicate will never be returned if the
// tree's AllowDuplicates flag is set.
func (t *PairTree) Add(o Pair) (val *Pair, isDupe bool) {
	v, d := t.ObjectTree.Add(o)
	if v != nil {
		x := v.(Pair)
		return &x, d
	}
	return nil, d
}

// Remove removes the element matching the given value.
func (t *PairTree) Remove(ptr string) *Pair {
	v := t.ObjectTree.Remove(Pair{ptr, nil})
	if v != nil {
		x := v.(Pair)
		return &x
	}
	return nil
}

// RemoveAt removes the element at the given index.
func (t *PairTree) RemoveAt(index int) *Pair {
	v := t.ObjectTree.RemoveAt(index)
	if v != nil {
		x := v.(Pair)
		return &x
	}
	return nil
}

// Print prints the values in the tree to the writer.
func (t *PairTree) Print(w io.Writer, f PairIterateFunc, itemSiz int) {
	t.ObjectTree.Print(w, func(v interface{}) bool { return f(v.(Pair)) }, itemSiz)
}
