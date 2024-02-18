package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerId: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "11011", DateOfBirth: "20202000"},
		{CustomerId: 1002, Name: "Ashish", City: "New Delhi", ZipCode: "11011", DateOfBirth: "20202000"},
	}
	return customerRepositoryMock{customers: customers}
}
func (r customerRepositoryMock) GetAll() ([]Customer, error) {

	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerId == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

//GetCustomers() ([]CustomerResponse, error)
//GetCustomer(int) (*CustomerResponse, error)
