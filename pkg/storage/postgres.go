package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/tjcain/example-subscription-api/pkg/adding"
	"github.com/tjcain/example-subscription-api/pkg/config"
)

// Postgres implements a PostgreSQL backed Store.
type Postgres struct {
	db *sqlx.DB
}

// NewPostgres creates a new PostgreSQL backed Store.
func NewPostgres(cfg config.DatabaseConfig) (Store, error) {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable",
		cfg.Name, cfg.Username, cfg.Password,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "could not open database")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	return &Postgres{db}, nil
}

func rollback(op string, tx *sqlx.Tx, err error) error {
	log.Printf("%s: transaction roll back due to error: %v", op, err)
	if err := tx.Rollback(); err != nil {
		log.Printf("%s: could not roll back transaction: %v", op, err)
	}
	return err
}

func commit(op string, tx *sqlx.Tx) error {
	if err := tx.Commit(); err != nil {
		log.Printf("%s: could not commit: %v", op, err)
		return err
	}
	return nil
}

// Migrate performs a database migration.
func (p *Postgres) Migrate() {}

// AddSubscription implements storage.Store.
func (p *Postgres) AddSubscription(adding.Subscription) (int, error) {

	return -1, nil
}

// UpdateSubscription implements storage.Store.
func (p *Postgres) UpdateSubscription(Subscription) error { return nil }

// DeleteSubscription implements storage.Store.
func (p *Postgres) DeleteSubscription(id int) error { return nil }

// AddProduct implements storage.Store.
func (p *Postgres) AddProduct(adding.Product) (int, error) {

	return -1, nil
}

// UpdateProduct implements storage.Store.
func (p *Postgres) UpdateProduct(Product) error { return nil }

// DeleteProduct implements storage.Store.
func (p *Postgres) DeleteProduct(id int) error { return nil }
