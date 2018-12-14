package main

import (
  "context"
  "flag"
  "fmt"
  "log"
  "encoding/json"

  "github.com/dgraph-io/dgo"
  "github.com/dgraph-io/dgo/protos/api"
  "google.golang.org/grpc"
)

var (
  dgraph = flag.String("d", "ec2-54-206-68-245.ap-southeast-2.compute.amazonaws.com:9080", "Dgraph server address")
)

func main() {
  flag.Parse()
  conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
  
  resp, err := dg.NewTxn().Query(context.Background(), `{
  bladerunner(func: eq(name@en, "Blade Runner")) {
    uid
    name@en
    initial_release_date
    netflix_id
  }
}`)
  
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s",string(resp.Json))
  //fmt.Printf("%c",resp.Json)

  //  note JSON output  {  name :  value } where value is an array of objects ie. [ object ]/

  // Response: {"bladerunner":[{"uid":"0xedde","name@en":"Blade Runner","initial_release_date":"1982-06-25T00:00:00Z"}]}

  //
  // Now decode into Go struct
  //
  type mapt map[string]string

//  type outt struct {  Name []mapt  `json:"bladerunner"` }          // go's json decode will return a pointer to array so we must define a slice of pointers.
//  type outt struct {  Bladerunner []mapt `json:"bladerunner"` }    //  this works but tag is uncessary (see next example). Note field names must be upper case to make visible.
    type outt struct {  Bladerunner []mapt }                         //  this works as Go will check for field name in a case insensitive manner. 
	
  var lines2 outt   

  if  err:=json.Unmarshal([]byte(resp.Json),&lines2); err != nil {      // pass in pointer so struct can be populated inplace.
	panic(err)
  }

  /*for i,v := range lines2.Bladerunner {                             // slice of maps
    for k2,v2 := range v {
	fmt.Printf("\nKey, value   %d  %s  %s",i,k2,v2)
    }
  } 
*/
}