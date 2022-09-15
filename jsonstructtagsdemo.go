//JSON:
/*
Only exported fields will be encoded/decoded in JSON. Fields must start with capital
letters to be exported.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name string
	//don't use keyword type
	//doesn't export because lowercase
	utype     string
	Password  string
	Subjects  []string
	CreatedAt time.Time
}

func main() {
	u := &User{
		Name:      "Bob Loblaw",
		utype:     "student",
		Password:  "abc123",
		CreatedAt: time.Now(),
	}

	//https://gobyexample.com/json
	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
