package main

import (
	"log"
	"sort"
	"time"
)


	type rowT  []interface{}		

        type customSort struct {
		tab  []rowT						// see example with []*rowT for more efficient copy in Swap.  This copies slice meta data [*data,len,cap]
		less  func(i,j rowT, order []int) bool
		order []int                    // column order of sort
	}

	func (tb customSort) Len() int           { return len(tb.tab) }					// methods must be defined at package level
	func (tb customSort) Swap(i,j int )      { tb.tab[i], tb.tab[j] = tb.tab[j], tb.tab[i] }        //  hence its associated type must be defined at package level
	func (tb customSort) Less(i,j int ) bool { return tb.less(tb.tab[i], tb.tab[j], tb.order) }

func main () {


	loc, _ := time.LoadLocation("Australia/Sydney") 

	var cs []rowT  = []rowT{ { 10, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 8, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 2, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 6, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "EBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "JBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "CBC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,14,13,14,29,loc) },
	                           { 13, "ABC", time.Date(2018,5,18,8,13,14,29,loc) },
	                           { 20, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 17, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 12, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
	                           { 4, "ABC", time.Date(2018,5,18,12,13,14,29,loc) },
				}
	c2:=cs								// copy struct into new var. NB its not a copy but a reference to cs, shares underlying array.
	c3:=make([]rowT,15)
	copy(c3,c2)
	//if c3==c2 { log.Println("c3==c2 true"); }			// cannot compare slices
 	for i,r := range c2 { log.Printf("%d   %#v", i, r) }
	log.Println()

	sort.Sort(customSort{ tab: cs, less: func(i,j rowT, order []int) bool { // use struct literal to initialise interface{}  arg in sort.Sort
	    for r := range order  {
		vi:=i[r]
		vj:=j[r]
		switch v1:=vi.(type) {
		case int:
			v2:=vj.(int)
			if  v1 != v2 { return v1<v2 }
 		case int64:
			v2:=vj.(int64)
			if  v1 != v2 { return v1<v2 }
		case float64:
			v2:=vj.(float64)
			if  v1 != v2 { return v1<v2 }
		case string:
			v2:=vj.(string)
			if  v1 != v2 { return v1<v2 }
		case time.Time:
			v2:=vj.(time.Time)
			if  v1 != v2 { return v1.Before(v2) }
		}
	    }
	    return false
	},order : []int{2,1,3}} )

 	for i,r := range cs { log.Printf("%d   %#v", i, r) }
	log.Println() 
/*
 	for i,r := range c2 { log.Printf("%d   %#v", i, r) }    // check values are not sorted as its a copy
	log.Println() 
 	for i,r := range c3 { log.Printf("%d   %#v", i, r) }
*/
}
