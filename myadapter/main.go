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
        http.HandleFunc("/", betterHandler)
        http.HandleFunc("/top", errorHandler(betterHandler2))      // errorHandler returns adapter 	http.HandlerFunc with method ServeHTTP
    }
     
    func doThis() error {
        return errors.New("Do this error");
        } 
    
    func doThat() error {
        return errors.New("Do that error");
        } 
    //
    //  betterHandler has no support for Handler interface. Must use adapter function (http.HandlerFunc conversion) to provide it.
    //  signature of betterHandler matches http.HandlerFunc so no need to provide additional wrapper/adapter function
    //  http.HandlerFunc is a function type with a method ServeHTTP so it satisfies interface http.Handler
    //  We use this fact to convert betterHandler to a http.HandlerFunc which we can do as their basic signature match.
    // 
    func betterHandler(w http.ResponseWriter, r *http.Request) {             // no error return value, hence signature exactly matches http.HandleFunc, hence can pass directly to
        if err := doThis(); err != nil {
            //return fmt.Errorf("doing this: %v", err)
            http.Error(w,err.Error(),22) 
            log.Printf("handling %q: %v", r.RequestURI, err)
    //        fmt.Fprintf(w,"doing this:  XXXXX - shouldn't write to w after http.error")
        }
    
        if err := doThat(); err != nil {
            http.Error(w,err.Error(),23) 
            log.Printf("handling %q: %v", r.RequestURI, err)
            fmt.Fprintf(w,"doing this: %v", err)
        }
    }
    //
    //  adapter function -  arg1: function-to-be-adapted  , returns http.HandlerFunc 
    //
    func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {   // http.HandlerFunc matches type of second arg in http.HandleFunc
        //
        // return anonymous function matching signature of http.HandlerFunc
        //
        return func(w http.ResponseWriter, r *http.Request) {      // pass anonymous function with no error return value (lets call it XX) returned as http.HandlerFunc 
                                                                   // HandleFunc will XX as is just like it was errorHandler2(..)
            err := f(w, r)                                         // f is betterHandle(..) 
                                                                   // 
            if err != nil {                                        //
                http.Error(w, err.Error(), http.StatusInternalServerError)
                fmt.Fprintf(w,"%s","Sent to ResponseWriter.....")
                log.Printf("handling %q: %v", r.RequestURI, err)
            }
        }
    }
    //
    // Following example shows the use of an adapter function that adapts an existing function betterHanler2 to an interface
    // betterHandler2 does not not have a Handler interface ie. it does not have an associated ServeHTTP method
    // It also doesn't match the ServeHTTP signature as it has a error return type
    // The adapter function must perform two tasks, provide a ServeHTTP method and manipulate the function to match the signature of ServerHTTP
    // In this case we have two adapter functions:  
    //       http.HandlerFunc   - which provides interface support ie. method ServerHTTP
    //       errorHandler       - which provides signature support ie. handles error return value and passes basic signature to http.HandlerFunc
    //                            accepts original function as argument and wraps it in an anonymous function that handles signature irregularity 
    //                            anonymous function is a closure - meaning it references objects outside its scope i.e. f in this case.
    //                            It creates an anonymous function and passes it to http.HandlerFunc for conversion. 
    //                            The anonymous functions provides a wrapper that handles the error return value.
    //
    // The executon flow:  -> represents CALLS
    //
    //  http.HandlerThingy -> a.http.ServeHTTP(w,r) -> (f(w,r)) i.e anonymousFunction.ServeHTTP(w,r) -> anonymousFunction(w,r) -> err:=f(w,r)
    //
    func betterHandler2(w http.ResponseWriter, r *http.Request) error {      // becomes f argument  in errorHandler.  Cannot pass directly to http.HandleFunc
                                                                             // as return value means it doesn't match http.HandleFunc signature.
                                                                             // require to be passed thru adapter function that handles betterHandler2 return value
                                                                             // and returns exact signature match to http.HandleFunc

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
