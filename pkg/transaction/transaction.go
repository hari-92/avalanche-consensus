package transaction

import "time"

type Status string

const (
	Pending   Status = "pending"
	Confirmed Status = "confirmed"
	Rejected  Status = "rejected"
)

type Transaction struct {
	ID     uint
	Data   interface{}
	Status Status
}

func NewTransaction(data interface{}) *Transaction {
	now := time.Now()
	tx := &Transaction{
		ID:     uint(now.Unix()),
		Data:   data,
		Status: Pending,
	}
	return tx
}
