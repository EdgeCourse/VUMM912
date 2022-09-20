/*
package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

// main function
func main() {

	// assigns value to struct members
	rect := Rectangle{7, 4}

	// call calculate() with struct variable rect
	calculate(rect)

}
*/

/*
Implement Go Interface by Multiple Structs
In Go, more than 1 struct can also implement a single interface
*/

/*
package main
import "fmt"

// interface
type Shape interface {
  area() float32
}

 // Rectangle struct implements the interface
type Rectangle struct {
  length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
  return r.length * r.breadth
}

// Triangle struct implements the interface
type Triangle struct {
  base, height float32
}

// Triangle provides implementation for area()
func (t Triangle) area() float32 {
    return 0.5 * t.base * t.height
}

// access method of the interface
func calculate(s Shape) float32 {
  return s.area()
}

// main function
func main() {

  // assigns value to struct members
  r := Rectangle{7, 4}
  t := Triangle{8, 12}

  // call calculate() with struct variable rect
  rectangleArea := calculate(r)
  fmt.Println("Area of Rectangle:", rectangleArea)

  triangleArea := calculate(t)
  fmt.Println("Area of Triangle:", triangleArea)

}
*/

/*
What happens if the struct doesn't implement all methods of interface?
When a struct implements an interface, it should provide an implementation for all the
methods of the interface. If it fails to implement any method, we will get an error.
*/

/*
package main
import "fmt"

// interface
type Shape interface {
  area() float32
  perimeter() float32
}

 // Rectangle struct implements the interface
type Rectangle struct {
  length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
  return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) float32 {
  return s.area()
}

// main function
func main() {

  // assigns value to struct members
  r := Rectangle{7, 4}

  // call calculate() with struct variable rect
  rectangleArea := calculate(r)
  fmt.Println("Area of Rectangle:", rectangleArea)
}
*/

