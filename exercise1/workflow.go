package exercise1

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// GreetingWorkflow is the workflow implementation.
type GreetingWorkflow struct{}

// Greet is the workflow method that orchestrates the greeting activity.
//
// TODO: Configure ActivityOptions with StartToCloseTimeout of 10 seconds
// TODO: Call the CreateGreeting activity and return the result
func (w *GreetingWorkflow) Greet(ctx workflow.Context, name string) (string, error) {
	return "", nil // Replace with activity call
}
