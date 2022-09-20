package main

import "fmt"

func main() {

	// defer the execution of Println() function
	defer fmt.Println("Three")

	//use defer
	//LIFO
	fmt.Println("One")
	fmt.Println("Two")

}
