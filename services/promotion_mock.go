package services

import "github.com/stretchr/testify/mock"

type promoServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *promoServiceMock {
	return &promoServiceMock{}
}

func (m *promoServiceMock) CalculateDiscount(amount int) (int, error) {
	args := m.Called(amount)
	return args.Int(0), args.Error(1)
}
