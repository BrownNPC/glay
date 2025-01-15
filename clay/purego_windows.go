//go:build windows && !nodynamic

package clay

import (
	"fmt"
	"syscall"
)

const (
	libname  = "clay.dll"
	libcname = "ucrtbase.dll"
)

func loadLibrary(name string) (uintptr, error) {
	handle, err := syscall.LoadLibrary(name)
	if err != nil {
		return 0, fmt.Errorf("cannot load library %s: %w", libname, err)
	}

	return uintptr(handle), nil
}
