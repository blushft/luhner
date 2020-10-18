package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blushft/luhner"
	"github.com/urfave/cli/v2"
)

var groups int
var grouplen int
var sep string

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "groups",
				Value:       3,
				Destination: &groups,
			},
			&cli.IntFlag{
				Name:        "len",
				Value:       4,
				Destination: &grouplen,
			},
			&cli.StringFlag{
				Name:        "sep",
				Value:       "-",
				Destination: &sep,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "generate",
				Action: cmdGenerate,
			},
			{
				Name:   "validate",
				Action: cmdValidate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func cmdGenerate(c *cli.Context) error {
	gen := newKeygen()

	parts := make([]string, gen.c.Groups)
	for i := 0; i < gen.c.Groups; i++ {
		gen.c.prefixFunc = func(s string) string {
			return selectPrefix(i)
		}

		p, err := gen.i.Generate()
		if err != nil {
			return err
		}

		parts[i] = p
	}

	fmt.Println(strings.Join(parts, gen.c.Separator))
	return nil
}

func cmdValidate(c *cli.Context) error {
	gen := newKeygen()
	key := c.Args().First()
	if key == "" {
		return errors.New("no key supplied")
	}

	parts := strings.Split(key, sep)
	for i, p := range parts {
		gen.c.prefixFunc = func(s string) string {
			return selectPrefix(i)
		}

		if !gen.i.Validate(p) {
			return errors.New("invalid key")
		}
	}

	fmt.Println("valid key")
	return nil
}

type keygen struct {
	c *keygenConfig
	i luhner.Instance
}

func newKeygen() keygen {
	kc := &keygenConfig{
		Groups:    groups,
		GroupLen:  grouplen,
		Separator: sep,
		Chars:     charset(),
	}

	inst := luhner.NewInstanceWithConfig(kc)

	return keygen{
		c: kc,
		i: inst,
	}
}

func charset() []string {
	return strings.Split("0123456798EFHIPXZL", "")
}

func selectPrefix(i int) string {
	pres := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	if i > len(pres) {
		return "X"
	}

	return pres[i]
}

type keygenConfig struct {
	Chars     []string
	Groups    int
	GroupLen  int
	Separator string

	prefixFunc func(string) string
}

func (kc *keygenConfig) Length() int {
	return kc.GroupLen
}

func (kc *keygenConfig) Charset() []string {
	return kc.Chars
}

func (kc *keygenConfig) Mod() int {
	return len(kc.Chars)
}

func (kc *keygenConfig) CodePoint(s string) (int, bool) {
	for i, v := range kc.Chars {
		if s == v {
			return i, true
		}
	}

	return -1, false
}

func (kc *keygenConfig) Prefix(v string) string {
	return kc.prefixFunc(v)
}
