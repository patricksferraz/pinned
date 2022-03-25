package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/lib/pq"
	"github.com/patricksferraz/menu/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Item struct {
	Base        `json:",inline" valid:"-"`
	Code        *int           `json:"code" gorm:"column:code;auto_increment;not_null" valid:"-"`
	Name        *string        `json:"name" gorm:"column:name;not null" valid:"required"`
	Description *string        `json:"description,omitempty" gorm:"column:description;type:varchar(500)" valid:"-"`
	Price       *float64       `json:"price" gorm:"column:price;not null" valid:"required"`
	Discount    *float64       `json:"discount,omitempty" gorm:"column:discount" valid:"-"`
	Tags        pq.StringArray `json:"tags,omitempty" gorm:"column:tags;type:text[]" valid:"-"`
	MenuID      *string        `json:"menu_id" gorm:"column:menu_id;type:uuid;not null" valid:"uuid"`
	Menu        *Menu          `json:"-" valid:"-"`
}

func NewItem(name, description *string, price, discount *float64, menu *Menu) (*Item, error) {
	e := Item{
		Name:        name,
		Price:       price,
		Discount:    discount,
		Description: description,
		MenuID:      menu.ID,
		Menu:        menu,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Item) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Item) AddTags(tag *[]string) error {
	e.Tags = *tag
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}
