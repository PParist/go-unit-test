package services

import (
	"unittest/repositories"
)

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	peomoRepo repositories.PromotionRepository
}

func NewPromotionService(promoRepo repositories.PromotionRepository) PromotionService {
	return &promotionService{peomoRepo: promoRepo}
}

func (s *promotionService) CalculateDiscount(amount int) (int, error) {

	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotions, err := s.peomoRepo.GetPromotions()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotions.PurchaseMin {
		return amount - (promotions.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}
