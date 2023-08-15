package repository

import (
	en "ddd/domain/entity"
)

type CRM interface {
	// GetCustomerByID returns a customer by ID
	GetCustomerByID(id int) (*en.Customer, error)
}

type RMS interface {
	// GetCustomerByID returns a customer by ID
	GetCustomerByID(id int) (*en.Customer, error)
}

type Integration struct {
	crm CRM
	rms RMS
}

func NewIntegration(crm CRM, rms RMS) *Integration {
	return &Integration{
		crm: crm,
		rms: rms,
	}
}
