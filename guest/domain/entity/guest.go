package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/paemuri/brdoc"
	"github.com/patricksferraz/guest/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.TagMap["doc"] = govalidator.Validator(func(str string) bool {
		return brdoc.IsCNPJ(str) || brdoc.IsCPF(str)
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type Guest struct {
	Base `json:",inline" valid:"-"`
	Name *string `json:"name" gorm:"column:name;not null" valid:"required"`
	Doc  *string `json:"doc" gorm:"column:doc;type:varchar(15)" valid:"doc,optional"`
}

func NewGuest(name, doc *string) (*Guest, error) {
	e := Guest{
		Name: name,
		Doc:  doc,
	}
	e.ID = utils.PString(uuid.NewV4().String())
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
