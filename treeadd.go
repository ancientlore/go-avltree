package avltree

// addData holds information used when adding nodes
type addData struct {
	lookingFor interface{}
	duplicate  interface{}
	tree       *Tree
}

func (d *addData) add(node **treeNode, taller *bool) interface{} {
	*taller = false // default: not taller
	var ptr interface{}
	ptr = nil

	if *node == nil {
		*node = newNode(d.lookingFor)
		*taller = true
		if *node != nil {
			ptr = (*node).value
		}
	} else {
		tallerSubTree := false

		code := d.tree.compare(d.lookingFor, (*node).value)

		if code == 0 && (d.tree.treeFlags&AllowDuplicates) != 0 {
			code = -1 // go left for duplicates
		}

		if code < 0 {
			ptr = d.add(&((*node).left), &tallerSubTree)
			if ptr != nil {
				if tallerSubTree {
					switch (*node).bal {
					case leftHigh:
						*node, *taller = leftBalance(*node, *taller)
					case equal:
						(*node).bal = leftHigh
						*taller = true
					case rightHigh:
						(*node).bal = equal
						*taller = false
					}
				}
			}
		} else if code > 0 {
			ptr = d.add(&((*node).right), &tallerSubTree)
			if ptr != nil {
				if tallerSubTree {
					switch (*node).bal {
					case leftHigh:
						(*node).bal = equal
						*taller = false
					case equal:
						(*node).bal = rightHigh
						*taller = true
					case rightHigh:
						*node, *taller = rightBalance(*node, *taller)
					}
				}
			}
		} else {
			d.duplicate = (*node).value // this node is the duplicate
		}

		if ptr != nil {
			(*node).size = (*node).leftSize() + (*node).rightSize() + 1
		}
	}

	return ptr
}

// Add adds an item to the tree, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found. A duplicate will never be returned if the
// tree's AllowDuplicates flag is set.
func (t *Tree) Add(o interface{}) (val interface{}, isDupe bool) {
	d := &addData{o, nil, t}
	taller := false
	isDupe = false
	val = d.add(&t.root, &taller)
	if val == nil {
		isDupe = true
		val = d.duplicate
	}
	return
}

func rightBalance(node *treeNode, taller bool) (*treeNode, bool) {
	var x *treeNode // right subtree of node
	var w *treeNode // left subtree of x

	x = node.right

	switch x.bal {
	case rightHigh:
		node.bal = equal
		x.bal = equal
		node = rotateLeft(node)
		taller = false
	case equal:
		// this should be impossible
	case leftHigh:
		w = x.left
		switch w.bal {
		case equal:
			node.bal = equal
			x.bal = equal
		case leftHigh:
			node.bal = equal
			x.bal = rightHigh
		case rightHigh:
			node.bal = leftHigh
			x.bal = equal
		}
		w.bal = equal
		x = rotateRight(x)
		node.right = x
		node = rotateLeft(node)
		taller = false
	}
	return node, taller
}

func leftBalance(node *treeNode, taller bool) (*treeNode, bool) {
	var x *treeNode // left subtree of node
	var w *treeNode // right subtree of x

	x = node.left

	switch x.bal {
	case leftHigh:
		node.bal = equal
		x.bal = equal
		node = rotateRight(node)
		taller = false
	case equal:
		// this should be impossible
	case rightHigh:
		w = x.right
		switch w.bal {
		case equal:
			node.bal = equal
			x.bal = equal
		case rightHigh:
			node.bal = equal
			x.bal = leftHigh
		case leftHigh:
			node.bal = rightHigh
			x.bal = equal
		}
		w.bal = equal
		x = rotateLeft(x)
		node.left = x
		node = rotateRight(node)
		taller = false
	}
	return node, taller
}

func rotateLeft(node *treeNode) *treeNode {
	if node != nil && node.right != nil {
		ptr := node.right
		node.size = node.leftSize() + ptr.leftSize() + 1
		ptr.size = node.size + ptr.rightSize() + 1
		node.right = ptr.left
		ptr.left = node
		node = ptr
	}
	return node
}

func rotateRight(node *treeNode) *treeNode {
	if node != nil && node.left != nil {
		ptr := node.left
		node.size = node.rightSize() + ptr.rightSize() + 1
		ptr.size = node.size + ptr.leftSize() + 1
		node.left = ptr.right
		ptr.right = node
		node = ptr
	}
	return node
}
