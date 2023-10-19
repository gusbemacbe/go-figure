package figure

import (
	"syscall"
	"unsafe"

	"github.com/fatih/color"
)

func init() {
	if majorVersion, _, _ := RtlGetNtVersionNumbers(); majorVersion < 10 {
		IsOldWindows = true

		colorColors = map[string]color.Attribute{
			ColorReset:  color.Reset,
			ColorRed:    color.FgRed,
			ColorGreen:  color.FgGreen,
			ColorYellow: color.FgYellow,
			ColorBlue:   color.FgBlue,
			ColorPurple: color.FgHiMagenta,
			ColorCyan:   color.FgCyan,
			ColorGray:   color.FgWhite,
			ColorWhite:  color.FgHiWhite,
		}
	}
}

func RtlGetNtVersionNumbers() (majorVersion, minorVersion, buildNumber uint32) {
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procRtlGetNtVersionNumbers := ntdll.NewProc("RtlGetNtVersionNumbers")
	_, _, _ = procRtlGetNtVersionNumbers.Call(
		uintptr(unsafe.Pointer(&majorVersion)),
		uintptr(unsafe.Pointer(&minorVersion)),
		uintptr(unsafe.Pointer(&buildNumber)),
	)
	buildNumber &= 0xffff
	return
}
