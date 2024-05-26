package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lfcifuentes/ddd-go/aggregate"
	"github.com/lfcifuentes/ddd-go/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}

	// create fake customer
	cust, err := aggregate.NewCustomer("Luis Cifuentes")
	if err != nil {
		t.Fatalf("failed to create customer: %v", err)
	}

	id := cust.GetID()

	// create repository
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:          "No Customer By ID",
			id:            uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedError: customer.ErrCustomerNotFound,
		}, {
			name:          "Customer By ID",
			id:            id,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}

			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
