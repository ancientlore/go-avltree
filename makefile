include $(GOROOT)/src/Make.inc

TARG=container/avltree
GOFILES=\
	node.go\
	tree.go\
	treeadd.go\
	treeprint.go\
	treeremove.go\
	stringtree.go

include $(GOROOT)/src/Make.pkg

