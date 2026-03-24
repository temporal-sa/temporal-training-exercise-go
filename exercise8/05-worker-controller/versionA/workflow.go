package versiona

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

var forceContinueAsNew bool

func TrackDriverLocation(ctx workflow.Context, driverID string, lastIter int) (string, error) {
	setupSignal(ctx)

	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var err error

	// Version A of this workflow pings every 60 seconds.
	// Version B of this workflow pings every 1 second.
	for i := lastIter; i < 30; i++ {
		if forceContinueAsNew {
			return "", workflow.NewContinueAsNewErrorWithOptions(ctx, workflow.ContinueAsNewErrorOptions{
				InitialVersioningBehavior: workflow.ContinueAsNewVersioningBehaviorAutoUpgrade,
			}, TrackDriverLocation, driverID, i)
		}

		var recordResult string
		err = workflow.ExecuteActivity(ctx, RecordDriverLocation, driverID).Get(ctx, &recordResult)

		workflow.Sleep(ctx, 60*time.Second)
	}

	return "Session complete", err
}

func setupSignal(ctx workflow.Context) {
	signalChan := workflow.GetSignalChannel(ctx, "force-continue-as-new")
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			selector := workflow.NewSelector(ctx)
			selector.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
				c.Receive(ctx, nil)
				forceContinueAsNew = true
			})
			selector.Select(ctx)
		}
	})
}
