package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

const numRoutins = 2
var ch = make(chan string, numRoutins)

func main() {
	for i := 1; i <= numRoutins; i++ {
		go worker(i)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		ch <- input
	}
}

func worker(i int)  {
	fmt.Printf("Start routine %d\n", i)
	for {
		input := <-ch
		fmt.Printf("Work routine %d\n", i)
		fmt.Printf("input: %s", input)
		time.Sleep(5000 * time.Millisecond)
		fmt.Printf("Finish routine with input: %s", input)
	}
}
