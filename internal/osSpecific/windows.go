//go:build windows
// +build windows

package osSpecific

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32           = windows.NewLazySystemDLL("kernel32.dll")
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
	procReadConsole    = kernel32.NewProc("ReadConsoleW")
)

// GetChar reads a single character from stdin without waiting for Enter.
func GetChar() (byte, error) {
	hIn := windows.Handle(syscall.Stdin)

	// Save old mode
	var oldMode uint32
	procGetConsoleMode.Call(uintptr(hIn), uintptr(unsafe.Pointer(&oldMode)))

	// Disable line buffering and echo
	procSetConsoleMode.Call(uintptr(hIn), uintptr(0))

	defer procSetConsoleMode.Call(uintptr(hIn), uintptr(oldMode))

	// Read one wide char
	var buf [1]uint16
	var read uint32
	ret, _, err := procReadConsole.Call(
		uintptr(hIn),
		uintptr(unsafe.Pointer(&buf[0])),
		1,
		uintptr(unsafe.Pointer(&read)),
		0,
	)
	if ret == 0 {
		return 0, err
	}
	return byte(buf[0]), nil
}
