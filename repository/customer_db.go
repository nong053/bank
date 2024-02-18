package repository

import "github.com/jmoiron/sqlx"

//adapter
type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	cusotomers := []Customer{}
	query := "select customer_id,name,date_of_birth,city,zipcode,status from customer"
	err := r.db.Select(&cusotomers, query)
	if err != nil {
		return nil, err
	}
	return cusotomers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {

	customer := Customer{}
	query := "select customer_id,name,date_of_birth,city,zipcode,status from customer where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
