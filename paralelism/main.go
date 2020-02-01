package main


import (
    "fmt"
    "time"
    "runtime"
)


func main() {
    fmt.Println("Number of CPUs found is ", runtime.NumCPU())
    fmt.Println("Setting the maximum number of threads to ", runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU()) // Set the maximum number of threads/processes


    d := make(chan string)
    fmt.Println(d)
    go boring("boring!", d, 1)
    go boring("boring!", d, 2)
    go boring("boring!", d, 3)
    go boring("boring!", d, 4)
    go boring("boring!", d, 5)
    go boring("boring!", d, 6)
    go boring("boring!", d, 7)
    go boring("boring!", d, 8)

    for i := 0; i < 100; i++ {
        time.Sleep(time.Second);
    }

    fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan string, id int) {
    for i := 0; ; i++ {

    }
}
