package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/guest-check/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Guest struct {
	Base `json:",inline" valid:"-"`
}

func NewGuest(id *string) (*Guest, error) {
	e := Guest{}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Guest) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
