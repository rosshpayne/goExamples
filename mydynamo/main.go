package main

import (
	"fmt"
	_ "log"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	)

    //  Marshaling is all about taking a Go datatype and building a dynamodb  AttributeValue -  which is  a struct containing all the dynamodb datatype plus values
    //   when populated. An AttributeValue represents the data for a dynamodb attribute, which inturn makes up a dynamodb item.
    //   
    //  type AttributeValue struct has these fields....
    //
    // Dynamo DataType        AttributeValue representation                 Struct field name and type 
    // -------------------   --------------------------------------------   ------------------------------------
    // Binary.               "B": "dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk"    B []byte `type:"blob"`
    //
    // Boolean.              "BOOL": true                                 BOOL *bool `type:"boolean"`
    //
    // Binary Set.           "BS": ["U3Vubnk=", "UmFpbnk=", "U25vd3k="]     BS [][]byte `type:"list"`
    //
    // List.                 "L": ["Cookies", "Coffee", 3.14159]             L []*AttributeValue `type:"list"`               All slices map to this
    //
    // Map.                  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}} where Name & Age are struct fields
    //                                                                       M map[string]*AttributeValue `type:"map"`       All structs map to this
    //
    // Number.               "N": "123.45"                                   N *string `type:"string"`
    // 
    // Number Set.           "NS": ["42.2", "-19", "7.5", "3.14"]           NS []*string `type:"list"`
    //
    // Null.                 "NULL": true                                  NULL *bool `type:"boolean"`
    //
    // String.               "S": "Hello"                                    S *string `type:"string"`
    //
    // String Set.           "SS": ["Giraffe", "Hippo" ,"Zebra"]
    //
    // ********  Struct Tags to modify Go datatype -> AtributeValue mapping  *******
    //
    // Field is ignored                          Field int      `dynamodbav:"-"`              -  ignore this field 
    // Field AttributeValue map key "myName"     Field int      `dynamodbav:"myName"`         -  rename field in attributeValue of map[<tagName>]
    // Field AttributeValue map key "myName",    Field int      `dynamodbav:"myName,omitempty"`
    //                                           Field int      `dynamodbav:",omitempty"`     -  omit field if empty
    // Field []string and map                                   `dynamodbav:",omitemptyelem"`
    // Field will be marshaled as a string       Field int      `dynamodbav:",string"`        -  only value for number types, (int,uint,float) 
    // Field will be marshaled as a binary set   Field [][]byte `dynamodbav:",binaryset"`
    // Field will be marshaled as a number set   Field []int    `dynamodbav:",numberset"`
    // Field will be marshaled as a string set   Field []string `dynamodbav:",stringset"`
    //                                           Field time.Time `dynamodbav:",unixtime"`

func main () {
	type Record struct {
	    Bytes   []byte
	    MyField string
	    Letters []string
	    A2Num   map[string]int
	}
	
	expect := Record{
	    Bytes:   []byte{48, 49},
	    MyField: "MyFieldValue",
	    Letters: []string{"a", "b", "c", "d"},
	    A2Num:   map[string]int{"a": 1, "b": 2, "c": 3},
	}
	//
	//  define AttributeValue - normally tis would be sourced from dynamodb	
	av := &dynamodb.AttributeValue{
	    M: map[string]*dynamodb.AttributeValue{
	        "Bytes":   {B: []byte{48, 49}},
	        "MyField": {S: aws.String("MyFieldValue")},
	        "Letters": {L: []*dynamodb.AttributeValue{
	            {S: aws.String("a")}, {S: aws.String("b")}, {S: aws.String("c")}, {S: aws.String("d")},
	        }},
	        "A2Num": {M: map[string]*dynamodb.AttributeValue{
	            "a": {N: aws.String("1")},
	            "b": {N: aws.String("2")},
	            "c": {N: aws.String("3")},
	        }},
	    },
	}
	
	actual := Record{}
	err := dynamodbattribute.Unmarshal(av, &actual)
	fmt.Println(err, reflect.DeepEqual(expect, actual))
	
}
