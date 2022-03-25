package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/guest-check/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Place struct {
	Base `json:",inline" valid:"-"`
}

func NewPlace(id *string) (*Place, error) {
	e := Place{}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Place) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
