package main

import (
    "log"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

    keys, ok := r.URL.Query()["key"]

    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items,
    // we only want the single item.
    key := keys[0]

    response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=" + string(key))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
        //fmt.Println("Url Param 'key' is: " + string(key) + " Resp: " + string(data))
        fmt.Fprint(w,"Url Param 'key' is: " + string(key) + " Resp: " + string(data))

    }

    //log.Println("Url Param 'key' is: " + string(key) + " Resp: " + string(data))
}
