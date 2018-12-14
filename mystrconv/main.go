// Parsing numbers from strings is a basic but common task
// in many programs; here's how to do it in Go.

package main

// The built-in package `strconv` provides the number
// parsing.
import  (
        "strconv"
        "fmt"
        "unsafe"
)

func main() {

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For `ParseInt`, the `0` means infer the base from
	// the string. `64` requires that the result fit in 64
	// bits.
	i, _ := strconv.ParseInt("123", 0, 16)
	fmt.Printf("strconv.ParseInt(123:   %T %v %d\n",i,i,unsafe.Sizeof(i))


	// `ParseInt` will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A `ParseUint` is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
        fmt.Printf("strconv.ParseUint(789:   %T %v %d\n",u,u,unsafe.Sizeof(u))

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.
	k, _ := strconv.Atoi("135")
        fmt.Printf("strconv.Atoi(135:   %T %v %d\n",k,k,unsafe.Sizeof(k))

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
