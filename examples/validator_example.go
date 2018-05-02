package main

import (
	"log"

	validator "bitbucket.org/rahulvramesh/robo-validator"
)

type User struct {
	Id    int    `validate:"create_booking:required,test_booking:required"`
	Name  int    `validate:"create_booking:number,test_booking:number"`
	Bio   string `validate:"test_booking:required"`
	Email string `validate:"-"`
}

func main() {

	user := User{
		Id:    0,
		Name:  0,
		Bio:   "ee",
		Email: "foobar",
	}

	err := validator.Validate(user, "test_booking")

	log.Println(len(err))

	for _, er := range err {
		log.Println(er.Error())
	}
}
