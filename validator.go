// Package robovalidator is package of validators and sanitizers for strings, structs and collections.
package robovalidator

import (
	"fmt"
	"reflect"
	"strings"
)

//tag used for setting valisation criterias
const tagName = "validate"

//Validator Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) (bool, error)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

type RequiredValidator struct {
}

//Validate default validator
func (v DefaultValidator) Validate(val interface{}) (bool, error) {

	//validate the value
	//string , float , int , time for now

	return true, nil
}

//Validate default validator
func (v RequiredValidator) Validate(val interface{}) (bool, error) {

	//validate the value
	//string , float , int , time for now

	typeOfVal := reflect.TypeOf(val).Kind()

	switch typeOfVal {

	case reflect.Int:
		if val.(int) <= 0 {
			return false, fmt.Errorf("should be greater than %v", 0)
		}
	case reflect.Int64:
		if val.(int) <= 0 {
			return false, fmt.Errorf("should be greater than %v", 0)
		}
	case reflect.String:
		if len(val.(string)) == 0 {
			return false, fmt.Errorf("string required")
		}
	}

	return true, nil
}

//Validate - validate the struct , later on will add more data types
func Validate(s interface{}, eventName string) []error {

	errs := []error{}

	//check its structs
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {

		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag, eventName)

		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string, eventName string) Validator {

	args := strings.Split(tag, ",")

	for _, arg := range args {

		condition := strings.Split(arg, ":")

		//check for event condition
		if condition[0] != eventName {
			//log.Println("not applicable here")
			continue
		}

		//check for value condition
		switch condition[1] {
		case "number":
			validator := NumberValidator{}
			return validator

		case "required":
			validator := RequiredValidator{}
			return validator
		}

	}

	//log.Println(args)

	return DefaultValidator{}
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {

	num := val.(int)

	//check non zero value

	if num <= 0 {
		return false, fmt.Errorf("should be greater than %v", 0)
	}

	return true, nil
}
