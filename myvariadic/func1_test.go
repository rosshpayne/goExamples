package testvar

import "testing"
import "log"

func TestVariadic(t *testing.T) {
	log.Println("Func()")
	Func1()
	log.Println("Func(2)")
	Func1(2)
	log.Println("Func(2,3,4)")
	Func1(2,3,4)
}



func TestVariadicChan(t *testing.T) {
	
	var done chan struct{}
        var xs []int
	switch {
	   case done == nil :
		log.Println("var done chan struct{} results in  done being nil")
	   default:
		log.Printf("** done is: %#v",done)
	}
	switch {
	   case xs == nil :
		log.Println("var xs []int results in xs being nil")
	   default:
		log.Printf("xs is not nil: %#v",done)
	}

	log.Println("Func()")
	Func2()

	log.Println("Func(done)")
	Func2(done)

	done = make(chan struct{});
	log.Printf("done is:  %#v",done)
	log.Println("Func(done)")
	Func2(done)

	log.Println("Func(nil)")
	Func2(nil)

	done2 := make(chan int)
	done3 := make(chan struct{})
	log.Println("Func2(done, done2, done3)")
	Func2(done, done2, done3)
}
