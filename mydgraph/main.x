package main

import (
  "context"
  "flag"
  "fmt"
  "log"
  "encoding/json"

  "github.com/dgraph-io/dgraph/client"
  "github.com/dgraph-io/dgraph/protos/api"
  "google.golang.org/grpc"
)

var (
  dgraph = flag.String("d", "ec2-54-252-186-201.ap-southeast-2.compute.amazonaws.com:9080", "Dgraph server address")
)

func main() {
  flag.Parse()
  conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  dg := client.NewDgraphClient(api.NewDgraphClient(conn))
  
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
  fmt.Printf("Response: %s\n", resp.Json)

// Response: {"bladerunner":[{"uid":"0xedde","name@en":"Blade Runner","initial_release_date":"1982-06-25T00:00:00Z"}]}

  var lines struct {
		  Uid   string   `json:"uid"`
		  //Name  string     `json:"name@en"`
                  //Rdate  string    `json:"initial_release_date"`
 	         }

  if  err:=json.Unmarshal(resp.Json,&lines); err != nil {
	panic(err)
  }

  fmt.Printf("%v",lines)
  /*for _,l := range lines {
	fmt.Printf("%v",l)
  } */
  
}
