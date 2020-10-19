<!-- Shelds -->
[![Go Report Card][go-reportcard-sheild]][go-reportcard-url]
![Go][go-status-url]
[![go.dev reference][godoc-shield]][godoc-url]
[![MIT License][license-shield]][license-url]

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]

<br/>
<p align="center">
    <a href="https://github.com/blushft/Luhmer">
     <h2 align="center">Luhner</h2>
    </a> 
    <p align="center">
        Generate and verify identifiers algorithmically
        </br>
    </p>
</p>

## About

Luhner (pronouced like lunar) is a go implementation of the [Luhn mod N algorithm](https://wikipedia.com/wiki/Luhn_mod_N_algorithm).

### Important Notice

The Luhn algorithm is useful to validate identifiers but is _not_ a secure cryptographic hash and **should not** be used to protect privieleged access or data.

## Usage

#### Basic

The simplest usage of luhner are the built in Generate() and Validate() functions:

```golang
id, err := luhner.Generate()
if err != nil {
    log.Fatal(err)
}

valid := luhner.Validate(id)
log.Printf("ID valid = %t", valid)
```

The default functions can be adjusted using package options to set the Length, Prefix, or Charset. Any options passed to Generate() must also be passed to Validate():

```golang
id, err := luhner.Generate(luhner.Length(10))
if err != nil {
    log.Fatal(err)
}

valid := luhner.Validate(id, luhner.Length(10))
```

#### Instance

An Instance contains a generator and validator using a common config:

```golang
inst := luhner.NewInstance(luhner.Length(8), luhner.Prefix("XYZ"))

id, err := inst.Generate()
if err != nil {
    log.Fatal(err)
}

valid := inst.Validate(id)
```

#### Config Interface

Advanced use cases can leverage the Config interface to provide custom implementations:

```golang
type Config interface {
    Length() int
    Mod() int
    Charset() []string
    CodePoint(s string) (int, bool)
    Prefix(string) string
}
```

An example implementation for a custom keygen is supplied in the [example](example/keygen.go) folder.

[go-reportcard-sheild]: https://goreportcard.com/badge/github.com/blushft/luhner
[go-reportcard-url]: https://goreportcard.com/report/github.com/blushft/luhner
[go-status-url]: https://github.com/blushft/luhner/workflows/Go/badge.svg
[godoc-shield]: https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square
[godoc-url]: https://pkg.go.dev/github.com/blushft/luhner
[license-shield]: https://img.shields.io/github/license/blushft/luhner.svg?style=flat-square
[license-url]: https://github.com/blushft/luhner/blob/master/LICENSE
[contributors-shield]: https://img.shields.io/github/contributors/blushft/luhner.svg?style=flat-square
[contributors-url]: https://github.com/blushft/luhner/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/blushft/luhner.svg?style=flat-square
[forks-url]: https://github.com/blushft/luhner/network/members
