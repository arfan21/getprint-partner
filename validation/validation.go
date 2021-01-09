package validation

import (
	"github.com/arfan21/getprint-partner/models"
	"github.com/go-ozzo/ozzo-validation/is"
	validator "github.com/go-ozzo/ozzo-validation/v4"
)

func Validate(p models.Partner) error {
	return validator.Errors{
		"user_id":      validator.Validate(p.UserID, validator.Required),
		"name":         validator.Validate(p.Name, validator.Required),
		"email":        validator.Validate(p.Email, validator.Required, is.Email),
		"phone_number": validator.Validate(p.PhoneNumber, validator.Required),
		"picture":      validator.Validate(p.Picture, validator.Required),
		"address":      validator.Validate(p.Address.Address, validator.Required),
		"lat":          validator.Validate(p.Address.Lat, validator.Required),
		"lng":          validator.Validate(p.Address.Lng, validator.Required),
	}.Filter()
}
