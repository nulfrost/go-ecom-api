package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nulfrost/ecom/types"
)

func TestProductServerHandlers(t *testing.T) {
	productStore := &mockProdctStore{}
	handler := NewHandler(productStore)

	t.Run("should fail if the product payload is invalid", func(t *testing.T) {
		payload := types.ProductPayload{
			Name:        "jordan 1",
			Description: "Sleek running shoe",
			Image:       "shoe.jpg",
			Quantity:    200,
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.createProduct)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should successfully create the product if the payload is valid", func(t *testing.T) {
		payload := types.ProductPayload{
			Name:        "jordan 1",
			Price:       235.00,
			Description: "Sleek running shoe",
			Image:       "shoe.jpg",
			Quantity:    200,
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.createProduct)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should return a list of products", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.getProducts)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should return a product by its ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products/{id}", handler.getProductById)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail if the product id is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/a", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products/{id}", handler.getProductById)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockProdctStore struct{}

func (m *mockProdctStore) GetProducts() ([]types.Product, error) {
	return []types.Product{}, nil
}

func (m *mockProdctStore) CreateProduct(types.Product) error {
	return nil
}

func (m *mockProdctStore) GetProductByID(id int) (*types.Product, error) {
	return &types.Product{}, nil
}
