//when to use pointers in Go
//https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac

//Go is not object-oriented programming, as it does not have inheritance
//composition, embedding and interfaces support code reuse and polymorphism in Go

/*
Decoupling
Golang does not have a concept of object-oriented programming (OOP), the concept of
class or a class-based inheritance. Golang has its own way to interpret the regular
object-oriented languages - decoupling.

In traditional OOP design, one class contains member variables, either public or private.
It also contains member functions, either public or private. In order to reuse the class
member function, another class has to inherit the class. Golang's approach makes design more
efficient: Decouple a concrete data (struct) with a set of member functions (interface).

Decoupling separates concrete data and behaviors. The concrete data can be defined as an
object data structure such as struct. The behaviors are defined as a set of member functions
in an interface.

We can implement the same interface methods for different concrete data, either with a
value receiver, or with a pointer receiver.
*/

//Go’s structs are typed collections of fields. They’re useful for grouping data together
//to form records.

//decoupling: struct v method
//composition: method

//a struct is like a structure, form, blueprint
/*
To define a new struct type, you list the names and types of each field.
The default zero value of a struct has all its fields zeroed.
You can access individual fields with dot notation.

Can use the keyword new https://yourbasic.org/golang/structs-explained/
*/

/*
Struct Field Exporting

Fields of a struct follow the same exporting rules as other identifiers within the Go
programming language. If a field name begins with a capital letter, it will be readable and
writeable by code outside of the package where the struct was defined. If the field begins
with a lowercase letter, only code within that struct’s package will be able to read and write
that field. This example defines fields that are exported and those that are not:
*/

//EXAMPLE STRUCT

package main

import "fmt"

//aligns when saved

type Person struct {
	Name string
	Type string

	subject string
}

func main() {
	//need comma at end
	//https://dave.cheney.net/2014/10/04/that-trailing-comma

	//because of the semicolon rule
	//https://go.dev/doc/effective_go#semicolons

	p := Person{
		Name: "Bob",
		Type: "student",

		subject: "Golang",
	}

	//access fields with dot notation
	fmt.Println(p.Name, "is a ", p.Type)
	fmt.Println("Subject is", p.subject)

	//This syntax creates a new struct.
	/*
		p2 := Person{"Emily", "student", "Golang"}

		fmt.Println(p2.Name, "is a", p2.Type, "subject is", p2.subject)

		//mutable
		p2.Name = "Amelia"
		//p2.subject = "Perl"
		fmt.Println(p2.Name, "is a", p2.Type, "subject is", p2.subject)
	*/
	//	we can name the fields when initializing a struct.
	/*
		p3 := Person{Type: "student", subject: "JSON", Name: "Alice"}

		fmt.Println(p3.Name, "is a", p3.Type, "subject is", p3.subject)
	*/
	//	Omitted fields will be zero-valued.

	/*
		p4 := Person{Name: "Fred"}

		fmt.Println(p4.Name, "is a", p4.Type, "subject is", p4.subject)

		p5 := Person{Name: "Fred"}
	*/

}
