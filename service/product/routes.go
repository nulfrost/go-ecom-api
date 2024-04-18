package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nulfrost/ecom/types"
	"github.com/nulfrost/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.getProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.createProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", h.getProductById).Methods(http.MethodGet)
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.ProductPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Price:       payload.Price,
		Description: payload.Description,
		Image:       payload.Image,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}

func (h *Handler) getProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	product, err := h.store.GetProductByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}
