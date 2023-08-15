package entity

// Customer is a struct that represents a customer
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// FullName returns the full name of the customer
func (c *Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}

// ChangeFirstName changes the first name of the customer
func (c *Customer) ChangeFirstName(newFirstName string) {
	c.FirstName = newFirstName
}
