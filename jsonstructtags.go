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

/*
JSON struct tags:
The JSON package can automatically encode your custom data types.
It will only include exported fields in the encoded output and will by default use those
names as the JSON keys.

You can use tags on struct field declarations to customize the encoded JSON key names.
*/

/*
Using Struct Tags to Control Encoding
You can modify the previous example to have exported fields that are properly
encoded with camel-cased field names by annotating each field with a struct tag.
The struct tag that encoding/json recognizes has a key of json and a value that controls
the output. By placing the camel-cased version of the field names as the value to the json
key, the encoder will use that name instead.
*/
/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	//What if I don't use a struct tag for something?
	Name string
	//what if I use a struct tag on an unexported value?
	utype     string    `json:"Type"`
	Password  string    `json:"password"`
	Subjects  []string  `json:"My Subjects"`
	CreatedAt time.Time `json:"time created"`
}

func main() {
	u := &User{
		Name:      "Bob Loblaw",
		Password:  "abc123",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
*/
/*
But now, the fields for some values were printed even though we did not set those values.
The JSON encoder can eliminate these fields as well, if necessary.

Removing Empty JSON Fields:
Suppress the outputting of fields that are unset in JSON. Since all types in Go have a
“zero value,” some default value that they are set to, the encoding/json package needs
additional information to be able to tell that some field should be considered unset when
it assumes this zero value. Within the value part of any json struct tag, you can suffix the
desired name of your field with ,omitempty to tell the JSON encoder to suppress the output of
this field when the field is set to the zero value. The following example fixes the previous
examples to no longer output empty fields:
*/

/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	//What if I don't use a struct tag for something?
	Name string
	//what if I use a struct tag on an unexported value?
	utype    string `json:"Type"`
	Password string `json:"password"`
	//add omitempty, no space
	Subjects  []string  `json:"My Subjects,omitempty"`
	CreatedAt time.Time `json:"time created"`
}

func main() {
	u := &User{
		Name:      "Bob Loblaw",
		Password:  "abc123",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
*/

/*
Ignoring Sensitive but Exported Fields When Outputting

Some fields must be exported from structs so that other packages can correctly interact
with the type. However, the nature of these fields may be sensitive, so in these
circumstances, we would like the JSON encoder to ignore the field entirely—even when it is
set. This is done using the special value - as the value argument to a json: struct tag.

This example fixes the issue of exposing the user’s password.
*/
/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	//What if I don't use a struct tag for something?
	Name string
	//what if I use a struct tag on an unexported value?
	utype string `json:"Type"`
	//use hyphen (not underscore) to suppress output of this field but still have it be available
	Password string `json:"-"`
	//add omitempty, no space
	Subjects  []string  `json:"My Subjects,omitempty"`
	CreatedAt time.Time `json:"time created"`
}

func main() {
	u := &User{
		Name:      "Bob Loblaw",
		Password:  "abc123",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
*/
