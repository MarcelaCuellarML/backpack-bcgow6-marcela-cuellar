package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type MockService struct {
// 	dataMock      []Product
// 	readWasCalled bool
// 	//writeWasCalled bool
// 	//errWrite       string
// 	errRead string
// }

// func (m *MockService) GetAllBySeller(data interface{}, sellerID string) error {
// 	m.readWasCalled = true
// 	if m.errRead != "" {
// 		return fmt.Errorf(m.errRead)
// 	}
// 	a := data.(*[]Product)
// 	*a = m.dataMock
// 	return nil
// }

func TestGetAllBySeller(t *testing.T) {
	//arrange
	expectedData := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	sellerSearch := "FEX112AC"
	//act
	repo := NewRepository()
	service := NewService(repo)
	response, err := service.GetAllBySeller(sellerSearch)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, response, expectedData)
	//assert.True(t, mockStorage.readWasCalled)
}
