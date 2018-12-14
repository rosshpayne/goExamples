package main

import ( 
	"fmt"
        "net/http"

	"github.com/fatih/structs"
       )



func main () {

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

server := &Server{
	Name:    "gopher",
	ID:      123456,
	Enabled: true,
}

// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := structs.Map(server)
fmt.Println("m is %v", m)

// Convert the values of a struct to a []interface{}
// => ["gopher", 123456, true]
v := structs.Values(server)
fmt.Println("v is %v", v)

// Convert the names of a struct to a []string
// (see "Names methods" for more info about fields)
n := structs.Names(server)
fmt.Println("n is %v", n)

// Convert the values of a struct to a []*Field
// (see "Field methods" for more info about fields)
f := structs.Fields(server)
fmt.Println("f is %v", f)

// Return the struct name => "Server"
n2 := structs.Name(server)
fmt.Println("n2 is %v", n2)

// Check if any field of a struct is initialized or not.
h := structs.HasZero(server)
fmt.Println("h is %v", h)

// Check if all fields of a struct is initialized or not.
z := structs.IsZero(server)
fmt.Println("z is %v", z)

// Check if server is a struct or a pointer to struct
i := structs.IsStruct(server)
fmt.Println("i is %v", i)

}
