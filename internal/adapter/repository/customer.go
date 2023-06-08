package repository

import "go-test-api/internal/port"

type customer client

var _ port.CustomerRepository = (*customer)(nil)

func (c *Client) GetCustomerRepository() port.CustomerRepository {
	return c.customer
}
