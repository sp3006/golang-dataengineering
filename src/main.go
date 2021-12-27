package main

import ( "fmt"
        "net/http")


func main(){

  // The "HandleFunc" method accepts a path and a function as arguments
  // (Yes, we can pass functions as arguments, and even treat them like variables in Go)
 // However, the handler function has to have the appropriate signature (as described by the "handler" function below)
   http.HandleFunc("/", handler
}



