package aggregate_test

import (
	"go-ddd/aggregate"
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			test:        "validvalues",
			name:        "name",
			description: "description",
			price:       9.99,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.description, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
