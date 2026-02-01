package solution5

type TransferRequest struct {
	FromAccount string
	ToAccount   string
	Amount      float64
	TransferID  string
}

type TransferStatus string

const (
	StatusPending   TransferStatus = "PENDING"
	StatusApproved  TransferStatus = "APPROVED"
	StatusCompleted TransferStatus = "COMPLETED"
	StatusFailed    TransferStatus = "FAILED"
	StatusCancelled TransferStatus = "CANCELLED"
)
