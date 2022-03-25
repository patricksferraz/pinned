package event

type createGuestMsg struct {
	ID *string `json:"id" valid:"uuid"`
}

type Guest struct {
	Event `json:",inline" valid:"required"`
	Msg   *createGuestMsg `json:"msg" valid:"required"`
}
