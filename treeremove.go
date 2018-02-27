package avltree

type removeData struct {
	lookingFor interface{}
	compare    CompareFunc
}

func findPredecessor(node *treeNode) *treeNode {
	if node != nil {
		pred := node.left
		if pred != nil {
			for pred.right != nil {
				pred = pred.right
			}
		}
		return pred
	}
	return nil
}

func remLeftSubBalance(node *treeNode, shorter bool) (*treeNode, bool) {
	q := node.right // q: root of taller subtree
	var w *treeNode

	switch q.bal {
	case equal:
		node.bal = rightHigh
		q.bal = leftHigh // q will be the new root node
		node = rotateLeft(node)
		shorter = false // next level not shorter
	case rightHigh:
		node.bal = equal
		q.bal = equal // q will be the new root node
		node = rotateLeft(node)
	case leftHigh:
		w = q.left
		if w.bal == leftHigh {
			q.bal = rightHigh
		} else {
			q.bal = equal
		}
		if w.bal == rightHigh {
			node.bal = leftHigh
		} else {
			node.bal = equal
		}
		w.bal = equal // w will be the new root node
		q = rotateRight(q)
		node.right = q
		node = rotateLeft(node)
	}

	return node, shorter
}

func remRightSubBalance(node *treeNode, shorter bool) (*treeNode, bool) {
	q := node.left // q: root of taller subtree
	var w *treeNode

	switch q.bal {
	case equal:
		node.bal = leftHigh
		q.bal = rightHigh // q will be the new root node
		node = rotateRight(node)
		shorter = false // next level not shorter
	case leftHigh:
		node.bal = equal
		q.bal = equal // q will be the new root node
		node = rotateRight(node)
	case rightHigh:
		w = q.right
		if w.bal == rightHigh {
			q.bal = leftHigh
		} else {
			q.bal = equal
		}
		if w.bal == leftHigh {
			node.bal = rightHigh
		} else {
			node.bal = equal
		}
		w.bal = equal // w will be the new root node
		q = rotateLeft(q)
		node.left = q
		node = rotateRight(node)
	}

	return node, shorter
}

func removePredecessor(node *treeNode, shorter bool) (*treeNode, bool) {
	if node.right != nil {
		node.right, shorter = removePredecessor(node.right, shorter)

		if shorter { // left subtree was shortened
			node, shorter = remRightBalance(node, shorter)
		}

		node.size = node.leftSize() + node.rightSize() + 1
	} else {
		node = remNode(node)
	}

	return node, shorter
}

func remLeftBalance(node *treeNode, shorter bool) (*treeNode, bool) {

	switch node.bal {
	case equal: // one subtree shortened
		node.bal = rightHigh // now it's right high
		shorter = false      // overall tree same height
	case leftHigh: // taller subtree shortened
		node.bal = equal // now it's equal
	case rightHigh: // shorter subtree shortened
		node, shorter = remLeftSubBalance(node, shorter)
	}
	return node, shorter
}

func remRightBalance(node *treeNode, shorter bool) (*treeNode, bool) {

	switch node.bal {
	case equal: // one subtree shortened
		node.bal = leftHigh // now it's left high
		shorter = false     // overall tree same height
	case rightHigh: // taller subtree shortened
		node.bal = equal // now it's equal
	case leftHigh: // shorter subtree shortened
		node, shorter = remRightSubBalance(node, shorter)
	}
	return node, shorter
}

func remNode(node *treeNode) *treeNode {
	if node.left != nil {
		node = node.left
	} else if node.right != nil {
		node = node.right
	} else {
		node = nil
	}
	return node
}

func (d *removeData) remove(node **treeNode, shorter *bool) interface{} {

	*shorter = true // default: shorter
	var ptr interface{}
	ptr = nil

	code := d.compare(d.lookingFor, (*node).value)

	if code < 0 {
		if (*node).left != nil {
			ptr = d.remove(&((*node).left), shorter)

			if *shorter && ptr != nil { // left subtree was shortened
				*node, *shorter = remLeftBalance(*node, *shorter)
			}
		}
	} else if code > 0 {
		if (*node).right != nil {
			ptr = d.remove(&((*node).right), shorter)

			if *shorter && ptr != nil { // left subtree was shortened
				*node, *shorter = remRightBalance(*node, *shorter)
			}
		}
	} else {
		ptr = (*node).value

		if (*node).left != nil && (*node).right != nil { // do the switch to find the prev.
			// node with only one subtree
			pred := findPredecessor(*node)
			(*node).value = pred.value
			(*node).left, *shorter = removePredecessor((*node).left, *shorter)

			if *shorter { // left subtree was shortened
				*node, *shorter = remLeftBalance(*node, *shorter)
			}
		} else { // we found the node; it has 1 subtree
			*node = remNode(*node)
		}
	}

	if ptr != nil && *node != nil {
		(*node).size = (*node).leftSize() + (*node).rightSize() + 1
	}

	return ptr
}

// Remove removes the element matching the given value.
func (t *Tree) Remove(ptr interface{}) interface{} {
	if ptr != nil && t.root != nil {
		d := &removeData{ptr, t.compare}
		var shorter bool
		return d.remove(&(t.root), &shorter)
	}

	return nil
}

func remove(node **treeNode, index int, shorter *bool) interface{} {

	*shorter = true // default: shorter
	var ptr interface{}
	ptr = nil

	if index < (*node).leftSize() {
		if (*node).left != nil {
			ptr = remove(&((*node).left), index, shorter)

			if *shorter && ptr != nil { // left subtree was shortened
				*node, *shorter = remLeftBalance(*node, *shorter)
			}
		}
	} else if index == (*node).leftSize() {
		ptr = (*node).value

		if (*node).left != nil && (*node).right != nil { // do the switch to find the prev.
			// node with only one subtree
			pred := findPredecessor(*node)
			(*node).value = pred.value
			(*node).left, *shorter = removePredecessor((*node).left, *shorter)

			if *shorter { // left subtree was shortened
				*node, *shorter = remLeftBalance(*node, *shorter)
			}
		} else { // we found the node; it has 1 subtree
			*node = remNode(*node)
		}
	} else {
		if (*node).right != nil {
			ptr = remove(&((*node).right), index-((*node).leftSize()+1), shorter)

			if *shorter && ptr != nil { // left subtree was shortened
				*node, *shorter = remRightBalance(*node, *shorter)
			}
		}
	}

	if ptr != nil && *node != nil {
		(*node).size = (*node).leftSize() + (*node).rightSize() + 1
	}

	return ptr
}

// RemoveAt removes the element at the given index.
func (t *Tree) RemoveAt(index int) interface{} {
	if t.root != nil && index < t.root.size && index >= 0 {
		var shorter bool
		return remove(&(t.root), index, &shorter)
	}

	return nil
}
