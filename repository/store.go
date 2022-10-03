package repository

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error
	AddNotification(payment types.Log) error
}
