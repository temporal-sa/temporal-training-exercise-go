package exercise4

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

// Transfer processes a money transfer with search attributes for filtering.
// TODO: Use workflow.UpsertSearchAttributes to set "AccountId" to request.FromAccount
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	// TODO: Set search attributes using workflow.UpsertSearchAttributes
	// Map key: "AccountId", value: request.FromAccount

	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var status TransferStatus = StatusPending
	var approved, approvalReceived bool

	err := workflow.SetQueryHandler(ctx, "getStatus", func() (TransferStatus, error) {
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

	err = workflow.ExecuteActivity(ctx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		status = StatusApproved
		err = workflow.ExecuteActivity(ctx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
		if err != nil {
			status = StatusFailed
			return "", err
		}
		status = StatusCompleted
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	err = workflow.ExecuteActivity(ctx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		status = StatusFailed
		return "", err
	}
	status = StatusCancelled
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
