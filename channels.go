/*
“Do not communicate by sharing memory; instead, share memory by communicating.”
This means that instead of struggling with complex mutex situations in shared memory, use channels to communicate goroutines.
But… why?
When sending a message to a channel, only one goroutine will receive it, so it is safe to access the data from there and no explicit synchronization is needed since it is handled by Go under the hood.
Following that approach, we will have a goroutine that keeps the state of our counter, and on the other side, other goroutines will send the messages that the first one will receive for interacting with the state. We have two types of messages:
incrementOp: Operation that requests incrementing the counter.
valueOp: Operation that requests the value of the counter

Note that all the operations have a res channel with the following purposes:
Receiving the operation’s response
Synchronizing goroutines
As you can see, this strategy is not based on sharing memory like the previous ones but instead relies on sending operations through channels.

use it when mutex is not an option

*/

package main

import (
	"log"
	"sync"
)

type op struct {
	res chan int
}

type incrementOp struct {
	op
}

type getValueOp struct {
	op
}

func newIncrementOp() incrementOp {
	return incrementOp{
		op: op{
			res: make(chan int),
		},
	}
}

func newGetValueOp() getValueOp {
	return getValueOp{
		op: op{
			res: make(chan int),
		},
	}
}

func increment(ops chan<- incrementOp, wg *sync.WaitGroup) {
	defer wg.Done()

	op := newIncrementOp()
	ops <- op
	<-op.res
}

func main() {
	incrementOps := make(chan incrementOp)
	getValueOps := make(chan getValueOp)

	go func() {
		counter := 0
		for {
			select {
			case op := <-incrementOps:
				counter++
				op.res <- counter
			case op := <-getValueOps:
				op.res <- counter
			}
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(incrementOps, &wg)
		go increment(incrementOps, &wg)
	}

	wg.Wait()

	getValueOp := newGetValueOp()
	getValueOps <- getValueOp

	log.Printf("Counter: %d", <-getValueOp.res)
}
