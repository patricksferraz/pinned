package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["guestCheckStatus"] = govalidator.Validator(func(str string) bool {
		res := str == GUEST_CHECK_PENDING.String()
		res = res || str == GUEST_CHECK_OPENED.String()
		res = res || str == GUEST_CHECK_CANCELED.String()
		res = res || str == GUEST_CHECK_AWAITING_PAYMENT.String()
		res = res || str == GUEST_CHECK_PAID.String()
		return res
	})
}

type GuestCheckStatus int

const (
	GUEST_CHECK_PENDING GuestCheckStatus = iota + 1
	GUEST_CHECK_OPENED
	GUEST_CHECK_CANCELED
	GUEST_CHECK_AWAITING_PAYMENT
	GUEST_CHECK_PAID
)

func (t GuestCheckStatus) String() string {
	switch t {
	case GUEST_CHECK_PENDING:
		return "PENDING"
	case GUEST_CHECK_OPENED:
		return "OPENED"
	case GUEST_CHECK_CANCELED:
		return "CANCELED"
	case GUEST_CHECK_AWAITING_PAYMENT:
		return "AWAITING_PAYMENT"
	case GUEST_CHECK_PAID:
		return "PAID"
	}
	return ""
}
