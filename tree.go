/*
Package avltree implements a height-balanced binary tree
with array-like indexing capability.

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
	"constraints"
	"context"
	"math"
)

// tree options
const (
	AllowDuplicates = 1
)

// compareFunc defines the function type used to compare values.
type compareFunc[T any] func(T, T) int

// iterateFunc defines the function type used for iterating a tree.
type iterateFunc[T any] func(T) bool

// Tree stores data about the binary tree.
type Tree[T any] struct {
	// root of the tree
	root *treeNode[T]

	// compare function
	compare compareFunc[T]

	// options controlling behavior
	treeFlags byte
}

// New returns an initialized tree.
func New[T any](c func(T, T) int, flags byte) *Tree[T] {
	 return &Tree[T]{
		 compare: c,
		 treeFlags: flags,
	 }
}

// NewOrdered returns an initialized tree using ordered types.
func NewOrdered[T constraints.Ordered](flags byte) *Tree[T] {
	 return &Tree[T]{
		compare: func (v1, v2 T) int {
			switch {
			case v1 < v2:
				return -1
			case v1 == v2:
				return 0
			default:
				return 1
			}
		},
		treeFlags: flags,
	}
}

// Clear removes all elements from the tree, keeping the
// current options and compare function.
func (t *Tree[T]) Clear() {
	t.root = nil
}

// calcHeightData contains information needed to compute the
// height of the tree
type calcHeightData[T any] struct {
	currentHeight int // current height of the tree
	maxHeight     int // maximum height of the tree we found
}

// calcHeight executes recursively to determine the height
// (number of levels) in the tree.
func (d *calcHeightData[T]) calcHeight(node *treeNode[T]) {
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
func (t *Tree[T]) Height() int {
	d := &calcHeightData[T]{0, 0}

	if t.root != nil {
		d.calcHeight(t.root)
	}

	return d.maxHeight
}

// Len returns the number of elements in the tree.
func (t *Tree[T]) Len() int {
	if t.root != nil {
		return t.root.size + 1
	}
	return 0
}

// Cap returns the capacity of the tree; that is, the
// maximum elements the tree can hold with at the
// current height. This is only useful as a measure
// of how skewed the tree is.
func (t *Tree[T]) Cap() int {

	var count, i int
	count = 0

	maxHeight := t.Height()

	for i = 0; i < maxHeight; i++ {
		count += int(math.Pow(2, float64(i)))
	}

	return count
}

// indexer recursively scans the tree to find the node
// at the given position.
func indexer[T any](node *treeNode[T], index int) *treeNode[T] {

	if index < node.leftSize() {
		return indexer(node.left, index)
	} else if index == node.leftSize() {
		return node
	} else if node.right != nil {
		return indexer(node.right, index-(node.leftSize()+1))
	}
	return nil
}

// At returns the value at the given index.
func (t *Tree[T]) At(index int) *T {

	if t.root != nil && index < t.root.size+1 && index >= 0 {
		node := indexer(t.root, index)
		if node != nil {
			return &node.value
		}
	}

	return nil
}

// findData[T] is used when searching the tree.
type findData[T any] struct {
	lookingFor T              // item we are searching for
	compare    compareFunc[T] // Comparision function
}

// finder recursively scans the tree to find the node with the
// value we're looking for.
func (d *findData[T]) finder(node *treeNode[T]) *treeNode[T] {

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
// the node's value and the given key value.
func (t *Tree[T]) Find(key T) *T {
	if t.root != nil {
		d := &findData[T]{key, t.compare}
		node := d.finder(t.root)
		if node != nil {
			return &node.value
		}
	}

	return nil
}

// iterate recursively traverses the tree and executes
// the iteration function.
func (d iterateFunc[T]) iterate(node *treeNode[T]) bool {
	var proceed bool

	if node.left != nil {
		proceed = d.iterate(node.left)
		if !proceed {
			return false
		}
	}

	proceed = d(node.value)
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
func (t *Tree[T]) Do(f func(T) bool) {

	if f != nil && t.root != nil {
		iterateFunc[T](f).iterate(t.root)
	}
}

// chanIterate should be used as a goroutine to produce all the values
// in the tree.
func (t *Tree[T]) chanIterate(ctx context.Context, c chan<- T) {
	t.Do(func(v T) bool {
		select {
		case c <- v:
			return true
		case <-ctx.Done():
			return false
		}
	})
	close(c)
}

// Iter returns a channel you can read through to fetch all the items.
func (t *Tree[T]) Iter() <-chan T {
	c := make(chan T)
	go t.chanIterate(context.Background(), c)
	return c
}

// IterContext returns a channel you can read through to fetch all the items.
func (t *Tree[T]) IterContext(ctx context.Context) <-chan T {
	c := make(chan T)
	go t.chanIterate(ctx, c)
	return c
}

// Data returns all the elements as a slice.
func (t *Tree[T]) Data() []T {
	arr := make([]T, t.Len())
	var i int
	i = 0
	t.Do(func(v T) bool {
		arr[i] = v
		i++
		return true
	})
	return arr
}
