package exercise6

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

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

	err = workflow.SetSignalChannel(ctx, "approve", func(ctx workflow.Context, value bool) {
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})
	if err != nil {
		return "", err
	}

	withdrawCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		Summary:             fmt.Sprintf("Withdrawing funds from account %s", request.FromAccount),
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
		depositCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 5 * time.Second,
			Summary:             fmt.Sprintf("Depositing funds to account %s", request.ToAccount),
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

	refundCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		Summary:             fmt.Sprintf("Refunding funds to account %s", request.FromAccount),
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
