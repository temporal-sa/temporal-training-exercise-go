package exercise2

import (
	"go.temporal.io/sdk/workflow"
)

type MoneyTransferWorkflow struct{}

// Transfer processes a money transfer with approval mechanism.
// TODO: Implement the complete transfer workflow:
// 1. Configure ActivityOptions with StartToCloseTimeout of 5 seconds
// 2. Set up signal channel for "approve" signal using workflow.GetSignalChannel
// 3. Use workflow.Go to receive signals in a goroutine
// 4. Withdraw from source account
// 5. Wait for approval signal using workflow.Await
// 6. If approved: deposit to target account
// 7. If not approved: refund to source account
// 8. Return appropriate status message
func (w *MoneyTransferWorkflow) Transfer(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting transfer", "request", request)

	// TODO: Implement workflow logic
	return "", nil
}
