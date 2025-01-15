//go:build darwin && !nodynamic

package clay

import (
	"fmt"

	"github.com/ebitengine/purego"
)

const (
	libname  = "libclay.dylib"
	libcname = "/usr/lib/libSystem.B.dylib"
)

func loadLibrary(name string) (uintptr, error) {
	handle, err := purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, fmt.Errorf("cannot load library: %w", err)
	}

	return uintptr(handle), nil
}
