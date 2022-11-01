package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/internal/domain"
)

type MockStorage struct {
	dataMock       []domain.Products
	readWasCalled  bool
	writeWasCalled bool
	errWrite       string
	errRead        string
}

func (m *MockStorage) Read(data interface{}) error {
	m.readWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Products)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	m.writeWasCalled = true
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Products)
	m.dataMock = a
	return nil
}

func TestServiceUpdate(t *testing.T) {
	//arrange
	product := domain.Products{
		Id:           1,
		Name:         "tele",
		Color:        "negro",
		Price:        2000,
		Stock:        50,
		Code:         "tele123",
		Published:    true,
		CreationDate: "10-2022",
	}
	updated := domain.Products{
		Id:           1,
		Name:         "televisor",
		Color:        "negro",
		Price:        2500,
		Stock:        100,
		Code:         "tele123",
		Published:    true,
		CreationDate: "10-2022",
	}
	//act
	database := []domain.Products{product}
	mockStorage := MockStorage{dataMock: database}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	result, err := service.UpdateProduct(
		updated.Id,
		updated.Stock,
		updated.Name,
		updated.Color,
		updated.Code,
		updated.CreationDate,
		updated.Price,
		updated.Published,
	)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, result, updated)
	assert.True(t, mockStorage.readWasCalled)
}

func TestServiceDelete(t *testing.T) {
	//arrange
	prod1 := domain.Products{
		Id:           1,
		Name:         "tele",
		Color:        "negro",
		Price:        2000,
		Stock:        50,
		Code:         "tele123",
		Published:    true,
		CreationDate: "10-2022",
	}
	prod2 := domain.Products{
		Name:         "nevera",
		Id:           1,
		Color:        "blanco",
		Price:        3000,
		Stock:        50,
		Code:         "nevera123",
		Published:    true,
		CreationDate: "10-2022",
	}
	database := []domain.Products{prod1, prod2}
	//act
	mockStorage := MockStorage{dataMock: database}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	result, err := service.DeleteElement(1)
	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readWasCalled)
	assert.True(t, mockStorage.writeWasCalled)
	assert.Equal(t, mockStorage.dataMock, result)

	// result, err = service.DeleteElement(2)
	// assert.NotNil(t, err)
	// assert.True(t, mockStorage.readWasCalled)
	// assert.True(t, mockStorage.writeWasCalled)
	// assert.Equal(t, mockStorage.dataMock, result)
}
