/*
An AVL tree (Adel'son-Vel'skii & Landis) is a binary search
tree in which the heights of the left and right subtrees
of the root differ by at most one and in which the left
and right subtrees are again AVL trees.

With each node of an AVL tree is associated a balance factor
that is Left High, Equal, or Right High according,
respectively, as the left subtree has height greater than,
equal to, or less than that of the right subtree.

The AVL tree is, in practice, balanced quite well.  It can
(at the worst case) become skewed to the left or right,
but never so much that it becomes inefficient.  The
balancing is done as items are added or deleted.

This version is enhanced to allow "indexing" of values in the
tree; however, the indexes are not stable as the tree could be
resorted as items are added or removed.

It is safe to iterate or search a tree from multiple threads
provided that no threads are modifying the tree.

See also:	Robert L. Kruse, Data Structures and Program Design, 2nd Ed., Prentice-Hall
*/
package avltree

import (
	"math"
)

// tree options
const (
	AllowDuplicates = 1
)

// Definition of a comparison function
type CompareFunc func(v1 interface{}, v2 interface{}) int

// Iterate function
type IterateFunc func(v interface{}) bool

// Tree object
type Tree struct {
	// root of the tree
	root *treeNode

	// compare function
	compare CompareFunc

	// options controlling behavior
	treeFlags byte
}

// Initialize or reset a Tree
func (t *Tree) Init(c CompareFunc, flags byte) *Tree {
	t.compare = c
	t.root = nil
	t.treeFlags = flags
	return t
}

// Return an initialized tree
func New(c CompareFunc, flags byte) *Tree { return new(Tree).Init(c, flags) }

// Clear removes all elements from the tree, keeping the
// current options and compare function
func (t *Tree) Clear() { t.Init(t.compare, t.treeFlags) }

// calcHeightData contains information needed to compute the
// height of the tree
type calcHeightData struct {
	currentHeight int
	maxHeight     int
}

// calcHeight executes recursively to determine the height
// (number of levels) in the tree.
func (d *calcHeightData) calcHeight(node *treeNode) {
	d.currentHeight++

	if node.left != nil {
		d.calcHeight(node.left)
	}

	if node.right != nil {
		d.calcHeight(node.right)
	}

	if d.currentHeight > d.maxHeight {
		d.maxHeight = d.currentHeight
	}
	d.currentHeight--
}

// Height returns the "height" of the tree, meaning the
// number of levels.
func (t *Tree) Height() int {
	d := &calcHeightData{0, 0}

	if t.root != nil {
		d.calcHeight(t.root)
	}

	return d.maxHeight
}

// Len returns the number of elements in the tree
func (t *Tree) Len() int {
	if t.root != nil {
		return t.root.size
	}
	return 0
}

// Cap returns the capacity of the tree; that is, the
// maximum elements the tree can hold with at the
// current height. This is only useful as a measure
// of how skewed the tree is.
func (t *Tree) Cap() int {

	var count, i int
	count = 0

	maxHeight := t.Height()

	for i = 0; i < maxHeight; i++ {
		count += int(math.Pow(2, float64(i)))
	}

	return count
}

// indexer recursively scans the tree to find the node
// at the given position
func indexer(node *treeNode, index int) *treeNode {

	if index < node.leftSize() {
		return indexer(node.left, index)
	} else if index == node.leftSize() {
		return node
	} else if node.right != nil {
		return indexer(node.right, index-(node.leftSize()+1))
	}
	return nil
}

// At returns the value at the given index
func (t *Tree) At(index int) interface{} {

	if t.root != nil && index < t.root.size && index >= 0 {
		node := indexer(t.root, index)
		if node != nil {
			return node.value
		}
	}

	return nil
}

// findData is used when searching the tree
type findData struct {
	lookingFor interface{}
	compare    CompareFunc
}

// finder recursively scans the tree to find the node with the
// value we're looking for
func (d *findData) finder(node *treeNode) *treeNode {

	if node != nil {
		code := d.compare(d.lookingFor, node.value)
		if code < 0 {
			return d.finder(node.left)
		} else if code > 0 {
			return d.finder(node.right)
		}
		return node
	}
	return nil
}

// Find returns the element where the comparison function matches
// the node's value and the given key value
func (t *Tree) Find(key interface{}) interface{} {
	if key != nil && t.root != nil {
		d := &findData{key, t.compare}
		node := d.finder(t.root)
		if node != nil {
			return node.value
		}
	}

	return nil
}

// iterData is used when iterating the tree
type iterData struct {
	iter IterateFunc
}

// iterate recursively traverses the tree and executes
// the iteration function
func (d *iterData) iterate(node *treeNode) bool {
	var proceed bool

	if node.left != nil {
		proceed = d.iterate(node.left)
		if !proceed {
			return false
		}
	}

	proceed = d.iter(node.value)
	if !proceed {
		return false
	}

	if node.right != nil {
		proceed = d.iterate(node.right)
		if !proceed {
			return false
		}
	}

	return true
}

// Do calls function f for each element of the tree, in order.
// The function should not change the structure of the tree underfoot.
func (t *Tree) Do(f IterateFunc) {

	if f != nil && t.root != nil {
		d := &iterData{f}
		d.iterate(t.root)
	}
}

// chanIterate should be used as a goroutine to produce all the values
// in the tree.
func (t *Tree) chanIterate(c chan<- interface{}) {
	t.Do(func(v interface{}) bool { c <- v; return true })
	close(c)
}

// Iter returns a channel you can read through to fetch all the items
func (t *Tree) Iter() <-chan interface{} {
	c := make(chan interface{})
	go t.chanIterate(c)
	return c
}

// Data returns all the elements as a slice.
func (t *Tree) Data() []interface{} {
	arr := make([]interface{}, t.Len())
	var i int
	i = 0
	t.Do(func(v interface{}) bool {
		arr[i] = v
		i++
		return true
	})
	return arr
}
