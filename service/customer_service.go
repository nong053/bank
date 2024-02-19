package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"net/http"
)

// adapter
type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

// GetCustomers implements CustomerService.
func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		//log.Println(err)
		logs.Error(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerId: customer.CustomerId,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

// GetCustomer implements CustomerService.
func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {

			//return nil, errors.New("customer not found")
			return nil, errs.AppError{
				Code:    http.StatusNotFound,
				Message: "customer not found",
			}
		}

		//log.Println(err)
		logs.Error(err)
		return nil, errs.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}
	}
	custReponse := CustomerResponse{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custReponse, nil
}
