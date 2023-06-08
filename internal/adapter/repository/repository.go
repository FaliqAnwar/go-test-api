package repository

import (
	"context"

	adapterPostgres "go-test-api/internal/adapter/postgres"
	"go-test-api/internal/model"
	"go-test-api/internal/port"

	"gorm.io/gorm"
)

type Client struct {
	customer *customer

	common client
	db     *gorm.DB
}

type client struct {
	c *Client
}

var _ port.Repositories = (*Client)(nil)

func NewRepository(ctx context.Context, cfg model.PostgresClient) *Client {
	db := adapterPostgres.New(cfg)
	c := &Client{
		db: db,
	}

	c.common.c = c
	c.customer = (*customer)(&c.common)

	return c
}
