package postgres

import (
	"client/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AnomalyStorer interface {
	StoreAnomaly(aml *domain.Anomaly) error
}

type Communicator struct {
	ctx    context.Context
	db     *pgxpool.Pool
	storer AnomalyStorer
}

func NewCommunicator(ctx context.Context, connURL string) (*Communicator, error) {
	pool, err := pgxpool.New(ctx, connURL)
	if err != nil {
		return nil, err
	}

	return &Communicator{
		ctx: ctx,
		db:  pool,
	}, nil
}

func (c *Communicator) Ping() error {
	return c.db.Ping(c.ctx)
}

func (c *Communicator) Store(aml *domain.Anomaly) error {
	return c.storer.StoreAnomaly(aml)
}
