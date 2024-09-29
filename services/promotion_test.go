package services_test

import (
	"errors"
	"testing"
	"unittest/entities"
	"unittest/repositories"
	"unittest/services"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{
			name:            "Test 1",
			purchaseMin:     100,
			discountPercent: 20,
			amount:          100,
			expected:        80,
		},
		{
			name:            "Test 2",
			purchaseMin:     100,
			discountPercent: 20,
			amount:          200,
			expected:        160,
		},
		{
			name:            "Test 3",
			purchaseMin:     100,
			discountPercent: 20,
			amount:          300,
			expected:        240,
		},
		{
			name:            "Test 4 failed case",
			purchaseMin:     100,
			discountPercent: 20,
			amount:          50,
			expected:        30,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			//TODO: Implement test by AAA pattern (Arrange, Act, Assert)
			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotions").Return(&entities.Promotion{
				ID: 1,
				//Name:            "Promo 1",
				PurchaseMin:     tc.purchaseMin,
				DiscountPercent: tc.discountPercent,
			}, nil)
			promoservice := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoservice.CalculateDiscount(tc.amount)
			expected := tc.expected

			// Assert
			assert.Equal(t, expected, discount)
		})

		t.Run(tc.name+" Zero Amount", func(t *testing.T) {
			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotions").Return(&entities.Promotion{
				ID: 1,
				//Name:            "Promo 1",
				PurchaseMin:     100,
				DiscountPercent: 20,
			}, nil)
			promoservice := services.NewPromotionService(promoRepo)

			// Act
			_, err := promoservice.CalculateDiscount(0)

			// Assert
			assert.ErrorIs(t, err, services.ErrZeroAmount)
			//TODO: check logic block call check zero value call after something
			promoRepo.AssertNotCalled(t, "GetPromotions")
		})

		t.Run(tc.name+" Repository Error", func(t *testing.T) {
			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotions").Return(&entities.Promotion{}, errors.New(""))
			promoservice := services.NewPromotionService(promoRepo)

			// Act
			_, err := promoservice.CalculateDiscount(100)

			// Assert
			assert.ErrorIs(t, err, services.ErrRepository)
		})
	}
}
