package contact

import "github.com/asaskevich/govalidator"

type RequestEmailMessage struct {
	From string `json:"from" valid:"required~El campo From es requerido,email~El formato del email no es válido"`
}

func (m *RequestEmailMessage) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
