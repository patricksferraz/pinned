package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/place/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Place struct {
	Base `json:",inline" valid:"-"`
	Name *string `json:"name" gorm:"column:name;not null" valid:"required"`
}

func NewPlace(name *string) (*Place, error) {
	e := Place{
		Name: name,
	}
	e.ID = utils.PString(uuid.NewV4().String())
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
