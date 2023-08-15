package repository

import (
	en "ddd/domain/entity"
)

// CRM is a customer relationship management system
type CRM interface {
	// GetCustomerByID returns a customer by ID
	GetCustomerByID(id int) (*en.Customer, error)

	// CreateCustomer creates a customer
	CreateCustomer(customer *en.Customer) error
}

// RMS is a retail management system
type RMS interface {
	// GetCustomerByID returns a customer by ID
	GetCustomerByID(id int) (*en.Customer, error)
}

// Integration is a integration of all the vendors
type Integration struct {
	crm CRM
	rms RMS
}

// NewIntegration creates a new integration
func NewIntegration(crm CRM, rms RMS) *Integration {
	return &Integration{
		crm: crm,
		rms: rms,
	}
}
