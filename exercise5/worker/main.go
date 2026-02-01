package main

import (
	"log"

	"github.com/temporal/training/exercise5"
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
	w.RegisterWorkflow((&exercise5.MoneyTransferWorkflow{}).Transfer)
	w.RegisterActivity(exercise5.Withdraw)
	w.RegisterActivity(exercise5.Deposit)
	w.RegisterActivity(exercise5.Refund)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
