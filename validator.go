// Package robo-validator is package of validators and sanitizers for strings, structs and collections.
package robovalidator

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

//tag used for setting validation criteria's
const tagName = "validate"

type DefaultValidator struct {}
type RequiredValidator struct {}

//Validator Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) (bool, error)
}

// Default Validator does not perform any validations.
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

// Required Validator
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
	case reflect.Float32:
		if val.(float32) <= 0 {
			return false, fmt.Errorf("should be greater than %v", 0)
		}
	case reflect.Float64:
		if val.(float64) <= 0 {
			return false, fmt.Errorf("should be greater than %v", 0)
		}
	case reflect.String:
		if len(val.(string)) == 0 {
			return false, fmt.Errorf("string required")
		}
	case reflect.Struct:
		if val.(time.Time).String() == "0001-01-01 00:00:00.000000" {
			return false, fmt.Errorf("datetime required")
		}
	}

	return true, nil
}

// Number Validator
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)
	//check non zero value
	if num <= 0 {
		return false, fmt.Errorf("should be greater than %v", 0)
	}

	return true, nil
}

//Validate - validate the struct , later on will add more data types
func Validate(s interface{}, eventName string) error {
	//check it's struct
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag, eventName)
		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Return error as result
		if !valid && err != nil {
			return fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error())
		}
	}

	return nil
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
