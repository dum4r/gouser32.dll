package win32

import (
	"gouser32/win32/w32"
	"image"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	// Protocols concurrent instance
	procGetMessage       = user32.NewProc("GetMessageW")
	procTranslateMessage = user32.NewProc("TranslateMessage")
	procDispatchMessage  = user32.NewProc("DispatchMessageW")
	procDefWindowProc    = user32.NewProc("DefWindowProcW")
	// procSendMessage      = user32.NewProc("SendMessageCallbackA")
)

func _PostQuitMessage(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
	hndl, err := Proc("PostQuitMessage", uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
	if err != nil {
		println(err.Error())
	}
	return uintptr(hndl)
}

// func _SendMessage(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
// 	hndl, err := ProcLazy(procSendMessage, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
// 	if err != nil {
// 		println(err.Error())
// 	}
// 	return uintptr(hndl)
// }

func _DefWindowProc(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
	hndl, _ := ProcLazy(procDefWindowProc, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
	return uintptr(hndl)
}

func _GetMessage(msg *w32.MSG, hwnd syscall.Handle, msgfiltermin uint32, msgfiltermax uint32) bool {
	r0, _ := ProcLazy(procGetMessage, uintptr(unsafe.Pointer(msg)))
	return int32(r0) > 0 && win.run
}

func _TranslateMessage(msg *w32.MSG) (done bool) {
	r0, _ := ProcLazy(procTranslateMessage, uintptr(unsafe.Pointer(msg)))
	return r0 != 0
}

func _DispatchMessage(msg *w32.MSG) int32 {
	r0, _ := ProcLazy(procDispatchMessage, uintptr(unsafe.Pointer(msg)))
	return int32(r0)
}

// func _PostMessage(hwnd syscall.Handle, uMsg uint32, wParam uint32, lParam uint32) bool {
// 	r0, err := Proc("PostMessageW", uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
// 	if err != nil {
// 		println("win32 GetMessage failed: %v", err)
// 		panic(err)
// 	}
// 	return int32(r0) > 0
// }

func RegisterClass(wcx *w32.WNDCLASSX) uint16 {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx))

	proc := user32.NewProc("RegisterClassExW")
	r0, _, e1 := proc.Call(uintptr(unsafe.Pointer(wcx)))
	var err error
	classAtom := uint16(r0) //

	if classAtom == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}

	if err != nil {
		panic(err)
	}
	return classAtom
}

func CreateWindow(lpClassName uintptr, opt *optsWin32) syscall.Handle {
	return ProcPanic("CreateWindowExW",
		uintptr(opt.Ex_Style),
		lpClassName,
		uintptr(unsafe.Pointer(UTF16PtrFromString(opt.TitleName))), // TitleWindows: lpWindowName
		uintptr(opt.Styles),

		uintptr(opt.Position.X),
		uintptr(opt.Position.Y),
		uintptr(opt.Dimension.X),
		uintptr(opt.Dimension.Y),

		uintptr(w32.HWND_DESKTOP), // no owner window: hWndParent
		0,                         // use class menu: hMenu
		uintptr(win.instance),
		0, // no window-creation data: lpParam
	)
}
func DestroyWindow(hndl syscall.Handle) {
	r0, err := Proc("DestroyWindow", uintptr(hndl))
	if r0 == 0 {
		println("Destroy Window: ", hndl, " -> ", err.Error())
	}
}

func ShowWindow(hwnd syscall.Handle, cmdshow int32) bool {
	r0, _ := Proc("ShowWindow", uintptr(hwnd), uintptr(cmdshow))
	return r0 != 0
}

func UpdateWindow(hwnd syscall.Handle) bool {
	r0, _ := Proc("UpdateWindow", uintptr(hwnd))
	return r0 != 0
}

func CreateIcon(img image.Image) syscall.Handle {
	// Convierte la imagen a un formato compatible con CreateIcon
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	data := make([]byte, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			data[y*width*4+x*4+0] = byte(b >> 8)
			data[y*width*4+x*4+1] = byte(g >> 8)
			data[y*width*4+x*4+2] = byte(r >> 8)
			data[y*width*4+x*4+3] = byte(a >> 8)
		}
	}

	return ProcPanic("CreateIcon",
		uintptr(win.instance),
		uintptr(width),
		uintptr(height),
		1, 32,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(unsafe.Pointer(&data[0])),
	)
}

func DestroyIcon(hndl syscall.Handle)   { ProcPanic("DestroyIcon", uintptr(hndl)) }
func DestroyCursor(hndl syscall.Handle) { ProcPanic("DestroyCursor", uintptr(hndl)) }
func ShowCursor(bShow bool)             { ProcPanic("ShowCursor", BoolToUintptr(bShow)) }
func DefaultIcon() syscall.Handle       { return ProcPanic("LoadIconW", 0, w32.IDI_APPLICATION) }
func DefaultCursor() syscall.Handle {
	return ProcPanic("LoadCursorW", uintptr(win.instance), w32.IDC_ARROW)
}

// System
func DimensionScreem() image.Point {
	proc := user32.NewProc("GetSystemMetrics")
	x, _, _ := proc.Call(0) // SM_CXSCREEN = 0
	y, _, _ := proc.Call(1) // SM_CYSCREEN = 1
	return image.Point{X: int(x), Y: int(y)}
}
func GetWindowRect(hwnd syscall.Handle, rect *w32.RECT) {
	_, err := Proc("GetWindowRect", uintptr(hwnd), uintptr(unsafe.Pointer(rect)))
	if err != nil {
		println(err.Error())
	}
}

func ProcLazy(proc *syscall.LazyProc, a ...uintptr) (hndl syscall.Handle, err error) {
	r0, _, e1 := proc.Call(a...)
	hndl = syscall.Handle(r0)
	if hndl == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// template protocols
func Proc(name string, a ...uintptr) (hndl syscall.Handle, err error) {
	proc := user32.NewProc(name)
	r0, _, e1 := proc.Call(a...)
	hndl = syscall.Handle(r0)
	if hndl == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ProcPanic(name string, a ...uintptr) syscall.Handle {
	proc := user32.NewProc(name)
	r0, _, e1 := proc.Call(a...)
	hndl := syscall.Handle(r0)

	var err error
	if hndl == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}

	if err != nil {
		panic(err)
	}
	defer func() {
		proc = nil

	}()
	return hndl
}

func UTF16PtrFromString(str string) *uint16 {
	pointer, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		panic(err)
	}
	return pointer
}

func BoolToUintptr(bln bool) uintptr {
	if bln {
		return 1
	}
	return 0
}
