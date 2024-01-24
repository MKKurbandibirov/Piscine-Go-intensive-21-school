package postgres

import (
	"client/internal/domain"
	"context"
	"fmt"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	anomalyTablename = "anomalies"
)

type Communicator struct {
	ctx context.Context
	db  *pgxpool.Pool
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
	sql, args := sqlbuilder.InsertInto(anomalyTablename).Values(aml.SessionID, aml.Frequency, aml.Time).Build()

	for tag := 1; tag <= 3; tag++ {
		sql = strings.Replace(sql, "?", fmt.Sprintf("$%v", tag), 1)
	}

	_, err := c.db.Exec(c.ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
