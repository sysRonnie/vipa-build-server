package customer

import (
	"context"
	"errors"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strconv"
)



type Service struct {
	store CustomerStore
}

func NewCustomerService(store CustomerStore) *Service {
	return &Service{
		store: store,
	}
}

type CustomerService interface {
	ProcessUpdateCustomer(ctx context.Context, customer CustomerInsertRequest) error
	ProcessCreateCustomer(ctx context.Context, customer CustomerInsertRequest) error
}

func (s *Service) ProcessCreateCustomer(
	ctx context.Context,
	customer CustomerInsertRequest,
) error {

	advisor := advisor.FromContext(ctx)

	fullName := customer.Customer.FirstName + " " + customer.Customer.LastName

	advisor.Log("processing_create_customer_request_for: " + fullName,)

	if customer.Customer.FirstName == "" || customer.Customer.LastName == "" {

		err := errors.New("customer first name or last name is empty")

		advisor.Error("customer_name_validation_failed",err)

		return err
	}

	exists, err := s.store.CheckCustomerExists(
		ctx,
		customer.Customer.FirstName,
		customer.Customer.LastName,
	)

	if err != nil {
		advisor.Error(
			"failed_to_check_customer_exists",
			err,
		)

		return ErrDatabaseFailure
	}

	if exists {
		return ErrCustomerAlreadyExists
	}

	recycledExists, err := s.store.CheckCustomerExistsRecycled(
		ctx,
		customer.Customer.FirstName,
		customer.Customer.LastName,
	)

	if err != nil {
		advisor.Error("failed_to_check_recycled_customer_exists",err)
		return ErrDatabaseFailure
	}

	if recycledExists {
		return ErrCustomerAlreadyExistsRecycled
	}

	return s.store.InsertCustomer(
		ctx,
		customer.Customer,
	)
}

func (s *Service) ProcessUpdateCustomer(ctx context.Context, customer CustomerInsertRequest) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("Processing update customer request with ID: " + strconv.Itoa(customer.Customer.ID))
	if customer.Customer.ID < 0 {
		advisor.Error("invalid customer ID: " + strconv.Itoa(customer.Customer.ID), ErrFullNameRequired)
		return network.ErrInvalidPayload

	}

	if customer.Customer.FirstName == "" || customer.Customer.LastName == "" {
		advisor.Error("invalid customer name for customer ID: " + strconv.Itoa(customer.Customer.ID), ErrFullNameRequired)
		return ErrFullNameRequired
	}

	return s.store.UpdateCustomer(ctx, customer)
}