package repository

import en "ddd/domain/entity"

type CustomerRepo struct { // note: could just keep this as a struct and remove the interface!
	integration *Integration
}

func NewCustomerRepo(integration *Integration) *CustomerRepo {
	return &CustomerRepo{
		integration: integration,
	}
}

func (r *CustomerRepo) FindById(id int) (*en.Customer, error) {
	return r.integration.crm.GetCustomerByID(id)
}
