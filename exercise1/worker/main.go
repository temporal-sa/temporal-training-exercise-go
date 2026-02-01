package main

import (
	"log"

	"github.com/temporal/training/exercise1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const TaskQueue = "hello-task-queue"

// Worker is responsible for executing workflow and activity code.
// It polls task queues for work and executes the registered implementations.
func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, TaskQueue, worker.Options{})
	
	// TODO: Register the workflow implementation
	// Use: w.RegisterWorkflow((&exercise1.GreetingWorkflow{}).Greet)
	
	// TODO: Register the activity implementation
	// Use: w.RegisterActivity(exercise1.CreateGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
