package avltree

import (
	"fmt"
	"io"
)

type printData struct {
	lineDepth int
	itemSize  int
	array     []byte
	index     int
	iter      IterateFunc
	w         io.Writer
}

func (d *printData) printer(node *treeNode) {
	d.iter(node.value)
	fmt.Fprintf(d.w, "-%03d", node.size)
	if node.bal == equal {
		fmt.Fprintf(d.w, "=")
	} else if node.bal == leftHigh {
		fmt.Fprintf(d.w, "/")
	} else {
		fmt.Fprintf(d.w, "\\")
	}

	d.lineDepth += (d.itemSize + 5)

	if node.left != nil || node.right != nil {
		fmt.Fprintf(d.w, "-+-")

		if node.left != nil {
			d.lineDepth += 3
			d.array[d.index] = '|'
			d.index++
			d.printer(node.left)
			d.index--
			d.lineDepth -= 3
		} else {
			fmt.Fprintf(d.w, "\n")
		}

		for i := 0; i < d.lineDepth; i++ {
			if ((i + 2) % (d.itemSize + 3 + 5)) == 0 {
				fmt.Fprintf(d.w, "%c", (d.array[((i+2)/(d.itemSize+3+5))-1]))
			} else {
				fmt.Fprintf(d.w, " ")
			}
		}

		fmt.Fprintf(d.w, " +-")

		if node.right != nil {
			d.lineDepth += 3
			d.array[d.index] = ' '
			d.index++

			d.printer(node.right)
			d.index--
			d.lineDepth -= 3
		} else {
			fmt.Fprintf(d.w, "\n")
		}
	} else {
		fmt.Fprintf(d.w, "\n")
	}

	d.lineDepth -= (d.itemSize + 5)
}

// Print prints the values of the Tree to the given writer.
func (t *Tree) Print(w io.Writer, f IterateFunc, itemSiz int) {

	fmt.Fprintf(w, "treeNode-+-Left \t / Left High\n")
	fmt.Fprintf(w, "     |      \t = Equal\n")
	fmt.Fprintf(w, "     +-Right\t \\ Right High\n\n")

	maxHeight := t.Height()

	if f != nil && t.root != nil {
		d := &printData{0, itemSiz, make([]byte, maxHeight), 0, f, w}
		d.printer(t.root)
	}
}
