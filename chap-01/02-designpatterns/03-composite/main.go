package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet:", leaf.name)
}

type Branch struct {
	name     string
	leafs    []Leaflet
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch:", branch.name)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) addLeaf(newleaf *Leaflet) {
	branch.leafs = append(branch.leafs, *newleaf)
}

func (branch *Branch) addBranch(newbranch *Branch) {
	branch.branches = append(branch.branches, *newbranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	branch1 := &Branch{
		name: "Branch 1",
	}

	leaf1 := &Leaflet{
		name: "Leaf 1",
	}

	leaf2 := &Leaflet{
		name: "Leaf 2",
	}

	branch2 := &Branch{
		name: "Branch 2",
	}

	branch1.addLeaf(leaf1)
	branch1.addLeaf(leaf2)

	branch1.addBranch(branch2)

	branch1.perform()
}
