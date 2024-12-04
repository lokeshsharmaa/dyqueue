# dyqueue

`dyqueue` is a Go library designed to handle dynamic load by scaling the number of goroutines based on the size of a queue. It provides a flexible and efficient mechanism to adjust the number of concurrent consumers and producers, ensuring optimal performance based on workload.

## Features

- **Dynamic Scaling**: Automatically adjusts the number of consumers based on the size of the queue.
- **Concurrency Management**: Spawns and terminates goroutines as needed to handle the load.
- **Flexible Interface**: Allows users to define custom `Produce` and `Consume` logic.

## How It Works

The `dyqueue` library defines an interface `Dyqueue`, which has the following methods:

- `Produce()`: Method to produce data and push it into the queue.
- `Consume(T)`: Method to consume data from the queue.
- `Start()`: Starts the queue processing and begins handling data.
- `Stop()`: Stops the queue and terminates all goroutines.

### AbstractDyqueue

The `AbstractDyqueue` struct implements the core queue logic, handling dynamic scaling of goroutines. It manages the message channel, exit channel, and provides a framework to start and stop consumers and producers.

### Concrete Implementation

To use `dyqueue`, you need to create a concrete struct that implements the `Produce` and `Consume` methods. These methods define the logic for producing and consuming data in your application.

## Usage

### 1. Install the Library

First, ensure that your Go workspace is set up and then run:

```bash
go get github.com/lokeshsharmaa/dyqueue
```

### 2. Implement Your Concrete Dyqueue (Refer main.go)


### 3. Run the Program

Once your program is set up, run the application with:

```bash
go run main.go
```

### 4. Customize for Your Needs

You can customize the `Produce` and `Consume` methods to handle any data or processing logic as required for your use case.

## License

This library is licensed under the MIT License. See the LICENSE file for more details.