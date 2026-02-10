package products

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
