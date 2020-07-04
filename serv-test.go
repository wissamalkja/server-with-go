package main

import (
    "fmt"
    "time"
    "log"
    "net/http"
        
)

func main() {


    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("#Active Server Online!  ", t)
            }
        }
    }()
   //just for testing time,,
    time.Sleep(3000000 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Serv stopped")
   
    ////server
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
        
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
        
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
    



}
