package repository

type Repositories struct {
	Customer *CustomerRepo
}

func NewRepositories(integration *Integration) *Repositories {
	return &Repositories{
		Customer: NewCustomerRepo(integration),
	}
}
