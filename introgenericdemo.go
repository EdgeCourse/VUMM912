package main

import (
	"fmt"
)

func PrintSlice(s []int) {
	fmt.Println("Generics")
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func main() {
	PrintSlice([]int{1, 2, 3})
	//PrintSlice([]string{"a", "b", "c"})
	//PrintSlice([]float64{1.2, -2.33, 4.55})

}
