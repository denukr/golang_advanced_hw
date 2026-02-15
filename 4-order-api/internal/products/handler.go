package products

import (
	"golang-adv/4-order-api/configs"
	"golang-adv/4-order-api/pkg/middleware"
	"golang-adv/4-order-api/pkg/req"
	"golang-adv/4-order-api/pkg/resp"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	Configs     *configs.Config
	ProductRepo *ProductRepository
}

type ProductHandlerDeps struct {
	Configs     *configs.Config
	ProductRepo *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) {
	h := ProductHandler{
		Configs:     deps.Configs,
		ProductRepo: deps.ProductRepo,
	}
	router.Handle("GET /product/{id}", middleware.IsAuthed(h.GetByID(), deps.Configs))
	router.Handle("POST /product", middleware.IsAuthed(h.Create(), deps.Configs))
	router.Handle("PATCH /product/{id}", middleware.IsAuthed(h.Update(), deps.Configs))
	router.Handle("DELETE /product/{id}", middleware.IsAuthed(h.Delete(), deps.Configs))

	router.Handle("GET /products", middleware.IsAuthed(h.GetAll(), deps.Configs))
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.ProductRepo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[ProductCreateRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product := NewProduct(payload.Name, payload.Description)
		createdProduct, err := h.ProductRepo.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp.JsonResp(w, createdProduct, http.StatusCreated)
	}
}

func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product *Product
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err = h.ProductRepo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.JsonResp(w, product, 200)
	}
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("User with phone number", r.Context().Value(middleware.Phone))
		products, err := h.ProductRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.JsonResp(w, products, 200)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[ProductUpdateRequest](w, r)
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedProduct, err := h.ProductRepo.Update(&Product{
			Model: gorm.Model{
				ID: uint(id),
			},
			Name:        payload.Name,
			Description: payload.Description,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.JsonResp(w, updatedProduct, http.StatusOK)
	}
}
