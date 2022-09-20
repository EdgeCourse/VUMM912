/*
Method
It contains a receiver.
Methods of the same name but different types can be defined in the program.
It cannot be used as a first-order object.

Function
It does not contain a receiver.
Functions of the same name but different type are not allowed to be defined in the program.
It can be used as first-order objects and can be passed
*/

/*
Method with struct type receiver
In Go language, you are allowed to define a method whose receiver is of a struct type.
This receiver is accessible inside the method as shown in the below example:
*/

// Go program to illustrate the
// method with struct type receiver
/*
package main

import "fmt"

// Author structure
type author struct {
	name	 string
	branch string
	articles int
	salary int
}

// Method with a receiver
// of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.articles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:	 "Bob",
		branch: "West",
		articles: 203,
		salary: 34000,
	}

	// Calling the method
	res.show()
}
*/

/*
Method with Non-Struct Type Receiver
In Go language, you are allowed to create a method with non-struct type receiver as
long as the type and the method definitions are present in the same package.
If they present in different packages like int, string, etc, then the compiler will give
an error because they are defined in different packages.

*/

/*
// Go program to illustrate the method
// with non-struct type receiver
package main

import "fmt"

// Type definition
type data int

// Defining a method with
// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}


// if you try to run this code,
// then compiler will throw an error
//func(d1 int)multiply(d2 int)int{
//return d1 * d2
//}


// Main function
func main() {
	value1 := data(23)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}

*/

/*
Methods with Pointer Receiver
In Go language, you are allowed to create a method with a pointer receiver.
With the help of a pointer receiver, if a change is made in the method, it will reflect
in the caller which is not possible with the value receiver methods.


*/

/*
package main

import "fmt"

// Author structure
type author struct {
	name	 string
	branch string
	articles int
}

// Method with a receiver of author type
func (a *author) show(abranch string) {
	(*a).branch = abranch
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name: "Susan",
		branch: "West",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(Before): ", res.branch)

	// Creating a pointer
	p := &res

	// Calling the show method
	p.show("East")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(After): ", res.branch)
}

*/

//Method Can Accept both Pointer and Value
//As we know that in Go, when a function has a value argument, then it will only accept the values of the parameter,
//and if you try to pass a pointer to a value function, then it will not accept and vice versa. But a Go method can accept both value and pointer, whether it is defined with pointer or value receiver.

/*
package main

import "fmt"

// Author structure
type author struct {
    name   string
    branch string
}

// Method with a pointer
// receiver of author type
func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}

// Method with a value
// receiver of author type
func (a author) show_2() {

    a.name = "Bob"
    fmt.Println("Author's name(Before) : ", a.name)
}

// Main function
func main() {

    // Initializing the values
    // of the author structure
    res := author{
        name:   "Susan",
        branch: "West",
    }

    fmt.Println("Branch Name(Before): ", res.branch)

    // Calling the show_1 method
    // (pointer method) with value
    res.show_1("East")
    fmt.Println("Branch Name(After): ", res.branch)

    // Calling the show_2 method
    // (value method) with a pointer
    (&res).show_2()
    fmt.Println("Author's name(After): ", res.name)
}
*/