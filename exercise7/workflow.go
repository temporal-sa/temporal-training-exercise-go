package exercise7

import (
	"strconv"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

// Transfer processes a money transfer with manual retry capability.
// TODO: Implement manual retry loops for Withdraw and Deposit activities:
//   1. Set MaximumAttempts to 1 in RetryPolicy to disable automatic retries
//   2. Wrap activity calls in a for loop
//   3. On error, set status to StatusRetrying and wait for retry signal
//   4. Handle retry signal to update request fields based on key/value
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	err := workflow.UpsertSearchAttributes(ctx, map[string]interface{}{
		"AccountId": request.FromAccount,
	})
	if err != nil {
		return "", err
	}

	// TODO: Configure ActivityOptions with MaximumAttempts: 1
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		// TODO: Add RetryPolicy with MaximumAttempts: 1
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var status TransferStatus = StatusPending
	var approved, approvalReceived bool
	var retryRequested bool
	updatedRequest := request

	err = workflow.SetQueryHandler(ctx, "getStatus", func() (TransferStatus, error) {
		return status, nil
	})
	if err != nil {
		return "", err
	}

	err = workflow.SetSignalChannel(ctx, "approve", func(ctx workflow.Context, value bool) {
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})
	if err != nil {
		return "", err
	}

	// TODO: Set up signal channel for "retry" that handles RetryUpdate
	// Use workflow.GetSignalChannel and workflow.Go with a loop
	// Update updatedRequest fields based on update.Key ("fromAccount", "toAccount", "amount")
	// Set retryRequested = true

	// TODO: Wrap Withdraw in retry loop
	err = workflow.ExecuteActivity(ctx, Withdraw, updatedRequest.FromAccount, updatedRequest.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		status = StatusApproved
		// TODO: Wrap Deposit in retry loop
		err = workflow.ExecuteActivity(ctx, Deposit, updatedRequest.ToAccount, updatedRequest.Amount).Get(ctx, nil)
		if err != nil {
			status = StatusFailed
			return "", err
		}
		status = StatusCompleted
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	err = workflow.ExecuteActivity(ctx, Refund, updatedRequest.FromAccount, updatedRequest.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	status = StatusCancelled
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
