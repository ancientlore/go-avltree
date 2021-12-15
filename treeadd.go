package avltree

// addData holds information used when adding nodes
type addData[T any] struct {
	lookingFor T        // Item to add
	duplicate  *T       // Duplicate found, if any
	tree       *Tree[T] // tree to add to
}

func (d *addData[T]) add(node **treeNode[T], taller *bool) *T {
	*taller = false // default: not taller
	var ptr *T

	if *node == nil {
		*node = &treeNode[T]{value: d.lookingFor}
		*taller = true
		if *node != nil {
			ptr = &(*node).value
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
			d.duplicate = &(*node).value // this node is the duplicate
		}

		if ptr != nil {
			(*node).size = (*node).leftSize() + (*node).rightSize()
		}
	}

	return ptr
}

// Add adds an item to the tree, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found. A duplicate will never be returned if the
// tree's AllowDuplicates flag is set.
func (t *Tree[T]) Add(o T) (val *T, isDupe bool) {
	d := &addData[T]{o, nil, t}
	taller := false
	isDupe = false
	val = d.add(&t.root, &taller)
	if val == nil {
		isDupe = true
		val = d.duplicate
	}
	return
}

func rightBalance[T any](node *treeNode[T], taller bool) (*treeNode[T], bool) {
	var x *treeNode[T] // right subtree of node
	var w *treeNode[T] // left subtree of x

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

func leftBalance[T any](node *treeNode[T], taller bool) (*treeNode[T], bool) {
	var x *treeNode[T] // left subtree of node
	var w *treeNode[T] // right subtree of x

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

func rotateLeft[T any](node *treeNode[T]) *treeNode[T] {
	if node != nil && node.right != nil {
		ptr := node.right
		node.size = node.leftSize() + ptr.leftSize()
		ptr.size = node.size + 1 + ptr.rightSize()
		node.right = ptr.left
		ptr.left = node
		node = ptr
	}
	return node
}

func rotateRight[T any](node *treeNode[T]) *treeNode[T] {
	if node != nil && node.left != nil {
		ptr := node.left
		node.size = node.rightSize() + ptr.rightSize()
		ptr.size = node.size + 1 + ptr.leftSize()
		node.left = ptr.right
		ptr.right = node
		node = ptr
	}
	return node
}
