package luhner

//Generator interface defines the method signature to generate valid Luhm mod N strings
type Generator interface {
	Generate() (string, error)
}

//Validator interface defines the method to validate generated strings
type Validator interface {
	Validate(string) bool
}

//Instance interface defines a generator/validator pair
type Instance interface {
	Generator
	Validator
}

//NewInstance returns the default Instance implementation using DefaultConfig configured by Options
func NewInstance(opts ...Option) Instance {
	c := NewDefaultConfig(opts...)
	return NewInstanceWithConfig(c)
}

//NewInstanceWithConfig returns the default Instance implementation using the passed Config interface
func NewInstanceWithConfig(c Config) Instance {
	return &instance{
		c: c,
	}
}

type instance struct {
	c Config
}

func (i *instance) Generate() (string, error) {
	return GenerateWithConfig(i.c)
}

func (i *instance) Validate(s string) bool {
	return ValidateWithConfig(s, i.c)
}
