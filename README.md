# Go Screenshot

another library for taking screenshot on your Linux+,Windows-,Mac-

there is many feature that still missing (like choosing display, set size, etc), and currently only support for linux

## Installation

```
go get -u github.com/codenoid/go-screenshot
```

## Usage

```
package main

import (
	"image/jpeg"
	"os"

	screenshot "github.com/codenoid/go-screenshot"
)

func main() {
	xgbconn, err := screenshot.NewSession()
	if err != nil {
		return
	}
	defer xgbconn.Close()

	ss := screenshot.Start{
		XgbConn: xgbconn,
	}

	img, err := ss.CaptureScreen()

	f, err := os.Create("screenshot.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	jpeg.Encode(f, img, nil)
}
```