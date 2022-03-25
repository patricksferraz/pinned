package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type CreateMenuRequest struct {
	Name string `json:"name"`
}

type Menu struct {
	Base              `json:",inline"`
	CreateMenuRequest `json:",inline"`
}

type CreateItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
}

type Item struct {
	Base              `json:",inline"`
	CreateItemRequest `json:",inline"`
}

type AddItemTagRequest struct {
	Tags []string `json:"tags"`
}
