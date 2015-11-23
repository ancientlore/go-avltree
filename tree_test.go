package avltree

import (
	"math/rand"
	"testing"
)

func compareInt(a interface{}, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) > b.(int) {
		return 1
	}
	return 0
}

func TestTree(t *testing.T) {

	// no duplicates tests

	tree := New(compareInt, 0)

	if tree.root != nil {
		t.Errorf("Initialized tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != 0 {
		t.Errorf("Initialized tree flags should be zero: %v\n", tree.treeFlags)
	}

	//if tree.compare != compareInt {
	//	t.Errorf("Initialized tree compare function not correct: %v\n", tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe := tree.Add(14)

	if v.(int) != 14 || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0).(int) != 14 {
		t.Errorf("Single value should be 14: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(14)

	if v.(int) != 14 || dupe != true {
		t.Errorf("Result of add should be 14/true: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(15)

	if v.(int) != 15 || dupe != false {
		t.Errorf("Result of add should be 15/false: %v/%v\n", v, dupe)
	}

	if tree.At(1).(int) != 15 {
		t.Errorf("Second value should be 15: %v\n", tree.At(0))
	}

	if tree.Len() != 2 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with two elements should be 2/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// reinitialize and allow duplicates

	tree.Init(compareInt, AllowDuplicates)

	if tree.root != nil {
		t.Errorf("Reinitialized tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != AllowDuplicates {
		t.Errorf("Reinitialized tree flags should be AllowDuplicates: %v\n", tree.treeFlags)
	}

	//if tree.compare != compareInt {
	//	t.Errorf("Reinitialized tree compare function not correct: %v\n", tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe = tree.Add(14)

	if v.(int) != 14 || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0).(int) != 14 {
		t.Errorf("Single value should be 14: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(14)

	if v.(int) != 14 || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(15)

	if v.(int) != 15 || dupe != false {
		t.Errorf("Result of add should be 15/false: %v/%v\n", v, dupe)
	}

	if tree.At(2).(int) != 15 {
		t.Errorf("Third value should be 15: %v\n", tree.At(0))
	}

	if tree.Len() != 3 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with three elements should be 3/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// test Find

	v = tree.Find(15)

	if v.(int) != 15 {
		t.Errorf("Find should locate 15: %v\n", v)
	}

	v = tree.Find(90)

	if v != nil {
		t.Errorf("Found an item that isn't there: %v\n", v)
	}

	v = tree.Find(nil)
	if v != nil {
		t.Errorf("Found an item that is not there: %v\n", v)
	}

	// test Data (slice)
	slice := tree.Data()

	if len(slice) != 3 || slice[2].(int) != 15 {
		t.Errorf("Slice is incorrect, expected 3/15: %v/%v\n", len(slice), slice[2].(int))
	}

	// test Do

	x := 0
	tree.Do(func(z interface{}) bool { x += z.(int); return true })

	if x != 43 {
		t.Errorf("Do function did not add up values correctly, expected 43: %d\n", x)
	}

	// test Remove

	v, dupe = tree.Add(16)

	if v.(int) != 16 || dupe != false {
		t.Errorf("Result of add should be 16/false: %v/%v\n", v, dupe)
	}

	v = tree.Remove(14)

	if v.(int) != 14 || tree.Len() != 3 {
		t.Errorf("Result of Remove should be 14/3: %v/%v\n", v, tree.Len())
	}

	// test Iter

	tree.Add(20)
	tree.Add(19)
	tree.Add(18)
	tree.Add(17)

	x = 14
	for v := range tree.Iter() {
		if v.(int) != x {
			t.Error("Iter expected", x, "got", v.(int))
		}
		x++
	}

	if x != 21 {
		t.Errorf("Iter ran wrong number of elements, expected 7, got %d\n", x-14)
	}

	// test clear

	tree.Clear()

	if tree.root != nil {
		t.Errorf("Cleared tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != AllowDuplicates {
		t.Errorf("Cleared tree flags should still be AllowDuplicates: %v\n", tree.treeFlags)
	}

	//if tree.compare != compareInt {
	//	t.Errorf("Cleared tree compare function not correct: %v\n", tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Cleared tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// one more test, larger data

	for j := 0; j < 100000; j++ {
		tree.Add(rand.Int())
	}

	if tree.Len() != 100000 {
		t.Errorf("Expected 100000 elements, got %d\n", tree.Len())
	}

	var prev int
	prev = -1

	// make sure elements sorted
	tree.Do(func(elem interface{}) bool {
		var cur int
		cur = elem.(int)
		if prev > cur {
			t.Errorf("Elements not in order, previous = %d, current = %d\n", prev, cur)
		}
		prev = cur
		return true
	})

}
