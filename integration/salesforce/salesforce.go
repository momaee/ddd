package salesforce

import (
	en "ddd/domain/entity"
	repo "ddd/domain/repository"
)

var _ repo.CRM = &Salesforce{}

type Salesforce struct {
}

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
