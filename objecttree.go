package avltree

import ()

// ObjectTree is a specialization of Tree that hides the wrapping of Elements around objects.
// The object just needs to implement Interface
type ObjectTree struct {
	Tree
}

// Implement Interface so that your object can be sorted in the tree
type Interface interface {
	// Return -1 if this < b, 0 if this == b, and 1 if this > b
	Compare(b Interface) int
}

func objectCompare(o1 interface{}, o2 interface{}) int {
	return o1.(Interface).Compare(o2.(Interface))
}

// Initialize or reset an ObjectTree
func (t *ObjectTree) Init(flags byte) *ObjectTree {
	t.Tree.Init(objectCompare, flags)
	return t
}

// Return an initialized ObjectTree
func NewObjectTree(flags byte) *ObjectTree { return new(ObjectTree).Init(flags) }

// Find returns the element where the comparison function matches
// the node's value and the given key value
func (t *ObjectTree) Find(key Interface) interface{} {
	return t.Tree.Find(key)
}

// Add adds an item to the tree, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found. A duplicate will never be returned if the
// tree's AllowDuplicates flag is set.
func (t *ObjectTree) Add(o Interface) (val interface{}, isDupe bool) {
	return t.Tree.Add(o)
}

// Remove removes the element matching the given value.
func (t *ObjectTree) Remove(ptr Interface) interface{} {
	return t.Tree.Remove(ptr)
}
