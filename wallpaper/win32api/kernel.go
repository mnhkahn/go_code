package win32api

import (
	"fmt"
	"unsafe"
)

// #define WIN32_LEAN_AND_MEAN
// #include <windows.h>
import "C"
import "syscall"

func GetCurrentDirectory() string {
	if bufLen := C.GetCurrentDirectoryW(0, nil); bufLen != 0 {
		buf := make([]uint16, bufLen)
		if bufLen := C.GetCurrentDirectoryW(bufLen, (*C.WCHAR)(&buf[0])); bufLen != 0 {
			return syscall.UTF16ToString(buf)
		}
	}
	return ""
}

func SetWallPaper(file_path string) error {
	path := []byte(file_path)
	result := int(C.SystemParametersInfo(C.SPI_SETDESKWALLPAPER, 0, C.PVOID(unsafe.Pointer(&path[0])), C.SPIF_UPDATEINIFILE))
	if result != C.TRUE {
		return fmt.Errorf("", C.GetLastError())
	}
	return nil
}
