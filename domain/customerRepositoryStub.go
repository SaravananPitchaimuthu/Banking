package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Saravanan", "Trichy", "620102", "04-08-1997", "active"},
		{"1002", "Saranya", "Trichy", "620102", "04-08-1997", "active"},
	}
	return CustomerRepositoryStub{customers}
}
