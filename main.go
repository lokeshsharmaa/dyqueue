package main

import (
	"log"
	"time"
)

type Data struct {
	first  int
	second int
}

func main() {

	messages := make(chan Data)
	result := make(chan int)
	exitChannel := make(chan int)
	hour := time.Minute * 60

	go printMessages(result)
	go printChannelLength(messages, result)

	go exitAfter(hour, exitChannel)
	generateMessages(messages)

	handleMessages(messages, result, exitChannel)

}

func exitAfter(duration time.Duration, exitChannel chan int) {
	time.Sleep(duration)
	exitChannel <- 1
}

func printChannelLength(messages chan Data, result chan int) {
	for {
		log.Println("Messages: ", len(messages))
		log.Println("Result: ", len(result))
		time.Sleep(1 * time.Second)
	}
}

func printMessages(result chan int) {
	for value := range result {
		log.Println(value)
	}
}

func handleMessages(messages chan Data, result chan int, exitChannel chan int) {
	go func() {
		for {
			select {
			case data := <-messages:
				result <- data.first + data.second
			case <-exitChannel:
				log.Println("Exiting")
				return
			}
		}
	}()
}

func generateMessages(messages chan Data) {
	go func() {
		for i := 0; i < 100; i++ {
			messages <- Data{first: i, second: i}
			time.Sleep(1 * time.Second)
		}
	}()
}
