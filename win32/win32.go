package win32

import (
	"gouser32/win32/w32"
	"image"
	"runtime"
	"syscall"
)

const clasName string = "WIN32_MAIN_APP"

type InterfaceWin interface {
	Update()
}

var (
	win struct {
		opts optsWin32
		ctrl InterfaceWin

		instance  syscall.Handle
		window    syscall.Handle
		icon      syscall.Handle
		cursor    syscall.Handle
		classAtom uint16
		run       bool
	}
)

func Run(opt *optsWin32) error {
	runtime.LockOSThread()
	// ctx := context.Background()

	if win.opts.Icon == nil {
		win.icon = DefaultIcon()
	} else {
		win.icon = CreateIcon(win.opts.Icon)
	}
	defer DestroyIcon(win.icon)

	if win.opts.ShowCursor {
		win.cursor = DefaultCursor()
	}

	win.classAtom = RegisterClass(&w32.WNDCLASSX{
		Style:         67, // w32.CLASSDC | w32.HREDRAW | w32.VREDRAW
		LpfnWndProc:   syscall.NewCallback(controllerClass),
		HInstance:     syscall.Handle(win.instance),
		HIcon:         win.icon,
		HCursor:       win.cursor,
		HbrBackground: syscall.Handle(w32.COLOR_BACKGROUND),
		LpszClassName: UTF16PtrFromString(clasName),
		HIconSm:       win.icon,
	})

	win.window = CreateWindow(uintptr(win.classAtom), &win.opts)
	defer DestroyWindow(win.window)

	win.run = true
	ShowWindow(win.window, w32.SW_SHOW)
	UpdateWindow(win.window)

	ShowCursor(win.opts.ShowCursor)
	if win.cursor != 0 {
		defer DestroyCursor(win.cursor)
	}

	runtime.UnlockOSThread()

	var m w32.MSG
	for _GetMessage(&m, win.window, 0, 0) {
		_TranslateMessage(&m)
		_DispatchMessage(&m)
	}
	return nil
}

func controllerClass(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	switch uMsg {
	case w32.WM_SIZE: // layoud windows
		var r w32.RECT
		GetWindowRect(hwnd, &r)
		win.opts.Position = image.Point{X: int(r.Left), Y: int(r.Top)}
		win.opts.Dimension = image.Point{X: r.Width(), Y: r.Height()}
	case w32.WM_QUIT, w32.WM_DESTROY, w32.WM_NCDESTROY, w32.WM_CLOSE:
		_PostQuitMessage(win.window, w32.WM_QUIT, 0, 0)
	}
	return _DefWindowProc(hwnd, uMsg, wParam, lParam)
}
