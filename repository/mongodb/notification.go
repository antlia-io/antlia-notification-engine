package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/antlia-io/antlia-notification-engine/models"
)

func (s *Store) AddNotification(swap *models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := s.db.Collection("Notification").InsertOne(ctx, swap)
	if err != nil {
		log.Printf("An error has occurred while storing swap history in store: %s", err)
		return err
	}
	return nil
}
