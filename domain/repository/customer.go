package repository

import en "ddd/domain/entity"

type CustomerRepo struct {
	integration *Integration
}

// NewCustomerRepo returns a new customer repository
func NewCustomerRepo(integration *Integration) *CustomerRepo {
	return &CustomerRepo{
		integration: integration,
	}
}

// FindById returns a customer by id
func (r *CustomerRepo) FindById(id int) (*en.Customer, error) { // CustomerByID
	// main logic is here
	return r.integration.crm.GetCustomerByID(id)
}

// Create creates a new customer
func (r *CustomerRepo) Create(customer *en.Customer) error {
	return r.integration.crm.CreateCustomer(customer)
}
