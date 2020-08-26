package avltree

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := NewMapOrdered[string, int]()

	if m.t.root != nil {
		t.Errorf("Initialized map root should be nil: %v\n", m.t.root)
	}

	if m.t.treeFlags != 0 {
		t.Errorf("Initialized map flags should be zero: %v\n", m.t.treeFlags)
	}

	if m.t.Len() != 0 || m.t.Cap() != 0 || m.t.Height() != 0 {
		t.Errorf("Initialized map sizes should all be zero: %d Len, %d Cap, %d Height\n",
			m.t.Len(), m.t.Cap(), m.t.Height())
	}

	v, dupe := m.Add("14", 14)

	if v == nil || *v != 14 || dupe != false {
		t.Errorf("Result of add should be 14/false: %v/%v\n", v, dupe)
	}

	if m.t.Len() != 1 || m.t.Cap() != 1 || m.t.Height() != 1 {
		t.Errorf("Map sizes with one element should all be one: %d Len, %d Cap, %d Height\n",
			m.t.Len(), m.t.Cap(), m.t.Height())
	}

	p := m.At(0)
	if p == nil || p.Key != "14" {
		t.Errorf("Single value should be 14: %v\n", p)
	}

	v, dupe = m.Add("14", 14)

	if v == nil || *v != 14 || dupe != true {
		t.Errorf("Result of add should be 14/true: %v/%v\n", v, dupe)
	}

	v, dupe = m.Add("15", 15)

	if v == nil || *v != 15 || dupe != false {
		t.Errorf("Result of add should be 15/false: %v/%v\n", v, dupe)
	}

	p = m.At(1)
	if p == nil || p.Key != "15" {
		t.Errorf("Second value should be 15: %v\n", p)
	}

	if m.t.Len() != 2 || m.t.Cap() != 3 || m.t.Height() != 2 {
		t.Errorf("Map sizes with two elements should be 2/3/2: %d Len, %d Cap, %d Height\n",
			m.t.Len(), m.t.Cap(), m.t.Height())
	}

	x := "13"
	for v := range m.Iter() {
		if v.Key <= x {
			t.Error("Iter expected", v, "to be >", x)
		}
		x = v.Key
	}

}

func TestPrintMap(t *testing.T) {
	var names = []string{
		"Bob",
		"Sue",
		"Billy",
		"Barbara",
		"X Æ A-12",
		"Atul",
		"Sandra",
		"Joseph",
		"Martin",
		"Ronald",
		"Lara",
		"Jennifer",
	}
	rand.Seed(time.Now().UnixNano())
	m := NewMapOrdered[rune, string]()
	for i := 0; i < 50; i++ {
		m.Add('A'+rune(rand.Intn(26)), names[rand.Intn(len(names))])
		m.Add('a'+rune(rand.Intn(26)), names[rand.Intn(len(names))])
		m.Add('0'+rune(rand.Intn(10)), names[rand.Intn(len(names))])
	}

	var buf bytes.Buffer
	PrintMap(m, &buf, func(c rune, s string) bool {
		fmt.Fprintf(&buf, "\"%c:%-8s\"", c, s)
		return true
	}, 12)

	t.Log(buf.String())
	fmt.Println(buf.String())
}
