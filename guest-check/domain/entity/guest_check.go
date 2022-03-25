package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/guest-check/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type GuestCheck struct {
	Base           `json:",inline" valid:"-"`
	TotalPrice     *float64          `json:"total_price,omitempty" gorm:"column:total_price" valid:"-"`
	TotalDiscount  *float64          `json:"total_discount,omitempty" gorm:"column:total_discount" valid:"-"`
	FinalPrice     *float64          `json:"final_price,omitempty" gorm:"column:final_price" valid:"-"`
	Status         GuestCheckStatus  `json:"status" gorm:"column:status;not null" valid:"guestCheckStatus"`
	CanceledReason *string           `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Local          *string           `json:"local" gorm:"column:local;type:varchar(255)" valid:"required"`
	GuestID        *string           `json:"guest_id" gorm:"column:guest_id;type:uuid;not null" valid:"uuid"`
	Guest          *Guest            `json:"-" valid:"-"`
	PlaceID        *string           `json:"place_id" gorm:"column:place_id;type:uuid;not null" valid:"uuid"`
	Place          *Place            `json:"-" valid:"-"`
	AttendantBy    *string           `json:"attendant_by" gorm:"column:attendant_by;type:uuid" valid:"uuid,optional"`
	Attendant      *Attendant        `json:"-" valid:"-"`
	Items          []*GuestCheckItem `json:"-" gorm:"ForeignKey:GuestCheckID" valid:"-"`
	items          []*GuestCheckItem `json:"-" gorm:"-" valid:"-"`
}

func NewGuestCheck(local *string, guest *Guest, place *Place) (*GuestCheck, error) {
	e := GuestCheck{
		Status:  GUEST_CHECK_PENDING,
		Local:   local,
		GuestID: guest.ID,
		Guest:   guest,
		PlaceID: place.ID,
		Place:   place,
	}

	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *GuestCheck) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *GuestCheck) processPrice() error {
	var totalPrice float64
	var totalDiscount float64

	for _, i := range e.items {
		if i.FinalPrice != nil {
			totalPrice += *i.TotalPrice
		}
		if i.Discount != nil {
			totalDiscount += *i.Discount
		}
	}

	e.TotalPrice = &totalPrice
	e.TotalDiscount = &totalDiscount
	e.FinalPrice = utils.PFloat64(*e.TotalPrice - *e.TotalDiscount)
	err := e.IsValid()
	return err
}

func (e *GuestCheck) WaitPayment() error {
	if e.Status == GUEST_CHECK_AWAITING_PAYMENT {
		return errors.New("the guest check has already been awaiting payment")
	}

	e.Status = GUEST_CHECK_AWAITING_PAYMENT
	e.UpdatedAt = utils.PTime(time.Now())

	if err := e.IsValid(); err != nil {
		return err
	}

	return nil
}

func (e *GuestCheck) Cancel(canceledReason *string) error {
	if e.Status == GUEST_CHECK_CANCELED {
		return errors.New("the guest check has already been canceled")
	}

	if e.Status == GUEST_CHECK_PAID {
		return errors.New("the paid guest check cannot be canceled")
	}

	// TODO: adds the best way
	if len(e.items) > 0 || len(e.Items) > 0 {
		return errors.New("the guest check cannot be canceled")
	}

	e.Status = GUEST_CHECK_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheck) Pay() error {
	if e.Status == GUEST_CHECK_CANCELED {
		return errors.New("the canceled guest check cannot be paid")
	}

	e.Status = GUEST_CHECK_PAID
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheck) AddItem(guestCheckItem *GuestCheckItem) error {
	if e.Status != GUEST_CHECK_OPENED {
		return errors.New("the guest check cannot be changed")
	}

	// NOTE: change if multi guest check items are added
	e.items = append(e.Items, guestCheckItem)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.processPrice()
	return err
}

func (e *GuestCheck) Open(attendant *Attendant) error {
	e.Status = GUEST_CHECK_OPENED
	e.AttendantBy = attendant.ID
	e.Attendant = attendant
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}
