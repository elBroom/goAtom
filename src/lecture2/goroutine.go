package main

import (
	"bufio"
	"os"
	"time"
	"log"
)

const numRoutins = 2
var ch = make(chan string, numRoutins)

func main() {
	//Создаем пулл горутин
	for i := 1; i <= numRoutins; i++ {
		go worker(i)
	}

	for {
		// Читаем из stdio
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		// Переедаем полученные данные в горутину
		ch <- input
	}
}

/* Основная логика */
func worker(i int)  {
	log.Printf("Start routine %d\n", i)
	for {
		// Принимаем данные
		input := <-ch

		// Обработка
		log.Printf("Work routine %d\n", i)
		log.Printf("input: %s", input)

		// Иммитация долгой работы - спим 5 сек
		time.Sleep(time.Duration(5000) * time.Millisecond)
		log.Printf("Finish routine %d with input: %s", i, input)
	}
}
