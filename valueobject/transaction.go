package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	// all values lowercase since they are immutable
	amount int
	// from is the sender of the transaction
	from uuid.UUID
	// to is the receiver of the transaction
	to uuid.UUID
	// createdAt is the time the transaction was created
	createdAt time.Time
}
