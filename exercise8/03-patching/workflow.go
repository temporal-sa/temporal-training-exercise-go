package patching

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func Greet(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	var err error

	// TODO: Use the patching API to introduce the Uppercase activity
	// v := workflow.GetVersion(ctx, "AddUppercase", workflow.DefaultVersion, 1)
	// if v == workflow.DefaultVersion {
	err = workflow.ExecuteActivity(ctx, Lowercase, name).Get(ctx, &name)
	// } else {
	// 	err = workflow.ExecuteActivity(ctx, Uppercase, name).Get(ctx, &name)
	// }

	for range 30 {
		err = workflow.ExecuteActivity(ctx, CreateGreeting, name).Get(ctx, &result)
		workflow.Sleep(ctx, 60*time.Second)
	}

	return result, err
}
