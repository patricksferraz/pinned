package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/lib/pq"
	"github.com/patricksferraz/guest-check/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type GuestCheckItem struct {
	Base           `json:",inline" valid:"-"`
	Status         GuestCheckItemStatus `json:"status" gorm:"column:status;not null" valid:"guestCheckItemStatus"`
	CanceledReason *string              `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Name           *string              `json:"name" gorm:"column:name;not null" valid:"required"`
	Code           *int                 `json:"code" gorm:"column:code;not null" valid:"required"`
	Quantity       *int                 `json:"quantity" gorm:"column:quantity;not null" valid:"required"`
	UnitPrice      *float64             `json:"unit_price" gorm:"column:unit_price;not null" valid:"required"`
	Discount       *float64             `json:"discount,omitempty" gorm:"column:discount" valid:"-"`
	TotalPrice     *float64             `json:"total_price,omitempty" gorm:"column:total_price" valid:"-"`
	FinalPrice     *float64             `json:"final_price" gorm:"column:final_price;not null" valid:"-"`
	Note           *string              `json:"note,omitempty" gorm:"column:note;type:varchar(255)" valid:"-"`
	Tags           pq.StringArray       `json:"tags" gorm:"column:tags;type:text[]" valid:"-"`
	GuestCheckID   *string              `json:"guest_check_id" gorm:"column:guest_check_id;type:uuid;not null" valid:"uuid"`
	GuestCheck     *GuestCheck          `json:"-" valid:"-"`
}

func NewGuestCheckItem(name *string, code, quantity *int, unitPrice *float64, discount *float64, note *string, tag *[]string, guestCheck *GuestCheck) (*GuestCheckItem, error) {
	e := GuestCheckItem{
		Name:         name,
		Code:         code,
		Status:       GUEST_CHECK_ITEM_PENDING,
		Quantity:     quantity,
		UnitPrice:    unitPrice,
		Discount:     discount,
		Note:         note,
		Tags:         *tag,
		GuestCheckID: guestCheck.ID,
		GuestCheck:   guestCheck,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())
	e.processPrice()

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *GuestCheckItem) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *GuestCheckItem) processPrice() error {
	e.TotalPrice = utils.PFloat64(*e.UnitPrice * float64(*e.Quantity))
	e.FinalPrice = e.TotalPrice
	if e.Discount != nil {
		e.FinalPrice = utils.PFloat64(*e.FinalPrice - *e.Discount)
	}
	err := e.IsValid()
	return err
}

func (e *GuestCheckItem) Cancel(canceledReason *string) error {
	e.Status = GUEST_CHECK_ITEM_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheckItem) Prepare() error {
	e.Status = GUEST_CHECK_ITEM_PREPARING
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheckItem) Ready() error {
	e.Status = GUEST_CHECK_ITEM_READY
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheckItem) Forward() error {
	e.Status = GUEST_CHECK_ITEM_ON_THE_WAY
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheckItem) Deliver() error {
	e.Status = GUEST_CHECK_ITEM_DELIVERED
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}
