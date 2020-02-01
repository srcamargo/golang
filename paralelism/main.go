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
    fmt.Println("Channel: ",d)
    go threadcreator("one", d, 1)
    go threadcreator("two", d, 2)
    go threadcreator("three", d, 3)
    go threadcreator("four", d, 4)
    go threadcreator("five", d, 5)
    go threadcreator("six", d, 6)
    go threadcreator("seven", d, 7)
    go threadcreator("eight", d, 8)

    for i := 0; i < 100; i++ {
        time.Sleep(time.Second);
    }

    fmt.Println("Finishing.")
}

func threadcreator(msg string, c chan string, id int) {
    for i := 0; ; i++ {

    }
}
