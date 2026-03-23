package exercise8

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func Greet(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	var err error

	// v := workflow.GetVersion(ctx, "AddUppercase", workflow.DefaultVersion, 1)
	// if v == workflow.DefaultVersion {

	// } else {
	// 	err = workflow.ExecuteActivity(ctx, Uppercase, name).Get(ctx, &name)
	// }

	for range 30 {
		err = workflow.ExecuteActivity(ctx, CreateGreeting, name).Get(ctx, &result)
		workflow.Sleep(ctx, time.Second)
	}

	return result, err
}
