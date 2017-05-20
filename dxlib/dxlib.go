package dxlib

// required:
// go get golang.org/x/text/encoding/japanese
import (
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func convertStringToCp932(str string) string {
	r, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return r
}

func pstr(str string) uintptr {
	cp932str := syscall.StringBytePtr(convertStringToCp932(str))
	return uintptr(unsafe.Pointer(cp932str))
}

func pint(i int) uintptr {
	return uintptr(i)
}

var (
	mod = syscall.NewLazyDLL("DxLib.dll")

	dx_ChangeWindowMode = mod.NewProc("dx_ChangeWindowMode")
	dx_DxLib_Init       = mod.NewProc("dx_DxLib_Init")
	dx_DxLib_End        = mod.NewProc("dx_DxLib_End")
	dx_ProcessMessage   = mod.NewProc("dx_ProcessMessage")
	dx_SetWindowText    = mod.NewProc("dx_SetWindowText")
	dx_DrawCircle       = mod.NewProc("dx_DrawCircle")
	dx_LoadGraph        = mod.NewProc("dx_LoadGraph")
	dx_DrawGraph        = mod.NewProc("dx_DrawGraph")
	dx_DrawExtendGraph  = mod.NewProc("dx_DrawExtendGraph")	
	dx_WaitTimer        = mod.NewProc("dx_WaitTimer")
	dx_ScreenFlip       = mod.NewProc("dx_ScreenFlip")
	dx_SetDrawScreen    = mod.NewProc("dx_SetDrawScreen")
	dx_ClearDrawScreen  = mod.NewProc("dx_ClearDrawScreen")
)

func ChangeWindowMode(mode int) int {
	var r, _, _ = dx_ChangeWindowMode.Call(pint(mode))
	return int(r)
}

func DxLib_Init() int {
	var r, _, _ = dx_DxLib_Init.Call()
	return int(r)
}

func DxLib_End() int {
	var r, _, _ = dx_DxLib_End.Call()
	return int(r)
}

func ProcessMessage() int {
	var r, _, _ = dx_ProcessMessage.Call()
	return int(r)
}

func SetWindowText(newTitle string) int {
	var r, _, _ = dx_SetWindowText.Call(pstr(newTitle))
	return int(r)
}

func DrawCircle(x int, y int, circleR int, color int, fillFlag int, lineThickness int) int {
	var r, _, _ = dx_DrawCircle.Call(pint(x), pint(y), pint(circleR), pint(color), pint(fillFlag), pint(lineThickness))
	return int(r)
}

func LoadGraph(fname string) int {
	var r, _, _ = dx_LoadGraph.Call(pstr(fname), 0)
	return int(r)
}

func DrawGraph(x int, y int, grHandle int, transFlag int) int {
	var r, _, _ = dx_DrawGraph.Call(pint(x), pint(y), pint(grHandle), pint(transFlag))
	return int(r)
}

func DrawExtendGraph(x1 int, y1 int, x2 int, y2 int, grHandle int, transFlag int) int {
	var r, _, _ = dx_DrawExtendGraph.Call(pint(x1), pint(y1), pint(x2), pint(y2), pint(grHandle), pint(transFlag))
	return int(r)
}

func WaitTimer(waitTime int) int {
	var r, _, _ = dx_WaitTimer.Call(pint(waitTime))
	return int(r)
}

func ScreenFlip() int {
	var r, _, _ = dx_ScreenFlip.Call()
	return int(r)
}

func SetDrawScreen(drawScreen int) int {
	var r, _, _ = dx_SetDrawScreen.Call(pint(drawScreen))
	return int(r)
}

func ClearDrawScreen() int {
	var r, _, _ = dx_ClearDrawScreen.Call(pint(0))
	return int(r)
}
