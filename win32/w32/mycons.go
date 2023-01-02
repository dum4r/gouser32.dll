package w32

// // Register class constants
// type CS_Style int32

// const (
// 	VREDRAW         CS_Style = 0x00000001
// 	HREDRAW         CS_Style = 0x00000002
// 	KEYCVTWINDOW    CS_Style = 0x00000004
// 	OWNDC           CS_Style = 0x00000020
// 	DBLCLKS         CS_Style = 0x00000008
// 	CLASSDC         CS_Style = 0x00000040
// 	PARENTDC        CS_Style = 0x00000080
// 	NOKEYCVT        CS_Style = 0x00000100
// 	NOCLOSE         CS_Style = 0x00000200
// 	SAVEBITS        CS_Style = 0x00000800
// 	BYTEALIGNCLIENT CS_Style = 0x00001000
// 	BYTEALIGNWINDOW CS_Style = 0x00002000
// 	GLOBALCLASS     CS_Style = 0x00004000
// 	IME             CS_Style = 0x00010000
// 	DROPSHADOW      CS_Style = 0x00020000
// )

// Window style constants
type WS_Style int

const (
	OVERLAPPED       WS_Style = 0x00000000
	POPUP            WS_Style = 0x80000000
	CHILD            WS_Style = 0x40000000
	MINIMIZE         WS_Style = 0x20000000
	VISIBLE          WS_Style = 0x10000000
	DISABLED         WS_Style = 0x08000000
	CLIPSIBLINGS     WS_Style = 0x04000000
	CLIPCHILDREN     WS_Style = 0x02000000
	MAXIMIZE         WS_Style = 0x01000000
	CAPTION          WS_Style = 0x00C00000
	BORDER           WS_Style = 0x00800000
	DLGFRAME         WS_Style = 0x00400000
	VSCROLL          WS_Style = 0x00200000
	HSCROLL          WS_Style = 0x00100000
	SYSMENU          WS_Style = 0x00080000
	THICKFRAME       WS_Style = 0x00040000 // No rezible window =>  OVERLAPPEDWINDOW &^ THICKFRAME
	GROUP            WS_Style = 0x00020000
	TABSTOP          WS_Style = 0x00010000
	MINIMIZEBOX      WS_Style = 0x00020000
	MAXIMIZEBOX      WS_Style = 0x00010000
	TILED            WS_Style = 0x00000000
	ICONIC           WS_Style = 0x20000000
	SIZEBOX          WS_Style = 0x00040000 // rezible window
	OVERLAPPEDWINDOW WS_Style = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	POPUPWINDOW      WS_Style = 0x80000000 | 0x00800000 | 0x00080000
	CHILDWINDOW      WS_Style = 0x40000000
)
