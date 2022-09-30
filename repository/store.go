package repository

import (
	"github.com/antlia-io/antlia-notification-engine/models"
)

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error
	AddNotification(payment *models.Notification) error
}
