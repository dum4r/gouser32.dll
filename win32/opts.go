package win32

import (
	"gouser32/win32/w32"
	"image"
)

type optsWin32 struct {
	Icon      image.Image
	TitleName string
	Position  image.Point
	Dimension image.Point

	ShowCursor bool
	Styles     w32.WS_Style
	Ex_Style   int
}

func OptionsWindow() *optsWin32 {

	// Set default Options for Window
	win.opts.Dimension.X = w32.CW_USEDEFAULT
	win.opts.Dimension.Y = w32.CW_USEDEFAULT

	win.opts.Position.X = w32.CW_USEDEFAULT
	win.opts.Position.Y = w32.CW_USEDEFAULT

	win.opts.TitleName = "Default Title"

	win.opts.ShowCursor = true

	win.opts.Styles = w32.OVERLAPPEDWINDOW | w32.SIZEBOX

	return &win.opts
}

// this function modifies the styles to make the window fullscreen
func (o *optsWin32) FullScreemStyle() {
	o.Position = image.Point{}
	o.Dimension = DimensionScreem()
	o.Styles = w32.POPUPWINDOW
}
