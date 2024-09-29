package repositories

import (
	"unittest/entities"

	"github.com/stretchr/testify/mock"
)

type protionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *protionRepositoryMock {
	return &protionRepositoryMock{}
}

func (m *protionRepositoryMock) GetPromotions() (*entities.Promotion, error) {
	args := m.Called()
	return args.Get(0).(*entities.Promotion), args.Error(1)
}
