package main

import (
	"log"
	"sort"
)

/* Consider the sort.Interface type

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

*/

// Fraction represents a fraction...
type Fraction struct {
	n, d int
}

func (f Fraction) print() {
	r := float64(f.n) / float64(f.d)
	log.Printf("%2d/%2d = %5.3f", f.n, f.d, r)
}

type Fractions []Fraction

func (s Fractions) Len() int {
	return len(s)
}

func (s Fractions) Less(i, j int) bool {
	a := s[i]
	b := s[j]
	return a.n*b.d < b.n*a.d
}

func (s Fractions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var s Fractions = make([]Fraction, 5)
	s[0] = Fraction{1, 2}
	s[1] = Fraction{1, 4}
	s[2] = Fraction{7, 22}
	s[3] = Fraction{3, 4}
	s[4] = Fraction{2, 3}

	sort.Sort(s)
	for _, v := range s {
		v.print()
	}
}
