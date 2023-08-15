package repository

// Repositories is a struct that contains all repositories
type Repositories struct {
	Customer *CustomerRepo
}

// NewRepositories is a function that returns a new Repositories struct
func NewRepositories(integration *Integration) *Repositories {
	return &Repositories{
		Customer: NewCustomerRepo(integration),
	}
}
