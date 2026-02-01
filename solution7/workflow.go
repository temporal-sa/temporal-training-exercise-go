package solution7

import (
	"strconv"
	"time"

	"go.temporal.io/sdk/temporal"
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

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
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

	approvalChannel := workflow.GetSignalChannel(ctx, "approve")
	workflow.Go(ctx, func(ctx workflow.Context) {
		var value bool
		approvalChannel.Receive(ctx, &value)
		logger.Info("Approval received", "approved", value)
		approved = value
		approvalReceived = true
	})

	retryChannel := workflow.GetSignalChannel(ctx, "retry")
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			var update RetryUpdate
			retryChannel.Receive(ctx, &update)
			logger.Info("Retry update received", "key", update.Key, "value", update.Value)
			switch update.Key {
			case "fromAccount":
				updatedRequest.FromAccount = update.Value
			case "toAccount":
				updatedRequest.ToAccount = update.Value
			case "amount":
				if amt, err := strconv.ParseFloat(update.Value, 64); err == nil {
					updatedRequest.Amount = amt
				}
			}
			retryRequested = true
		}
	})

	for {
		err = workflow.ExecuteActivity(ctx, Withdraw, updatedRequest.FromAccount, updatedRequest.Amount).Get(ctx, nil)
		if err == nil {
			break
		}
		status = StatusRetrying
		retryRequested = false
		workflow.Await(ctx, func() bool { return retryRequested })
	}
	logger.Info("Withdrawal completed")

	logger.Info("Waiting for approval...")
	workflow.Await(ctx, func() bool { return approvalReceived })

	if approved {
		status = StatusApproved
		for {
			err = workflow.ExecuteActivity(ctx, Deposit, updatedRequest.ToAccount, updatedRequest.Amount).Get(ctx, nil)
			if err == nil {
				break
			}
			status = StatusRetrying
			retryRequested = false
			workflow.Await(ctx, func() bool { return retryRequested })
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
