package avltree

import (
	"fmt"
	"math/rand"
	"testing"
)

type MyObject struct {
	Key   string
	Value string
}

func (o MyObject) Compare(b Interface) int {
	if o.Key < b.(MyObject).Key {
		return -1
	}
	if o.Key > b.(MyObject).Key {
		return 1
	}
	return 0
}

func TestObjectTree(t *testing.T) {

	// no duplicates tests

	tree := NewObjectTree(0)

	if tree.Tree.root != nil {
		t.Errorf("Initialized tree root should be nil: %v\n", tree.Tree.root)
	}

	if tree.Tree.treeFlags != 0 {
		t.Errorf("Initialized tree flags should be zero: %v\n", tree.Tree.treeFlags)
	}

	//if tree.Tree.compare != objectCompare {
	//	t.Errorf("Initialized tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe := tree.Add(MyObject{"foo", "bar"})

	if v.(MyObject).Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0).(MyObject).Key != "foo" {
		t.Errorf("Single value should be foo: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v.(MyObject).Key != "foo" || dupe != true {
		t.Errorf("Result of add should be foo/true: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(MyObject{"bar", "baz"})

	if v.(MyObject).Key != "bar" || dupe != false {
		t.Errorf("Result of add should be bar/false: %v/%v\n", v, dupe)
	}

	if tree.At(1).(MyObject).Key != "foo" {
		t.Errorf("Second value should be foo: %v\n", tree.At(0))
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

	//if tree.Tree.compare != objectCompare {
	//	t.Errorf("Reinitialized tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v.(MyObject).Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	if tree.At(0).(MyObject).Key != "foo" {
		t.Errorf("Single value should be foo: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v.(MyObject).Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(MyObject{"bar", "baz"})

	if v.(MyObject).Key != "bar" || dupe != false {
		t.Errorf("Result of add should be bar/false: %v/%v\n", v, dupe)
	}

	if tree.At(2).(MyObject).Key != "foo" {
		t.Errorf("Third value should be foo: %v\n", tree.At(0))
	}

	if tree.Len() != 3 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with three elements should be 3/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// test Find
	v = tree.Find(MyObject{"bar", ""})

	if v.(MyObject).Key != "bar" {
		t.Errorf("Find should locate bar: %v\n", v)
	}

	v = tree.Find(MyObject{"foobar", ""})

	if v != nil {
		t.Errorf("Found an item that isn't there: %v\n", v)
	}

	// test Data (slice)
	slice := tree.Data()

	if len(slice) != 3 || slice[2].(MyObject).Key != "foo" {
		t.Errorf("Slice is incorrect, expected 3/foo: %v/%v\n", len(slice), slice[2])
	}

	// test Do

	x := ""
	tree.Do(func(z interface{}) { x += z.(MyObject).Key })

	if x != "barfoofoo" {
		t.Errorf("Do function did not concat values correctly, expected barfoofoo: %d\n", x)
	}

	// test Remove

	v, dupe = tree.Add(MyObject{"zoo", "berlin"})

	if v.(MyObject).Key != "zoo" || dupe != false {
		t.Errorf("Result of add should be zoo/false: %v/%v\n", v, dupe)
	}

	v = tree.Remove(MyObject{"zoo", ""})

	if v.(MyObject).Key != "zoo" || tree.Len() != 3 {
		t.Errorf("Result of Remove should be zoo/3: %v/%v\n", v, tree.Len())
	}

	// test Iter

	tree.Add(MyObject{"zeek", ""})
	tree.Add(MyObject{"jim", ""})
	tree.Add(MyObject{"zope", ""})
	tree.Add(MyObject{"yikes", ""})

	x = "aaaa"
	for v := range tree.Iter() {
		if v.(MyObject).Key < x {
			t.Error("Iter expected", v.(MyObject).Key, "to be >", x)
		}
		x = v.(MyObject).Key
	}

	if x != "zope" {
		t.Errorf("Iter ran wrong number of elements, expected last element of zope, got %s\n", x)
	}

	// test clear

	tree.Clear()

	if tree.Tree.root != nil {
		t.Errorf("Cleared tree root should be nil: %v\n", tree.Tree.root)
	}

	if tree.Tree.treeFlags != AllowDuplicates {
		t.Errorf("Cleared tree flags should still be AllowDuplicates: %v\n", tree.Tree.treeFlags)
	}

	//if tree.Tree.compare != objectCompare {
	//	t.Errorf("Cleared tree compare function not correct: %v\n", tree.Tree.compare)
	//}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Cleared tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// one more test, larger data

	for j := 0; j < 100000; j++ {
		tree.Add(MyObject{fmt.Sprintf("%d", rand.Int()), "Hello"})
	}

	if tree.Len() != 100000 {
		t.Errorf("Expected 100000 elements, got %d\n", tree.Len())
	}

	var prev string
	prev = ""

	// make sure elements sorted
	tree.Do(func(elem interface{}) {
		var cur string
		cur = elem.(MyObject).Key
		if prev > cur {
			t.Errorf("Elements not in order, previous = %d, current = %d\n", prev, cur)
		}
		prev = cur
	})

}
