package event

type openGuestCheckMsg struct {
	GuestCheckID *string `json:"guest_check_id" valid:"uuid"`
	AttendantID  *string `json:"attendant_id" valid:"uuid"`
}

type OpenGuestCheck struct {
	Event `json:",inline" valid:"required"`
	Msg   *openGuestCheckMsg `json:"msg" valid:"required"`
}
