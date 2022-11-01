package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	repo := NewRepository()
	service := NewService(repo)
	p := NewHandler(service)
	r := gin.Default()
	pr := r.Group("/api/v1//products")
	pr.GET("", p.GetProducts)
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetAllStatusOk(t *testing.T) {
	// arrange
	var resp []Product
	r := createServer()
	req, responseRecorder := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC", "")
	// act
	r.ServeHTTP(responseRecorder, req)
	// assert
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, responseRecorder.Code)
	//assert.Equal(t, , resp)
}
