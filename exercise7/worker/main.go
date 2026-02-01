package main

import (
	"log"

	"github.com/temporal/training/exercise7"
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
	w.RegisterWorkflow((&exercise7.MoneyTransferWorkflow{}).Transfer)
	w.RegisterActivity(exercise7.Withdraw)
	w.RegisterActivity(exercise7.Deposit)
	w.RegisterActivity(exercise7.Refund)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
