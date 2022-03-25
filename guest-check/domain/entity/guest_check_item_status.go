package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["guestCheckItemStatus"] = govalidator.Validator(func(str string) bool {
		res := str == GUEST_CHECK_ITEM_PENDING.String()
		res = res || str == GUEST_CHECK_ITEM_CANCELED.String()
		res = res || str == GUEST_CHECK_ITEM_PREPARING.String()
		res = res || str == GUEST_CHECK_ITEM_READY.String()
		res = res || str == GUEST_CHECK_ITEM_ON_THE_WAY.String()
		res = res || str == GUEST_CHECK_ITEM_DELIVERED.String()
		return res
	})
}

type GuestCheckItemStatus int

const (
	GUEST_CHECK_ITEM_PENDING GuestCheckItemStatus = iota + 1
	GUEST_CHECK_ITEM_CANCELED
	GUEST_CHECK_ITEM_PREPARING
	GUEST_CHECK_ITEM_READY
	GUEST_CHECK_ITEM_ON_THE_WAY
	GUEST_CHECK_ITEM_DELIVERED
)

func (t GuestCheckItemStatus) String() string {
	switch t {
	case GUEST_CHECK_ITEM_PENDING:
		return "PENDING"
	case GUEST_CHECK_ITEM_CANCELED:
		return "CANCELED"
	case GUEST_CHECK_ITEM_PREPARING:
		return "PREPARING"
	case GUEST_CHECK_ITEM_READY:
		return "READY"
	case GUEST_CHECK_ITEM_ON_THE_WAY:
		return "ON_THE_WAY"
	case GUEST_CHECK_ITEM_DELIVERED:
		return "DELIVERED"
	}
	return ""
}
