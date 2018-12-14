package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3, 4}

	b := Map(a, func(val interface{}) interface{} {
		return val.(int) * 2
	})

//	fmt.Println("Map:", a, b)

	c := Reduce(b, 0, func(val interface{}, memo interface{}) interface{} {
		return memo.(int) + val.(int)
	})

	fmt.Println("Reduce:", b, c)

	d := Filter(b, func(val interface{}) bool {
		return val.(int)%4 == 0
	})

	fmt.Println("Filter:", b, d)
}

func Map(in interface{}, fn func(interface{}) interface{}) interface{} {
	val := reflect.ValueOf(in)
        fmt.Printf("%T\n",val);
	out := make([]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface())
	}

	return out
}

func Reduce(in interface{}, memo interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}

	return memo
}

func Filter(in interface{}, fn func(interface{}) bool) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()

		if fn(current) {
			out = append(out, current)
		}
	}

	return out
}
