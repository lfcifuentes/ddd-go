package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := NewCustomer(tc.name)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			// Check if the error matches the expected error
			assert.Equal(t, tc.expectedErr, err, "unexpected error for name: %s", tc.name)
		})
	}
}

func TestCustomer_GetID(t *testing.T) {
	// Create a new customer
	customer, err := NewCustomer("Luis")
	assert.NoError(t, err, "unexpected error creating customer")
	// Check if the ID is not empty
	assert.NotEmpty(t, customer.GetID().String(), "expected non-empty ID")
}

func TestCustomer_SetID(t *testing.T) {
	// Create a new customer
	customer, err := NewCustomer("Luis")
	assert.NoError(t, err, "unexpected error creating customer")
	// current ID
	currentID := customer.GetID()
	// Set a new ID
	newID := uuid.New()
	customer.SetID(newID)

	// Check if the ID is not empty and different from the previous one
	assert.NotEmpty(t, customer.GetID().String(), "expected non-empty ID")
	assert.NotEqual(t, currentID, customer.GetID(), "expected different ID")
	assert.Equal(t, newID, customer.GetID(), "expected the same new ID set")
}

func TestCustomer_SetName(t *testing.T) {
	// Create a new customer
	customer, err := NewCustomer("Luis")
	assert.NoError(t, err, "unexpected error creating customer")

	// Set a new name
	newName := "Fernando"
	customer.SetName(newName)

	// Check if the name is not empty and different from the previous one
	assert.NotEmpty(t, customer.GetName(), "expected non-empty name")
	assert.Equal(t, newName, customer.GetName(), "expected the new name set")
}

func TestCustomer_EmptyID(t *testing.T) {
	// Create a new customer
	customer, _ := NewCustomer("")

	// Set a new ID
	newID := uuid.New()
	customer.SetID(newID)

	assert.Empty(t, customer.person.Name, "expected empty name")
}

func TestCustomer_EmptyName(t *testing.T) {
	// Create a new customer
	customer, _ := NewCustomer("")
	// Set a new ID
	name := "Luis"
	customer.SetName(name)

	assert.Equal(t, uuid.Nil, customer.person.ID, "expected non-nil ID")
}
