package main

import (
	"fmt"
	"sort"
)

// A couple of type definitions to make the units clear.
type earthMass float64
type au float64

					// A Planet defines the properties of a solar system object.
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

					// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Planet) bool

					// Sort is a method on the function type, By, that sorts the argument slice according to the 
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,            // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)                   // run sort.Sort(<item>) -   func Sort(data Interface) where ps must support Interface interface.
}

					// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

					// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

					// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

					// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func main() {
	// Closures that order the Planet structure.
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return !distance(p1, p2)
	}

	// Sort the planets by the various criteria.
	By(name).Sort(planets)                       // convert name -> By which can be done as name is a function with same signature as type By.
 						     // this is a common pattern to enable types to inherit methods by conversion to another similar type with those methods
	fmt.Println("By name:", planets)             // name is now of type By, has now has method Sort in which we pass in a planet slice
						     // the rest is handled by calling the method Sort on By. THis method inturn calls sort.Sort and passes in its parameter.

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)

}