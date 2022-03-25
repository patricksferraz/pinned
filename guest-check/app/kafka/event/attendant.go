package event

type createAttendantMsg struct {
	ID *string `json:"id" valid:"uuid"`
}

type Attendant struct {
	Event `json:",inline" valid:"required"`
	Msg   *createAttendantMsg `json:"msg" valid:"required"`
}
