package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/menu/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Menu struct {
	Base  `json:",inline" valid:"-"`
	Name  *string `json:"name" gorm:"column:name;not null" valid:"required"`
	Items []*Item `json:"-" gorm:"ForeignKey:MenuID" valid:"-"`
}

func NewMenu(name *string) (*Menu, error) {
	e := Menu{
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

func (e *Menu) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
