[![Build Status](https://api.travis-ci.com/ancientlore/go-avltree.svg?branch=v2)](https://app.travis-ci.com/ancientlore/go-avltree)
[![Coverage Status](https://coveralls.io/repos/ancientlore/go-avltree/badge.svg?branch=v2)](https://coveralls.io/github/ancientlore/go-avltree?branch=v2)
[![pkg.go.dev](https://pkg.go.dev/github.com/ancientlore/go-avltree/v2?status.png)](https://pkg.go.dev/github.com/ancientlore/go-avltree/v2)
[gocover](http://gocover.io/github.com/ancientlore/go-avltree/v2)

An [AVL tree](http://en.wikipedia.org/wiki/AVL_tree) (Adel'son-Vel'skii & Landis) is a binary search tree in which the heights of the left and right subtrees of the root differ by at most one and in which the left and right subtrees are again AVL trees.

With each node of an AVL tree is associated a balance factor that is Left High, Equal, or Right High according, respectively, as the left subtree has height greater than, equal to, or less than that of the right subtree.

The AVL tree is, in practice, balanced quite well. It can (at the worst case) become skewed to the left or right, but never so much that it becomes inefficient. The balancing is done as items are added or deleted.

This version is enhanced to allow "indexing" of values in the tree; however, the indexes are not stable as the tree could be resorted as items are added or removed.

It is safe to iterate or search a tree from multiple threads provided that no threads are modifying the tree.

The tree works on generic types and there is also a specialization for maps. Additionally, the tree supports iteration and a channel iterator.

	t.Do(func(z int) bool {
		if z % 3333 == 0 {
			fmt.Printf("%d ", z)
		}
		return true
	})

	for v := range t.Iter() {
        	if v % 3333 == 0 {
                	fmt.Printf("%d ", v);
        	}
	}

To install, you can use:

	go get github.com/ancientlore/go-avltree/v2@latest

For more information on generics see:

* https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md
* https://go2goplay.golang.org/
* https://blog.golang.org/generics-next-step
* https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md
