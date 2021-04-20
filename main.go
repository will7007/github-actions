// Launch microservice server- main.go
package main

import (
    "lab12/microservice"
    "log"
)

func main() {
    s := microservice.NewServer("", "8000")
    log.Fatal(s.ListenAndServe())
}

