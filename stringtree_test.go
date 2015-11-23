package avltree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestStringTree(t *testing.T) {

	// no duplicates tests

	tree := NewStringTree(0)

	if tree.Tree.root != nil {
		t.Errorf("Initialized tree root should be nil: %v\n", tree.Tree.root)
	}

	if tree.Tree.treeFlags != 0 {
		t.Errorf("Initialized tree flags should be zero: %v\n", tree.Tree.treeFlags)
	}

	//if tree.Tree.compare != stringCompare {
	//	t.Errorf("Initialized tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe := tree.Add("14")

	if v != "14" || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0) != "14" {
		t.Errorf("Single value should be 14: %v\n", tree.At(0))
	}

	v, dupe = tree.Add("14")

	if v != "14" || dupe != true {
		t.Errorf("Result of add should be 14/true: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add("15")

	if v != "15" || dupe != false {
		t.Errorf("Result of add should be 15/false: %v/%v\n", v, dupe)
	}

	if tree.At(1) != "15" {
		t.Errorf("Second value should be 15: %v\n", tree.At(0))
	}

	if tree.Len() != 2 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with two elements should be 2/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// reinitialize and allow duplicates

	tree.Init(AllowDuplicates)

	if tree.Tree.root != nil {
		t.Errorf("Reinitialized tree root should be nil: %v\n", tree.Tree.root)
	}

	if tree.Tree.treeFlags != AllowDuplicates {
		t.Errorf("Reinitialized tree flags should be AllowDuplicates: %v\n", tree.Tree.treeFlags)
	}

	//if tree.Tree.compare != stringCompare {
	//	t.Errorf("Reinitialized tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe = tree.Add("14")

	if v != "14" || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0) != "14" {
		t.Errorf("Single value should be 14: %v\n", tree.At(0))
	}

	v, dupe = tree.Add("14")

	if v != "14" || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add("15")

	if v != "15" || dupe != false {
		t.Errorf("Result of add should be 15/false: %v/%v\n", v, dupe)
	}

	if tree.At(2) != "15" {
		t.Errorf("Third value should be 15: %v\n", tree.At(0))
	}

	if tree.Len() != 3 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with three elements should be 3/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// test Find
	var st string
	if st != "" {
		t.Errorf("Init string != empty string\n", st)
	}

	v = tree.Find("15")

	if v != "15" {
		t.Errorf("Find should locate 15: %v\n", v)
	}

	v = tree.Find("90")

	if v != "" {
		t.Errorf("Found an item that isn't there: %v\n", v)
	}

	// test Data (slice)
	slice := tree.Data()

	if len(slice) != 3 || slice[2] != "15" {
		t.Errorf("Slice is incorrect, expected 3/15: %v/%v\n", len(slice), slice[2])
	}

	// test Do

	x := ""
	tree.Do(func(z string) bool { x += z; return true })

	if x != "141415" {
		t.Errorf("Do function did not concatvalues correctly, expected 141415: %d\n", x)
	}

	// test Remove

	v, dupe = tree.Add("16")

	if v != "16" || dupe != false {
		t.Errorf("Result of add should be 16/false: %v/%v\n", v, dupe)
	}

	v = tree.Remove("14")

	if v != "14" || tree.Len() != 3 {
		t.Errorf("Result of Remove should be 14/3: %v/%v\n", v, tree.Len())
	}

	// test Iter

	tree.Add("20")
	tree.Add("19")
	tree.Add("18")
	tree.Add("17")

	x = "13"
	for v := range tree.Iter() {
		if v <= x {
			t.Error("Iter expected", v, "to be >", x)
		}
		x = v
	}

	if x != "20" {
		t.Errorf("Iter ran wrong number of elements, expected last element of 20, got %d\n", x)
	}

	// test clear

	tree.Clear()

	if tree.Tree.root != nil {
		t.Errorf("Cleared tree root should be nil: %v\n", tree.Tree.root)
	}

	if tree.Tree.treeFlags != AllowDuplicates {
		t.Errorf("Cleared tree flags should still be AllowDuplicates: %v\n", tree.Tree.treeFlags)
	}

	//if tree.Tree.compare != stringCompare {
	//	t.Errorf("Cleared tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Cleared tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// one more test, larger data

	for j := 0; j < 100000; j++ {
		tree.Add(fmt.Sprintf("%d", rand.Int()))
	}

	if tree.Len() != 100000 {
		t.Errorf("Expected 100000 elements, got %d\n", tree.Len())
	}

	var prev string
	prev = ""

	// make sure elements sorted
	tree.Do(func(elem string) bool {
		var cur string
		cur = elem
		if prev > cur {
			t.Errorf("Elements not in order, previous = %d, current = %d\n", prev, cur)
		}
		prev = cur
		return true
	})

}
