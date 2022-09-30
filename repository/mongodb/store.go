package mongodb

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/antlia-io/antlia-notification-engine/config"
	"github.com/antlia-io/antlia-notification-engine/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var gdb *mongo.Database
var storeOnce sync.Once
var store repository.Store

type Store struct {
	db *mongo.Database
}

func SharedStore() repository.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})

	return store
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := config.LoadConfiguration("config/config.go")
	ctx := context.Background()
	clint := options.Client().ApplyURI(cfg.MongoURL)

	// Access payment service from the default app
	client, err := mongo.Connect(ctx, clint)
	if err != nil {
		return fmt.Errorf("error initializing mongo app: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	gdb = client.Database("Antlia")
	gdb.Collection("Notification")
	return nil
}

//BeginTx return store instance
func (s *Store) BeginTx() (repository.Store, error) {

	return NewStore(gdb), nil
}

//Rollback
func (s *Store) Rollback() error {
	return nil
}

//CommitTx
func (s *Store) CommitTx() error {
	return nil
}
