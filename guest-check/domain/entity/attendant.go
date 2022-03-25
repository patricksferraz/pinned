package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/guest-check/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Attendant struct {
	Base        `json:",inline" valid:"-"`
	GuestChecks []*GuestCheck `json:"guest_checks" gorm:"ForeignKey:AttendantBy" valid:"-"`
}

func NewAttendant(id *string) (*Attendant, error) {
	e := Attendant{}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Attendant) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
