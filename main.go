package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	//Example mock
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("John", 30, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("customer not found"))
	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Name: %s, Age: %d", name, age)

}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {

	args := m.Called(id)
	//Warring: check type before return
	return args.String(0), args.Int(1), args.Error(2)

}
