package main

import (
	"log"

	"github.com/temporal/training/exercise2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const TaskQueue = "money-transfer-task-queue"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, TaskQueue, worker.Options{})
	
	// TODO: Register workflow and activities
	// w.RegisterWorkflow((&exercise2.MoneyTransferWorkflow{}).Transfer)
	// w.RegisterActivity(exercise2.Withdraw)
	// w.RegisterActivity(exercise2.Deposit)
	// w.RegisterActivity(exercise2.Refund)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
