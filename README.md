# pisvalidator
Brazilian pis validator (PIS) - validation package in Golang.

It is an essential package to validate PIS number in your application.


## Information about PIS

- [Caixa - PIS](http://www.caixa.gov.br/beneficios-trabalhador/pis/Paginas/default.aspx)

## Installation
Use the `go tool` for do that:
```bash
$ go get github.com/Cegielkowski/pisvalidator
```

## Usage

You need to import the *package* `github.com/Cegielkowski/pisvalidator` like that:

```go
import (
    "github.com/Cegielkowski/pisvalidator"
)
```

## Example

```go
package main

import (
	"fmt"
	"github.com/Cegielkowski/pisvalidator"
)

func main() {

	// This will validate the PIS and clean the formatting.
	pis, err := pisvalidator.ValidatePis("477.11617.27-5")

	// Verifies if it is a valid PIS
	if err != nil {
		panic(fmt.Errorf("It isn't valid: %v", err))
	}

	// Formatted output
	fmt.Println(pis)
	// Output: 47711617275
}
```

## License
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](LICENSE)