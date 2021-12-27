package main

import ( "fmt"
        "net/http")


func main(){

  // The "HandleFunc" method accepts a path and a function as arguments
  // (Yes, we can pass functions as arguments, and even treat them like variables in Go)
 // However, the handler function has to have the appropriate signature (as described by the "handler" function below)
   http.HandleFunc("/", handler)

   // After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}



