package main

import (
	"log"
	workflowtypeversioning "versioning/01-workflow-type-versioning"

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
	w.RegisterWorkflow(workflowtypeversioning.Greet)
	// TODO: Once you've made non-deterministic changes to ../workflow.go
	//		 and seen them show up in the Temporal UI, copy the Workflow
	//		 and register it here.
	// w.RegisterWorkflow(workflowtypeversioning.GreetV2)
	w.RegisterActivity(workflowtypeversioning.CreateGreeting)
	w.RegisterActivity(workflowtypeversioning.Uppercase)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
