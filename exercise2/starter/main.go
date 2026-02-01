package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/temporal/training/exercise2"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	request := exercise2.TransferRequest{
		FromAccount: "account-123",
		ToAccount:   "account-456",
		Amount:      100.0,
		TransferID:  "transfer-1",
	}

	options := client.StartWorkflowOptions{
		ID:        "money-transfer-workflow",
		TaskQueue: "money-transfer-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, (&exercise2.MoneyTransferWorkflow{}).Transfer, request)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	// TODO: Send approval signal after delay
	// time.Sleep(2 * time.Second)
	// err = c.SignalWorkflow(context.Background(), we.GetID(), we.GetRunID(), "approve", true)

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	fmt.Println(result)
}
