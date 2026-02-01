package exercise3

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

// Transfer processes a money transfer with status tracking and query support.
// TODO: Implement query handler for "getStatus" that returns current status
// TODO: Update status at each step:
//   - StatusPending at start
//   - StatusApproved when approved
//   - StatusCompleted when deposit succeeds
//   - StatusCancelled when refunded
//   - StatusFailed on errors
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	ao := workflow.ActivityOptions{StartToCloseTimeout: 5 * time.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var status TransferStatus = StatusPending
	var approved, approvalReceived bool

	// TODO: Set up query handler for "getStatus"

	err := workflow.SetSignalChannel(ctx, "approve", func(ctx workflow.Context, value bool) {
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, Withdraw, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		// TODO: Update status to StatusFailed
		return "", err
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		// TODO: Update status to StatusApproved
		err = workflow.ExecuteActivity(ctx, Deposit, request.ToAccount, request.Amount).Get(ctx, nil)
		if err != nil {
			// TODO: Update status to StatusFailed
			return "", err
		}
		// TODO: Update status to StatusCompleted
		logger.Info("Transfer approved and completed")
		return "Transfer completed successfully", nil
	}

	err = workflow.ExecuteActivity(ctx, Refund, request.FromAccount, request.Amount).Get(ctx, nil)
	if err != nil {
		// TODO: Update status to StatusFailed
		return "", err
	}
	// TODO: Update status to StatusCancelled
	logger.Info("Transfer rejected and refunded")
	return "Transfer rejected and refunded", nil
}
