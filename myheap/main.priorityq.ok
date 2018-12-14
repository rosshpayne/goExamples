// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int //
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item                                    // slice of pointers to struct  p === (*p) in Go compiler when accessing struct via pointer.
							      // access fields of struct using p.field or (*p).field.
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here. The larger the value the greater the priority
	return pq[i].priority >  pq[j].priority                 
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)                     // pass pointer pq as receiver is pointer which satisfes Interface. 
}

func indirect(x *PriorityQueue, item *Item, value string, priority int) {
        fmt.Printf("indirect x address: %p\n",&x)
        x.update(item , value , priority )            // Go auto &x  x is []*item.  Different pointer by same slice, len,cap & *data
        *x=(*x)[:len(*x)-2]
        fmt.Println("x.Len = ",x.Len())
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 4, "apple": 6, "pear": 2,"fruit": 9,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))        // declare pq and allocate memory
	i := 0
	for value, priority := range items {
		pq[i] = &Item{                       // initialise  pq.  
			value:    value,
			priority: priority,
			index:    i,                 // index entry in slice. Why?
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 5,
	}
        for i:= range pq {
                    fmt.Printf("  %d  %#v\n", i, pq[i] )
        }
        fmt.Printf("\n Less (1,2) through non-pointer:  %v\n",pq.Less(1,2))
        pqp:=&pq 
        fmt.Printf("Pointer: %p\n",pqp)
        fmt.Printf(" Less (1,2) through     pointer:  %v\n",pqp.Less(1,2))    // will GO auto dereference ie. (*p).Less.  Yes it does.
        fmt.Printf(" Len through non-pointer %d  pointer %d \n\n",pq.Len(),pqp.Len())

	heap.Push( pq, item)                         // pass in pointer as Interface receiver is pointer on some methods.
                                                     // note: Go will automatically &pq if receiver is pointer & you pass in non-pointer, BUT for arguments
                                                     //       this is not the case as arguments passed by value and &value is a different pointer.
				 		     //       so this auto &pq does not work when passed through function arguments.
                                                     //       similarly Go will *pq if you pass in pointer and receiver is non-pointer. Pass pointer is the more generic case.
        for i:= range pq { 			     // pointer receiver as slice len is changed and there is no return value to update caller.
                    fmt.Printf("after push priority 5    %d  %#v\n", i, pq[i] )
        }
	pq.update(item, item.value, 8)               // Go will &pq automatically - as receiver must be a pointer to pq.
        for i:= range pq {
                    fmt.Printf(" update priority 5->8       %d  %#v\n", i, pq[i] )
        }
	pqp.update(item, item.value, 10)               // using pointer receiver this time - as it should be.
        indirect(pqp,item,item.value,11)
        fmt.Println("pq.Len = ",pq.Len())

	// Take the items out; they arrive in decreasing priority order.
	for pqp.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Pop %.2d:%s  %d\n ", item.priority, item.value, item.index)
                for i:= range pq {
                    fmt.Printf("  %d  %#v\n", i, pq[i] )
                }
        }
}
