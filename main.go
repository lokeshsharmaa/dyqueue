package main

import (
	"log"
	"sync"
	"time"
)

type Data struct {
	first  int
	second int
}

func main() {
	var wg sync.WaitGroup

	messages := make(chan Data)
	result := make(chan int)
	exitChannel := make(chan int)
	hour := time.Second * 5
	log.Println("Starting")

	wg.Add(1)
	go printMessages(result)
	go printChannelLength(messages, result)
	go exitAfter(hour, exitChannel)

	generateMessages(messages)
	go handleMessages(messages, result, exitChannel, &wg)

	wg.Wait()
	log.Println("Program exited")
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

func handleMessages(messages chan Data, result chan int, exitChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg := <-messages:
			// Handle the message
			result <- msg.first + msg.second
		case <-exitChannel:
			log.Print("Exiting")
			close(messages)
			close(result)
			return
		}
	}
}

func generateMessages(messages chan Data) {
	go func() {
		for i := 0; i < 100; i++ {
			messages <- Data{first: i, second: i}
			time.Sleep(1 * time.Second)
		}
	}()
}
