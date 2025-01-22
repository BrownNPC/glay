package clay

import (
	"errors"
	"unsafe"

	"github.com/ebitengine/purego"
)

// trademark: chat gippite
type String struct {
	Length int
	Chars  *byte
}

type StringArray struct {
	Capacity      uint32
	Length        uint32
	InternalArray []String
}

type Context struct {
	maxElementCount int32
}

type Arena struct {
	NextAllocation uintptr
	Capacity       int
	Memory         unsafe.Pointer
}

type Dimensions struct {
	Width, Height float32
}

type Vector2 struct {
	X float32
	Y float32
}

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

type BoundingBox struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type ElementId struct {
	ID       uint32
	Offset   uint32
	BaseID   uint32
	StringID String
}

type CornerRadius struct {
	TopLeft     float32
	TopRight    float32
	BottomLeft  float32
	BottomRight float32
}

type ElementConfigType int

const (
	ElementConfigTypeNone ElementConfigType = iota
	ElementConfigTypeRectangle
	ElementConfigTypeBorderContainer
	ElementConfigTypeFloatingContainer
	ElementConfigTypeScrollContainer
	ElementConfigTypeImage
	ElementConfigTypeText
	ElementConfigTypeCustom
)

type LayoutDirection int

const (
	LeftToRight LayoutDirection = iota
	TopToBottom
)

type LayoutAlignmentX int

const (
	AlignXLeft LayoutAlignmentX = iota
	AlignXRight
	AlignXCenter
)

type LayoutAlignmentY int

const (
	AlignYTop LayoutAlignmentY = iota
	AlignYBottom
	AlignYCenter
)

type SizingType int

const (
	SizingTypeFit SizingType = iota
	SizingTypeGrow
	SizingTypePercent
	SizingTypeFixed
)

type ChildAlignment struct {
	X LayoutAlignmentX
	Y LayoutAlignmentY
}

type SizingMinMax struct {
	Min float32
	Max float32
}

type SizingAxis struct {
	Size struct {
		MinMax  SizingMinMax
		Percent float32
	}
	Type SizingType
}

type Sizing struct {
	Width  SizingAxis
	Height SizingAxis
}

type Padding struct {
	X uint16
	Y uint16
}

type LayoutConfig struct {
	Sizing          Sizing
	Padding         Padding
	ChildGap        uint16
	ChildAlignment  ChildAlignment
	LayoutDirection LayoutDirection
}

var LayoutDefault LayoutConfig

type RectangleElementConfig struct {
	Color        Color
	CornerRadius CornerRadius
}

type TextElementConfigWrapMode int

const (
	TextWrapWords TextElementConfigWrapMode = iota
	TextWrapNewlines
	TextWrapNone
)

type TextElementConfig struct {
	TextColor     Color
	FontID        uint16
	FontSize      uint16
	LetterSpacing uint16
	LineHeight    uint16
	WrapMode      TextElementConfigWrapMode
}

type ImageElementConfig struct {
	ImageData        interface{}
	SourceDimensions Dimensions
}

type FloatingAttachPointType int

const (
	AttachPointLeftTop FloatingAttachPointType = iota
	AttachPointLeftCenter
	AttachPointLeftBottom
	AttachPointCenterTop
	AttachPointCenterCenter
	AttachPointCenterBottom
	AttachPointRightTop
	AttachPointRightCenter
	AttachPointRightBottom
)

type FloatingAttachPoints struct {
	Element FloatingAttachPointType
	Parent  FloatingAttachPointType
}

type PointerCaptureMode int

const (
	PointerCaptureModeCapture PointerCaptureMode = iota
	PointerCaptureModePassthrough
)

type FloatingElementConfig struct {
	Offset             Vector2
	Expand             Dimensions
	ZIndex             uint16
	ParentID           uint32
	Attachment         FloatingAttachPoints
	PointerCaptureMode PointerCaptureMode
}

type CustomElementConfig struct {
	CustomData interface{}
}

type ScrollElementConfig struct {
	Horizontal bool
	Vertical   bool
}

type Border struct {
	Width uint32
	Color Color
}

type BorderElementConfig struct {
	Left            Border
	Right           Border
	Top             Border
	Bottom          Border
	BetweenChildren Border
	CornerRadius    CornerRadius
}

type ElementConfigUnion struct {
	RectangleElementConfig *RectangleElementConfig
	TextElementConfig      *TextElementConfig
	ImageElementConfig     *ImageElementConfig
	FloatingElementConfig  *FloatingElementConfig
	CustomElementConfig    *CustomElementConfig
	ScrollElementConfig    *ScrollElementConfig
	BorderElementConfig    *BorderElementConfig
}

type ElementConfig struct {
	Type   ElementConfigType
	Config ElementConfigUnion
}

type ScrollContainerData struct {
	ScrollPosition            *Vector2
	ScrollContainerDimensions Dimensions
	ContentDimensions         Dimensions
	Config                    ScrollElementConfig
	Found                     bool
}

type RenderCommandType int

const (
	RenderCommandTypeNone RenderCommandType = iota
	RenderCommandTypeRectangle
	RenderCommandTypeBorder
	RenderCommandTypeText
	RenderCommandTypeImage
	RenderCommandTypeScissorStart
	RenderCommandTypeScissorEnd
	RenderCommandTypeCustom
)

type RenderCommand struct {
	BoundingBox BoundingBox
	Config      ElementConfigUnion
	Text        String
	ID          uint32
	CommandType RenderCommandType
}

type RenderCommandArray struct {
	Capacity      uint32
	Length        uint32
	InternalArray []RenderCommand
}

type PointerDataInteractionState int

const (
	PointerDataPressedThisFrame PointerDataInteractionState = iota
	PointerDataPressed
	PointerDataReleasedThisFrame
	PointerDataReleased
)

type PointerData struct {
	Position Vector2
	State    PointerDataInteractionState
}

type ErrorType int

const (
	ErrorTypeTextMeasurementFunctionNotProvided ErrorType = iota
	ErrorTypeArenaCapacityExceeded
	ErrorTypeElementsCapacityExceeded
	ErrorTypeTextMeasurementCapacityExceeded
	ErrorTypeDuplicateID
	ErrorTypeFloatingContainerParentNotFound
	ErrorTypeInternalError
)

type ErrorData struct {
	ErrorType ErrorType
	ErrorText String
	UserData  uintptr
}

type ErrorHandler struct {
	ErrorHandlerFunction uintptr
	UserData             uintptr
}

// Function Forward Declarations
func MinMemorySize() uint32 {
	if dynamic {
		return _MinMemorySize()
	} else {
		panic("wtf")
	}
	return 0
}

func NewArena() *Arena {
	if dynamic {
		var arena Arena
		_NewArena(uintptr(unsafe.Pointer(&arena)))
		return &arena
	}
	return nil
}

// func CreateArenaWithCapacityAndMemory(capacity uint32, offset unsafe.Pointer) Arena {
// 	// Implementation required
// 	return Arena{}
// }

func SetPointerState(position Vector2, pointerDown bool) {
	// Implementation required
}

func Initialize(arena *Arena, layoutDimensions Dimensions, handleErrorFunc func(err error)) *Context {
	if dynamic {
		errorHandlerFunc := purego.NewCallback(func(error unsafe.Pointer) uintptr {
			errData := (*ErrorData)(error)
			err := errors.New(clayStringToGoString(errData.ErrorText))
			handleErrorFunc(err)
			return 0
		})

		ctx := _Initialize(arena, uintptr(unsafe.Pointer(&layoutDimensions)),
			uintptr(unsafe.Pointer(&ErrorHandler{ErrorHandlerFunction: errorHandlerFunc})))
		return ctx
	}
	return nil
}

func UpdateScrollContainers(enableDragScrolling bool, scrollDelta Vector2, deltaTime float64) {
	// Implementation required
}

func SetLayoutDimensions(dimensions Dimensions) {
	// Implementation required
}

func BeginLayout() {
	// Implementation required
}

func EndLayout() RenderCommandArray {
	// Implementation required
	return RenderCommandArray{}
}

func GetElementId(idString String) ElementId {
	// Implementation required
	return ElementId{}
}

func GetElementIdWithIndex(idString String, index uint32) ElementId {
	// Implementation required
	return ElementId{}
}

func Hovered() bool {
	// Implementation required
	return false
}

func OnHover(onHoverFunction func(ElementId, PointerData, int), userData int) {
	// Implementation required
}

func PointerOver(elementId ElementId) bool {
	// Implementation required
	return false
}

func GetScrollContainerData(id ElementId) ScrollContainerData {
	// Implementation required
	return ScrollContainerData{}
}

func SetMeasureTextFunction(measureTextFunction func(*String, *TextElementConfig) Dimensions) {
	// Implementation required
}

func SetQueryScrollOffsetFunction(queryScrollOffsetFunction func(uint32) Vector2) {
	// Implementation required
}

func RenderCommandArrayGet(array *RenderCommandArray, index int32) *RenderCommand {
	// Implementation required
	return nil
}

func SetDebugModeEnabled(enabled bool) {
	// Implementation required
}

func IsDebugModeEnabled() bool {
	// Implementation required
	return false
}

func SetCullingEnabled(enabled bool) {
	// Implementation required
}

func SetMaxElementCount(maxElementCount uint32) {
	// Implementation required
}

func SetMaxMeasureTextCacheWordCount(maxMeasureTextCacheWordCount uint32) {
	// Implementation required
}

// Internal API functions required by macros
func OpenElement() {
	// Implementation required
}

func CloseElement() {
	// Implementation required
}

func StoreLayoutConfig(config LayoutConfig) *LayoutConfig {
	// Implementation required
	return nil
}

func ElementPostConfiguration() {
	// Implementation required
}

func AttachId(id ElementId) {
	// Implementation required
}

func AttachLayoutConfig(config *LayoutConfig) {
	// Implementation required
}

func AttachElementConfig(config ElementConfigUnion, configType ElementConfigType) {
	// Implementation required
}

func StoreRectangleElementConfig(config RectangleElementConfig) *RectangleElementConfig {
	// Implementation required
	return nil
}

func StoreTextElementConfig(config TextElementConfig) *TextElementConfig {
	// Implementation required
	return nil
}

func StoreImageElementConfig(config ImageElementConfig) *ImageElementConfig {
	// Implementation required
	return nil
}

func StoreFloatingElementConfig(config FloatingElementConfig) *FloatingElementConfig {
	// Implementation required
	return nil
}

func StoreCustomElementConfig(config CustomElementConfig) *CustomElementConfig {
	// Implementation required
	return nil
}

func StoreScrollElementConfig(config ScrollElementConfig) *ScrollElementConfig {
	// Implementation required
	return nil
}

func StoreBorderElementConfig(config BorderElementConfig) *BorderElementConfig {
	// Implementation required
	return nil
}

func HashString(key String, offset uint32, seed uint32) ElementId {
	// Implementation required
	return ElementId{}
}

func Noop() {
	// Implementation required
}

func OpenTextElement(text String, textConfig *TextElementConfig) {
	// Implementation required
}
