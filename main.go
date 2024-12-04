package main

import (
	"dyqueue/client/dyqueue"
	"log"
	"time"
)

type Data struct {
	first  int
	second int
}

type ConcreteDyqueue struct {
	*dyqueue.AbstractDyqueue[Data]
	ResultChannel chan int
}

func (c *ConcreteDyqueue) Consume(message Data) {
	c.ResultChannel <- message.first + message.second
}

func (c *ConcreteDyqueue) Produce() {
	for i := 0; i < 1000; i++ {
		for i := 0; i < 10; i++ {
			log.Println("Producing Message: ", i)
			c.AbstractDyqueue.MessageChannel <- Data{first: i, second: i}
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	messages := make(chan Data, 100)
	result := make(chan int, 100)
	dyqueue := ConcreteDyqueue{
		AbstractDyqueue: dyqueue.NewAbstractDyqueue[Data](1, messages),
		ResultChannel:   result,
	}
	dyqueue.SetConcrete(&dyqueue) // Need to fix this line

	hour := time.Second * 10

	go printMessages(result)
	go printChannelLength(messages, result)
	go dyqueue.exitAfter(hour)

	dyqueue.Start()

	log.Println("Program Exiting")
}

func (c *ConcreteDyqueue) exitAfter(duration time.Duration) {
	time.Sleep(duration)
	c.Stop()
	close(c.ResultChannel)
	close(c.MessageChannel)
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
		log.Println("Result: ", value)
	}
}
