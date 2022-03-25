package event

type createPlaceMsg struct {
	ID *string `json:"id" valid:"uuid"`
}

type Place struct {
	Event `json:",inline" valid:"required"`
	Msg   *createPlaceMsg `json:"msg" valid:"required"`
}
