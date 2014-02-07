package avltree

import (
	"io"
)

// StringTree is a specialization of Tree that hides the wrapping of Elements around strings.
type StringTree struct {
	Tree
}

func stringCompare(s1 interface{}, s2 interface{}) int {
	if s1.(string) < s2.(string) {
		return -1
	} else if s1.(string) > s2.(string) {
		return 1
	}
	return 0
}

// Iterate function
type StringIterateFunc func(v string)

// Initialize or reset a StringTree
func (t *StringTree) Init(flags byte) *StringTree {
	t.Tree.Init(stringCompare, flags)
	return t
}

// Return an initialized StringTree
func NewStringTree(flags byte) *StringTree { return new(StringTree).Init(flags) }

// At returns the value at the given index
func (t *StringTree) At(index int) string {
	v := t.Tree.At(index)
	if v != nil {
		return v.(string)
	}
	return ""
}

// Find returns the element where the comparison function matches
// the node's value and the given key value
func (t *StringTree) Find(key string) string {
	v := t.Tree.Find(key)
	if v != nil {
		return v.(string)
	}
	return ""
}

// Do calls function f for each element of the tree, in order.
// The function should not change the structure of the tree underfoot.
func (t *StringTree) Do(f StringIterateFunc) { t.Tree.Do(func(v interface{}) { f(v.(string)) }) }

// chanIterate should be used as a goroutine to produce all the values
// in the tree.
func (t *StringTree) chanIterate(c chan<- string) {
	t.Do(func(v string) { c <- v })
	close(c)
}

// Iter returns a channel you can read through to fetch all the items
func (t *StringTree) Iter() <-chan string {
	c := make(chan string)
	go t.chanIterate(c)
	return c
}

// Data returns all the elements as a slice.
func (t *StringTree) Data() []string {
	arr := make([]string, t.Len())
	var i int
	i = 0
	t.Do(func(v string) {
		arr[i] = v
		i++
	})
	return arr
}

// Add adds an item to the tree, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found. A duplicate will never be returned if the
// tree's AllowDuplicates flag is set.
func (t *StringTree) Add(o string) (val string, isDupe bool) {
	v, d := t.Tree.Add(o)
	if v != nil {
		return v.(string), d
	}
	return "", d
}

// Remove removes the element matching the given value.
func (t *StringTree) Remove(ptr string) string {
	v := t.Tree.Remove(ptr)
	if v != nil {
		return v.(string)
	}
	return ""
}

// Remove removes the element at the given index
func (t *StringTree) RemoveAt(index int) string {
	v := t.Tree.RemoveAt(index)
	if v != nil {
		return v.(string)
	}
	return ""
}

func (t *StringTree) Print(w io.Writer, f StringIterateFunc, itemSiz int) {
	t.Tree.Print(w, func(v interface{}) { f(v.(string)) }, itemSiz)
}
