package avltree

import "testing"

func TestBalanceFactorValues(t *testing.T) {
	if equal != 0 {
		t.Errorf("equal should be 0 but was %d", equal)
	}
}

func TestNodeInit(t *testing.T) {
    var nd treeNode[int64]

    if nd.bal != equal {
        t.Errorf("Balance should initialize to zero: %d", nd.bal)
    }

    if nd.size != 0 {
        t.Errorf("Size should initialize to zero: %d", nd.size)
    }
}
