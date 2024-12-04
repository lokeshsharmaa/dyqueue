package dyqueue

import (
	"log"
	"sync"
	"time"
)

type Dyqueue[T any] interface {
	Produce()
	Consume(T)
	Start()
	Stop()
	SetConcrete(Dyqueue[T])
}

type AbstractDyqueue[T any] struct {
	Dyqueue[T]
	wg             sync.WaitGroup
	NumOfRoutine   int
	MessageChannel chan T
	exitChannel    chan int
}

func (d *AbstractDyqueue[T]) SetConcrete(concrete Dyqueue[T]) {
	d.Dyqueue = concrete
}

func NewAbstractDyqueue[T any](numOfRoutine int, messageChannel chan T) *AbstractDyqueue[T] {
	return &AbstractDyqueue[T]{NumOfRoutine: numOfRoutine, MessageChannel: messageChannel, exitChannel: make(chan int)}
}

func (d *AbstractDyqueue[T]) Done() {
	d.wg.Done()
}

func (d *AbstractDyqueue[T]) Stop() {
	log.Println("Stopping Dyqueue")
	d.stopAllConsumers()
}

func (d *AbstractDyqueue[T]) Start() {
	log.Println("Starting Dyqueue")

	go d.startProducer()

	for i := 0; i < d.NumOfRoutine; i++ {
		d.startConsumer()
	}
	d.wg.Wait()
	log.Println("Exiting Dyqueue")
}

func (d *AbstractDyqueue[T]) startProducer() {
	defer d.stopAllConsumers()
	log.Println("Starting Producer")
	d.startMonitoring()
	d.Produce()
}

func (d *AbstractDyqueue[T]) startMonitoring() {
	go func() {
		maxConsumers := 10 // Set a maximum limit
		currentConsumers := 0
		for {
			if len(d.MessageChannel) > 3 && currentConsumers < maxConsumers {
				currentConsumers++
				d.startConsumer()
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

var once sync.Once

func (d *AbstractDyqueue[T]) stopAllConsumers() {
	once.Do(func() {
		log.Println("Stopping all consumers")
		close(d.exitChannel)
	})
}

func (d *AbstractDyqueue[T]) startConsumer() {
	d.wg.Add(1)

	// Implement the consumer logic here
	log.Println("Starting Consumer")
	go func() {
		defer d.Done()
		for {
			select {
			case msg := <-d.MessageChannel:
				log.Println("Consuming Message: ", msg)
				d.Consume(msg)
			case <-d.exitChannel:
				log.Print("Exiting Consumer")
				return
			}
		}
	}()

}
