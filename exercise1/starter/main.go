package main

import (
	"context"
	"fmt"
	"log"

	"github.com/temporal/training/exercise1"
	"go.temporal.io/sdk/client"
)

// Starter creates and executes a workflow.
// This is the client that starts workflow executions.
func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "hello-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, (&exercise1.GreetingWorkflow{}).Greet, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	fmt.Println(result)
}
