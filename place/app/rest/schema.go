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

type CreatePlaceRequest struct {
	Name string `json:"name"`
}

type Place struct {
	Base `json:",inline"`
	Name string `json:"name"`
}
