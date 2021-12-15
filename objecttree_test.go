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

func TestObjectTree(t *testing.T) {

	// no duplicates tests

	tree := New[MyObject](func(a, b MyObject) int {
		if a.Key < b.Key {
			return -1
		}
		if a.Key > b.Key {
			return 1
		}
		return 0

	}, 0)

	if tree.root != nil {
		t.Errorf("Initialized tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != 0 {
		t.Errorf("Initialized tree flags should be zero: %v\n", tree.treeFlags)
	}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe := tree.Add(MyObject{"foo", "bar"})

	if v == nil || v.Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v = tree.At(0)
	if v == nil || v.Key != "foo" {
		t.Errorf("Single value should be foo: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v == nil || v.Key != "foo" || dupe != true {
		t.Errorf("Result of add should be foo/true: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(MyObject{"bar", "baz"})

	if v == nil || v.Key != "bar" || dupe != false {
		t.Errorf("Result of add should be bar/false: %v/%v\n", v, dupe)
	}

	v = tree.At(1)
	if v == nil || v.Key != "foo" {
		t.Errorf("Second value should be foo: %v\n", tree.At(0))
	}

	if tree.Len() != 2 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with two elements should be 2/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// reinitialize and allow duplicates

	tree = New[MyObject](func(a, b MyObject) int {
		if a.Key < b.Key {
			return -1
		}
		if a.Key > b.Key {
			return 1
		}
		return 0

	}, AllowDuplicates)

	if tree.root != nil {
		t.Errorf("Reinitialized tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != AllowDuplicates {
		t.Errorf("Reinitialized tree flags should be AllowDuplicates: %v\n", tree.treeFlags)
	}

	if tree.Len() != 0 || tree.Cap() != 0 || tree.Height() != 0 {
		t.Errorf("Initialized tree sizes should all be zero: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v == nil || v.Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	if tree.Len() != 1 || tree.Cap() != 1 || tree.Height() != 1 {
		t.Errorf("Tree sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	v = tree.At(0)
	if v == nil || v.Key != "foo" {
		t.Errorf("Single value should be foo: %v\n", tree.At(0))
	}

	v, dupe = tree.Add(MyObject{"foo", "baz"})

	if v == nil || v.Key != "foo" || dupe != false {
		t.Errorf("Result of add should be foo/false: %v/%v\n", v, dupe)
	}

	v, dupe = tree.Add(MyObject{"bar", "baz"})

	if v == nil || v.Key != "bar" || dupe != false {
		t.Errorf("Result of add should be bar/false: %v/%v\n", v, dupe)
	}

	v = tree.At(2)
	if v == nil || v.Key != "foo" {
		t.Errorf("Third value should be foo: %v\n", tree.At(0))
	}

	if tree.Len() != 3 || tree.Cap() != 3 || tree.Height() != 2 {
		t.Errorf("Tree sizes with three elements should be 3/3/2: %d Len, %d Cap, %d Height\n",
			tree.Len(), tree.Cap(), tree.Height())
	}

	// test Find
	v = tree.Find(MyObject{"bar", ""})

	if v == nil || v.Key != "bar" {
		t.Errorf("Find should locate bar: %v\n", v)
	}

	v = tree.Find(MyObject{"foobar", ""})

	if v != nil {
		t.Errorf("Found an item that isn't there: %v\n", v)
	}

	// test Data (slice)
	slice := tree.Data()

	if len(slice) != 3 || slice[2].Key != "foo" {
		t.Errorf("Slice is incorrect, expected 3/foo: %v/%v\n", len(slice), slice[2])
	}

	// test Do

	x := ""
	tree.Do(func(z MyObject) bool { x += z.Key; return true })

	if x != "barfoofoo" {
		t.Errorf("Do function did not concat values correctly, expected barfoofoo: %s\n", x)
	}

	// test Remove

	v, dupe = tree.Add(MyObject{"zoo", "berlin"})

	if v == nil || v.Key != "zoo" || dupe != false {
		t.Errorf("Result of add should be zoo/false: %v/%v\n", v, dupe)
	}

	v = tree.Remove(MyObject{"zoo", ""})

	if v == nil || v.Key != "zoo" || tree.Len() != 3 {
		t.Errorf("Result of Remove should be zoo/3: %v/%v\n", v, tree.Len())
	}

	// test Iter

	tree.Add(MyObject{"zeek", ""})
	tree.Add(MyObject{"jim", ""})
	tree.Add(MyObject{"zope", ""})
	tree.Add(MyObject{"yikes", ""})

	x = "aaaa"
	for v := range tree.Iter() {
		if v.Key < x {
			t.Error("Iter expected", v.Key, "to be >", x)
		}
		x = v.Key
	}

	if x != "zope" {
		t.Errorf("Iter ran wrong number of elements, expected last element of zope, got %s\n", x)
	}

	// test clear

	tree.Clear()

	if tree.root != nil {
		t.Errorf("Cleared tree root should be nil: %v\n", tree.root)
	}

	if tree.treeFlags != AllowDuplicates {
		t.Errorf("Cleared tree flags should still be AllowDuplicates: %v\n", tree.treeFlags)
	}

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
	tree.Do(func(elem MyObject) bool {
		var cur string
		cur = elem.Key
		if prev > cur {
			t.Errorf("Elements not in order, previous = %s, current = %s\n", prev, cur)
			return false
		}
		prev = cur
		return true
	})

}
