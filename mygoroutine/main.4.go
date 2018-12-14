package main

import (
        "time"
        "fmt"
      )
type Ball struct{ hits int }

func main() {
    done := make(chan struct{})
    defer close(done)
    table := make(chan *Ball)
    go player("ping", table, done)
    go player("pong", table, done)

    table <- new(Ball) // game on; toss the ball
    time.Sleep(1 * time.Second)
    <-table // game over; grab the ball

     panic("show me the stacks")
}

func player(name string, table chan *Ball, done chan struct{}) {
    var ball *Ball
    for {
	select {
	case <-done: 
		return
        case ball = <-table:
	}
        ball.hits++
        fmt.Println(name, ball.hits)
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}
