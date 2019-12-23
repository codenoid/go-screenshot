# Go Screenshot

another library for taking screenshot on your Linux+,Windows-,Mac-

there is many feature that still missing (like choosing display, set size, etc), and currently only support for linux, and don't worry about [this bug](https://github.com/BurntSushi/xgb/issues/32) because you can open once and reuse xgb connection

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
	// start xgb session, you can reuse this for anything else
	xgbconn, err := screenshot.NewSession()
	if err != nil {
		return
	}
	defer xgbconn.Close()

	// call screenshot method
	ss := screenshot.Start{
		XgbConn: xgbconn,
	}

	// capture current screen frame
	img, err := ss.CaptureScreen()

	f, err := os.Create("screenshot.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	jpeg.Encode(f, img, nil)
}
```