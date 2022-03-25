package entity

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/guest/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Message interface {
	IsValid() error
}

type Event struct {
	Base `json:",inline" valid:"required"`
	Msg  Message `json:"msg" valid:"required"`
}

func NewEvent(msg Message) (*Event, error) {
	e := Event{
		Msg: msg,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	if err := msg.IsValid(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Event) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Event) ToJson() ([]byte, error) {
	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(e)
	if err != nil {
		return nil, nil
	}

	return result, nil
}
