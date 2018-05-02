package robovalidator

type Config struct {
	EventName string
}

// NumberValidator performs numerical value validation.
// Its limited to int type for simplicity.
type NumberValidator struct {
	Min int
	Max int
}
