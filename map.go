package avltree

import (
	"constraints"
	"context"
	"io"
)

type Pair[K, V any] struct {
	Key K
	Value V
}

type Map[K, V any] struct {
	t Tree[Pair[K, V]]
}

// NewMap returns an initialized map.
func NewMap[K, V any](c func(K, K) int) *Map[K, V] {
    return &Map[K, V]{
		t: Tree[Pair[K, V]]{
       		compare: func(v1, v2 Pair[K, V]) int {
				return c(v1.Key, v2.Key)
			},
		},	
    }
}

// NewMapOrdered returns an initialized map using ordered types.
func NewMapOrdered[K constraints.Ordered, V any]() *Map[K, V] {
	 return &Map[K, V]{
		t: Tree[Pair[K, V]]{
			compare: func (v1, v2 Pair[K, V]) int {
				switch {
				case v1.Key < v2.Key:
					return -1
				case v1.Key == v2.Key:
					return 0
				default:
					return 1
				}
			},
		},
	}
}

// Clear removes all elements from the map, keeping the
// current options and compare function.
func (m *Map[K, V]) Clear() {
	m.t.Clear()
}

// Height returns the "height" of the map, meaning the
// number of levels.
func (m *Map[K, V]) Height() int {
	return m.t.Height()
}

// Len returns the number of elements in the map.
func (m *Map[K, V]) Len() int {
	return m.t.Len()
}

// Cap returns the capacity of the map; that is, the
// maximum elements the tree can hold with at the
// current height. This is only useful as a measure
// of how skewed the map is.
func (m *Map[K, V]) Cap() int {
	return m.t.Cap()
}

// At returns the value at the given index.
func (m *Map[K, V]) At(index int) *Pair[K, V] {
	return m.t.At(index)
}

// Find returns the element where the comparison function matches
// the node's value and the given key value.
func (m *Map[K, V]) Find(key K) *V {
	e := m.t.Find(Pair[K, V]{Key: key})
	if e != nil {
		return &e.Value
	}
	return nil
}

// Do calls function f for each element of the map, in order.
// The function should not change the structure of the map underfoot.
func (m *Map[K, V]) Do(f func(K, V) bool) {

	if f != nil && m.t.root != nil {
		iterateFunc[Pair[K, V]](func(e Pair[K, V]) bool {
			return f(e.Key, e.Value)
		}).iterate(m.t.root)
	}
}

// Iter returns a channel you can read through to fetch all the items.
func (m *Map[K, V]) Iter() <-chan Pair[K, V] {
	c := make(chan (Pair[K, V]))
	go m.t.chanIterate(context.Background(), c)
	return c
}

// IterContext returns a channel you can read through to fetch all the items.
func (m *Map[K, V]) IterContext(ctx context.Context) <-chan Pair[K, V] {
	c := make(chan (Pair[K, V]))
	go m.t.chanIterate(ctx, c)
	return c
}

// Keys returns all the keys as a slice.
func (m *Map[K, V]) Keys() []K {
	arr := make([]K, m.t.Len())
	var i int
	i = 0
	m.t.Do(func(v Pair[K, V]) bool {
		arr[i] = v.Key
		i++
		return true
	})
	return arr
}

// Values returns all the values as a slice.
func (m *Map[K, V]) Values() []V {
	arr := make([]V, m.t.Len())
	var i int
	i = 0
	m.t.Do(func(v Pair[K, V]) bool {
		arr[i] = v.Value
		i++
		return true
	})
	return arr
}

// Add adds an item to the map, returning a pair indicating the added
// (or duplicate) item, and a flag indicating whether the item is the
// duplicate that was found.
func (m *Map[K, V]) Add(k K, v V) (*V, bool) {
	p, dupe := m.t.Add(Pair[K, V]{Key: k, Value: v})
	if p != nil {
		return &p.Value, dupe
	}
	return nil, dupe
}

// Remove removes the element matching the given value.
func (m *Map[K, V]) Remove(key K) *V {
	p := m.t.Remove(Pair[K, V]{Key: key})
	if p != nil {
		return &p.Value
	}
	return nil
}

// RemoveAt removes the element at the given index.
func (m *Map[K, V]) RemoveAt(index int) *Pair[K, V] {
	return m.t.RemoveAt(index)
}

// PrintMap prints the values of the Map to the given writer.
func PrintMap[K, V any](m *Map[K, V], w io.Writer, f func(K, V) bool, itemSiz int) {
	Print(&m.t, w, func(p Pair[K, V]) bool {
		return f(p.Key, p.Value)
	}, itemSiz)
}
