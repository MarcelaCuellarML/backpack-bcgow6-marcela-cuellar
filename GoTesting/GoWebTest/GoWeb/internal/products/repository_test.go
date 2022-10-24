package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/GoWebTest/GoWeb/internal/domain"
)

type StubReadProducts struct {
	mockData      []domain.Products
	readWasCalled bool
}

func (s *StubReadProducts) Read(data interface{}) error {
	s.readWasCalled = true
	a, ok := data.(*[]domain.Products)
	if !ok {
		return errors.New("no ha sido posible hacer una lectura")
	}
	*a = s.mockData
	return nil
}

func (s *StubReadProducts) Write(data interface{}) error {
	return nil
}

func TestGetAllProds(t *testing.T) {
	//arrange
	products := []domain.Products{
		{
			Name:         "tele",
			Id:           1,
			Color:        "negro",
			Price:        1000,
			Stock:        20,
			Code:         "tele123",
			Published:    true,
			CreationDate: "10-2022",
		},
		{
			Name:         "nevera",
			Id:           1,
			Color:        "blanco",
			Price:        3000,
			Stock:        50,
			Code:         "nevera123",
			Published:    true,
			CreationDate: "10-2022",
		},
	}
	//act
	db := &StubReadProducts{mockData: products}
	repo := NewRepository(db)
	response, err := repo.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, products, response)
}

func TestUpdateName(t *testing.T) {
	//arrange
	data := []domain.Products{
		{
			Id:           1,
			Name:         "Televisor",
			Color:        "black",
			Price:        2500,
			Stock:        100,
			Code:         "tele123",
			Published:    true,
			CreationDate: "10-2022",
		},
	}
	productExpected := data[0]
	updateNameData := domain.Products{
		Name:         "Televisor",
		Color:        "black",
		Price:        2500,
		Stock:        100,
		Code:         "tele123",
		Published:    true,
		CreationDate: "10-2022",
	}
	//act
	productExpected.Name = updateNameData.Name
	db := &StubReadProducts{mockData: data, readWasCalled: false}
	r := NewRepository(db)
	response, err := r.UpdateProduct(1, updateNameData)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, productExpected.Name, response.Name)
	assert.True(t, db.readWasCalled)
}

func TestDelete(t *testing.T) {
	//arrange
	data := []domain.Products{
		{
			Name:         "tele",
			Id:           1,
			Color:        "negro",
			Price:        1000,
			Stock:        20,
			Code:         "tele123",
			Published:    true,
			CreationDate: "10-2022",
		},
		{
			Name:         "nevera",
			Id:           1,
			Color:        "blanco",
			Price:        3000,
			Stock:        50,
			Code:         "nevera123",
			Published:    true,
			CreationDate: "10-2022",
		},
	}
	expectedData := []domain.Products{
		{
			Name:         "nevera",
			Id:           1,
			Color:        "blanco",
			Price:        3000,
			Stock:        50,
			Code:         "nevera123",
			Published:    true,
			CreationDate: "10-2022",
		},
	}
	idToDelete := 1
	//act
	db := &StubReadProducts{mockData: data, readWasCalled: false}
	r := NewRepository(db)
	response, err := r.DeleteElement(idToDelete)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, response)
	assert.True(t, db.readWasCalled)
}

func TestCreate(t *testing.T) {
	//arrange
	data := []domain.Products{
		{
			Name:         "tele",
			Id:           1,
			Color:        "negro",
			Price:        1000,
			Stock:        20,
			Code:         "tele123",
			Published:    true,
			CreationDate: "10-2022",
		},
	}
	expectedData := domain.Products{
		Name:         "nevera",
		Id:           2,
		Color:        "blanco",
		Price:        3000,
		Stock:        50,
		Code:         "nevera123",
		Published:    true,
		CreationDate: "10-2022",
	}
	createNewData := domain.Products{
		Name:         "nevera",
		Id:           2,
		Color:        "blanco",
		Price:        3000,
		Stock:        50,
		Code:         "nevera123",
		Published:    true,
		CreationDate: "10-2022",
	}
	//act
	db := &StubReadProducts{mockData: data, readWasCalled: false}
	r := NewRepository(db)
	response, err := r.AgregarProducto(createNewData)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, response)
	assert.True(t, db.readWasCalled)
}
