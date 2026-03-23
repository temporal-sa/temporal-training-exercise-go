package main

import (
	"log"
	patching "versioning/03-patching"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const TaskQueue = "hello-task-queue"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, TaskQueue, worker.Options{})
	w.RegisterWorkflow(patching.Greet)
	w.RegisterActivity(patching.CreateGreeting)
	w.RegisterActivity(patching.Uppercase)
	w.RegisterActivity(patching.Lowercase)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
