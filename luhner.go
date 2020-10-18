//Package luhner is an implementation of the Luhn mod N algorithm (https://wikipedia.com/wiki/Luhn_mod_N).
package luhner

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

//ErrInvalidCodePoint is returned when calculating the checksum if a character in the string is not present in the configured charset
var ErrInvalidCodePoint = errors.New("code string contains invalid code point")

//ErrInvalidCharset is returned when the charset is nil or a zero length slice
var ErrInvalidCharset = errors.New("invalid charset or bad config")

//Generate returns a Luhn mod N verifiable string from a Config created by the supplied options.
//If no options are defined, default values will be used.
func Generate(opts ...Option) (string, error) {
	c := NewDefaultConfig(opts...)
	return GenerateWithConfig(c)
}

//GenerateWithConfig returns a Luhn mod N verifiable string using the supplied config.
func GenerateWithConfig(c Config) (string, error) {
	clen := DefaultLength
	if c.Length() > 0 {
		clen = c.Length()
	}

	charset := c.Charset()
	if charset == nil || len(charset) == 0 {
		return "", ErrInvalidCharset
	}

	pre := c.Prefix("")
	length := clen - 1
	if pre != "" {
		length = length - len(pre)
	}

	r := randomString(length, charset)

	ctrl, err := control(r, 2, c)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{pre, r, ctrl}, ""), nil
}

//Validate returns true when the supplied string is verifiable against the Luhn mod N algorithm using a config built from the
//supplied options.
func Validate(s string, opts ...Option) bool {
	c := NewDefaultConfig(opts...)
	return ValidateWithConfig(s, c)
}

//ValidateWithConfig returns true when the supplied string is verified against the Luhn mod N algorithm using the supplied config.
func ValidateWithConfig(s string, c Config) bool {
	pre := c.Prefix(s)
	if pre != "" {
		s = strings.TrimLeft(s, pre)
	}

	chk, err := checksum(s, 1, c)
	if err != nil {
		return false
	}

	v := chk % c.Mod()

	return v == 0
}

func control(s string, f int, c Config) (string, error) {
	mod := c.Mod()
	if mod == 0 {
		return "", ErrInvalidCharset
	}

	cs, err := checksum(s, f, c)
	if err != nil {
		return "", err
	}

	rem := cs % mod
	ctrl := (mod - rem) % mod

	return c.Charset()[ctrl], nil
}

func checksum(s string, factor int, c Config) (int, error) {
	src := strings.Split(s, "")
	sum := 0
	mod := c.Mod()

	if mod == 0 {
		return -1, ErrInvalidCharset
	}

	for i := len(src) - 1; i > -1; i-- {
		n, ok := c.CodePoint(src[i])
		if !ok {
			return -1, nil
		}

		addend := factor * n
		if factor == 2 {
			factor = 1
		} else {
			factor = 2
		}

		addend = (addend / mod) + (addend % mod)
		sum += addend
	}

	return sum, nil
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func randomString(length int, charset []string) string {
	chars := strings.Join(charset, "")
	b := make([]byte, length)

	for i := range b {
		b[i] = chars[seededRand.Intn(len(charset))]
	}

	return string(b)
}
