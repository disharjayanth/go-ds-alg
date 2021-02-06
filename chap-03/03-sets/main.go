package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

func (s *Set) Add(element int) {
	if !s.integerMap[element] {
		s.integerMap[element] = true
	}
}

func (s *Set) Contains(element int) bool {
	_, exists := s.integerMap[element]
	return exists
}

func (s *Set) Delete(element int) {
	if s.integerMap[element] {
		delete(s.integerMap, element)
	}
}

func (s *Set) Intersect(anotherSet *Set) *Set {
	newSet := &Set{}
	newSet.New()

	for key, _ := range s.integerMap {
		if anotherSet.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

func (s *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()

	for key, _ := range s.integerMap {
		unionSet.Add(key)
	}

	for key, _ := range anotherSet.integerMap {
		unionSet.Add(key)
	}

	return unionSet
}

func main() {
	set1 := &Set{}
	set1.New()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)
	set1.Add(5)

	fmt.Println("Set1:")
	fmt.Println(set1)

	if set1.Contains(1) {
		fmt.Println("Set1 contains given element 1")
	}

	set1.Delete(3)

	fmt.Println("Set1 after deleting 3:")
	fmt.Println(set1)

	fmt.Println("******************************************")

	set2 := &Set{}
	set2.New()

	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)
	set2.Add(7)

	fmt.Println("Set2:")
	fmt.Println(set2)

	interSectionSet := set1.Intersect(set2)
	fmt.Println("Interset of set1 and set2")
	fmt.Println(interSectionSet)

	unionSet := set1.Union(set2)
	fmt.Println(unionSet, len(unionSet.integerMap))
}
