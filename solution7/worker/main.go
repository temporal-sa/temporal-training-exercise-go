package main

import (
	"log"

	"github.com/temporal/training/solution7"
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
	w.RegisterWorkflow((&solution7.MoneyTransferWorkflow{}).Transfer)
	w.RegisterActivity(solution7.Withdraw)
	w.RegisterActivity(solution7.Deposit)
	w.RegisterActivity(solution7.Refund)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
