package salesforce

import en "ddd/domain/entity"

var customers = make(map[int]*Customer)

// Customer represents a customer
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// GetId returns the customer id
// note: should be GetID or GetId. we should have a convention and use it everywhere
func (c *Customer) GetId() int {
	return c.ID
}

// GetFirstName returns the customer first name
func (c *Customer) GetFirstName() string {
	return c.FirstName
}

// GetCustomerByID returns a customer by id
func GetCustomerByID(id int) (*Customer, error) {
	if customer, ok := customers[id]; ok {
		return customer, nil
	}
	if id == 1 {
		return &Customer{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Phone:     "",
		}, nil
	}

	return &Customer{}, nil
}

// Create creates a new customer
func (c *Customer) Create() error {
	customers[c.ID] = c
	return nil
}

// ------------------------ < Adapter > --------------------------------

// ToEntityCustomer converts a Customer to an entity.Customer
func (c *Customer) ToEntityCustomer() *en.Customer {
	return &en.Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
	}
}

// ToSalesforceCustomer converts an entity.Customer to a Customer
func ToSalesforceCustomer(c *en.Customer) *Customer {
	return &Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
	}
}
