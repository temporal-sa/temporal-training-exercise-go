package solution7

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func Withdraw(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Withdrawing", "amount", amount, "account", account)
	if account == "invalid-account" {
		return fmt.Errorf("withdrawal failed - invalid account")
	}
	return nil
}

func Deposit(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Depositing", "amount", amount, "account", account)
	if account == "invalid-account" {
		return fmt.Errorf("deposit failed - invalid account")
	}
	return nil
}

func Refund(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Refunding", "amount", amount, "account", account)
	return nil
}
