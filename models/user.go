package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// User A user is a user of 3DSIM products
// swagger:model user
type User struct {

	// blocked
	Blocked bool `json:"blocked,omitempty"`

	// email
	// Required: true
	Email *strfmt.Email `json:"email"`

	// email verified
	EmailVerified bool `json:"emailVerified,omitempty"`

	// first name
	// Required: true
	FirstName *string `json:"firstName"`

	// id
	ID string `json:"id,omitempty"`

	// last IP
	LastIP string `json:"lastIP,omitempty"`

	// last login
	LastLogin strfmt.DateTime `json:"lastLogin,omitempty"`

	// last name
	// Required: true
	LastName *string `json:"lastName"`

	// logins count
	LoginsCount int64 `json:"loginsCount,omitempty"`

	// password
	// Required: true
	Password *strfmt.Password `json:"password"`

	// phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// phone verified
	PhoneVerified bool `json:"phoneVerified,omitempty"`

	// roles
	Roles []string `json:"roles"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *User) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *User) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var userRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Admin","User"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRolesItemsEnum = append(userRolesItemsEnum, v)
	}
}

func (m *User) validateRolesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userRolesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *User) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {

		// value enum
		if err := m.validateRolesItemsEnum("roles"+"."+strconv.Itoa(i), "body", m.Roles[i]); err != nil {
			return err
		}

	}

	return nil
}