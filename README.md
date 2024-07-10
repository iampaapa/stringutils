This is basically to help me understand the right way to publish a module to `pkg.go.dev` so that I can do the same for another module I am writing. 

# stringutils

[![Go Reference](https://pkg.go.dev/badge/github.com/iampaapa/stringutils.svg)](https://pkg.go.dev/github.com/iampaapa/stringutils)

`stringutils` is a simple Go package providing utility functions for basic string manipulation.

## Installation

To install the package, run:

```sh
go get github.com/iampaapa/stringutils
```

## Usage

Import the package in your Go code:

```go
import "github.com/iampaapa/stringutils"
```

### Functions

#### Reverse

Reverses the input string.

```go
func Reverse(s string) string
```

Example:

```go
package main

import (
    "fmt"
    "github.com/iampaapa/stringutils"
)

func main() {
    fmt.Println(stringutils.Reverse("hello")) // Output: "olleh"
}
```

#### ToUpper

Converts the input string to uppercase.

```go
func ToUpper(s string) string
```

Example:

```go
package main

import (
    "fmt"
    "github.com/iampaapa/stringutils"
)

func main() {
    fmt.Println(stringutils.ToUpper("hello")) // Output: "HELLO"
}
```

#### ToLower

Converts the input string to lowercase.

```go
func ToLower(s string) string
```

Example:

```go
package main

import (
    "fmt"
    "github.com/iampaapa/stringutils"
)

func main() {
    fmt.Println(stringutils.ToLower("HELLO")) // Output: "hello"
}
```

## Testing

To run tests for the package, execute:

```sh
go test
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the BSD 3-Clause License. See the [LICENSE](LICENSE) file for details.
