package main

import (
    "fmt"
)

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s { // _ = key, v = values
		sum += v
	}
	c <- sum // send sum to c
}

func main(){
    s := []int{327, 112, 82, 239, 214, 100}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //sum 1st half of array
	go sum(s[len(s)/2:], c) //sum 2nd half of array
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
