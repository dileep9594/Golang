package main

// Questinon : How can you sort a slice of custom structs with the help of an example?
// Answer : We can sort slices of custom structs by using sort.Sort and sort.Stable functions. These methods sort any collection that implements sort.Interface interface that has Len(), Less() and Swap() methods as shown in the code below:
type Human struct {
	name string
	age  int
}

type ageFactor []Human

func (a ageFactor) Len() int {
	return len(a)
}

func (a ageFactor) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a ageFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
