voice
=====

Go package that is just calling the OS X say command, tweet sized.

## Installation

```bash
$ go get github.com/peterhellberg/voice
```

## Usage example

```go
package main

import . "github.com/peterhellberg/voice"

func main() {
	Say("Hello", 42, 3.14)
}
```
