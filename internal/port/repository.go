package port

type Repositories interface {
	GetCustomerRepository() CustomerRepository
}
