package workflowtypeversioning

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// TODO: create a copy of this function and name it `GreetV2`
func Greet(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	var err error

	// TODO: After you've started a `Greet` workflow, uncomment the following
	//		 line to introduce a non-deterministic error.
	// TODO: After you've copied this workflow to `GreetV2`, uncomment this line
	// err = workflow.ExecuteActivity(ctx, Uppercase, name).Get(ctx, &name)

	// run for 5 minutes
	for range 20 {
		err = workflow.ExecuteActivity(ctx, CreateGreeting, name).Get(ctx, &result)
		workflow.Sleep(ctx, 15*time.Second)
	}

	return result, err
}
