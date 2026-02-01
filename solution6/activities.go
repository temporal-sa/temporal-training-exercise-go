package solution6

import (
	"context"
	"fmt"
	"math/rand"

	"go.temporal.io/sdk/activity"
)

func Withdraw(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Withdrawing", "amount", amount, "account", account)
	if rand.Float64() < 0.1 {
		return fmt.Errorf("withdrawal failed - insufficient funds")
	}
	return nil
}

func Deposit(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Depositing", "amount", amount, "account", account)
	if rand.Float64() < 0.05 {
		return fmt.Errorf("deposit failed - account not found")
	}
	return nil
}

func Refund(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Refunding", "amount", amount, "account", account)
	return nil
}
