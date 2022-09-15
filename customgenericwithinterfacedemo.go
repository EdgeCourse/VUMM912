//generic for a general numeric

//before generics were introduced in Go 1.18, people used work-arounds:
//https://yourbasic.org/golang/generics/

package main

import (
	"fmt"
)

//add name of generic
type  interface {
	//add rules
	//~int //anything derived from int
}

//add generic type here
func Add {
	return a + b
}

func main() {
	fmt.Println("4 + 3 =", Add(4, 3))
	fmt.Println("4.1 + 3.2 =", Add(4.1, 3.2))
}
