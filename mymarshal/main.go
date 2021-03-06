package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//
	//   Single JSON object marshaled into map[string]string
	//  
	var jsonBlob = []byte(`
		{"Name": "Platypus", "Order": "Monotremata"}
	`)
        
	type output map[string]string       // **********************  Marshal Go type
	//
	// unmarshal will take a array of JSON objects and create a []Animal structs.  
	//  We supply the struct and unmarshal will do the rest, ie. create a slice create an element for each object in the JSON array.
	//
	var animals output
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	fmt.Printf("\n%#v %[1]T", animals)

	for r,i := range animals {
	   fmt.Printf("\n %s %s",r,i)
	}
	//
	//   Single JSON object marshaled into map[string]interface{}
	//  
	var jsonBlob2 = []byte(`
		{"Name": "Platypus", "Order": "Monotremata"}
	`)
        
	type output2 map[string]interface{}    // ******************** Marshal Go type
	//
	var animals2 output2
	err  = json.Unmarshal(jsonBlob2, &animals2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n\n%+v", animals2)
	fmt.Printf("\n%#v %[1]T", animals2)

	for r,i := range animals2 {
	   if s,ok:=i.(string) ; ok {
	   	fmt.Printf("\nLoop  %s %s",r,s)
	   } 
	}
	//
	// JSON  Array of objects
	//
	var jsonBlob3 = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	//
	// unmarshal will take a array of JSON objects and create a []Animal structs.  
	//  We supply the struct and unmarshal will do the rest, ie. create a slice create an element for each object in the JSON array.
	//
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	fmt.Printf("\n%#v %[1]T", animals)
}

	type mapT map[string]*string
	type mapiT map[string]interface{}
	
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	
	m:=make(mapT)
	name:="ross payne"
	m["name"]=&name
	age:="61"
	m["age"] =&age
	place:="canberra"
	m["place"]=&place
	
	fmt.Println()
	c,err:=json.Marshal(m)
		if err != nil {
		fmt.Println("error:", err)
	}
	
	os.Stdout.Write(c)
	
	i:=make(mapiT)
	i["name"]=&name
	i["name2"]="payl apyne"
	i["age"]=53
	i["Place"]=&place
	
	fmt.Println()
	d,err:=json.Marshal(i)
		if err != nil {
		fmt.Println("error:", err)
	}
	
	os.Stdout.Write(d)
	
}
}
