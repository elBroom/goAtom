package main

import (
	"fmt"
	"bufio"
	"os"
)

func worker(id int, jobs <-chan string, results chan <- tmp) {
	for j := range jobs {
		results <- tmp{id, j}
	}
}

type tmp struct {
	id int
	name string
}

func main() {
	jobs := make(chan string, 10)
	results := make(chan tmp, 10)
	nWorkers := 3

	for w := 1; w <= nWorkers; w++ {
		go worker(w, jobs, results)
	}

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= 3; i++ {
		input, error := reader.ReadString('\n')
		if error != nil {
			panic("AAAAAAAAAAA")
		}
		jobs <- input
	}

	for t := 0; t<= 3; t++ {
		variable := <- results
		fmt.Print(variable)
	}


}
