package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePacksNeeded(t *testing.T) {
    testPackSizes := []int{5000, 2000, 1000, 500, 250}

	tests := []struct {
		name          string
		orderQuantity int
		packSizes     []int
		expectedPacks map[int]int
		expectError   bool
	}{
		{
			name:          "Test for 1 item",
			orderQuantity: 1,
			packSizes:     testPackSizes,
			expectedPacks: map[int]int{250: 1},
			expectError:   false,
		},
		{
			name:          "Test for negative order quantity",
			orderQuantity: -1,
			packSizes:     testPackSizes,
			expectedPacks: nil,
			expectError:   true,
		},
		{
			name:          "Test for 0 items",
			orderQuantity: 0,
			packSizes:     testPackSizes,
			expectedPacks: nil,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculatePacksNeeded(tt.orderQuantity, tt.packSizes)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPacks, result)
			}
		})
	}
}
