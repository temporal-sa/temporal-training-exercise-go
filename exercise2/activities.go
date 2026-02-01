package exercise2

import (
	"context"
	"fmt"
	"math/rand/v2"

	"go.temporal.io/sdk/activity"
)

// Withdraw removes funds from an account.
// TODO: Implement withdraw logic with logging and simulate occasional failures
// Use rand.Float64() < 0.1 to simulate network failures and return an error
func Withdraw(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Withdrawing", "amount", amount, "account", account)
	// TODO: Add failure simulation
	return nil
}

// Deposit adds funds to an account.
// TODO: Implement deposit logic with logging and simulate occasional failures
// Use rand.Float64() < 0.05 to simulate network failures and return an error
func Deposit(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Depositing", "amount", amount, "account", account)
	if rand.Float64() < 0.05 {
		return fmt.Errorf("deposit failed - account not found")
	}
	return nil
}

// Refund returns funds to an account (compensation operation).
func Refund(ctx context.Context, account string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Refunding", "amount", amount, "account", account)
	return nil
}
