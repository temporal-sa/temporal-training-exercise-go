package exercise5

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

// Transfer processes a money transfer with descriptive activity summaries.
// TODO: Add Summary field to ActivityOptions for each activity execution:
//   - Withdraw: "Withdrawing funds from account {fromAccount}"
//   - Deposit: "Depositing funds to account {toAccount}"
//   - Refund: "Refunding funds to account {fromAccount}"
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	err := workflow.UpsertSearchAttributes(ctx, map[string]interface{}{
		"AccountId": request.FromAccount,
	})
	if err != nil {
		return "", err
	}

	var status TransferStatus = StatusPending
	var approved, approvalReceived bool

	err = workflow.SetQueryHandler(ctx, "getStatus", func() (TransferStatus, error) {
		return status, nil
	})
	if err != nil {
		return "", err
	}

	approvalChannel := workflow.GetSignalChannel(ctx, "approve")
	workflow.Go(ctx, func(ctx workflow.Context) {
		var value bool
		approvalChannel.Receive(ctx, &value)
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})

	// TODO: Create activity context with Summary for withdraw
	withdrawCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		// TODO: Add Summary field
	})
	err = workflow.ExecuteActivity(withdrawCtx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		status = StatusApproved
		// TODO: Create activity context with Summary for deposit
		depositCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 5 * time.Second,
			// TODO: Add Summary field
		})
		err = workflow.ExecuteActivity(depositCtx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
		if err != nil {
			status = StatusFailed
			return "", err
		}
		status = StatusCompleted
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	// TODO: Create activity context with Summary for refund
	refundCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		// TODO: Add Summary field
	})
	err = workflow.ExecuteActivity(refundCtx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	status = StatusCancelled
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
