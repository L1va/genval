//This file was automatically generated by the genval generator v1.2
//Please don't modify it manually. Edit your entity tags and then
//run go generate

package overriding

import (
	"github.com/gojuno/genval/errlist"
)

type validatable interface {
	Validate() error
}

func validate(i interface{}) error {
	if v, ok := i.(validatable); ok {
		return v.Validate()
	}
	return nil
}

// Validate validates Age1
func (r Age1) Validate() error {
	var errs errlist.ErrList
	if r.Value < 3 {
		errs.AddFieldErrf("Value", "less than 3")
	}
	if r.Value > 64 {
		errs.AddFieldErrf("Value", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Age2
func (r Age2) Validate() error {
	var errs errlist.ErrList
	if r.Value < 3 {
		errs.AddFieldErrf("Value", "less than 3")
	}
	if r.Value > 64 {
		errs.AddFieldErrf("Value", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Age3
func (r Age3) Validate() error {
	var errs errlist.ErrList
	if r.Value < 3 {
		errs.AddFieldErrf("Value", "less than 3")
	}
	if r.Value > 64 {
		errs.AddFieldErrf("Value", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Age4
func (r Age4) Validate() error {
	var errs errlist.ErrList
	if r.Value < 3 {
		errs.AddFieldErrf("Value", "less than 3")
	}
	if r.Value > 64 {
		errs.AddFieldErrf("Value", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Age5
func (r Age5) Validate() error {
	var errs errlist.ErrList
	if r.Value < 3 {
		errs.AddFieldErrf("Value", "less than 3")
	}
	if r.Value > 64 {
		errs.AddFieldErrf("Value", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Request2
func (r Request2) Validate() error {
	var errs errlist.ErrList
	if err := r.Age.ValidateMin10(); err != nil {
		errs.AddFieldErr("Age", err)
	}
	if r.Some < 3 {
		errs.AddFieldErrf("Some", "less than 3")
	}
	if r.Some > 64 {
		errs.AddFieldErrf("Some", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Request3
func (r Request3) Validate() error {
	var errs errlist.ErrList
	if err := validateMin10(r.Age); err != nil {
		errs.AddFieldErr("Age", err)
	}
	if r.Some < 3 {
		errs.AddFieldErrf("Some", "less than 3")
	}
	if r.Some > 64 {
		errs.AddFieldErrf("Some", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Request4
func (r Request4) Validate() error {
	var errs errlist.ErrList
	if err := r.Age.ValidateMin10(); err != nil {
		errs.AddFieldErr("Age", err)
	}
	if err := validateMax128(r.Age); err != nil {
		errs.AddFieldErr("Age", err)
	}
	if r.Some < 3 {
		errs.AddFieldErrf("Some", "less than 3")
	}
	if r.Some > 64 {
		errs.AddFieldErrf("Some", "more than 64")
	}
	return errs.ErrorOrNil()
}

// Validate validates Request5
func (r Request5) Validate() error {
	var errs errlist.ErrList
	if err := r.Age.Validate(); err != nil {
		errs.AddFieldErr("Age", err)
	}
	if r.Some < 3 {
		errs.AddFieldErrf("Some", "less than 3")
	}
	if r.Some > 64 {
		errs.AddFieldErrf("Some", "more than 64")
	}
	errs.Add(r.validate())
	return errs.ErrorOrNil()
}
