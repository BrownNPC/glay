//go:build (unix || darwin || windows) && !nodynamic

package clay

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	_MinMemorySize                    func() uint32
	_CreateArenaWithCapacityAndMemory func(capacity uint32, offset **byte) *Arena
	_NewArena                         func(arena uintptr) *Arena
	_SetPointerState                  func(position Vector2, pointerDown bool)
	_Initialize                       func(arena uintptr, layoutDimensions uintptr, errorHandler uintptr) *Context
	_UpdateScrollContainers           func(enableDragScrolling bool, scrollDelta Vector2, deltaTime float64)
	_SetLayoutDimensions              func(dimensions Dimensions)
	_BeginLayout                      func()
	_EndLayout                        func() RenderCommandArray
	_GetElementId                     func(idString string) ElementId
	_GetElementIdWithIndex            func(idString string, index uint32) ElementId
	_Hovered                          func() bool
	_OnHover                          func(onHoverFunction func(ElementId, PointerData, int), userData int)
	_PointerOver                      func(elementId ElementId) bool
	_GetScrollContainerData           func(id ElementId) ScrollContainerData
	_SetMeasureTextFunction           func(measureTextFunction func(string, *TextElementConfig) Dimensions)
	_SetQueryScrollOffsetFunction     func(queryScrollOffsetFunction func(uint32) Vector2)
	_RenderCommandArrayGet            func(array *RenderCommandArray, index int32) *RenderCommand
	_SetDebugModeEnabled              func(enabled bool)
	_IsDebugModeEnabled               func() bool
	_SetCullingEnabled                func(enabled bool)
	_SetMaxElementCount               func(maxElementCount uint32)
	_SetMaxMeasureTextCacheWordCount  func(maxMeasureTextCacheWordCount uint32)
	// Internal API functions required
	_OpenElement                 func()
	_CloseElement                func()
	_StoreLayoutConfig           func(config LayoutConfig) *LayoutConfig
	_ElementPostConfiguration    func()
	_AttachId                    func(id ElementId)
	_AttachLayoutConfig          func(config *LayoutConfig)
	_AttachElementConfig         func(config ElementConfigUnion, configType ElementConfigType)
	_StoreRectangleElementConfig func(config RectangleElementConfig) *RectangleElementConfig
	_StoreTextElementConfig      func(config TextElementConfig) *TextElementConfig
	_StoreImageElementConfig     func(config ImageElementConfig) *ImageElementConfig
	_StoreFloatingElementConfig  func(config FloatingElementConfig) *FloatingElementConfig
	_StoreCustomElementConfig    func(config CustomElementConfig) *CustomElementConfig
	_StoreScrollElementConfig    func(config ScrollElementConfig) *ScrollElementConfig
	_StoreBorderElementConfig    func(config BorderElementConfig) *BorderElementConfig
	_HashString                  func(key string, offset uint32, seed uint32) ElementId
	_Noop                        func()
	_OpenTextElement             func(text string, textConfig *TextElementConfig)
)

var (
	malloc func(size int) *byte
	free   func(ptr unsafe.Pointer)
)

var (
	libclay    uintptr
	libc       uintptr
	dynamic    bool
	dynamicErr error
)

func init() {
	var err error
	defer func() {
		if r := recover(); r != nil {
			dynamic = false
			dynamicErr = fmt.Errorf("%v", r)
		}
	}()

	libclay, err = loadLibrary(libname)
	if err == nil {
		dynamic = true
	} else {
		dynamicErr = err
	}
	log.Println(dynamicErr)
	if !dynamic {
		return
	}

	libc, err = loadLibrary(libcname)
	if err != nil {
		dynamicErr = err
	}
	purego.RegisterLibFunc(&malloc, libc, "malloc")
	purego.RegisterLibFunc(&_NewArena, libclay, "NewArena")
	purego.RegisterLibFunc(&_MinMemorySize, libclay, "Clay_MinMemorySize")

	// purego.RegisterLibFunc(&_CreateArenaWithCapacityAndMemory, libclay, "Clay_CreateArenaWithCapacityAndMemory")
	// purego.RegisterLibFunc(&_SetPointerState, libclay, "Clay_SetPointerState")
	purego.RegisterLibFunc(&_Initialize, libclay, "Clay_Initialize")
	// purego.RegisterLibFunc(&_UpdateScrollContainers, libclay, "Clay_UpdateScrollContainers")
	// purego.RegisterLibFunc(&_SetLayoutDimensions, libclay, "Clay_SetLayoutDimensions")
	purego.RegisterLibFunc(&_BeginLayout, libclay, "Clay_BeginLayout")
	// purego.RegisterLibFunc(&_EndLayout, libclay, "Clay_EndLayout")
	// purego.RegisterLibFunc(&_GetElementId, libclay, "Clay_GetElementId")
	// purego.RegisterLibFunc(&_GetElementIdWithIndex, libclay, "Clay_GetElementIdWithIndex")
	// purego.RegisterLibFunc(&_Hovered, libclay, "Clay_Hovered")
	// purego.RegisterLibFunc(&_OnHover, libclay, "Clay_OnHover")
	// purego.RegisterLibFunc(&_PointerOver, libclay, "Clay_PointerOver")
	// purego.RegisterLibFunc(&_GetScrollContainerData, libclay, "Clay_GetScrollContainerData")
	// purego.RegisterLibFunc(&_SetMeasureTextFunction, libclay, "Clay_SetMeasureTextFunction")
	// purego.RegisterLibFunc(&_SetQueryScrollOffsetFunction, libclay, "Clay_SetQueryScrollOffsetFunction")
	// purego.RegisterLibFunc(&_RenderCommandArrayGet, libclay, "Clay_RenderCommandArrayGet")
	// purego.RegisterLibFunc(&_SetDebugModeEnabled, libclay, "Clay_SetDebugModeEnabled")
	// purego.RegisterLibFunc(&_IsDebugModeEnabled, libclay, "Clay_IsDebugModeEnabled")
	// purego.RegisterLibFunc(&_SetCullingEnabled, libclay, "Clay_SetCullingEnabled")
	// purego.RegisterLibFunc(&_SetMaxElementCount, libclay, "Clay_SetMaxElementCount")
	// purego.RegisterLibFunc(&_SetMaxMeasureTextCacheWordCount, libclay, "Clay_SetMaxMeasureTextCacheWordCount")

	//Private Api
	// purego.RegisterLibFunc(&_OpenElement, libclay, "_OpenElement")
	// purego.RegisterLibFunc(&_CloseElement, libclay, "_CloseElement")
	// purego.RegisterLibFunc(&_StoreLayoutConfig, libclay, "_StoreLayoutConfig")
	// purego.RegisterLibFunc(&_ElementPostConfiguration, libclay, "_ElementPostConfiguration")
	// purego.RegisterLibFunc(&_AttachLayoutConfig, libclay, "_AttachLayoutConfig")
	// purego.RegisterLibFunc(&_AttachElementConfig, libclay, "_AttachElementConfig")
	// purego.RegisterLibFunc(&_StoreRectangleElementConfig, libclay, "_StoreRectangleElementConfig")
	// purego.RegisterLibFunc(&_StoreTextElementConfig, libclay, "_StoreTextElementConfig")
	// purego.RegisterLibFunc(&_StoreImageElementConfig, libclay, "_StoreImageElementConfig")
	// purego.RegisterLibFunc(&_StoreFloatingElementConfig, libclay, "_StoreFloatingElementConfig")
	// purego.RegisterLibFunc(&_StoreCustomElementConfig, libclay, "_StoreCustomElementConfig")
	// purego.RegisterLibFunc(&_StoreScrollElementConfig, libclay, "_StoreScrollElementConfig")
	// purego.RegisterLibFunc(&_StoreBorderElementConfig, libclay, "_StoreBorderElementConfig")
	// purego.RegisterLibFunc(&_HashString, libclay, "_HashString")
	// purego.RegisterLibFunc(&_Noop, libclay, "_Noop")
	// purego.RegisterLibFunc(&_OpenTextElement, libclay, "_OpenTextElement")

}
func clayStringToGoString(cString String) string {
	// Check for nil pointer
	if cString.Chars == nil || cString.Length <= 0 {
		return ""
	}

	// Create a Go slice backed by the C string memory
	charSlice := unsafe.Slice(cString.Chars, cString.Length)

	// Convert the slice to a Go string
	return string(charSlice)
}
