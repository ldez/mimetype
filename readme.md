# mimetype

This package provides mimetypes as constants.

The constants are generated for the IONA website: https://www.iana.org/assignments/media-types/media-types.xhtml

Detailed documentation: https://pkg.go.dev/github.com/ldez/mimetype

```go
package main

import (
	"fmt"

	"github.com/ldez/mimetype"
)

func main() {
	fmt.Println(mimetype.AudioAac)
	fmt.Println(mimetype.ImageBmp)

	fmt.Println(mimetype.IsApplication("foobar"))
	fmt.Println(mimetype.IsAudio("foobar"))
	fmt.Println(mimetype.IsFont("foobar"))
	fmt.Println(mimetype.IsImage("foobar"))
	fmt.Println(mimetype.IsMessage("foobar"))
	fmt.Println(mimetype.IsModel("foobar"))
	fmt.Println(mimetype.IsMultipart("foobar"))
	fmt.Println(mimetype.IsText("foobar"))
	fmt.Println(mimetype.IsVideo("foobar"))
}

```

