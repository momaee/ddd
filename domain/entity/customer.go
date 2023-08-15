package entity

type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (c *Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c *Customer) ChangeFirstName(newFirstName string) {
	c.FirstName = newFirstName
}
