package avltree

// Balance factor
const (
	leftHigh = iota
	equal
	rightHigh
)

// treeNode is a node in the tree
type treeNode struct {
	// left and right nodes
	left, right *treeNode

	// The contents of this node
	value interface{}

	// The balance factor of this node
	bal byte

	// The number of nodes in the subtree
	size int
}

// Init initializes a node with the given value
func (n *treeNode) init(val interface{}) *treeNode {
	n.left = nil
	n.right = nil
	n.bal = equal
	n.size = 1
	n.value = val
	return n
}

// newNode returns an initialized treeNode
func newNode(val interface{}) *treeNode { return new(treeNode).init(val) }

// leftSize returns the size of the left subtree
// of the node
func (n *treeNode) leftSize() int {
	if n.left != nil {
		return n.left.size
	}
	return 0
}

// rightSize returns the size of the right subtree
// of the node
func (n treeNode) rightSize() int {
	if n.right != nil {
		return n.right.size
	}
	return 0
}
