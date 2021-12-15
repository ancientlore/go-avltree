package avltree

// Balance factor
const (
	equal = iota
	leftHigh
	rightHigh
)

// treeNode is a node in the tree
type treeNode[T any] struct {
	// Left and right nodes.
	left, right *treeNode[T]

	// The number of nodes in the left and right subtrees
	// (excludes this node).
	size int

	// The contents of this node.
	value T

	// The balance factor of this node.
	bal byte
}

// leftSize returns the size of the left subtree
// of the node
func (n *treeNode[T]) leftSize() int {
	if n.left != nil {
		return n.left.size + 1
	}
	return 0
}

// rightSize returns the size of the right subtree
// of the node
func (n treeNode[T]) rightSize() int {
	if n.right != nil {
		return n.right.size + 1
	}
	return 0
}
