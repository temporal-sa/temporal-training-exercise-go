package solution6

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (s *WorkflowTestSuite) TestSuccessfulTransfer() {
	env := s.NewTestWorkflowEnvironment()
	env.RegisterWorkflow((&MoneyTransferWorkflow{}).Transfer)

	env.OnActivity(Withdraw, mock.Anything, "account-123", 100.0).Return(nil)
	env.OnActivity(Deposit, mock.Anything, "account-456", 100.0).Return(nil)

	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("approve", true)
	}, time.Second)

	request := TransferRequest{
		FromAccount: "account-123",
		ToAccount:   "account-456",
		Amount:      100.0,
		TransferID:  "transfer-1",
	}

	env.ExecuteWorkflow((&MoneyTransferWorkflow{}).Transfer, request)
	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())

	var result string
	s.NoError(env.GetWorkflowResult(&result))
	s.Equal("Transfer completed successfully", result)
}

func (s *WorkflowTestSuite) TestRejectedTransfer() {
	env := s.NewTestWorkflowEnvironment()
	env.RegisterWorkflow((&MoneyTransferWorkflow{}).Transfer)

	env.OnActivity(Withdraw, mock.Anything, "account-123", 100.0).Return(nil)
	env.OnActivity(Refund, mock.Anything, "account-123", 100.0).Return(nil)

	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("approve", false)
	}, time.Second)

	request := TransferRequest{
		FromAccount: "account-123",
		ToAccount:   "account-456",
		Amount:      100.0,
		TransferID:  "transfer-2",
	}

	env.ExecuteWorkflow((&MoneyTransferWorkflow{}).Transfer, request)
	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())

	var result string
	s.NoError(env.GetWorkflowResult(&result))
	s.Equal("Transfer rejected and refunded", result)
}

func (s *WorkflowTestSuite) TestQueryStatus() {
	env := s.NewTestWorkflowEnvironment()
	env.RegisterWorkflow((&MoneyTransferWorkflow{}).Transfer)

	env.OnActivity(Withdraw, mock.Anything, "account-123", 100.0).Return(nil)
	env.OnActivity(Deposit, mock.Anything, "account-456", 100.0).Return(nil)

	env.RegisterDelayedCallback(func() {
		var status TransferStatus
		val, err := env.QueryWorkflow("getStatus")
		s.NoError(err)
		s.NoError(val.Get(&status))
		s.Equal(StatusPending, status)
	}, 500*time.Millisecond)

	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("approve", true)
	}, time.Second)

	request := TransferRequest{
		FromAccount: "account-123",
		ToAccount:   "account-456",
		Amount:      100.0,
		TransferID:  "transfer-3",
	}

	env.ExecuteWorkflow((&MoneyTransferWorkflow{}).Transfer, request)
	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())
}
