package screenshot

import (
	"image"

	"github.com/BurntSushi/xgb/xproto"
)

func (ss *Start) ScreenRect() (image.Rectangle, error) {
	screen := xproto.Setup(ss.XgbConn).DefaultScreen(ss.XgbConn)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func (ss *Start) CaptureScreen() (*image.RGBA, error) {
	r, e := ss.ScreenRect()
	if e != nil {
		return nil, e
	}
	return ss.CaptureRect(r)
}

func (ss *Start) CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	screen := xproto.Setup(ss.XgbConn).DefaultScreen(ss.XgbConn)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(ss.XgbConn, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}