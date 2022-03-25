package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Event struct {
	ID *string `json:"id" valid:"uuid"`
}

func (e *Event) IsValid(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}

func (e *Event) ParseJson(data []byte, i interface{}) error {
	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	err = e.IsValid(i)
	if err != nil {
		return err
	}

	return nil
}
