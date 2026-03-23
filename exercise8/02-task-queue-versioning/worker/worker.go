package main

import (
	"log"
	taskqueueversioning "versioning/02-task-queue-versioning"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// TODO: Once you've copied this worker, change the task queue name
const TaskQueue = "hello-task-queue"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, TaskQueue, worker.Options{})
	// TODO: Once you've copied this worker, delete this line, and uncomment
	//		 the three lines below it (ie we want to use GreetV2)
	w.RegisterWorkflowWithOptions(taskqueueversioning.Greet, workflow.RegisterOptions{})
	// w.RegisterWorkflowWithOptions(taskqueueversioning.GreetV2, workflow.RegisterOptions{
	// 	// We want to invoke our workflow using the same name (`Greet`), because both
	// 	// Workflow definitions are in the same package we have to mannually set the
	// 	// Workflow name to `Greet`.
	// 	Name: "Greet",
	// })
	w.RegisterActivity(taskqueueversioning.CreateGreeting)
	w.RegisterActivity(taskqueueversioning.Uppercase)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
