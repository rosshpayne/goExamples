package main

import (
        "net/http"
        "fmt"
        "errors"
        "log"
    )

    
    func init() {
        //
        // func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
        //
        http.HandleFunc("/", errorHandler(betterHandler))      // errorHandler returns adapter 	http.HandlerFunc with method ServeHTTP
        http.HandleFunc("/top", errorHandler(betterHandler2))      // errorHandler returns adapter 	http.HandlerFunc with method ServeHTTP
    }
     
    func doThis() error {
        return errors.New("Do this error");
        } 
    
    func doThat() error {
        return errors.New("Do that error");
        } 
    //
    //   errorHandler performs function adapter role.  
    //   HandlerFunc.serveHTTP
    //
    func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {   // http.HandlerFunc supports Handler interface meaning serveHTTP method
        //
        return func(w http.ResponseWriter, r *http.Request) {      // anonymous function (lets call it XX) returned as http.HandlerFunc (which has method ServeHTTP)
                                                                   // HandleFunc will execute ServeHTTP via HandlerFunc
            err := f(w, r)                                         // f is betterHandle(..) 
                                                                   // HandleFunc executes ServerHTTP  which executes { XX(w,r) } which executes f(w,r) using lexical scoping
            if err != nil {                                        //
                http.Error(w, err.Error(), http.StatusInternalServerError)
                fmt.Fprintf(w,"%s","Sent to ResponseWriter.....")
                log.Printf("handling %q: %v", r.RequestURI, err)
            }
        }
    }
    
    func betterHandler(w http.ResponseWriter, r *http.Request) error {             // becomes f argument  in errorHandler
        if err := doThis(); err != nil {
            return fmt.Errorf("doing this: %v", err)
        }
    
        if err := doThat(); err != nil {
            return fmt.Errorf("doing that: %v", err)
        }
        return nil
    }

    func betterHandler2(w http.ResponseWriter, r *http.Request) error {             // becomes f argument  in errorHandler
        if err := doThis(); err != nil {
            return fmt.Errorf("doing this:n Handler2 %v", err)
        }
    
        if err := doThat(); err != nil {
            return fmt.Errorf("doing that: %v", err)
        }
        return nil
    }

    func main () {
    log.Fatal(http.ListenAndServe(":9030",nil) )
     }
