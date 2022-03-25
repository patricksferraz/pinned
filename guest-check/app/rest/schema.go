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

type Guest struct {
	Base `json:",inline"`
}

type Place struct {
	Base `json:",inline"`
}

type CreateGuestCheckRequest struct {
	Local   string `json:"local"`
	GuestID string `json:"guest_id"`
	PlaceID string `json:"place_id"`
}

type GuestCheck struct {
	Base           `json:",inline"`
	TotalPrice     float64 `json:"total_price,omitempty"`
	TotalDiscount  float64 `json:"total_discount,omitempty"`
	FinalPrice     float64 `json:"final_price,omitempty"`
	Status         int     `json:"status"`
	CanceledReason string  `json:"canceled_reason,omitempty"`
	Local          string  `json:"local"`
	GuestID        string  `json:"guest_id"`
	PlaceID        string  `json:"place_id"`
}

type CancelGuestCheckRequest struct {
	CanceledReason string `json:"canceled_reason"`
}

type AddGuestCheckItemRequest struct {
	Name      string   `json:"name"`
	Code      int      `json:"code"`
	Quantity  int      `json:"quantity"`
	UnitPrice float64  `json:"unit_price"`
	Discount  float64  `json:"discount,omitempty"`
	Note      string   `json:"note,omitempty"`
	Tags      []string `json:"tags"`
}

type CancelGuestCheckItemRequest struct {
	CanceledReason string `json:"canceled_reason"`
}

type GuestCheckItem struct {
	Base           `json:",inline"`
	Status         int     `json:"status"`
	CanceledReason string  `json:"canceled_reason,omitempty"`
	Name           string  `json:"name"`
	Code           int     `json:"code"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	Discount       float64 `json:"discount,omitempty"`
	TotalPrice     float64 `json:"total_price,omitempty"`
	FinalPrice     float64 `json:"final_price"`
	Note           string  `json:"note,omitempty"`
	Tag            string  `json:"tag"`
	GuestCheckID   string  `json:"guest_check_id"`
}

type Attendant struct {
	Base `json:",inline"`
}
