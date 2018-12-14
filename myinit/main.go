
package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)

func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")   // (*p string-var, flag-name, default, usage )
}

func main () {
    flag.Parse()                      //  ./main -gopath /xxx 
    fmt.Println(user, home, gopath)
}
