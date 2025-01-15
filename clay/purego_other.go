//go:build (!unix && !darwin && !windows) || nodynamic

package clay

import (
	"fmt"
)

var (
	dynamic    = false
	dynamicErr = fmt.Errorf("webp: dynamic disabled")
)

func loadLibrary(name string) (uintptr, error) {
	return 0, dynamicErr
}
