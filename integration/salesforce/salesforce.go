package salesforce

import (
	en "ddd/domain/entity"
	repo "ddd/domain/repository"
)

// Salesforce implements the CRM interface
var _ repo.CRM = &Salesforce{}

// Salesforce is a CRM
type Salesforce struct {
	// TODO: add salesforce client
}

// New returns a new Salesforce CRM
func New() *Salesforce {
	return &Salesforce{}
}

// GetCustomerByID returns a customer by ID
func (s *Salesforce) GetCustomerByID(id int) (*en.Customer, error) {
	sfCustomer, err := GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	return sfCustomer.ToEntityCustomer(), nil
}

// CreateCustomer creates a customer
func (s *Salesforce) CreateCustomer(customer *en.Customer) error {
	sfCustomer := ToSalesforceCustomer(customer)
	return sfCustomer.Create()
}
