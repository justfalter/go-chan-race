package main

import (
        "fmt"
        "time"
        "os"
)

func main() {
        for i := 1; i < 10000; i++ {
                chUnbuffered := make(chan bool, 0)
                chBuffered := make(chan string, 1)
                chDone := make(chan bool, 0)

                go func() {
                        for true {
                                select {
                                case _ = <-chUnbuffered:
                                        // nop
                                case <- time.After(1 * time.Second):
                                        continue
                                }

                                select {
                                case msg := <-chBuffered:
                                        _ = msg
                                        //fmt.Printf("Got the message: %s\n", msg)
                                        chDone <- true

                                case <- time.After(0):
                                        fmt.Printf("No data in chBuffered in iter %d\n", i)
                                        os.Exit(1)

                                }
                        }
                }()

                chBuffered <- fmt.Sprintf("This is my message: %d", i)
                chUnbuffered <- true
                <-chDone
        }
        fmt.Printf("All done\n")

}
