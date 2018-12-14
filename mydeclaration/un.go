package main

import (
	"encoding/json"
	"fmt"
)
type mystruct struct { fred string
	   age int
         }

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal		// animals is allocated metastruct of slice - but no underlying array so animals is nil, but &animals exists.
	var animalsp *[]Animal		// pointer to animals.  No allocation exists for pointer things (must use make/new) so no [] metastructure is created
	var animals2p *[]Animal		// pointer to animals.  No allocation exists for pointer things (must use make/new) so no [] metastructure is created

        if animals == nil {
		fmt.Printf("animals is Nil ie. no underlying array - but metadata struct exists as it has a pointer value to it %p\n",&animals)
	}
	fmt.Printf("but has a pointer values so something is allocated namely metadata :  %p\n",&animals)
	fmt.Printf("Type animals: %T\n\n",animals)
	//
	// unmarshal jsonBlob into animals
	//
	fmt.Printf("animals before unmarshal: %#v  %p\n",animals,&animals)
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("animals after unmarshal:  %#v %p\n",animals,&animals)

        if animalsp == nil {
		fmt.Println("\n\nanimalsp is nil")
 	}
	animalsp=new([]Animal)					// allocate pointer 
	//  pass in nil pointer and Go will allocate it pointer to []Animal and populate with elements of the decoded JSON.
	err = json.Unmarshal(jsonBlob, animalsp)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", *animalsp)

	err = json.Unmarshal(jsonBlob, &animals2p)		// note: pointer to pointer to []Animal, passed in as arg. Unmarshal handles the double ref
        if err != nil {
                fmt.Println("error:", err)
        }
        fmt.Printf("%+v\n", *animals2p)				// deref animals2p 

        var st mystruct			// allocates a st struct and populates with zero values.
	sta:=mystruct{}                 // T{}  - type literal - populates a type instance.  Same as declaration:  var sta mystruct

	fmt.Printf("st pointer = %p\n",&st)
	if st.age == 0 {
		fmt.Println("st is allocated")
   	}
	st.age = 44
	if sta.age == 0 {
		fmt.Println("sta is allocated")
   	}

}
