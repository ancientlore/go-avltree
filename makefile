include $(GOROOT)/src/Make.$(GOARCH)

TARG=container/avltree
GOFILES=\
	node.go\
	tree.go\
	treeadd.go\
	treeprint.go\
	treeremove.go\
	stringtree.go

include $(GOROOT)/src/Make.pkg

