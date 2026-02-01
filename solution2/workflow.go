package solution2

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var approved, approvalReceived bool

	approvalChannel := workflow.GetSignalChannel(ctx, "approve")
	workflow.Go(ctx, func(ctx workflow.Context) {
		var value bool
		approvalChannel.Receive(ctx, &value)
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})

	err := workflow.ExecuteActivity(ctx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		err = workflow.ExecuteActivity(ctx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
		if err != nil {
			return "", err
		}
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	err = workflow.ExecuteActivity(ctx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		return "", err
	}
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
