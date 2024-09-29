package repositories

import "unittest/entities"

type PromotionRepository interface {
	GetPromotions() (*entities.Promotion, error)
}
