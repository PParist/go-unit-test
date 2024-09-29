package handlers

import (
	"strconv"
	"unittest/services"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return &promotionHandler{promoService: promoService}
}

func (h *promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 	"message": "Invalid amount",
		// })
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discountedAmount, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		// return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		// 	"message": "NotFound",
		// })
		return c.SendStatus(fiber.StatusNotFound)
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"amount":           amount,
	// 	"discountedAmount": discountedAmount,
	// })
	return c.SendString(strconv.Itoa(discountedAmount))
}
