package main

import (
    "net/http"
    "os"
    "log"
    "fmt"
)

func main() {
	// put your go webserver code here!
    addr := os.Getenv("PORT")
    fmt.Println(addr)
    if len(addr) == 0 {
        addr = ":80"
    }
    http.Handle("/", http.FileServer(http.Dir("./static")))
    log.Printf("listening on %s...", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
