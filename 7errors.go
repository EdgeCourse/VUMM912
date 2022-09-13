package main

import (
	"errors"
	"fmt"
)

func somethingHappened() error {
	return errors.New("error with running function")
}

func main() {
	if err := somethingHappened(); err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Let's go!")
}

/*

package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}

*/

/*
package main

import (
	"errors"
	"fmt"
)

func somethingHappened() error {
	return errors.New("this is the error")
}

func main() {
	err := somethingHappened()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Let's Go!")
}

*/
/*
package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func main() {
	name, err := capitalize("bobby")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}

	fmt.Println("Capitalized name:", name)
}

*/

/*
//use anonymous function to reduce boilerplate code 
//The following program modifies the last example to include the length of the name that we’re capitalizing. It has three values to return

package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}

	if name == "" {
		return handle(errors.New("no name provided"))
	}

	return strings.Title(name), len(name), nil
}

func main() {
	name, size, err := capitalize("bobby")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Printf("Capitalized name: %s, length: %d", name, size)
}





*/

/*

//multiple return values

package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.Title(name), nil
}

func main() {
	_, err := capitalize("")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}
	fmt.Println("Success!")
}


*/
