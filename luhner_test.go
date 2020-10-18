package luhner

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LuhnerSuite struct {
	suite.Suite
}

func TestRunLuhnerSuite(t *testing.T) {
	suite.Run(t, new(LuhnerSuite))
}

func (s *LuhnerSuite) TestNew() {
	code, err := Generate()
	s.Require().NoError(err)
	s.Assert().Len(code, 6)
	s.Assert().True(Validate(code))
}

func (s *LuhnerSuite) TestPrintValues() {
	c := NewDefaultConfig(Length(6), Prefix("XYZ"))
	code, err := GenerateWithConfig(c)
	s.Require().NoError(err)
	fmt.Println(toID(code))

	valid := ValidateWithConfig(code, c)
	s.Assert().True(valid)
}

func (s *LuhnerSuite) TestNewOptions() {
	table := []struct {
		opts []Option
	}{
		{
			opts: []Option{
				Length(12),
			},
		},
		{
			opts: []Option{
				Length(9),
				Prefix("abc"),
			},
		},
	}

	for _, tt := range table {
		code, err := Generate(tt.opts...)
		s.Require().NoError(err)
		s.Assert().True(Validate(code, tt.opts...))
	}
}

func (s *LuhnerSuite) TestNewConfig() {
	table := []struct {
		config Config
	}{
		{
			config: DefaultConfig{Len: 10, Chars: DefaultCharset},
		},
		{config: DefaultConfig{Pre: "XYZ", Chars: DefaultCharset}},
	}

	for _, tt := range table {
		code, err := GenerateWithConfig(tt.config)
		s.Require().NoError(err)
		s.Assert().True(ValidateWithConfig(code, tt.config))
	}
}

func (s *LuhnerSuite) TestValidate() {
	table := []struct {
		code  string
		valid bool
	}{
		{
			code:  "43371C2322",
			valid: false,
		},
	}

	for _, tt := range table {
		s.Assert().Equal(tt.valid, Validate(tt.code))
	}
}

func toID(s string) string {
	var sp []string

	switch len(s) {
	case 6:
		sp = splitGroup(s, 3)
	case 8, 12:
		sp = splitGroup(s, 4)
	default:
		return s
	}

	return strings.Join(sp, "-")
}

func splitGroup(s string, size int) []string {
	chunks := len(s) / size
	res := make([]string, chunks)

	idx := 0
	for i := 0; i < len(s); i += size {
		start := i
		end := minInt(len(s), i+size)
		res[idx] = s[start:end]
		idx++
	}

	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
