package main

import (
	"log"

	validator "bitbucket.org/rahulvramesh/robo-validator"
)

type User struct {
	Id    int    `validate:"create_booking:number,test_booking:number"`
	Name  string `validate:"-"`
	Bio   string `validate:"-"`
	Email string `validate:"-"`
}

func main() {

	user := User{
		Id:    0,
		Name:  "superlongstring",
		Bio:   "",
		Email: "foobar",
	}

	err := validator.Validate(user, "test_booking")

	log.Println(len(err))

	for _, er := range err {
		log.Println(er.Error())
	}
}
