package main

import "fmt"

func chanDome( r chan int,n int ){
    c := make(chan int)
    go func(){
        for{
            n := <-c
            fmt.Println(n)
        }
    }
    c <- 1
    c <- 2
}

func main(){
    chanDome()
}
