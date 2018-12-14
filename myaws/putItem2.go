package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Create structs to hold info about new item
type ItemInfo struct {
    Plot string			`json:"plot"`
    Rating float64		`json:"rating"`
}

type Item struct {
    Year int			`json:"year"`
    Title string		`json:"title"`
    Info ItemInfo		`json:"info"`			// embedded struct - missing some atributes compared to data in first item
}

// Get table items from JSON file
func getItems() []Item {
    raw, err := ioutil.ReadFile("./movie_data_2.json")

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var items []Item						// decode destination. Slice of struct maps to an array of json objects.
    json.Unmarshal(raw, &items)
    return items
}


/* Conditional Insert example.  Do not overwrite (default behaviour) if key=2013 present

*/

func main() {

	sess, err := session.NewSession(&aws.Config{
	Region: aws.String("ap-southeast-2"),
	Credentials: credentials.NewSharedCredentials("", "bob-developer"),
	})
    	if err != nil {
        	fmt.Println("Error creating session:")
		log.Panic(err)
    	}
    	// Create DynamoDB client
    	svc := dynamodb.New(sess)
	_ = svc

    // Get table items from movie_data.json
    items := getItems()



    // Add each item to Movies table:
    for _, item := range items {
        av, err := dynamodbattribute.MarshalMap(item)			// encode go struct item Dynamodb attributevalue (AV) -> Map[string]*AV
        if err != nil {
            fmt.Println("Got error marshalling map:")
            fmt.Println(err.Error())
            os.Exit(1)
        }

	// Create a KeyBuilder representing the item key in this case
	//cond := expression.Equal(expression.Name("year"),expression.Value("2013")).Not()          //  this doesn't work as expected.  Items are still added.
	cond := expression.AttributeNotExists(expression.Name("year"))
	fmt.Printf("\n%#v",cond)
        expr,err:= expression.NewBuilder().WithCondition(cond).Build()

	fmt.Println("Condition :  ",*(expr.Condition()));
	for k,v := range expr.Names() {
		fmt.Printf("\n %s  %s",k,*v)
	}
	fmt.Printf("\n Condition : %s \n",expr.Values());

	input := &dynamodb.PutItemInput{
    		ConditionExpression:    expr.Condition(),		// PutItem does not have a keyCondition
  		ExpressionAttributeNames:  expr.Names(),		
  		ExpressionAttributeValues: expr.Values(),
		Item:                   av,
    		TableName:              aws.String("Movies"),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		}

        outp, err := svc.PutItem(input)					// insert item & create attributes as required.
        if err != nil {
            fmt.Println("*****  Got error calling PutItem:")
            fmt.Println(err.Error())
        }  else {
	    fmt.Println(outp)
            fmt.Println("\n\nSuccessfully added '",item.Title,"' (",item.Year,") to Movies table")
	}

    }
}
