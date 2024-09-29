//go:build integration

package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"unittest/entities"
	handlers "unittest/handler"
	"unittest/repositories"
	"unittest/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	//TODO: Implement test by AAA pattern (Arrange, Act, Assert)
	t.Run("success", func(t *testing.T) {
		// Arrange
		amount := 100
		expected := 80
		promorepo := repositories.NewPromotionRepositoryMock()
		promorepo.On("GetPromotions").Return(&entities.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)
		promoService := services.NewPromotionService(promorepo)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusOK, res.StatusCode)

		body, _ := io.ReadAll(res.Body)
		assert.Equal(t, strconv.Itoa(expected), string(body))
	})
}
