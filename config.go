package luhner

//DefaultLength will be used by the default Config and methods if no alternate value is supplied
var DefaultLength = 6

//DefaultCharset will be used by the default Config and methods if no alternate charset is supplied
// The default set is Base16 characters 0 through F
var DefaultCharset = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

//Config interface defines the methods used by the generator and validator fuctions
type Config interface {
	Length() int
	Mod() int
	Charset() []string
	CodePoint(s string) (int, bool)
	Prefix(string) string
}

//DefaultConfig is the basic implementation of the Config interface
type DefaultConfig struct {
	Chars []string
	Pre   string
	Len   int
}

//Length returns the total string length of the generated string inclusive of the prefix and control character
func (c DefaultConfig) Length() int {
	if c.Len < 1 {
		return DefaultLength
	}

	return c.Len
}

//Mod returns the modulous (artihmetic base) which is typically the length of the Charset
func (c DefaultConfig) Mod() int {
	return len(c.Charset())
}

//Charset returns a slice of strings representing the valid characters used to generate IDs.
//The code value of each character is equal to it's numbered position in the slice.
func (c DefaultConfig) Charset() []string {
	if c.Chars == nil {
		return DefaultCharset
	}

	return c.Chars
}

//CodePoint checks for the index value of a string in the charset and returns that value and a bool
//to indicate its existence in the set. If a character is not present in the charset, values of -1 and false will be returned.
func (c DefaultConfig) CodePoint(s string) (int, bool) {
	for i, v := range c.Charset() {
		if s == v {
			return i, true
		}
	}

	return -1, false
}

//Prefix returns an optional prefix to be added to generated strings
func (c DefaultConfig) Prefix(string) string {
	return c.Pre
}

//NewDefaultConfig returns a Config backed by the default implementation and values
func NewDefaultConfig(opts ...Option) Config {
	c := defaultConfig()

	for _, o := range opts {
		o(&c)
	}

	return c
}

func defaultConfig() DefaultConfig {
	return DefaultConfig{
		Chars: DefaultCharset,
		Len:   DefaultLength,
	}
}

//Option updates values in DefaultConfig
type Option func(*DefaultConfig)

//Charset sets the DefaultConfig charset to the supplied slice
func Charset(s []string) Option {
	return func(c *DefaultConfig) {
		c.Chars = s
	}
}

//Length sets the DefaultConfig length to the supplied value
func Length(l int) Option {
	return func(c *DefaultConfig) {
		c.Len = l
	}
}

//Prefix sets the Defaultconfig prefix to the supplied value
func Prefix(p string) Option {
	return func(c *DefaultConfig) {
		c.Pre = p
	}
}
