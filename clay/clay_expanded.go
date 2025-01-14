package clay

import (
	"math"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"
)

const CLAY__NULL = 0
const CLAY__MAXFLOAT = 3.4028234663852886e+38

var CLAY__ELEMENT_DEFINITION_LATCH uint8

type Clay_String string
type Clay__StringArray Arena[string]

type Clay_Context struct {
	maxElementCount                    int32
	maxMeasureTextCacheWordCount       int32
	warningsEnabled                    bool
	errorHandler                       Clay_ErrorHandler
	booleanWarnings                    Clay_BooleanWarnings
	warnings                           Clay__WarningArray
	pointerInfo                        Clay_PointerData
	layoutDimensions                   Clay_Dimensions
	dynamicElementIndexBaseHash        Clay_ElementId
	dynamicElementIndex                uint32
	debugModeEnabled                   bool
	disableCulling                     bool
	externalScrollHandlingEnabled      bool
	debugSelectedElementId             uint32
	generation                         uint32
	arenaResetOffset                   uint64
	internalArena                      Clay_Arena
	layoutElements                     Clay_LayoutElementArray
	renderCommands                     Clay_RenderCommandArray
	openLayoutElementStack             Clay__int32_tArray
	layoutElementChildren              Clay__int32_tArray
	layoutElementChildrenBuffer        Clay__int32_tArray
	textElementData                    Clay__TextElementDataArray
	imageElementPointers               Clay__LayoutElementPointerArray
	reusableElementIndexBuffer         Clay__int32_tArray
	layoutElementClipElementIds        Clay__int32_tArray
	layoutConfigs                      Clay__LayoutConfigArray
	elementConfigBuffer                Clay__ElementConfigArray
	elementConfigs                     Clay__ElementConfigArray
	rectangleElementConfigs            Clay__RectangleElementConfigArray
	textElementConfigs                 Clay__TextElementConfigArray
	imageElementConfigs                Clay__ImageElementConfigArray
	floatingElementConfigs             Clay__FloatingElementConfigArray
	scrollElementConfigs               Clay__ScrollElementConfigArray
	customElementConfigs               Clay__CustomElementConfigArray
	borderElementConfigs               Clay__BorderElementConfigArray
	layoutElementIdStrings             Clay__StringArray
	wrappedTextLines                   Clay__WrappedTextLineArray
	layoutElementTreeNodeArray1        Clay__LayoutElementTreeNodeArray
	layoutElementTreeRoots             Clay__LayoutElementTreeRootArray
	layoutElementsHashMapInternal      Clay__LayoutElementHashMapItemArray
	layoutElementsHashMap              Clay__int32_tArray
	measureTextHashMapInternal         Clay__MeasureTextCacheItemArray
	measureTextHashMapInternalFreeList Clay__int32_tArray
	measureTextHashMap                 Clay__int32_tArray
	measuredWords                      Clay__MeasuredWordArray
	measuredWordsFreeList              Clay__int32_tArray
	openClipElementStack               Clay__int32_tArray
	pointerOverIds                     Clay__ElementIdArray
	scrollContainerDatas               Clay__ScrollContainerDataInternalArray
	treeNodeVisited                    Clay__BoolArray
	dynamicStringData                  Clay__CharArray
	debugElementData                   Clay__DebugElementDataArray
}
type Clay_Arena struct {
	nextAllocation uint64
	capacity       uint64
	memory         *byte
}

type Clay_Dimensions struct {
	width  float32
	height float32
}

type Clay_Vector2 struct {
	x float32
	y float32
}

type Clay_Color struct {
	r float32
	g float32
	b float32
	a float32
}

type Clay_BoundingBox struct {
	x      float32
	y      float32
	width  float32
	height float32
}

type Clay_ElementId Clay_String

type Clay_CornerRadius struct {
	topLeft     float32
	topRight    float32
	bottomLeft  float32
	bottomRight float32
}

type Clay__ElementConfigType int

const (
	CLAY__ELEMENT_CONFIG_TYPE_NONE               Clay__ElementConfigType = 0
	CLAY__ELEMENT_CONFIG_TYPE_RECTANGLE          Clay__ElementConfigType = 1
	CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER   Clay__ElementConfigType = 2
	CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER Clay__ElementConfigType = 4
	CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER   Clay__ElementConfigType = 8
	CLAY__ELEMENT_CONFIG_TYPE_IMAGE              Clay__ElementConfigType = 16
	CLAY__ELEMENT_CONFIG_TYPE_TEXT               Clay__ElementConfigType = 32
	CLAY__ELEMENT_CONFIG_TYPE_CUSTOM             Clay__ElementConfigType = 64
)

type Clay_LayoutDirection int

const (
	CLAY_LEFT_TO_RIGHT = Clay_LayoutDirection(iota)
	CLAY_TOP_TO_BOTTOM
)

type Clay_LayoutAlignmentX int

const (
	CLAY_ALIGN_X_LEFT = Clay_LayoutAlignmentX(iota)
	CLAY_ALIGN_X_RIGHT
	CLAY_ALIGN_X_CENTER
)

type Clay_LayoutAlignmentY int

const (
	CLAY_ALIGN_Y_TOP = Clay_LayoutAlignmentY(iota)
	CLAY_ALIGN_Y_BOTTOM
	CLAY_ALIGN_Y_CENTER
)

type Clay__SizingType int

const (
	CLAY__SIZING_TYPE_FIT = Clay__SizingType(iota)
	CLAY__SIZING_TYPE_GROW
	CLAY__SIZING_TYPE_PERCENT
	CLAY__SIZING_TYPE_FIXED
)

type Clay_ChildAlignment struct {
	x Clay_LayoutAlignmentX
	y Clay_LayoutAlignmentY
}

type Clay_SizingMinMax struct {
	min float32
	max float32
}

type Clay_SizingAxis struct {
	size struct {
		// union
		minMax  Clay_SizingMinMax
		percent float32
	}
	type_ Clay__SizingType
}

type Clay_Sizing struct {
	width  Clay_SizingAxis
	height Clay_SizingAxis
}

type Clay_Padding struct {
	left   uint16
	right  uint16
	top    uint16
	bottom uint16
}

type Clay_LayoutConfig struct {
	sizing          Clay_Sizing
	padding         Clay_Padding
	childGap        uint16
	childAlignment  Clay_ChildAlignment
	layoutDirection Clay_LayoutDirection
}

type Clay_RectangleElementConfig struct {
	color        Clay_Color
	cornerRadius Clay_CornerRadius
}

type Clay_TextElementConfigWrapMode int

const (
	CLAY_TEXT_WRAP_WORDS = Clay_TextElementConfigWrapMode(iota)
	CLAY_TEXT_WRAP_NEWLINES
	CLAY_TEXT_WRAP_NONE
)

type Clay_TextElementConfig struct {
	textColor     Clay_Color
	fontId        uint16
	fontSize      uint16
	letterSpacing uint16
	lineHeight    uint16
	wrapMode      Clay_TextElementConfigWrapMode
}

type Clay_ImageElementConfig struct {
	imageData        unsafe.Pointer
	sourceDimensions Clay_Dimensions
}

type Clay_FloatingAttachPointType int

const (
	CLAY_ATTACH_POINT_LEFT_TOP = Clay_FloatingAttachPointType(iota)
	CLAY_ATTACH_POINT_LEFT_CENTER
	CLAY_ATTACH_POINT_LEFT_BOTTOM
	CLAY_ATTACH_POINT_CENTER_TOP
	CLAY_ATTACH_POINT_CENTER_CENTER
	CLAY_ATTACH_POINT_CENTER_BOTTOM
	CLAY_ATTACH_POINT_RIGHT_TOP
	CLAY_ATTACH_POINT_RIGHT_CENTER
	CLAY_ATTACH_POINT_RIGHT_BOTTOM
)

type Clay_FloatingAttachPoints struct {
	element Clay_FloatingAttachPointType
	parent  Clay_FloatingAttachPointType
}

type Clay_PointerCaptureMode int

const (
	CLAY_POINTER_CAPTURE_MODE_CAPTURE = Clay_PointerCaptureMode(iota)
	CLAY_POINTER_CAPTURE_MODE_PASSTHROUGH
)

type Clay_FloatingElementConfig struct {
	offset             Clay_Vector2
	expand             Clay_Dimensions
	zIndex             uint16
	parentId           uint32
	attachment         Clay_FloatingAttachPoints
	pointerCaptureMode Clay_PointerCaptureMode
}

type Clay_CustomElementConfig struct {
	customData unsafe.Pointer
}

type Clay_ScrollElementConfig struct {
	horizontal bool
	vertical   bool
}

type Clay_Border struct {
	width uint32
	color Clay_Color
}

type Clay_BorderElementConfig struct {
	left            Clay_Border
	right           Clay_Border
	top             Clay_Border
	bottom          Clay_Border
	betweenChildren Clay_Border
	cornerRadius    Clay_CornerRadius
}

type Clay_ElementConfigUnion struct {
	// union
	rectangleElementConfig *Clay_RectangleElementConfig
	textElementConfig      *Clay_TextElementConfig
	imageElementConfig     *Clay_ImageElementConfig
	floatingElementConfig  *Clay_FloatingElementConfig
	customElementConfig    *Clay_CustomElementConfig
	scrollElementConfig    *Clay_ScrollElementConfig
	borderElementConfig    *Clay_BorderElementConfig
}

type Clay_ElementConfig struct {
	type_  Clay__ElementConfigType
	config Clay_ElementConfigUnion
}

type Clay_ScrollContainerData struct {
	scrollPosition            *Clay_Vector2
	scrollContainerDimensions Clay_Dimensions
	contentDimensions         Clay_Dimensions
	config                    Clay_ScrollElementConfig
	found                     bool
}

type Clay_RenderCommandType int

const (
	CLAY_RENDER_COMMAND_TYPE_NONE = Clay_RenderCommandType(iota)
	CLAY_RENDER_COMMAND_TYPE_RECTANGLE
	CLAY_RENDER_COMMAND_TYPE_BORDER
	CLAY_RENDER_COMMAND_TYPE_TEXT
	CLAY_RENDER_COMMAND_TYPE_IMAGE
	CLAY_RENDER_COMMAND_TYPE_SCISSOR_START
	CLAY_RENDER_COMMAND_TYPE_SCISSOR_END
	CLAY_RENDER_COMMAND_TYPE_CUSTOM
)

type Clay_RenderCommand struct {
	boundingBox Clay_BoundingBox
	config      Clay_ElementConfigUnion
	text        Clay_String
	id          uint32
	commandType Clay_RenderCommandType
}

type Clay_RenderCommandArray *Clay_RenderCommand

type Clay_PointerDataInteractionState int

const (
	CLAY_POINTER_DATA_PRESSED_THIS_FRAME = Clay_PointerDataInteractionState(iota)
	CLAY_POINTER_DATA_PRESSED
	CLAY_POINTER_DATA_RELEASED_THIS_FRAME
	CLAY_POINTER_DATA_RELEASED
)

type Clay_PointerData struct {
	position Clay_Vector2
	state    Clay_PointerDataInteractionState
}

type Clay_ErrorType int

const (
	CLAY_ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED = Clay_ErrorType(iota)
	CLAY_ERROR_TYPE_ARENA_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED
	CLAY_ERROR_TYPE_DUPLICATE_ID
	CLAY_ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND
	CLAY_ERROR_TYPE_INTERNAL_ERROR
)

type Clay_ErrorData struct {
	errorType Clay_ErrorType
	errorText Clay_String
	userData  uint64
}

type Clay_ErrorHandler struct {
	errorHandlerFunction func(errorText Clay_ErrorData)
	userData             uint64
}

var Clay__currentContext *Clay_Context
var Clay__defaultMaxElementCount int32 = 8192
var Clay__defaultMaxMeasureTextWordCacheCount int32 = 16384

func Clay__ErrorHandlerFunctionDefault(errorText Clay_ErrorData) {
	_ = errorText
}

var CLAY__SPACECHAR Clay_String = Clay_String(" ")
var CLAY__STRING_DEFAULT Clay_String

type Clay_BooleanWarnings struct {
	maxElementsExceeded         bool
	maxRenderCommandsExceeded   bool
	maxTextMeasureCacheExceeded bool
}

type Clay__Warning struct {
	baseMessage    Clay_String
	dynamicMessage Clay_String
}

var CLAY__WARNING_DEFAULT Clay__Warning

type Clay__WarningArray *Clay__Warning

type Clay__BoolArray Arena[bool]

var CLAY__ELEMENT_ID_DEFAULT Clay_ElementId

type Clay__ElementIdArray Arena[Clay_ElementId]

var CLAY__ELEMENT_CONFIG_DEFAULT Clay_ElementConfig = Clay_ElementConfig{type_: CLAY__ELEMENT_CONFIG_TYPE_NONE, config: Clay_ElementConfigUnion{}}

type Clay__ElementConfigArray *Clay_ElementConfig

type Clay__ElementConfigArraySlice struct {
	length        int32
	internalArray *Clay_ElementConfig
}

var CLAY_LAYOUT_DEFAULT Clay_LayoutConfig = Clay_LayoutConfig{sizing: Clay_Sizing{width: Clay_SizingAxis{size: struct {
	// union
	minMax  Clay_SizingMinMax
	percent float32
}{minMax: Clay_SizingMinMax{min: 0, max: CLAY__MAXFLOAT}}, type_: CLAY__SIZING_TYPE_FIT}, height: Clay_SizingAxis{size: struct {
	// union
	minMax  Clay_SizingMinMax
	percent float32
}{minMax: Clay_SizingMinMax{min: 0, max: CLAY__MAXFLOAT}}, type_: CLAY__SIZING_TYPE_FIT}}}

type Clay__LayoutConfigArray *Clay_LayoutConfig

var CLAY__RECTANGLE_ELEMENT_CONFIG_DEFAULT Clay_RectangleElementConfig = Clay_RectangleElementConfig{}

type Clay__RectangleElementConfigArray *Clay_RectangleElementConfig

var CLAY__TEXT_ELEMENT_CONFIG_DEFAULT Clay_TextElementConfig = Clay_TextElementConfig{}

type Clay__TextElementConfigArray *Clay_TextElementConfig

var CLAY__IMAGE_ELEMENT_CONFIG_DEFAULT Clay_ImageElementConfig = Clay_ImageElementConfig{}

type Clay__ImageElementConfigArray *Clay_ImageElementConfig

var CLAY__FLOATING_ELEMENT_CONFIG_DEFAULT Clay_FloatingElementConfig = Clay_FloatingElementConfig{}

type Clay__FloatingElementConfigArray *Clay_FloatingElementConfig

var CLAY__CUSTOM_ELEMENT_CONFIG_DEFAULT Clay_CustomElementConfig = Clay_CustomElementConfig{}

type Clay__CustomElementConfigArray *Clay_CustomElementConfig

var CLAY__SCROLL_ELEMENT_CONFIG_DEFAULT Clay_ScrollElementConfig = Clay_ScrollElementConfig{horizontal: false}

type Clay__ScrollElementConfigArray *Clay_ScrollElementConfig

type Clay__StringArraySlice Arena[*Clay_String]

type Clay__WrappedTextLine struct {
	dimensions Clay_Dimensions
	line       Clay_String
}

var CLAY__WRAPPED_TEXT_LINE_DEFAULT Clay__WrappedTextLine = Clay__WrappedTextLine{}

type Clay__WrappedTextLineArray *Clay__WrappedTextLine

type Clay__WrappedTextLineArraySlice Arena[*Clay__WrappedTextLine]

type Clay__TextElementData struct {
	text                Clay_String
	preferredDimensions Clay_Dimensions
	elementIndex        int32
	wrappedLines        Clay__WrappedTextLineArraySlice
}

var CLAY__TEXT_ELEMENT_DATA_DEFAULT Clay__TextElementData = Clay__TextElementData{}

type Clay__TextElementDataArray *Clay__TextElementData

var CLAY__BORDER_ELEMENT_CONFIG_DEFAULT Clay_BorderElementConfig = Clay_BorderElementConfig{}

type Clay__BorderElementConfigArray *Clay_BorderElementConfig

type Clay__LayoutElementChildren struct {
	elements *int32
	length   uint16
}

type Clay_LayoutElement struct {
	childrenOrTextContent struct {
		// union
		children        Clay__LayoutElementChildren
		textElementData *Clay__TextElementData
	}
	dimensions     Clay_Dimensions
	minDimensions  Clay_Dimensions
	layoutConfig   *Clay_LayoutConfig
	elementConfigs Clay__ElementConfigArraySlice
	configsEnabled uint32
	id             uint32
}

var CLAY__LAYOUT_ELEMENT_DEFAULT Clay_LayoutElement = Clay_LayoutElement{}

type Clay_LayoutElementArray Arena[*Clay_LayoutElement]

type Clay__LayoutElementPointerArray Arena[*Clay_LayoutElement]

var CLAY__RENDER_COMMAND_DEFAULT Clay_RenderCommand = Clay_RenderCommand{}

type Clay__ScrollContainerDataInternal struct {
	layoutElement       *Clay_LayoutElement
	boundingBox         Clay_BoundingBox
	contentSize         Clay_Dimensions
	scrollOrigin        Clay_Vector2
	pointerOrigin       Clay_Vector2
	scrollMomentum      Clay_Vector2
	scrollPosition      Clay_Vector2
	previousDelta       Clay_Vector2
	momentumTime        float32
	elementId           uint32
	openThisFrame       bool
	pointerScrollActive bool
}

var CLAY__SCROLL_CONTAINER_DEFAULT Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternal{}

type Clay__ScrollContainerDataInternalArray Arena[*Clay__ScrollContainerDataInternal]

type Clay__DebugElementData struct {
	collision bool
	collapsed bool
}

var CLAY__DEBUG_ELEMENT_DATA_DEFAULT Clay__DebugElementData = Clay__DebugElementData{collision: false}

type Clay__DebugElementDataArray Arena[*Clay__DebugElementData]

type Clay_LayoutElementHashMapItem struct {
	boundingBox           Clay_BoundingBox
	elementId             Clay_ElementId
	layoutElement         *Clay_LayoutElement
	onHoverFunction       func(elementId Clay_ElementId, pointerInfo Clay_PointerData, userData int64)
	hoverFunctionUserData int64
	nextIndex             int32
	generation            uint32
	debugData             *Clay__DebugElementData
}

var CLAY__LAYOUT_ELEMENT_HASH_MAP_ITEM_DEFAULT Clay_LayoutElementHashMapItem = Clay_LayoutElementHashMapItem{layoutElement: &CLAY__LAYOUT_ELEMENT_DEFAULT}

type Clay__LayoutElementHashMapItemArray *Clay_LayoutElementHashMapItem

type Clay__MeasuredWord struct {
	startOffset int32
	length      int32
	width       float32
	next        int32
}

var CLAY__MEASURED_WORD_DEFAULT Clay__MeasuredWord = Clay__MeasuredWord{next: -1}

type Clay__MeasuredWordArray *Clay__MeasuredWord

type Clay__MeasureTextCacheItem struct {
	unwrappedDimensions     Clay_Dimensions
	measuredWordsStartIndex int32
	containsNewlines        bool
	id                      uint32
	nextIndex               int32
	generation              uint32
}

var CLAY__MEASURE_TEXT_CACHE_ITEM_DEFAULT Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItem{measuredWordsStartIndex: -1}

type Clay__MeasureTextCacheItemArray *Clay__MeasureTextCacheItem

type Clay__int32_tArray Arena[int32]

type Clay__LayoutElementTreeNode struct {
	layoutElement   *Clay_LayoutElement
	position        Clay_Vector2
	nextChildOffset Clay_Vector2
}

var CLAY__LAYOUT_ELEMENT_TREE_NODE_DEFAULT Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNode{}

type Clay__LayoutElementTreeNodeArray *Clay__LayoutElementTreeNode

type Clay__LayoutElementTreeRoot struct {
	layoutElementIndex int32
	parentId           uint32
	clipElementId      uint32
	zIndex             int32
	pointerOffset      Clay_Vector2
}

var CLAY__LAYOUT_ELEMENT_TREE_ROOT_DEFAULT Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRoot{}

type Clay__LayoutElementTreeRootArray *Clay__LayoutElementTreeRoot

type Clay__CharArray *uint8

var Clay__MeasureText func(text *Clay_String, config *Clay_TextElementConfig) Clay_Dimensions
var Clay__QueryScrollOffset func(elementId uint32) Clay_Vector2

func Clay__GetOpenLayoutElement() *Clay_LayoutElement {
	var context *Clay_Context = Clay_GetCurrentContext()
	return Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-1)))
}
func Clay__GetParentElementId() uint32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-2))).id
}
func Clay__ElementHasConfig(element *Clay_LayoutElement, type_ Clay__ElementConfigType) bool {
	return Clay__ElementConfigType(element.configsEnabled)&type_ != 0
}
func Clay__FindElementConfigWithType(element *Clay_LayoutElement, type_ Clay__ElementConfigType) Clay_ElementConfigUnion {
	for i := int32(0); int(i) < int(element.elementConfigs.length); i++ {
		var config *Clay_ElementConfig = Clay__ElementConfigArraySlice_Get(&element.elementConfigs, i)
		if config.type_ == type_ {
			return config.config
		}
	}
	return Clay_ElementConfigUnion{}
}
func Clay__HashNumber(offset uint32, seed uint32) Clay_ElementId {
	var hash uint32 = seed
	hash += uint32(int32(int(offset) + 48))
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(int32(int(hash) << 3))
	hash ^= uint32(int32(int(hash) >> 11))
	hash += uint32(int32(int(hash) << 15))
	return Clay_ElementId{id: uint32(int32(int(hash) + 1)), offset: offset, baseId: seed, stringId: CLAY__STRING_DEFAULT}
}
func Clay__HashString(key Clay_String, offset uint32, seed uint32) Clay_ElementId {
	var (
		hash uint32 = 0
		base uint32 = seed
	)
	for i := int32(0); int(i) < int(key.length); i++ {
		base += uint32(*(*byte)(unsafe.Add(unsafe.Pointer(key.chars), i)))
		base += uint32(int32(int(base) << 10))
		base ^= uint32(int32(int(base) >> 6))
	}
	hash = base
	hash += offset
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(int32(int(hash) << 3))
	base += uint32(int32(int(base) << 3))
	hash ^= uint32(int32(int(hash) >> 11))
	base ^= uint32(int32(int(base) >> 11))
	hash += uint32(int32(int(hash) << 15))
	base += uint32(int32(int(base) << 15))
	return Clay_ElementId{id: uint32(int32(int(hash) + 1)), offset: offset, baseId: uint32(int32(int(base) + 1)), stringId: key}
}
func Clay__Rehash(elementId Clay_ElementId, number uint32) Clay_ElementId {
	var id uint32 = elementId.baseId
	id += number
	id += uint32(int32(int(id) << 10))
	id ^= uint32(int32(int(id) >> 6))
	id += uint32(int32(int(id) << 3))
	id ^= uint32(int32(int(id) >> 11))
	id += uint32(int32(int(id) << 15))
	return Clay_ElementId{id: id, offset: number, baseId: elementId.baseId, stringId: elementId.stringId}
}
func Clay__RehashWithNumber(id uint32, number uint32) uint32 {
	id += number
	id += uint32(int32(int(id) << 10))
	id ^= uint32(int32(int(id) >> 6))
	id += uint32(int32(int(id) << 3))
	id ^= uint32(int32(int(id) >> 11))
	id += uint32(int32(int(id) << 15))
	return id
}
func Clay__HashTextWithConfig(text *Clay_String, config *Clay_TextElementConfig) uint32 {
	var (
		hash            uint32 = 0
		pointerAsNumber uint64 = uint64(uintptr(unsafe.Pointer(text.chars)))
	)
	hash += uint32(pointerAsNumber)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(text.length)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(config.fontId)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(config.fontSize)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(config.lineHeight)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(config.letterSpacing)
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(int32(config.wrapMode))
	hash += uint32(int32(int(hash) << 10))
	hash ^= uint32(int32(int(hash) >> 6))
	hash += uint32(int32(int(hash) << 3))
	hash ^= uint32(int32(int(hash) >> 11))
	hash += uint32(int32(int(hash) << 15))
	return uint32(int32(int(hash) + 1))
}
func Clay__AddMeasuredWord(word Clay__MeasuredWord, previousWord *Clay__MeasuredWord) *Clay__MeasuredWord {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int(context.measuredWordsFreeList.length) > 0 {
		var newItemIndex uint32 = uint32(Clay__int32_tArray_Get(&context.measuredWordsFreeList, int32(int(context.measuredWordsFreeList.length)-1)))
		context.measuredWordsFreeList.length--
		Clay__MeasuredWordArray_Set(&context.measuredWords, int32(int(newItemIndex)), word)
		previousWord.next = int32(newItemIndex)
		return Clay__MeasuredWordArray_Get(&context.measuredWords, int32(int(newItemIndex)))
	} else {
		previousWord.next = context.measuredWords.length
		return Clay__MeasuredWordArray_Add(&context.measuredWords, word)
	}
}
func Clay__MeasureTextCached(text *Clay_String, config *Clay_TextElementConfig) *Clay__MeasureTextCacheItem {
	var context *Clay_Context = Clay_GetCurrentContext()
	if Clay__MeasureText == nil {
		context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_TEXT_MEASUREMENT_FUNCTION_NOT_PROVIDED, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay's internal MeasureText function is null. You may have forgotten to call Clay_SetMeasureTextFunction(), or passed a NULL function pointer by mistake.")}, userData: context.errorHandler.userData})
		return nil
	}
	var id uint32 = Clay__HashTextWithConfig(text, config)
	var hashBucket uint32 = uint32(int32(int(id) % (int(context.maxMeasureTextCacheWordCount) / 32)))
	var elementIndexPrevious int32 = 0
	var elementIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.measureTextHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for int(elementIndex) != 0 {
		var hashEntry *Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItemArray_Get(&context.measureTextHashMapInternal, elementIndex)
		if int(hashEntry.id) == int(id) {
			hashEntry.generation = context.generation
			return hashEntry
		}
		if int(context.generation)-int(hashEntry.generation) > 2 {
			var nextWordIndex int32 = hashEntry.measuredWordsStartIndex
			for int(nextWordIndex) != -1 {
				var measuredWord *Clay__MeasuredWord = Clay__MeasuredWordArray_Get(&context.measuredWords, nextWordIndex)
				Clay__int32_tArray_Add(&context.measuredWordsFreeList, nextWordIndex)
				nextWordIndex = measuredWord.next
			}
			var nextIndex int32 = hashEntry.nextIndex
			Clay__MeasureTextCacheItemArray_Set(&context.measureTextHashMapInternal, elementIndex, Clay__MeasureTextCacheItem{measuredWordsStartIndex: -1})
			Clay__int32_tArray_Add(&context.measureTextHashMapInternalFreeList, elementIndex)
			if int(elementIndexPrevious) == 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(context.measureTextHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = nextIndex
			} else {
				var previousHashEntry *Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItemArray_Get(&context.measureTextHashMapInternal, elementIndexPrevious)
				previousHashEntry.nextIndex = nextIndex
			}
			elementIndex = nextIndex
		} else {
			elementIndexPrevious = elementIndex
			elementIndex = hashEntry.nextIndex
		}
	}
	var newItemIndex int32 = 0
	var newCacheItem Clay__MeasureTextCacheItem = Clay__MeasureTextCacheItem{measuredWordsStartIndex: -1, id: id, generation: context.generation}
	var measured *Clay__MeasureTextCacheItem = nil
	if int(context.measureTextHashMapInternalFreeList.length) > 0 {
		newItemIndex = Clay__int32_tArray_Get(&context.measureTextHashMapInternalFreeList, int32(int(context.measureTextHashMapInternalFreeList.length)-1))
		context.measureTextHashMapInternalFreeList.length--
		Clay__MeasureTextCacheItemArray_Set(&context.measureTextHashMapInternal, newItemIndex, newCacheItem)
		measured = Clay__MeasureTextCacheItemArray_Get(&context.measureTextHashMapInternal, newItemIndex)
	} else {
		if int(context.measureTextHashMapInternal.length) == int(context.measureTextHashMapInternal.capacity)-1 {
			if context.booleanWarnings.maxTextMeasureCacheExceeded {
				context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay ran out of capacity while attempting to measure text elements. Try using Clay_SetMaxElementCount() with a higher value.")}, userData: context.errorHandler.userData})
				context.booleanWarnings.maxTextMeasureCacheExceeded = true
			}
			return &CLAY__MEASURE_TEXT_CACHE_ITEM_DEFAULT
		}
		measured = Clay__MeasureTextCacheItemArray_Add(&context.measureTextHashMapInternal, newCacheItem)
		newItemIndex = int32(int(context.measureTextHashMapInternal.length) - 1)
	}
	var start int32 = 0
	var end int32 = 0
	var lineWidth float32 = 0
	var measuredWidth float32 = 0
	var measuredHeight float32 = 0
	var spaceWidth float32 = Clay__MeasureText(&CLAY__SPACECHAR, config).width
	var tempWord Clay__MeasuredWord = Clay__MeasuredWord{next: -1}
	var previousWord *Clay__MeasuredWord = &tempWord
	for int(end) < int(text.length) {
		if int(context.measuredWords.length) == int(context.measuredWords.capacity)-1 {
			if !context.booleanWarnings.maxTextMeasureCacheExceeded {
				context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_TEXT_MEASUREMENT_CAPACITY_EXCEEDED, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay has run out of space in it's internal text measurement cache. Try using Clay_SetMaxMeasureTextCacheWordCount() (default 16384, with 1 unit storing 1 measured word).")}, userData: context.errorHandler.userData})
				context.booleanWarnings.maxTextMeasureCacheExceeded = true
			}
			return &CLAY__MEASURE_TEXT_CACHE_ITEM_DEFAULT
		}
		var current int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(text.chars), end)))
		if int(current) == ' ' || int(current) == '\n' {
			var (
				length     int32           = int32(int(end) - int(start))
				word       Clay_String     = Clay_String{length: length, chars: (*byte)(unsafe.Add(unsafe.Pointer(text.chars), start))}
				dimensions Clay_Dimensions = Clay__MeasureText(&word, config)
			)
			if measuredHeight > dimensions.height {
				measuredHeight = measuredHeight
			} else {
				measuredHeight = dimensions.height
			}
			if int(current) == ' ' {
				dimensions.width += spaceWidth
				previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{startOffset: start, length: int32(int(length) + 1), width: dimensions.width, next: -1}, previousWord)
				lineWidth += dimensions.width
			}
			if int(current) == '\n' {
				if int(length) > 0 {
					previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{startOffset: start, length: length, width: dimensions.width, next: -1}, previousWord)
				}
				previousWord = Clay__AddMeasuredWord(Clay__MeasuredWord{startOffset: int32(int(end) + 1), length: 0, width: 0, next: -1}, previousWord)
				lineWidth += dimensions.width
				if lineWidth > measuredWidth {
					measuredWidth = lineWidth
				} else {
					measuredWidth = measuredWidth
				}
				measured.containsNewlines = true
				lineWidth = 0
			}
			start = int32(int(end) + 1)
		}
		end++
	}
	if int(end)-int(start) > 0 {
		var (
			lastWord   Clay_String     = Clay_String{length: int32(int(end) - int(start)), chars: (*byte)(unsafe.Add(unsafe.Pointer(text.chars), start))}
			dimensions Clay_Dimensions = Clay__MeasureText(&lastWord, config)
		)
		Clay__AddMeasuredWord(Clay__MeasuredWord{startOffset: start, length: int32(int(end) - int(start)), width: dimensions.width, next: -1}, previousWord)
		lineWidth += dimensions.width
		if measuredHeight > dimensions.height {
			measuredHeight = measuredHeight
		} else {
			measuredHeight = dimensions.height
		}
	}
	if lineWidth > measuredWidth {
		measuredWidth = lineWidth
	} else {
		measuredWidth = measuredWidth
	}
	measured.measuredWordsStartIndex = tempWord.next
	measured.unwrappedDimensions.width = measuredWidth
	measured.unwrappedDimensions.height = measuredHeight
	if int(elementIndexPrevious) != 0 {
		Clay__MeasureTextCacheItemArray_Get(&context.measureTextHashMapInternal, elementIndexPrevious).nextIndex = newItemIndex
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.measureTextHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = newItemIndex
	}
	return measured
}
func Clay__PointIsInsideRect(point Clay_Vector2, rect Clay_BoundingBox) bool {
	return point.x >= rect.x && point.x <= rect.x+rect.width && point.y >= rect.y && point.y <= rect.y+rect.height
}
func Clay__AddHashMapItem(elementId Clay_ElementId, layoutElement *Clay_LayoutElement) *Clay_LayoutElementHashMapItem {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int(context.layoutElementsHashMapInternal.length) == int(context.layoutElementsHashMapInternal.capacity)-1 {
		return nil
	}
	var item Clay_LayoutElementHashMapItem = Clay_LayoutElementHashMapItem{elementId: elementId, layoutElement: layoutElement, nextIndex: -1, generation: uint32(int32(int(context.generation) + 1))}
	var hashBucket uint32 = uint32(int32(int(elementId.id) % int(context.layoutElementsHashMap.capacity)))
	var hashItemPrevious int32 = -1
	var hashItemIndex int32 = *(*int32)(unsafe.Add(unsafe.Pointer(context.layoutElementsHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	for int(hashItemIndex) != -1 {
		var hashItem *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Get(&context.layoutElementsHashMapInternal, hashItemIndex)
		if int(hashItem.elementId.id) == int(elementId.id) {
			item.nextIndex = hashItem.nextIndex
			if int(hashItem.generation) <= int(context.generation) {
				hashItem.generation = uint32(int32(int(context.generation) + 1))
				hashItem.layoutElement = layoutElement
				hashItem.debugData.collision = false
			} else {
				context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_DUPLICATE_ID, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("An element with this ID was already previously declared during this layout.")}, userData: context.errorHandler.userData})
				if context.debugModeEnabled {
					hashItem.debugData.collision = true
				}
			}
			return hashItem
		}
		hashItemPrevious = hashItemIndex
		hashItemIndex = hashItem.nextIndex
	}
	var hashItem *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Add(&context.layoutElementsHashMapInternal, item)
	hashItem.debugData = Clay__DebugElementDataArray_Add(&context.debugElementData, Clay__DebugElementData{collision: false})
	if int(hashItemPrevious) != -1 {
		Clay__LayoutElementHashMapItemArray_Get(&context.layoutElementsHashMapInternal, hashItemPrevious).nextIndex = int32(int(context.layoutElementsHashMapInternal.length) - 1)
	} else {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.layoutElementsHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket))) = int32(int(context.layoutElementsHashMapInternal.length) - 1)
	}
	return hashItem
}
func Clay__GetHashMapItem(id uint32) *Clay_LayoutElementHashMapItem {
	var (
		context      *Clay_Context = Clay_GetCurrentContext()
		hashBucket   uint32        = uint32(int32(int(id) % int(context.layoutElementsHashMap.capacity)))
		elementIndex int32         = *(*int32)(unsafe.Add(unsafe.Pointer(context.layoutElementsHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(hashBucket)))
	)
	for int(elementIndex) != -1 {
		var hashEntry *Clay_LayoutElementHashMapItem = Clay__LayoutElementHashMapItemArray_Get(&context.layoutElementsHashMapInternal, elementIndex)
		if int(hashEntry.elementId.id) == int(id) {
			return hashEntry
		}
		elementIndex = hashEntry.nextIndex
	}
	return &CLAY__LAYOUT_ELEMENT_HASH_MAP_ITEM_DEFAULT
}
func Clay__GenerateIdForAnonymousElement(openLayoutElement *Clay_LayoutElement) {
	var (
		context       *Clay_Context       = Clay_GetCurrentContext()
		parentElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-2)))
		elementId     Clay_ElementId      = Clay__HashNumber(uint32(parentElement.childrenOrTextContent.children.length), parentElement.id)
	)
	openLayoutElement.id = elementId.id
	Clay__AddHashMapItem(elementId, openLayoutElement)
	Clay__StringArray_Add(&context.layoutElementIdStrings, elementId.stringId)
}
func Clay__ElementPostConfiguration() {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	if int(openLayoutElement.id) == 0 {
		Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	if openLayoutElement.layoutConfig == nil {
		openLayoutElement.layoutConfig = &CLAY_LAYOUT_DEFAULT
	}
	openLayoutElement.elementConfigs.internalArray = (*Clay_ElementConfig)(unsafe.Add(unsafe.Pointer(context.elementConfigs.internalArray), unsafe.Sizeof(Clay_ElementConfig{})*uintptr(context.elementConfigs.length)))
	for elementConfigIndex := int32(0); int(elementConfigIndex) < int(openLayoutElement.elementConfigs.length); elementConfigIndex++ {
		var config *Clay_ElementConfig = Clay__ElementConfigArray_Add(&context.elementConfigs, *Clay__ElementConfigArray_Get(&context.elementConfigBuffer, int32(int(context.elementConfigBuffer.length)-int(openLayoutElement.elementConfigs.length)+int(elementConfigIndex))))
		openLayoutElement.configsEnabled |= uint32(int32(config.type_))
		switch config.type_ {
		case CLAY__ELEMENT_CONFIG_TYPE_RECTANGLE:
			fallthrough
		case CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER:
		case CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER:
			var (
				floatingConfig     *Clay_FloatingElementConfig = config.config.floatingElementConfig
				hierarchicalParent *Clay_LayoutElement         = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-2)))
			)
			if hierarchicalParent == nil {
				break
			}
			var clipElementId uint32 = 0
			if int(floatingConfig.parentId) == 0 {
				var newConfig Clay_FloatingElementConfig = *floatingConfig
				newConfig.parentId = hierarchicalParent.id
				floatingConfig = Clay__FloatingElementConfigArray_Add(&context.floatingElementConfigs, newConfig)
				config.config.floatingElementConfig = floatingConfig
				if int(context.openClipElementStack.length) > 0 {
					clipElementId = uint32(Clay__int32_tArray_Get(&context.openClipElementStack, int32(int(context.openClipElementStack.length)-1)))
				}
			} else {
				var parentItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(floatingConfig.parentId)
				if parentItem == nil {
					context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_FLOATING_CONTAINER_PARENT_NOT_FOUND, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("A floating element was declared with a parentId, but no element with that ID was found.")}, userData: context.errorHandler.userData})
				} else {
					clipElementId = uint32(Clay__int32_tArray_Get(&context.layoutElementClipElementIds, int32(int64(uintptr(unsafe.Pointer(parentItem.layoutElement))-uintptr(unsafe.Pointer(context.layoutElements.internalArray))))))
				}
			}
			Clay__LayoutElementTreeRootArray_Add(&context.layoutElementTreeRoots, Clay__LayoutElementTreeRoot{layoutElementIndex: Clay__int32_tArray_Get(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-1)), parentId: floatingConfig.parentId, clipElementId: clipElementId, zIndex: int32(floatingConfig.zIndex)})
		case CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER:
			Clay__int32_tArray_Add(&context.openClipElementStack, int32(int(openLayoutElement.id)))
			var scrollOffset *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
			for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
				var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
				if int(openLayoutElement.id) == int(mapping.elementId) {
					scrollOffset = mapping
					scrollOffset.layoutElement = openLayoutElement
					scrollOffset.openThisFrame = true
				}
			}
			if scrollOffset == nil {
				scrollOffset = Clay__ScrollContainerDataInternalArray_Add(&context.scrollContainerDatas, Clay__ScrollContainerDataInternal{layoutElement: openLayoutElement, scrollOrigin: Clay_Vector2{x: -1, y: -1}, elementId: openLayoutElement.id, openThisFrame: true})
			}
			if context.externalScrollHandlingEnabled {
				scrollOffset.scrollPosition = Clay__QueryScrollOffset(scrollOffset.elementId)
			}
		case CLAY__ELEMENT_CONFIG_TYPE_CUSTOM:
		case CLAY__ELEMENT_CONFIG_TYPE_IMAGE:
			Clay__LayoutElementPointerArray_Add(&context.imageElementPointers, openLayoutElement)
		case CLAY__ELEMENT_CONFIG_TYPE_TEXT:
			fallthrough
		default:
		}
	}
	context.elementConfigBuffer.length -= openLayoutElement.elementConfigs.length
}
func Clay__CloseElement() {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	var layoutConfig *Clay_LayoutConfig = openLayoutElement.layoutConfig
	var elementHasScrollHorizontal bool = false
	var elementHasScrollVertical bool = false
	if Clay__ElementHasConfig(openLayoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER) {
		var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(openLayoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
		elementHasScrollHorizontal = scrollConfig.horizontal
		elementHasScrollVertical = scrollConfig.vertical
		context.openClipElementStack.length--
	}
	openLayoutElement.childrenOrTextContent.children.elements = (*int32)(unsafe.Add(unsafe.Pointer(context.layoutElementChildren.internalArray), unsafe.Sizeof(int32(0))*uintptr(context.layoutElementChildren.length)))
	if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
		openLayoutElement.dimensions.width = float32(int(layoutConfig.padding.left) + int(layoutConfig.padding.right))
		for i := int32(0); int(i) < int(openLayoutElement.childrenOrTextContent.children.length); i++ {
			var (
				childIndex int32               = Clay__int32_tArray_Get(&context.layoutElementChildrenBuffer, int32(int(context.layoutElementChildrenBuffer.length)-int(openLayoutElement.childrenOrTextContent.children.length)+int(i)))
				child      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, childIndex)
			)
			openLayoutElement.dimensions.width += child.dimensions.width
			if openLayoutElement.dimensions.height > (child.dimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)) {
				openLayoutElement.dimensions.height = openLayoutElement.dimensions.height
			} else {
				openLayoutElement.dimensions.height = child.dimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)
			}
			if !elementHasScrollHorizontal {
				openLayoutElement.minDimensions.width += child.minDimensions.width
			}
			if !elementHasScrollVertical {
				if openLayoutElement.minDimensions.height > (child.minDimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)) {
					openLayoutElement.minDimensions.height = openLayoutElement.minDimensions.height
				} else {
					openLayoutElement.minDimensions.height = child.minDimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)
				}
			}
			Clay__int32_tArray_Add(&context.layoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int {
			if (int(openLayoutElement.childrenOrTextContent.children.length) - 1) > 0 {
				return int(openLayoutElement.childrenOrTextContent.children.length) - 1
			}
			return 0
		}()) * int(layoutConfig.childGap))
		openLayoutElement.dimensions.width += childGap
		openLayoutElement.minDimensions.width += childGap
	} else if layoutConfig.layoutDirection == CLAY_TOP_TO_BOTTOM {
		openLayoutElement.dimensions.height = float32(int(layoutConfig.padding.top) + int(layoutConfig.padding.bottom))
		for i := int32(0); int(i) < int(openLayoutElement.childrenOrTextContent.children.length); i++ {
			var (
				childIndex int32               = Clay__int32_tArray_Get(&context.layoutElementChildrenBuffer, int32(int(context.layoutElementChildrenBuffer.length)-int(openLayoutElement.childrenOrTextContent.children.length)+int(i)))
				child      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, childIndex)
			)
			openLayoutElement.dimensions.height += child.dimensions.height
			if openLayoutElement.dimensions.width > (child.dimensions.width + float32(layoutConfig.padding.left) + float32(layoutConfig.padding.right)) {
				openLayoutElement.dimensions.width = openLayoutElement.dimensions.width
			} else {
				openLayoutElement.dimensions.width = child.dimensions.width + float32(layoutConfig.padding.left) + float32(layoutConfig.padding.right)
			}
			if !elementHasScrollVertical {
				openLayoutElement.minDimensions.height += child.minDimensions.height
			}
			if !elementHasScrollHorizontal {
				if openLayoutElement.minDimensions.width > (child.minDimensions.width + float32(layoutConfig.padding.left) + float32(layoutConfig.padding.right)) {
					openLayoutElement.minDimensions.width = openLayoutElement.minDimensions.width
				} else {
					openLayoutElement.minDimensions.width = child.minDimensions.width + float32(layoutConfig.padding.left) + float32(layoutConfig.padding.right)
				}
			}
			Clay__int32_tArray_Add(&context.layoutElementChildren, childIndex)
		}
		var childGap float32 = float32((func() int {
			if (int(openLayoutElement.childrenOrTextContent.children.length) - 1) > 0 {
				return int(openLayoutElement.childrenOrTextContent.children.length) - 1
			}
			return 0
		}()) * int(layoutConfig.childGap))
		openLayoutElement.dimensions.height += childGap
		openLayoutElement.minDimensions.height += childGap
	}
	context.layoutElementChildrenBuffer.length -= int32(openLayoutElement.childrenOrTextContent.children.length)
	if layoutConfig.sizing.width.type_ != CLAY__SIZING_TYPE_PERCENT {
		if layoutConfig.sizing.width.size.minMax.max <= 0 {
			layoutConfig.sizing.width.size.minMax.max = CLAY__MAXFLOAT
		}
		if (func() float32 {
			if openLayoutElement.dimensions.width > layoutConfig.sizing.width.size.minMax.min {
				return openLayoutElement.dimensions.width
			}
			return layoutConfig.sizing.width.size.minMax.min
		}()) < layoutConfig.sizing.width.size.minMax.max {
			if openLayoutElement.dimensions.width > layoutConfig.sizing.width.size.minMax.min {
				openLayoutElement.dimensions.width = openLayoutElement.dimensions.width
			} else {
				openLayoutElement.dimensions.width = layoutConfig.sizing.width.size.minMax.min
			}
		} else {
			openLayoutElement.dimensions.width = layoutConfig.sizing.width.size.minMax.max
		}
		if (func() float32 {
			if openLayoutElement.minDimensions.width > layoutConfig.sizing.width.size.minMax.min {
				return openLayoutElement.minDimensions.width
			}
			return layoutConfig.sizing.width.size.minMax.min
		}()) < layoutConfig.sizing.width.size.minMax.max {
			if openLayoutElement.minDimensions.width > layoutConfig.sizing.width.size.minMax.min {
				openLayoutElement.minDimensions.width = openLayoutElement.minDimensions.width
			} else {
				openLayoutElement.minDimensions.width = layoutConfig.sizing.width.size.minMax.min
			}
		} else {
			openLayoutElement.minDimensions.width = layoutConfig.sizing.width.size.minMax.max
		}
	} else {
		openLayoutElement.dimensions.width = 0
	}
	if layoutConfig.sizing.height.type_ != CLAY__SIZING_TYPE_PERCENT {
		if layoutConfig.sizing.height.size.minMax.max <= 0 {
			layoutConfig.sizing.height.size.minMax.max = CLAY__MAXFLOAT
		}
		if (func() float32 {
			if openLayoutElement.dimensions.height > layoutConfig.sizing.height.size.minMax.min {
				return openLayoutElement.dimensions.height
			}
			return layoutConfig.sizing.height.size.minMax.min
		}()) < layoutConfig.sizing.height.size.minMax.max {
			if openLayoutElement.dimensions.height > layoutConfig.sizing.height.size.minMax.min {
				openLayoutElement.dimensions.height = openLayoutElement.dimensions.height
			} else {
				openLayoutElement.dimensions.height = layoutConfig.sizing.height.size.minMax.min
			}
		} else {
			openLayoutElement.dimensions.height = layoutConfig.sizing.height.size.minMax.max
		}
		if (func() float32 {
			if openLayoutElement.minDimensions.height > layoutConfig.sizing.height.size.minMax.min {
				return openLayoutElement.minDimensions.height
			}
			return layoutConfig.sizing.height.size.minMax.min
		}()) < layoutConfig.sizing.height.size.minMax.max {
			if openLayoutElement.minDimensions.height > layoutConfig.sizing.height.size.minMax.min {
				openLayoutElement.minDimensions.height = openLayoutElement.minDimensions.height
			} else {
				openLayoutElement.minDimensions.height = layoutConfig.sizing.height.size.minMax.min
			}
		} else {
			openLayoutElement.minDimensions.height = layoutConfig.sizing.height.size.minMax.max
		}
	} else {
		openLayoutElement.dimensions.height = 0
	}
	var elementIsFloating bool = Clay__ElementHasConfig(openLayoutElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER)
	var closingElementIndex int32 = Clay__int32_tArray_RemoveSwapback(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-1))
	openLayoutElement = Clay__GetOpenLayoutElement()
	if !elementIsFloating && int(context.openLayoutElementStack.length) > 1 {
		openLayoutElement.childrenOrTextContent.children.length++
		Clay__int32_tArray_Add(&context.layoutElementChildrenBuffer, closingElementIndex)
	}
}
func Clay__OpenElement() {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int(context.layoutElements.length) == int(context.layoutElements.capacity)-1 || context.booleanWarnings.maxElementsExceeded {
		context.booleanWarnings.maxElementsExceeded = true
		return
	}
	var layoutElement Clay_LayoutElement = Clay_LayoutElement{}
	Clay_LayoutElementArray_Add(&context.layoutElements, layoutElement)
	Clay__int32_tArray_Add(&context.openLayoutElementStack, int32(int(context.layoutElements.length)-1))
	if int(context.openClipElementStack.length) > 0 {
		Clay__int32_tArray_Set(&context.layoutElementClipElementIds, int32(int(context.layoutElements.length)-1), Clay__int32_tArray_Get(&context.openClipElementStack, int32(int(context.openClipElementStack.length)-1)))
	} else {
		Clay__int32_tArray_Set(&context.layoutElementClipElementIds, int32(int(context.layoutElements.length)-1), 0)
	}
}
func Clay__OpenTextElement(text Clay_String, textConfig *Clay_TextElementConfig) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int(context.layoutElements.length) == int(context.layoutElements.capacity)-1 || context.booleanWarnings.maxElementsExceeded {
		context.booleanWarnings.maxElementsExceeded = true
		return
	}
	var parentElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	parentElement.childrenOrTextContent.children.length++
	Clay__OpenElement()
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	Clay__int32_tArray_Add(&context.layoutElementChildrenBuffer, int32(int(context.layoutElements.length)-1))
	var textMeasured *Clay__MeasureTextCacheItem = Clay__MeasureTextCached(&text, textConfig)
	var elementId Clay_ElementId = Clay__HashString(Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Text")}, uint32(parentElement.childrenOrTextContent.children.length), parentElement.id)
	openLayoutElement.id = elementId.id
	Clay__AddHashMapItem(elementId, openLayoutElement)
	Clay__StringArray_Add(&context.layoutElementIdStrings, elementId.stringId)
	var textDimensions Clay_Dimensions = Clay_Dimensions{width: textMeasured.unwrappedDimensions.width, height: func() float32 {
		if int(textConfig.lineHeight) > 0 {
			return float32(textConfig.lineHeight)
		}
		return textMeasured.unwrappedDimensions.height
	}()}
	openLayoutElement.dimensions = textDimensions
	openLayoutElement.minDimensions = Clay_Dimensions{width: textMeasured.unwrappedDimensions.height, height: textDimensions.height}
	openLayoutElement.childrenOrTextContent.textElementData = Clay__TextElementDataArray_Add(&context.textElementData, Clay__TextElementData{text: text, preferredDimensions: textMeasured.unwrappedDimensions, elementIndex: int32(int(context.layoutElements.length) - 1)})
	openLayoutElement.elementConfigs = Clay__ElementConfigArraySlice{length: 1, internalArray: Clay__ElementConfigArray_Add(&context.elementConfigs, Clay_ElementConfig{type_: CLAY__ELEMENT_CONFIG_TYPE_TEXT, config: Clay_ElementConfigUnion{textElementConfig: textConfig}})}
	openLayoutElement.configsEnabled |= uint32(int32(CLAY__ELEMENT_CONFIG_TYPE_TEXT))
	openLayoutElement.layoutConfig = &CLAY_LAYOUT_DEFAULT
	Clay__int32_tArray_RemoveSwapback(&context.openLayoutElementStack, int32(int(context.openLayoutElementStack.length)-1))
}
func Clay__InitializeEphemeralMemory(context *Clay_Context) {
	var (
		maxElementCount int32       = context.maxElementCount
		arena           *Clay_Arena = &context.internalArena
	)
	arena.nextAllocation = context.arenaResetOffset
	context.layoutElementChildrenBuffer = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElements = Clay_LayoutElementArray_Allocate_Arena(maxElementCount, arena)
	context.warnings = Clay__WarningArray_Allocate_Arena(100, arena)
	context.layoutConfigs = Clay__LayoutConfigArray_Allocate_Arena(maxElementCount, arena)
	context.elementConfigBuffer = Clay__ElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.elementConfigs = Clay__ElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.rectangleElementConfigs = Clay__RectangleElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.textElementConfigs = Clay__TextElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.imageElementConfigs = Clay__ImageElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.floatingElementConfigs = Clay__FloatingElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.scrollElementConfigs = Clay__ScrollElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.customElementConfigs = Clay__CustomElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.borderElementConfigs = Clay__BorderElementConfigArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementIdStrings = Clay__StringArray_Allocate_Arena(maxElementCount, arena)
	context.wrappedTextLines = Clay__WrappedTextLineArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementTreeNodeArray1 = Clay__LayoutElementTreeNodeArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementTreeRoots = Clay__LayoutElementTreeRootArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementChildren = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.openLayoutElementStack = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.textElementData = Clay__TextElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.imageElementPointers = Clay__LayoutElementPointerArray_Allocate_Arena(maxElementCount, arena)
	context.renderCommands = Clay_RenderCommandArray_Allocate_Arena(maxElementCount, arena)
	context.treeNodeVisited = Clay__BoolArray_Allocate_Arena(maxElementCount, arena)
	context.treeNodeVisited.length = context.treeNodeVisited.capacity
	context.openClipElementStack = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.reusableElementIndexBuffer = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementClipElementIds = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.dynamicStringData = Clay__CharArray_Allocate_Arena(maxElementCount, arena)
}
func Clay__InitializePersistentMemory(context *Clay_Context) {
	var (
		maxElementCount              int32       = context.maxElementCount
		maxMeasureTextCacheWordCount int32       = context.maxMeasureTextCacheWordCount
		arena                        *Clay_Arena = &context.internalArena
	)
	context.scrollContainerDatas = Clay__ScrollContainerDataInternalArray_Allocate_Arena(10, arena)
	context.layoutElementsHashMapInternal = Clay__LayoutElementHashMapItemArray_Allocate_Arena(maxElementCount, arena)
	context.layoutElementsHashMap = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.measureTextHashMapInternal = Clay__MeasureTextCacheItemArray_Allocate_Arena(maxElementCount, arena)
	context.measureTextHashMapInternalFreeList = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.measuredWordsFreeList = Clay__int32_tArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.measureTextHashMap = Clay__int32_tArray_Allocate_Arena(maxElementCount, arena)
	context.measuredWords = Clay__MeasuredWordArray_Allocate_Arena(maxMeasureTextCacheWordCount, arena)
	context.pointerOverIds = Clay__ElementIdArray_Allocate_Arena(maxElementCount, arena)
	context.debugElementData = Clay__DebugElementDataArray_Allocate_Arena(maxElementCount, arena)
	context.arenaResetOffset = arena.nextAllocation
}
func Clay__CompressChildrenAlongAxis(xAxis bool, totalSizeToDistribute float32, resizableContainerBuffer Clay__int32_tArray) {
	var (
		context           *Clay_Context      = Clay_GetCurrentContext()
		largestContainers Clay__int32_tArray = context.openClipElementStack
	)
	largestContainers.length = 0
	for totalSizeToDistribute > 0.1 {
		var (
			largestSize float32 = 0
			targetSize  float32 = 0
		)
		for i := int32(0); int(i) < int(resizableContainerBuffer.length); i++ {
			var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&resizableContainerBuffer, i))
			if !xAxis && Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE) {
				continue
			}
			var childSize float32
			if xAxis {
				childSize = childElement.dimensions.width
			} else {
				childSize = childElement.dimensions.height
			}
			if (childSize-largestSize) < 0.1 && (childSize-largestSize) > -0.1 {
				Clay__int32_tArray_Add(&largestContainers, Clay__int32_tArray_Get(&resizableContainerBuffer, i))
			} else if childSize > largestSize {
				targetSize = largestSize
				largestSize = childSize
				largestContainers.length = 0
				Clay__int32_tArray_Add(&largestContainers, Clay__int32_tArray_Get(&resizableContainerBuffer, i))
			} else if childSize > targetSize {
				targetSize = childSize
			}
		}
		targetSize = (func() float32 {
			if targetSize > ((largestSize * float32(largestContainers.length)) - totalSizeToDistribute) {
				return targetSize
			}
			return (largestSize * float32(largestContainers.length)) - totalSizeToDistribute
		}()) / float32(largestContainers.length)
		for childOffset := int32(0); int(childOffset) < int(largestContainers.length); childOffset++ {
			var (
				childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&largestContainers, childOffset))
				childSize    *float32
			)
			if xAxis {
				childSize = &childElement.dimensions.width
			} else {
				childSize = &childElement.dimensions.height
			}
			var childMinSize float32
			if xAxis {
				childMinSize = childElement.minDimensions.width
			} else {
				childMinSize = childElement.minDimensions.height
			}
			var oldChildSize float32 = *childSize
			if childMinSize > targetSize {
				*childSize = childMinSize
			} else {
				*childSize = targetSize
			}
			totalSizeToDistribute -= oldChildSize - *childSize
			if *childSize == childMinSize {
				Clay__int32_tArray_RemoveSwapback(&largestContainers, childOffset)
				childOffset--
			}
		}
		if int(largestContainers.length) == 0 {
			break
		}
	}
}
func Clay__SizeContainersAlongAxis(xAxis bool) {
	var (
		context                  *Clay_Context      = Clay_GetCurrentContext()
		bfsBuffer                Clay__int32_tArray = context.layoutElementChildrenBuffer
		resizableContainerBuffer Clay__int32_tArray = context.openLayoutElementStack
	)
	for rootIndex := int32(0); int(rootIndex) < int(context.layoutElementTreeRoots.length); rootIndex++ {
		bfsBuffer.length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, rootIndex)
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, int32(int(root.layoutElementIndex)))
		Clay__int32_tArray_Add(&bfsBuffer, root.layoutElementIndex)
		if Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER) {
			var (
				floatingElementConfig *Clay_FloatingElementConfig    = Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER).floatingElementConfig
				parentItem            *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(floatingElementConfig.parentId)
			)
			if parentItem != nil {
				var parentLayoutElement *Clay_LayoutElement = parentItem.layoutElement
				if rootElement.layoutConfig.sizing.width.type_ == CLAY__SIZING_TYPE_GROW {
					rootElement.dimensions.width = parentLayoutElement.dimensions.width
				}
				if rootElement.layoutConfig.sizing.height.type_ == CLAY__SIZING_TYPE_GROW {
					rootElement.dimensions.height = parentLayoutElement.dimensions.height
				}
			}
		}
		if (func() float32 {
			if rootElement.dimensions.width > rootElement.layoutConfig.sizing.width.size.minMax.min {
				return rootElement.dimensions.width
			}
			return rootElement.layoutConfig.sizing.width.size.minMax.min
		}()) < rootElement.layoutConfig.sizing.width.size.minMax.max {
			if rootElement.dimensions.width > rootElement.layoutConfig.sizing.width.size.minMax.min {
				rootElement.dimensions.width = rootElement.dimensions.width
			} else {
				rootElement.dimensions.width = rootElement.layoutConfig.sizing.width.size.minMax.min
			}
		} else {
			rootElement.dimensions.width = rootElement.layoutConfig.sizing.width.size.minMax.max
		}
		if (func() float32 {
			if rootElement.dimensions.height > rootElement.layoutConfig.sizing.height.size.minMax.min {
				return rootElement.dimensions.height
			}
			return rootElement.layoutConfig.sizing.height.size.minMax.min
		}()) < rootElement.layoutConfig.sizing.height.size.minMax.max {
			if rootElement.dimensions.height > rootElement.layoutConfig.sizing.height.size.minMax.min {
				rootElement.dimensions.height = rootElement.dimensions.height
			} else {
				rootElement.dimensions.height = rootElement.layoutConfig.sizing.height.size.minMax.min
			}
		} else {
			rootElement.dimensions.height = rootElement.layoutConfig.sizing.height.size.minMax.max
		}
		for i := int32(0); int(i) < int(bfsBuffer.length); i++ {
			var (
				parentIndex        int32               = Clay__int32_tArray_Get(&bfsBuffer, i)
				parent             *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, parentIndex)
				parentStyleConfig  *Clay_LayoutConfig  = parent.layoutConfig
				growContainerCount int32               = 0
				parentSize         float32
			)
			if xAxis {
				parentSize = parent.dimensions.width
			} else {
				parentSize = parent.dimensions.height
			}
			var parentPadding float32 = float32(func() int {
				if xAxis {
					return int(parent.layoutConfig.padding.left) + int(parent.layoutConfig.padding.right)
				}
				return int(parent.layoutConfig.padding.top) + int(parent.layoutConfig.padding.bottom)
			}())
			var innerContentSize float32 = 0
			var growContainerContentSize float32 = 0
			var totalPaddingAndChildGaps float32 = parentPadding
			var sizingAlongAxis bool = xAxis && parentStyleConfig.layoutDirection == CLAY_LEFT_TO_RIGHT || !xAxis && parentStyleConfig.layoutDirection == CLAY_TOP_TO_BOTTOM
			resizableContainerBuffer.length = 0
			var parentChildGap float32 = float32(parentStyleConfig.childGap)
			for childOffset := int32(0); int(childOffset) < int(parent.childrenOrTextContent.children.length); childOffset++ {
				var (
					childElementIndex int32               = *(*int32)(unsafe.Add(unsafe.Pointer(parent.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, childElementIndex)
					childSizing       Clay_SizingAxis
				)
				if xAxis {
					childSizing = childElement.layoutConfig.sizing.width
				} else {
					childSizing = childElement.layoutConfig.sizing.height
				}
				var childSize float32
				if xAxis {
					childSize = childElement.dimensions.width
				} else {
					childSize = childElement.dimensions.height
				}
				if !Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) && int(childElement.childrenOrTextContent.children.length) > 0 {
					Clay__int32_tArray_Add(&bfsBuffer, childElementIndex)
				}
				if childSizing.type_ != CLAY__SIZING_TYPE_PERCENT && childSizing.type_ != CLAY__SIZING_TYPE_FIXED && (!Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) || Clay__FindElementConfigWithType(childElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT).textElementConfig.wrapMode == CLAY_TEXT_WRAP_WORDS) {
					Clay__int32_tArray_Add(&resizableContainerBuffer, childElementIndex)
				}
				if sizingAlongAxis {
					if childSizing.type_ == CLAY__SIZING_TYPE_PERCENT {
						innerContentSize += 0
					} else {
						innerContentSize += childSize
					}
					if childSizing.type_ == CLAY__SIZING_TYPE_GROW {
						growContainerContentSize += childSize
						growContainerCount++
					}
					if int(childOffset) > 0 {
						innerContentSize += parentChildGap
						totalPaddingAndChildGaps += parentChildGap
					}
				} else {
					if childSize > innerContentSize {
						innerContentSize = childSize
					} else {
						innerContentSize = innerContentSize
					}
				}
			}
			for childOffset := int32(0); int(childOffset) < int(parent.childrenOrTextContent.children.length); childOffset++ {
				var (
					childElementIndex int32               = *(*int32)(unsafe.Add(unsafe.Pointer(parent.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(childOffset)))
					childElement      *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, childElementIndex)
					childSizing       Clay_SizingAxis
				)
				if xAxis {
					childSizing = childElement.layoutConfig.sizing.width
				} else {
					childSizing = childElement.layoutConfig.sizing.height
				}
				var childSize *float32
				if xAxis {
					childSize = &childElement.dimensions.width
				} else {
					childSize = &childElement.dimensions.height
				}
				if childSizing.type_ == CLAY__SIZING_TYPE_PERCENT {
					*childSize = (parentSize - totalPaddingAndChildGaps) * childSizing.size.percent
					if sizingAlongAxis {
						innerContentSize += *childSize
						if int(childOffset) > 0 {
							innerContentSize += parentChildGap
							totalPaddingAndChildGaps += parentChildGap
						}
					} else {
						if (*childSize) > innerContentSize {
							innerContentSize = *childSize
						} else {
							innerContentSize = innerContentSize
						}
					}
				}
			}
			if sizingAlongAxis {
				var sizeToDistribute float32 = parentSize - parentPadding - innerContentSize
				if sizeToDistribute < 0 {
					if Clay__ElementHasConfig(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER) {
						var scrollElementConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
						if xAxis && scrollElementConfig.horizontal || !xAxis && scrollElementConfig.vertical {
							continue
						}
					}
					Clay__CompressChildrenAlongAxis(xAxis, -sizeToDistribute, resizableContainerBuffer)
				} else if sizeToDistribute > 0 && int(growContainerCount) > 0 {
					var targetSize float32 = (sizeToDistribute + growContainerContentSize) / float32(growContainerCount)
					for childOffset := int32(0); int(childOffset) < int(resizableContainerBuffer.length); childOffset++ {
						var (
							childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&resizableContainerBuffer, childOffset))
							childSizing  Clay_SizingAxis
						)
						if xAxis {
							childSizing = childElement.layoutConfig.sizing.width
						} else {
							childSizing = childElement.layoutConfig.sizing.height
						}
						if childSizing.type_ == CLAY__SIZING_TYPE_GROW {
							var childSize *float32
							_ = childSize
							if xAxis {
								childSize = &childElement.dimensions.width
							} else {
								childSize = &childElement.dimensions.height
							}
							var minSize *float32
							if xAxis {
								minSize = &childElement.minDimensions.width
							} else {
								minSize = &childElement.minDimensions.height
							}
							if targetSize < *minSize {
								growContainerContentSize -= *minSize
								Clay__int32_tArray_RemoveSwapback(&resizableContainerBuffer, childOffset)
								growContainerCount--
								targetSize = (sizeToDistribute + growContainerContentSize) / float32(growContainerCount)
								childOffset = -1
								continue
							}
							*childSize = targetSize
						}
					}
				}
			} else {
				for childOffset := int32(0); int(childOffset) < int(resizableContainerBuffer.length); childOffset++ {
					var (
						childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&resizableContainerBuffer, childOffset))
						childSizing  Clay_SizingAxis
					)
					if xAxis {
						childSizing = childElement.layoutConfig.sizing.width
					} else {
						childSizing = childElement.layoutConfig.sizing.height
					}
					var childSize *float32
					if xAxis {
						childSize = &childElement.dimensions.width
					} else {
						childSize = &childElement.dimensions.height
					}
					if !xAxis && Clay__ElementHasConfig(childElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE) {
						continue
					}
					var maxSize float32 = parentSize - parentPadding
					if Clay__ElementHasConfig(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER) {
						var scrollElementConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(parent, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
						if xAxis && scrollElementConfig.horizontal || !xAxis && scrollElementConfig.vertical {
							if maxSize > innerContentSize {
								maxSize = maxSize
							} else {
								maxSize = innerContentSize
							}
						}
					}
					if childSizing.type_ == CLAY__SIZING_TYPE_FIT {
						if childSizing.size.minMax.min > (func() float32 {
							if (*childSize) < maxSize {
								return *childSize
							}
							return maxSize
						}()) {
							*childSize = childSizing.size.minMax.min
						} else if (*childSize) < maxSize {
							*childSize = *childSize
						} else {
							*childSize = maxSize
						}
					} else if childSizing.type_ == CLAY__SIZING_TYPE_GROW {
						if maxSize < childSizing.size.minMax.max {
							*childSize = maxSize
						} else {
							*childSize = childSizing.size.minMax.max
						}
					}
				}
			}
		}
	}
}
func Clay__IntToString(integer int32) Clay_String {
	if int(integer) == 0 {
		return Clay_String{length: 1, chars: libc.CString("0")}
	}
	var context *Clay_Context = Clay_GetCurrentContext()
	var chars *byte = (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(context.dynamicStringData.internalArray), context.dynamicStringData.length))))
	var length int32 = 0
	var sign int32 = integer
	if int(integer) < 0 {
		integer = -integer
	}
	for int(integer) > 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = byte(int8(int(integer)%10 + '0'))
		integer /= 10
	}
	if int(sign) < 0 {
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), func() int32 {
			p := &length
			x := *p
			*p++
			return x
		}())) = '-'
	}
	for j, k := int32(0), int32(int32(int(length)-1)); int(j) < int(k); func() int32 {
		j++
		return func() int32 {
			p := &k
			x := *p
			*p--
			return x
		}()
	}() {
		var temp int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(chars), j)))
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), j)) = *(*byte)(unsafe.Add(unsafe.Pointer(chars), k))
		*(*byte)(unsafe.Add(unsafe.Pointer(chars), k)) = byte(temp)
	}
	context.dynamicStringData.length += length
	return Clay_String{length: length, chars: chars}
}
func Clay__AddRenderCommand(renderCommand Clay_RenderCommand) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if int(context.renderCommands.length) < int(context.renderCommands.capacity)-1 {
		Clay_RenderCommandArray_Add(&context.renderCommands, renderCommand)
	} else {
		if !context.booleanWarnings.maxRenderCommandsExceeded {
			context.booleanWarnings.maxRenderCommandsExceeded = true
			context.errorHandler.errorHandlerFunction(Clay_ErrorData{errorType: CLAY_ERROR_TYPE_ELEMENTS_CAPACITY_EXCEEDED, errorText: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay ran out of capacity while attempting to create render commands. This is usually caused by a large amount of wrapping text elements while close to the max element capacity. Try using Clay_SetMaxElementCount() with a higher value.")}, userData: context.errorHandler.userData})
		}
	}
}
func Clay__ElementIsOffscreen(boundingBox *Clay_BoundingBox) bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.disableCulling {
		return false
	}
	return boundingBox.x > context.layoutDimensions.width || boundingBox.y > context.layoutDimensions.height || boundingBox.x+boundingBox.width < 0 || boundingBox.y+boundingBox.height < 0
}
func Clay__CalculateFinalLayout() {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__SizeContainersAlongAxis(true)
	for textElementIndex := int32(0); int(textElementIndex) < int(context.textElementData.length); textElementIndex++ {
		var textElementData *Clay__TextElementData = Clay__TextElementDataArray_Get(&context.textElementData, textElementIndex)
		textElementData.wrappedLines = Clay__WrappedTextLineArraySlice{length: 0, internalArray: (*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(context.wrappedTextLines.internalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(context.wrappedTextLines.length)))}
		var containerElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, int32(int(textElementData.elementIndex)))
		var textConfig *Clay_TextElementConfig = Clay__FindElementConfigWithType(containerElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT).textElementConfig
		var measureTextCacheItem *Clay__MeasureTextCacheItem = Clay__MeasureTextCached(&textElementData.text, textConfig)
		var lineWidth float32 = 0
		var lineHeight float32
		if int(textConfig.lineHeight) > 0 {
			lineHeight = float32(textConfig.lineHeight)
		} else {
			lineHeight = textElementData.preferredDimensions.height
		}
		var lineLengthChars int32 = 0
		var lineStartOffset int32 = 0
		if !measureTextCacheItem.containsNewlines && textElementData.preferredDimensions.width <= containerElement.dimensions.width {
			Clay__WrappedTextLineArray_Add(&context.wrappedTextLines, Clay__WrappedTextLine{dimensions: containerElement.dimensions, line: textElementData.text})
			textElementData.wrappedLines.length++
			continue
		}
		var wordIndex int32 = measureTextCacheItem.measuredWordsStartIndex
		for int(wordIndex) != -1 {
			if int(context.wrappedTextLines.length) > int(context.wrappedTextLines.capacity)-1 {
				break
			}
			var measuredWord *Clay__MeasuredWord = Clay__MeasuredWordArray_Get(&context.measuredWords, wordIndex)
			if int(lineLengthChars) == 0 && lineWidth+measuredWord.width > containerElement.dimensions.width {
				Clay__WrappedTextLineArray_Add(&context.wrappedTextLines, Clay__WrappedTextLine{dimensions: Clay_Dimensions{width: measuredWord.width, height: lineHeight}, line: Clay_String{length: measuredWord.length, chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.text.chars), measuredWord.startOffset))}})
				textElementData.wrappedLines.length++
				wordIndex = measuredWord.next
				lineStartOffset = int32(int(measuredWord.startOffset) + int(measuredWord.length))
			} else if int(measuredWord.length) == 0 || lineWidth+measuredWord.width > containerElement.dimensions.width {
				Clay__WrappedTextLineArray_Add(&context.wrappedTextLines, Clay__WrappedTextLine{dimensions: Clay_Dimensions{width: lineWidth, height: lineHeight}, line: Clay_String{length: lineLengthChars, chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.text.chars), lineStartOffset))}})
				textElementData.wrappedLines.length++
				if int(lineLengthChars) == 0 || int(measuredWord.length) == 0 {
					wordIndex = measuredWord.next
				}
				lineWidth = 0
				lineLengthChars = 0
				lineStartOffset = measuredWord.startOffset
			} else {
				lineWidth += measuredWord.width
				lineLengthChars += measuredWord.length
				wordIndex = measuredWord.next
			}
		}
		if int(lineLengthChars) > 0 {
			Clay__WrappedTextLineArray_Add(&context.wrappedTextLines, Clay__WrappedTextLine{dimensions: Clay_Dimensions{width: lineWidth, height: lineHeight}, line: Clay_String{length: lineLengthChars, chars: (*byte)(unsafe.Add(unsafe.Pointer(textElementData.text.chars), lineStartOffset))}})
			textElementData.wrappedLines.length++
		}
		containerElement.dimensions.height = lineHeight * float32(textElementData.wrappedLines.length)
	}
	for i := int32(0); int(i) < int(context.imageElementPointers.length); i++ {
		var (
			imageElement *Clay_LayoutElement      = Clay__LayoutElementPointerArray_Get(&context.imageElementPointers, i)
			config       *Clay_ImageElementConfig = Clay__FindElementConfigWithType(imageElement, CLAY__ELEMENT_CONFIG_TYPE_IMAGE).imageElementConfig
		)
		imageElement.dimensions.height = (config.sourceDimensions.height / (func() float32 {
			if config.sourceDimensions.width > 1 {
				return config.sourceDimensions.width
			}
			return 1
		}())) * imageElement.dimensions.width
	}
	var dfsBuffer Clay__LayoutElementTreeNodeArray = context.layoutElementTreeNodeArray1
	dfsBuffer.length = 0
	for i := int32(0); int(i) < int(context.layoutElementTreeRoots.length); i++ {
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, i)
		*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), dfsBuffer.length)) = false
		Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{layoutElement: Clay_LayoutElementArray_Get(&context.layoutElements, int32(int(root.layoutElementIndex)))})
	}
	for int(dfsBuffer.length) > 0 {
		var (
			currentElementTreeNode *Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNodeArray_Get(&dfsBuffer, int32(int(dfsBuffer.length)-1))
			currentElement         *Clay_LayoutElement          = currentElementTreeNode.layoutElement
		)
		if !*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) {
			*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) = true
			if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) || int(currentElement.childrenOrTextContent.children.length) == 0 {
				dfsBuffer.length--
				continue
			}
			for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), dfsBuffer.length)) = false
				Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{layoutElement: Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))})
			}
			continue
		}
		dfsBuffer.length--
		var layoutConfig *Clay_LayoutConfig = currentElement.layoutConfig
		if layoutConfig.sizing.height.type_ == CLAY__SIZING_TYPE_PERCENT {
			continue
		}
		if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
			for j := int32(0); int(j) < int(currentElement.childrenOrTextContent.children.length); j++ {
				var (
					childElement           *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(j))))
					childHeightWithPadding float32             = (func() float32 {
						if (childElement.dimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)) > currentElement.dimensions.height {
							return childElement.dimensions.height + float32(layoutConfig.padding.top) + float32(layoutConfig.padding.bottom)
						}
						return currentElement.dimensions.height
					}())
				)
				if (func() float32 {
					if childHeightWithPadding > layoutConfig.sizing.height.size.minMax.min {
						return childHeightWithPadding
					}
					return layoutConfig.sizing.height.size.minMax.min
				}()) < layoutConfig.sizing.height.size.minMax.max {
					if childHeightWithPadding > layoutConfig.sizing.height.size.minMax.min {
						currentElement.dimensions.height = childHeightWithPadding
					} else {
						currentElement.dimensions.height = layoutConfig.sizing.height.size.minMax.min
					}
				} else {
					currentElement.dimensions.height = layoutConfig.sizing.height.size.minMax.max
				}
			}
		} else if layoutConfig.layoutDirection == CLAY_TOP_TO_BOTTOM {
			var contentHeight float32 = float32(int(layoutConfig.padding.top) + int(layoutConfig.padding.bottom))
			for j := int32(0); int(j) < int(currentElement.childrenOrTextContent.children.length); j++ {
				var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(j))))
				contentHeight += childElement.dimensions.height
			}
			contentHeight += float32((func() int {
				if (int(currentElement.childrenOrTextContent.children.length) - 1) > 0 {
					return int(currentElement.childrenOrTextContent.children.length) - 1
				}
				return 0
			}()) * int(layoutConfig.childGap))
			if (func() float32 {
				if contentHeight > layoutConfig.sizing.height.size.minMax.min {
					return contentHeight
				}
				return layoutConfig.sizing.height.size.minMax.min
			}()) < layoutConfig.sizing.height.size.minMax.max {
				if contentHeight > layoutConfig.sizing.height.size.minMax.min {
					currentElement.dimensions.height = contentHeight
				} else {
					currentElement.dimensions.height = layoutConfig.sizing.height.size.minMax.min
				}
			} else {
				currentElement.dimensions.height = layoutConfig.sizing.height.size.minMax.max
			}
		}
	}
	Clay__SizeContainersAlongAxis(false)
	var sortMax int32 = int32(int(context.layoutElementTreeRoots.length) - 1)
	for int(sortMax) > 0 {
		for i := int32(0); int(i) < int(sortMax); i++ {
			var (
				current Clay__LayoutElementTreeRoot = *Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, i)
				next    Clay__LayoutElementTreeRoot = *Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, int32(int(i)+1))
			)
			if int(next.zIndex) < int(current.zIndex) {
				Clay__LayoutElementTreeRootArray_Set(&context.layoutElementTreeRoots, i, next)
				Clay__LayoutElementTreeRootArray_Set(&context.layoutElementTreeRoots, int32(int(i)+1), current)
			}
		}
		sortMax--
	}
	context.renderCommands.length = 0
	dfsBuffer.length = 0
	for rootIndex := int32(0); int(rootIndex) < int(context.layoutElementTreeRoots.length); rootIndex++ {
		dfsBuffer.length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, rootIndex)
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, int32(int(root.layoutElementIndex)))
		var rootPosition Clay_Vector2 = Clay_Vector2{}
		var parentHashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(root.parentId)
		if Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER) && parentHashMapItem != nil {
			var (
				config               *Clay_FloatingElementConfig = Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER).floatingElementConfig
				rootDimensions       Clay_Dimensions             = rootElement.dimensions
				parentBoundingBox    Clay_BoundingBox            = parentHashMapItem.boundingBox
				targetAttachPosition Clay_Vector2                = Clay_Vector2{}
			)
			switch config.attachment.parent {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				targetAttachPosition.x = parentBoundingBox.x
			case CLAY_ATTACH_POINT_CENTER_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.x = parentBoundingBox.x + parentBoundingBox.width/2
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.x = parentBoundingBox.x + parentBoundingBox.width
			}
			switch config.attachment.element {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
			case CLAY_ATTACH_POINT_CENTER_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				targetAttachPosition.x -= rootDimensions.width / 2
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.x -= rootDimensions.width
			}
			switch config.attachment.parent {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_TOP:
				targetAttachPosition.y = parentBoundingBox.y
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.y = parentBoundingBox.y + parentBoundingBox.height/2
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.y = parentBoundingBox.y + parentBoundingBox.height
			}
			switch config.attachment.element {
			case CLAY_ATTACH_POINT_LEFT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_TOP:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_TOP:
			case CLAY_ATTACH_POINT_LEFT_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_CENTER:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_CENTER:
				targetAttachPosition.y -= rootDimensions.height / 2
			case CLAY_ATTACH_POINT_LEFT_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_CENTER_BOTTOM:
				fallthrough
			case CLAY_ATTACH_POINT_RIGHT_BOTTOM:
				targetAttachPosition.y -= rootDimensions.height
			}
			targetAttachPosition.x += config.offset.x
			targetAttachPosition.y += config.offset.y
			rootPosition = targetAttachPosition
		}
		if int(root.clipElementId) != 0 {
			var clipHashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(root.clipElementId)
			if clipHashMapItem != nil {
				if context.externalScrollHandlingEnabled {
					var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(clipHashMapItem.layoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
					for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
						if mapping.layoutElement == clipHashMapItem.layoutElement {
							root.pointerOffset = mapping.scrollPosition
							if scrollConfig.horizontal {
								rootPosition.x += mapping.scrollPosition.x
							}
							if scrollConfig.vertical {
								rootPosition.y += mapping.scrollPosition.y
							}
							break
						}
					}
				}
				Clay__AddRenderCommand(Clay_RenderCommand{boundingBox: clipHashMapItem.boundingBox, config: Clay_ElementConfigUnion{scrollElementConfig: Clay__StoreScrollElementConfig(Clay_ScrollElementConfig{horizontal: false})}, id: Clay__RehashWithNumber(rootElement.id, 10), commandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_START})
			}
		}
		Clay__LayoutElementTreeNodeArray_Add(&dfsBuffer, Clay__LayoutElementTreeNode{layoutElement: rootElement, position: rootPosition, nextChildOffset: Clay_Vector2{x: float32(rootElement.layoutConfig.padding.left), y: float32(rootElement.layoutConfig.padding.top)}})
		*context.treeNodeVisited.internalArray = false
		for int(dfsBuffer.length) > 0 {
			var (
				currentElementTreeNode *Clay__LayoutElementTreeNode = Clay__LayoutElementTreeNodeArray_Get(&dfsBuffer, int32(int(dfsBuffer.length)-1))
				currentElement         *Clay_LayoutElement          = currentElementTreeNode.layoutElement
				layoutConfig           *Clay_LayoutConfig           = currentElement.layoutConfig
				scrollOffset           Clay_Vector2                 = Clay_Vector2{}
			)
			if !*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) {
				*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) = true
				var currentElementBoundingBox Clay_BoundingBox = Clay_BoundingBox{x: currentElementTreeNode.position.x, y: currentElementTreeNode.position.y, width: currentElement.dimensions.width, height: currentElement.dimensions.height}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER) {
					var (
						floatingElementConfig *Clay_FloatingElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER).floatingElementConfig
						expand                Clay_Dimensions             = floatingElementConfig.expand
					)
					currentElementBoundingBox.x -= expand.width
					currentElementBoundingBox.width += expand.width * 2
					currentElementBoundingBox.y -= expand.height
					currentElementBoundingBox.height += expand.height * 2
				}
				var scrollContainerData *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER) {
					var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
					for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
						if mapping.layoutElement == currentElement {
							scrollContainerData = mapping
							mapping.boundingBox = currentElementBoundingBox
							if scrollConfig.horizontal {
								scrollOffset.x = mapping.scrollPosition.x
							}
							if scrollConfig.vertical {
								scrollOffset.y = mapping.scrollPosition.y
							}
							if context.externalScrollHandlingEnabled {
								scrollOffset = Clay_Vector2{}
							}
							break
						}
					}
				}
				var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.id)
				if hashMapItem != nil {
					hashMapItem.boundingBox = currentElementBoundingBox
				}
				var sortedConfigIndexes [20]int32
				for elementConfigIndex := int32(0); int(elementConfigIndex) < int(currentElement.elementConfigs.length); elementConfigIndex++ {
					sortedConfigIndexes[elementConfigIndex] = elementConfigIndex
				}
				sortMax = int32(int(currentElement.elementConfigs.length) - 1)
				for int(sortMax) > 0 {
					for i := int32(0); int(i) < int(sortMax); i++ {
						var (
							current     int32                   = sortedConfigIndexes[i]
							next        int32                   = sortedConfigIndexes[int(i)+1]
							currentType Clay__ElementConfigType = Clay__ElementConfigArraySlice_Get(&currentElement.elementConfigs, current).type_
							nextType    Clay__ElementConfigType = Clay__ElementConfigArraySlice_Get(&currentElement.elementConfigs, next).type_
						)
						if nextType == CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER || currentType == CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER {
							sortedConfigIndexes[i] = next
							sortedConfigIndexes[int(i)+1] = current
						}
					}
					sortMax--
				}
				for elementConfigIndex := int32(0); int(elementConfigIndex) < int(currentElement.elementConfigs.length); elementConfigIndex++ {
					var (
						elementConfig *Clay_ElementConfig = Clay__ElementConfigArraySlice_Get(&currentElement.elementConfigs, sortedConfigIndexes[elementConfigIndex])
						renderCommand Clay_RenderCommand  = Clay_RenderCommand{boundingBox: currentElementBoundingBox, config: elementConfig.config, id: currentElement.id}
						offscreen     bool                = Clay__ElementIsOffscreen(&currentElementBoundingBox)
						shouldRender  bool                = !offscreen
					)
					switch elementConfig.type_ {
					case CLAY__ELEMENT_CONFIG_TYPE_RECTANGLE:
						renderCommand.commandType = CLAY_RENDER_COMMAND_TYPE_RECTANGLE
					case CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER:
						shouldRender = false
					case CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER:
						renderCommand.commandType = CLAY_RENDER_COMMAND_TYPE_NONE
						shouldRender = false
					case CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER:
						renderCommand.commandType = CLAY_RENDER_COMMAND_TYPE_SCISSOR_START
						shouldRender = true
					case CLAY__ELEMENT_CONFIG_TYPE_IMAGE:
						renderCommand.commandType = CLAY_RENDER_COMMAND_TYPE_IMAGE
					case CLAY__ELEMENT_CONFIG_TYPE_TEXT:
						if !shouldRender {
							break
						}
						shouldRender = false
						var configUnion Clay_ElementConfigUnion = elementConfig.config
						var textElementConfig *Clay_TextElementConfig = configUnion.textElementConfig
						var naturalLineHeight float32 = currentElement.childrenOrTextContent.textElementData.preferredDimensions.height
						var finalLineHeight float32
						if int(textElementConfig.lineHeight) > 0 {
							finalLineHeight = float32(textElementConfig.lineHeight)
						} else {
							finalLineHeight = naturalLineHeight
						}
						var lineHeightOffset float32 = (finalLineHeight - naturalLineHeight) / 2
						var yPosition float32 = lineHeightOffset
						for lineIndex := int32(0); int(lineIndex) < int(currentElement.childrenOrTextContent.textElementData.wrappedLines.length); lineIndex++ {
							var wrappedLine Clay__WrappedTextLine = *(*Clay__WrappedTextLine)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.textElementData.wrappedLines.internalArray), unsafe.Sizeof(Clay__WrappedTextLine{})*uintptr(lineIndex)))
							if int(wrappedLine.line.length) == 0 {
								yPosition += finalLineHeight
								continue
							}
							Clay__AddRenderCommand(Clay_RenderCommand{boundingBox: Clay_BoundingBox{x: currentElementBoundingBox.x, y: currentElementBoundingBox.y + yPosition, width: wrappedLine.dimensions.width, height: wrappedLine.dimensions.height}, config: configUnion, text: wrappedLine.line, id: Clay__HashNumber(uint32(lineIndex), currentElement.id).id, commandType: CLAY_RENDER_COMMAND_TYPE_TEXT})
							yPosition += finalLineHeight
							if !context.disableCulling && currentElementBoundingBox.y+yPosition > context.layoutDimensions.height {
								break
							}
						}
					case CLAY__ELEMENT_CONFIG_TYPE_CUSTOM:
						renderCommand.commandType = CLAY_RENDER_COMMAND_TYPE_CUSTOM
					default:
					}
					if shouldRender {
						Clay__AddRenderCommand(renderCommand)
					}
					if offscreen {
					}
				}
				if !Clay__ElementHasConfig(currentElementTreeNode.layoutElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
					var contentSize Clay_Dimensions = Clay_Dimensions{}
					if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
						for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
							var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							contentSize.width += childElement.dimensions.width
							if contentSize.height > childElement.dimensions.height {
								contentSize.height = contentSize.height
							} else {
								contentSize.height = childElement.dimensions.height
							}
						}
						contentSize.width += float32((func() int {
							if (int(currentElement.childrenOrTextContent.children.length) - 1) > 0 {
								return int(currentElement.childrenOrTextContent.children.length) - 1
							}
							return 0
						}()) * int(layoutConfig.childGap))
						var extraSpace float32 = currentElement.dimensions.width - float32(int(layoutConfig.padding.left)+int(layoutConfig.padding.right)) - contentSize.width
						switch layoutConfig.childAlignment.x {
						case CLAY_ALIGN_X_LEFT:
							extraSpace = 0
						case CLAY_ALIGN_X_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.nextChildOffset.x += extraSpace
					} else {
						for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
							var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
							if contentSize.width > childElement.dimensions.width {
								contentSize.width = contentSize.width
							} else {
								contentSize.width = childElement.dimensions.width
							}
							contentSize.height += childElement.dimensions.height
						}
						contentSize.height += float32((func() int {
							if (int(currentElement.childrenOrTextContent.children.length) - 1) > 0 {
								return int(currentElement.childrenOrTextContent.children.length) - 1
							}
							return 0
						}()) * int(layoutConfig.childGap))
						var extraSpace float32 = currentElement.dimensions.height - float32(int(layoutConfig.padding.top)+int(layoutConfig.padding.bottom)) - contentSize.height
						switch layoutConfig.childAlignment.y {
						case CLAY_ALIGN_Y_TOP:
							extraSpace = 0
						case CLAY_ALIGN_Y_CENTER:
							extraSpace /= 2
						default:
						}
						currentElementTreeNode.nextChildOffset.y += extraSpace
					}
					if scrollContainerData != nil {
						scrollContainerData.contentSize = Clay_Dimensions{width: contentSize.width + float32(int(layoutConfig.padding.left)+int(layoutConfig.padding.right)), height: contentSize.height + float32(int(layoutConfig.padding.top)+int(layoutConfig.padding.bottom))}
					}
				}
			} else {
				var closeScrollElement bool = false
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER) {
					closeScrollElement = true
					var scrollConfig *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
					for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
						var mapping *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
						if mapping.layoutElement == currentElement {
							if scrollConfig.horizontal {
								scrollOffset.x = mapping.scrollPosition.x
							}
							if scrollConfig.vertical {
								scrollOffset.y = mapping.scrollPosition.y
							}
							if context.externalScrollHandlingEnabled {
								scrollOffset = Clay_Vector2{}
							}
							break
						}
					}
				}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER) {
					var (
						currentElementData        *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.id)
						currentElementBoundingBox Clay_BoundingBox               = currentElementData.boundingBox
					)
					if !Clay__ElementIsOffscreen(&currentElementBoundingBox) {
						var (
							borderConfig  *Clay_BorderElementConfig = Clay__FindElementConfigWithType(currentElement, CLAY__ELEMENT_CONFIG_TYPE_BORDER_CONTAINER).borderElementConfig
							renderCommand Clay_RenderCommand        = Clay_RenderCommand{boundingBox: currentElementBoundingBox, config: Clay_ElementConfigUnion{borderElementConfig: borderConfig}, id: Clay__RehashWithNumber(currentElement.id, 4), commandType: CLAY_RENDER_COMMAND_TYPE_BORDER}
						)
						Clay__AddRenderCommand(renderCommand)
						if int(borderConfig.betweenChildren.width) > 0 && borderConfig.betweenChildren.color.a > 0 {
							var (
								rectangleConfig *Clay_RectangleElementConfig = Clay__StoreRectangleElementConfig(Clay_RectangleElementConfig{color: borderConfig.betweenChildren.color})
								halfGap         float32                      = float32(int(layoutConfig.childGap) / 2)
								borderOffset    Clay_Vector2                 = Clay_Vector2{x: float32(layoutConfig.padding.left) - halfGap, y: float32(layoutConfig.padding.top) - halfGap}
							)
							if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
								for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
									var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if int(i) > 0 {
										Clay__AddRenderCommand(Clay_RenderCommand{boundingBox: Clay_BoundingBox{x: currentElementBoundingBox.x + borderOffset.x + scrollOffset.x, y: currentElementBoundingBox.y + scrollOffset.y, width: float32(borderConfig.betweenChildren.width), height: currentElement.dimensions.height}, config: Clay_ElementConfigUnion{rectangleElementConfig: rectangleConfig}, id: Clay__RehashWithNumber(currentElement.id, uint32(int32(int(i)+5))), commandType: CLAY_RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.x += childElement.dimensions.width + float32(layoutConfig.childGap)
								}
							} else {
								for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
									var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
									if int(i) > 0 {
										Clay__AddRenderCommand(Clay_RenderCommand{boundingBox: Clay_BoundingBox{x: currentElementBoundingBox.x + scrollOffset.x, y: currentElementBoundingBox.y + borderOffset.y + scrollOffset.y, width: currentElement.dimensions.width, height: float32(borderConfig.betweenChildren.width)}, config: Clay_ElementConfigUnion{rectangleElementConfig: rectangleConfig}, id: Clay__RehashWithNumber(currentElement.id, uint32(int32(int(i)+5))), commandType: CLAY_RENDER_COMMAND_TYPE_RECTANGLE})
									}
									borderOffset.y += childElement.dimensions.height + float32(layoutConfig.childGap)
								}
							}
						}
					}
				}
				if closeScrollElement {
					Clay__AddRenderCommand(Clay_RenderCommand{id: Clay__RehashWithNumber(currentElement.id, 11), commandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_END})
				}
				dfsBuffer.length--
				continue
			}
			if !Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
				dfsBuffer.length += int32(currentElement.childrenOrTextContent.children.length)
				for i := int32(0); int(i) < int(currentElement.childrenOrTextContent.children.length); i++ {
					var childElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
						currentElementTreeNode.nextChildOffset.y = float32(currentElement.layoutConfig.padding.top)
						var whiteSpaceAroundChild float32 = currentElement.dimensions.height - float32(int(layoutConfig.padding.top)+int(layoutConfig.padding.bottom)) - childElement.dimensions.height
						switch layoutConfig.childAlignment.y {
						case CLAY_ALIGN_Y_TOP:
						case CLAY_ALIGN_Y_CENTER:
							currentElementTreeNode.nextChildOffset.y += whiteSpaceAroundChild / 2
						case CLAY_ALIGN_Y_BOTTOM:
							currentElementTreeNode.nextChildOffset.y += whiteSpaceAroundChild
						}
					} else {
						currentElementTreeNode.nextChildOffset.x = float32(currentElement.layoutConfig.padding.left)
						var whiteSpaceAroundChild float32 = currentElement.dimensions.width - float32(int(layoutConfig.padding.left)+int(layoutConfig.padding.right)) - childElement.dimensions.width
						switch layoutConfig.childAlignment.x {
						case CLAY_ALIGN_X_LEFT:
						case CLAY_ALIGN_X_CENTER:
							currentElementTreeNode.nextChildOffset.x += whiteSpaceAroundChild / 2
						case CLAY_ALIGN_X_RIGHT:
							currentElementTreeNode.nextChildOffset.x += whiteSpaceAroundChild
						}
					}
					var childPosition Clay_Vector2 = Clay_Vector2{x: currentElementTreeNode.position.x + currentElementTreeNode.nextChildOffset.x + scrollOffset.x, y: currentElementTreeNode.position.y + currentElementTreeNode.nextChildOffset.y + scrollOffset.y}
					var newNodeIndex uint32 = uint32(int32(int(dfsBuffer.length) - 1 - int(i)))
					*(*Clay__LayoutElementTreeNode)(unsafe.Add(unsafe.Pointer(dfsBuffer.internalArray), unsafe.Sizeof(Clay__LayoutElementTreeNode{})*uintptr(newNodeIndex))) = Clay__LayoutElementTreeNode{layoutElement: childElement, position: Clay_Vector2{x: childPosition.x, y: childPosition.y}, nextChildOffset: Clay_Vector2{x: float32(childElement.layoutConfig.padding.left), y: float32(childElement.layoutConfig.padding.top)}}
					*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), newNodeIndex)) = false
					if layoutConfig.layoutDirection == CLAY_LEFT_TO_RIGHT {
						currentElementTreeNode.nextChildOffset.x += childElement.dimensions.width + float32(layoutConfig.childGap)
					} else {
						currentElementTreeNode.nextChildOffset.y += childElement.dimensions.height + float32(layoutConfig.childGap)
					}
				}
			}
		}
		if int(root.clipElementId) != 0 {
			Clay__AddRenderCommand(Clay_RenderCommand{id: Clay__RehashWithNumber(rootElement.id, 11), commandType: CLAY_RENDER_COMMAND_TYPE_SCISSOR_END})
		}
	}
}
func Clay__AttachId(elementId Clay_ElementId) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	openLayoutElement.id = elementId.id
	Clay__AddHashMapItem(elementId, openLayoutElement)
	Clay__StringArray_Add(&context.layoutElementIdStrings, elementId.stringId)
}
func Clay__AttachLayoutConfig(config *Clay_LayoutConfig) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	Clay__GetOpenLayoutElement().layoutConfig = config
}
func Clay__AttachElementConfig(config Clay_ElementConfigUnion, type_ Clay__ElementConfigType) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	openLayoutElement.elementConfigs.length++
	Clay__ElementConfigArray_Add(&context.elementConfigBuffer, Clay_ElementConfig{type_: type_, config: config})
}
func Clay__StoreLayoutConfig(config Clay_LayoutConfig) *Clay_LayoutConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY_LAYOUT_DEFAULT
	}
	return Clay__LayoutConfigArray_Add(&Clay_GetCurrentContext().layoutConfigs, config)
}
func Clay__StoreRectangleElementConfig(config Clay_RectangleElementConfig) *Clay_RectangleElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__RECTANGLE_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__RectangleElementConfigArray_Add(&Clay_GetCurrentContext().rectangleElementConfigs, config)
}
func Clay__StoreTextElementConfig(config Clay_TextElementConfig) *Clay_TextElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__TEXT_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__TextElementConfigArray_Add(&Clay_GetCurrentContext().textElementConfigs, config)
}
func Clay__StoreImageElementConfig(config Clay_ImageElementConfig) *Clay_ImageElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__IMAGE_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__ImageElementConfigArray_Add(&Clay_GetCurrentContext().imageElementConfigs, config)
}
func Clay__StoreFloatingElementConfig(config Clay_FloatingElementConfig) *Clay_FloatingElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__FLOATING_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__FloatingElementConfigArray_Add(&Clay_GetCurrentContext().floatingElementConfigs, config)
}
func Clay__StoreCustomElementConfig(config Clay_CustomElementConfig) *Clay_CustomElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__CUSTOM_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__CustomElementConfigArray_Add(&Clay_GetCurrentContext().customElementConfigs, config)
}
func Clay__StoreScrollElementConfig(config Clay_ScrollElementConfig) *Clay_ScrollElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__SCROLL_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__ScrollElementConfigArray_Add(&Clay_GetCurrentContext().scrollElementConfigs, config)
}
func Clay__StoreBorderElementConfig(config Clay_BorderElementConfig) *Clay_BorderElementConfig {
	if Clay_GetCurrentContext().booleanWarnings.maxElementsExceeded {
		return &CLAY__BORDER_ELEMENT_CONFIG_DEFAULT
	}
	return Clay__BorderElementConfigArray_Add(&Clay_GetCurrentContext().borderElementConfigs, config)
}

var Clay__debugViewWidth uint32 = 400
var Clay__debugViewHighlightColor Clay_Color = Clay_Color{r: 168, g: 66, b: 28, a: 100}

func Clay_MinMemorySize() uint32 {
	var (
		fakeContext    Clay_Context  = Clay_Context{maxElementCount: Clay__defaultMaxElementCount, maxMeasureTextCacheWordCount: Clay__defaultMaxMeasureTextWordCacheCount, internalArena: Clay_Arena{capacity: math.MaxUint64, memory: nil}}
		currentContext *Clay_Context = Clay_GetCurrentContext()
	)
	if currentContext != nil {
		fakeContext.maxElementCount = currentContext.maxElementCount
		fakeContext.maxMeasureTextCacheWordCount = currentContext.maxElementCount
	}
	Clay__Context_Allocate_Arena(&fakeContext.internalArena)
	Clay__InitializePersistentMemory(&fakeContext)
	Clay__InitializeEphemeralMemory(&fakeContext)
	return uint32(fakeContext.internalArena.nextAllocation)
}
func Clay_CreateArenaWithCapacityAndMemory(capacity uint32, offset unsafe.Pointer) Clay_Arena {
	var arena Clay_Arena = Clay_Arena{capacity: uint64(capacity), memory: (*byte)(offset)}
	return arena
}
func Clay_SetMeasureTextFunction(measureTextFunction func(text *Clay_String, config *Clay_TextElementConfig) Clay_Dimensions) {
	Clay__MeasureText = measureTextFunction
}
func Clay_SetQueryScrollOffsetFunction(queryScrollOffsetFunction func(elementId uint32) Clay_Vector2) {
	Clay__QueryScrollOffset = queryScrollOffsetFunction
}
func Clay_SetLayoutDimensions(dimensions Clay_Dimensions) {
	Clay_GetCurrentContext().layoutDimensions = dimensions
}
func Clay_SetPointerState(position Clay_Vector2, isPointerDown bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	context.pointerInfo.position = position
	context.pointerOverIds.length = 0
	var dfsBuffer Clay__int32_tArray = context.layoutElementChildrenBuffer
	for rootIndex := int32(int32(int(context.layoutElementTreeRoots.length) - 1)); int(rootIndex) >= 0; rootIndex-- {
		dfsBuffer.length = 0
		var root *Clay__LayoutElementTreeRoot = Clay__LayoutElementTreeRootArray_Get(&context.layoutElementTreeRoots, rootIndex)
		Clay__int32_tArray_Add(&dfsBuffer, root.layoutElementIndex)
		*context.treeNodeVisited.internalArray = false
		var found bool = false
		for int(dfsBuffer.length) > 0 {
			if *(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) {
				dfsBuffer.length--
				continue
			}
			*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) = true
			var currentElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, Clay__int32_tArray_Get(&dfsBuffer, int32(int(dfsBuffer.length)-1)))
			var mapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(currentElement.id)
			var elementBox Clay_BoundingBox = mapItem.boundingBox
			elementBox.x -= root.pointerOffset.x
			elementBox.y -= root.pointerOffset.y
			if mapItem != nil {
				if Clay__PointIsInsideRect(position, elementBox) {
					if mapItem.onHoverFunction != nil {
						mapItem.onHoverFunction(mapItem.elementId, context.pointerInfo, mapItem.hoverFunctionUserData)
					}
					Clay__ElementIdArray_Add(&context.pointerOverIds, mapItem.elementId)
					found = true
				}
				if Clay__ElementHasConfig(currentElement, CLAY__ELEMENT_CONFIG_TYPE_TEXT) {
					dfsBuffer.length--
					continue
				}
				for i := int32(int32(int(currentElement.childrenOrTextContent.children.length) - 1)); int(i) >= 0; i-- {
					Clay__int32_tArray_Add(&dfsBuffer, *(*int32)(unsafe.Add(unsafe.Pointer(currentElement.childrenOrTextContent.children.elements), unsafe.Sizeof(int32(0))*uintptr(i))))
					*(*bool)(unsafe.Add(unsafe.Pointer(context.treeNodeVisited.internalArray), int(dfsBuffer.length)-1)) = false
				}
			} else {
				dfsBuffer.length--
			}
		}
		var rootElement *Clay_LayoutElement = Clay_LayoutElementArray_Get(&context.layoutElements, root.layoutElementIndex)
		if found && Clay__ElementHasConfig(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER) && Clay__FindElementConfigWithType(rootElement, CLAY__ELEMENT_CONFIG_TYPE_FLOATING_CONTAINER).floatingElementConfig.pointerCaptureMode == CLAY_POINTER_CAPTURE_MODE_CAPTURE {
			break
		}
	}
	if isPointerDown {
		if context.pointerInfo.state == CLAY_POINTER_DATA_PRESSED_THIS_FRAME {
			context.pointerInfo.state = CLAY_POINTER_DATA_PRESSED
		} else if context.pointerInfo.state != CLAY_POINTER_DATA_PRESSED {
			context.pointerInfo.state = CLAY_POINTER_DATA_PRESSED_THIS_FRAME
		}
	} else {
		if context.pointerInfo.state == CLAY_POINTER_DATA_RELEASED_THIS_FRAME {
			context.pointerInfo.state = CLAY_POINTER_DATA_RELEASED
		} else if context.pointerInfo.state != CLAY_POINTER_DATA_RELEASED {
			context.pointerInfo.state = CLAY_POINTER_DATA_RELEASED_THIS_FRAME
		}
	}
}
func Clay_Initialize(arena Clay_Arena, layoutDimensions Clay_Dimensions, errorHandler Clay_ErrorHandler) *Clay_Context {
	var context *Clay_Context = Clay__Context_Allocate_Arena(&arena)
	if context == nil {
		return nil
	}
	var oldContext *Clay_Context = Clay_GetCurrentContext()
	*context = Clay_Context{maxElementCount: int32(func() int {
		if oldContext != nil {
			return int(oldContext.maxElementCount)
		}
		return int(Clay__defaultMaxElementCount)
	}()), maxMeasureTextCacheWordCount: int32(func() int {
		if oldContext != nil {
			return int(oldContext.maxMeasureTextCacheWordCount)
		}
		return int(Clay__defaultMaxElementCount) * 2
	}()), errorHandler: func() Clay_ErrorHandler {
		if errorHandler.errorHandlerFunction != nil {
			return errorHandler
		}
		return Clay_ErrorHandler{errorHandlerFunction: Clay__ErrorHandlerFunctionDefault}
	}(), layoutDimensions: layoutDimensions, internalArena: arena}
	Clay_SetCurrentContext(context)
	Clay__InitializePersistentMemory(context)
	Clay__InitializeEphemeralMemory(context)
	for i := int32(0); int(i) < int(context.layoutElementsHashMap.capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.layoutElementsHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = -1
	}
	for i := int32(0); int(i) < int(context.measureTextHashMap.capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.measureTextHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.measureTextHashMapInternal.length = 1
	context.layoutDimensions = layoutDimensions
	return context
}
func Clay_GetCurrentContext() *Clay_Context {
	return Clay__currentContext
}
func Clay_SetCurrentContext(context *Clay_Context) {
	Clay__currentContext = context
}
func Clay_UpdateScrollContainers(enableDragScrolling bool, scrollDelta Clay_Vector2, deltaTime float32) {
	var (
		context                     *Clay_Context                      = Clay_GetCurrentContext()
		isPointerActive             bool                               = enableDragScrolling && (context.pointerInfo.state == CLAY_POINTER_DATA_PRESSED || context.pointerInfo.state == CLAY_POINTER_DATA_PRESSED_THIS_FRAME)
		highestPriorityElementIndex int32                              = -1
		highestPriorityScrollData   *Clay__ScrollContainerDataInternal = (*Clay__ScrollContainerDataInternal)(unsafe.Pointer(uintptr(CLAY__NULL)))
	)
	for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
		var scrollData *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
		if !scrollData.openThisFrame {
			Clay__ScrollContainerDataInternalArray_RemoveSwapback(&context.scrollContainerDatas, i)
			continue
		}
		scrollData.openThisFrame = false
		var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(scrollData.elementId)
		if hashMapItem == nil {
			Clay__ScrollContainerDataInternalArray_RemoveSwapback(&context.scrollContainerDatas, i)
			continue
		}
		if !isPointerActive && scrollData.pointerScrollActive {
			var xDiff float32 = scrollData.scrollPosition.x - scrollData.scrollOrigin.x
			if xDiff < -10 || xDiff > 10 {
				scrollData.scrollMomentum.x = (scrollData.scrollPosition.x - scrollData.scrollOrigin.x) / (scrollData.momentumTime * 25)
			}
			var yDiff float32 = scrollData.scrollPosition.y - scrollData.scrollOrigin.y
			if yDiff < -10 || yDiff > 10 {
				scrollData.scrollMomentum.y = (scrollData.scrollPosition.y - scrollData.scrollOrigin.y) / (scrollData.momentumTime * 25)
			}
			scrollData.pointerScrollActive = false
			scrollData.pointerOrigin = Clay_Vector2{}
			scrollData.scrollOrigin = Clay_Vector2{}
			scrollData.momentumTime = 0
		}
		scrollData.scrollPosition.x += scrollData.scrollMomentum.x
		scrollData.scrollMomentum.x *= 0.95
		var scrollOccurred bool = scrollDelta.x != 0 || scrollDelta.y != 0
		if scrollData.scrollMomentum.x > -0.1 && scrollData.scrollMomentum.x < 0.1 || scrollOccurred {
			scrollData.scrollMomentum.x = 0
		}
		if (func() float32 {
			if scrollData.scrollPosition.x > (-func() float32 {
				if (scrollData.contentSize.width - scrollData.layoutElement.dimensions.width) > 0 {
					return scrollData.contentSize.width - scrollData.layoutElement.dimensions.width
				}
				return 0
			}()) {
				return scrollData.scrollPosition.x
			}
			return -func() float32 {
				if (scrollData.contentSize.width - scrollData.layoutElement.dimensions.width) > 0 {
					return scrollData.contentSize.width - scrollData.layoutElement.dimensions.width
				}
				return 0
			}()
		}()) < 0 {
			if scrollData.scrollPosition.x > (-func() float32 {
				if (scrollData.contentSize.width - scrollData.layoutElement.dimensions.width) > 0 {
					return scrollData.contentSize.width - scrollData.layoutElement.dimensions.width
				}
				return 0
			}()) {
				scrollData.scrollPosition.x = scrollData.scrollPosition.x
			} else if (scrollData.contentSize.width - scrollData.layoutElement.dimensions.width) > 0 {
				scrollData.scrollPosition.x = -(scrollData.contentSize.width - scrollData.layoutElement.dimensions.width)
			} else {
				scrollData.scrollPosition.x = 0
			}
		} else {
			scrollData.scrollPosition.x = 0
		}
		scrollData.scrollPosition.y += scrollData.scrollMomentum.y
		scrollData.scrollMomentum.y *= 0.95
		if scrollData.scrollMomentum.y > -0.1 && scrollData.scrollMomentum.y < 0.1 || scrollOccurred {
			scrollData.scrollMomentum.y = 0
		}
		if (func() float32 {
			if scrollData.scrollPosition.y > (-func() float32 {
				if (scrollData.contentSize.height - scrollData.layoutElement.dimensions.height) > 0 {
					return scrollData.contentSize.height - scrollData.layoutElement.dimensions.height
				}
				return 0
			}()) {
				return scrollData.scrollPosition.y
			}
			return -func() float32 {
				if (scrollData.contentSize.height - scrollData.layoutElement.dimensions.height) > 0 {
					return scrollData.contentSize.height - scrollData.layoutElement.dimensions.height
				}
				return 0
			}()
		}()) < 0 {
			if scrollData.scrollPosition.y > (-func() float32 {
				if (scrollData.contentSize.height - scrollData.layoutElement.dimensions.height) > 0 {
					return scrollData.contentSize.height - scrollData.layoutElement.dimensions.height
				}
				return 0
			}()) {
				scrollData.scrollPosition.y = scrollData.scrollPosition.y
			} else if (scrollData.contentSize.height - scrollData.layoutElement.dimensions.height) > 0 {
				scrollData.scrollPosition.y = -(scrollData.contentSize.height - scrollData.layoutElement.dimensions.height)
			} else {
				scrollData.scrollPosition.y = 0
			}
		} else {
			scrollData.scrollPosition.y = 0
		}
		for j := int32(0); int(j) < int(context.pointerOverIds.length); j++ {
			if int(scrollData.layoutElement.id) == int(Clay__ElementIdArray_Get(&context.pointerOverIds, j).id) {
				highestPriorityElementIndex = j
				highestPriorityScrollData = scrollData
			}
		}
	}
	if int(highestPriorityElementIndex) > -1 && highestPriorityScrollData != nil {
		var (
			scrollElement         *Clay_LayoutElement       = highestPriorityScrollData.layoutElement
			scrollConfig          *Clay_ScrollElementConfig = Clay__FindElementConfigWithType(scrollElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig
			canScrollVertically   bool                      = scrollConfig.vertical && highestPriorityScrollData.contentSize.height > scrollElement.dimensions.height
			canScrollHorizontally bool                      = scrollConfig.horizontal && highestPriorityScrollData.contentSize.width > scrollElement.dimensions.width
		)
		if canScrollVertically {
			highestPriorityScrollData.scrollPosition.y = highestPriorityScrollData.scrollPosition.y + scrollDelta.y*10
		}
		if canScrollHorizontally {
			highestPriorityScrollData.scrollPosition.x = highestPriorityScrollData.scrollPosition.x + scrollDelta.x*10
		}
		if isPointerActive {
			highestPriorityScrollData.scrollMomentum = Clay_Vector2{}
			if !highestPriorityScrollData.pointerScrollActive {
				highestPriorityScrollData.pointerOrigin = context.pointerInfo.position
				highestPriorityScrollData.scrollOrigin = highestPriorityScrollData.scrollPosition
				highestPriorityScrollData.pointerScrollActive = true
			} else {
				var (
					scrollDeltaX float32 = 0
					scrollDeltaY float32 = 0
				)
				if canScrollHorizontally {
					var oldXScrollPosition float32 = highestPriorityScrollData.scrollPosition.x
					highestPriorityScrollData.scrollPosition.x = highestPriorityScrollData.scrollOrigin.x + (context.pointerInfo.position.x - highestPriorityScrollData.pointerOrigin.x)
					if (func() float32 {
						if highestPriorityScrollData.scrollPosition.x < 0 {
							return highestPriorityScrollData.scrollPosition.x
						}
						return 0
					}()) > (-(highestPriorityScrollData.contentSize.width - highestPriorityScrollData.boundingBox.width)) {
						if highestPriorityScrollData.scrollPosition.x < 0 {
							highestPriorityScrollData.scrollPosition.x = highestPriorityScrollData.scrollPosition.x
						} else {
							highestPriorityScrollData.scrollPosition.x = 0
						}
					} else {
						highestPriorityScrollData.scrollPosition.x = -(highestPriorityScrollData.contentSize.width - highestPriorityScrollData.boundingBox.width)
					}
					scrollDeltaX = highestPriorityScrollData.scrollPosition.x - oldXScrollPosition
				}
				if canScrollVertically {
					var oldYScrollPosition float32 = highestPriorityScrollData.scrollPosition.y
					highestPriorityScrollData.scrollPosition.y = highestPriorityScrollData.scrollOrigin.y + (context.pointerInfo.position.y - highestPriorityScrollData.pointerOrigin.y)
					if (func() float32 {
						if highestPriorityScrollData.scrollPosition.y < 0 {
							return highestPriorityScrollData.scrollPosition.y
						}
						return 0
					}()) > (-(highestPriorityScrollData.contentSize.height - highestPriorityScrollData.boundingBox.height)) {
						if highestPriorityScrollData.scrollPosition.y < 0 {
							highestPriorityScrollData.scrollPosition.y = highestPriorityScrollData.scrollPosition.y
						} else {
							highestPriorityScrollData.scrollPosition.y = 0
						}
					} else {
						highestPriorityScrollData.scrollPosition.y = -(highestPriorityScrollData.contentSize.height - highestPriorityScrollData.boundingBox.height)
					}
					scrollDeltaY = highestPriorityScrollData.scrollPosition.y - oldYScrollPosition
				}
				if scrollDeltaX > -0.1 && scrollDeltaX < 0.1 && scrollDeltaY > -0.1 && scrollDeltaY < 0.1 && highestPriorityScrollData.momentumTime > 0.15 {
					highestPriorityScrollData.momentumTime = 0
					highestPriorityScrollData.pointerOrigin = context.pointerInfo.position
					highestPriorityScrollData.scrollOrigin = highestPriorityScrollData.scrollPosition
				} else {
					highestPriorityScrollData.momentumTime += deltaTime
				}
			}
		}
		if canScrollVertically {
			if (func() float32 {
				if highestPriorityScrollData.scrollPosition.y < 0 {
					return highestPriorityScrollData.scrollPosition.y
				}
				return 0
			}()) > (-(highestPriorityScrollData.contentSize.height - scrollElement.dimensions.height)) {
				if highestPriorityScrollData.scrollPosition.y < 0 {
					highestPriorityScrollData.scrollPosition.y = highestPriorityScrollData.scrollPosition.y
				} else {
					highestPriorityScrollData.scrollPosition.y = 0
				}
			} else {
				highestPriorityScrollData.scrollPosition.y = -(highestPriorityScrollData.contentSize.height - scrollElement.dimensions.height)
			}
		}
		if canScrollHorizontally {
			if (func() float32 {
				if highestPriorityScrollData.scrollPosition.x < 0 {
					return highestPriorityScrollData.scrollPosition.x
				}
				return 0
			}()) > (-(highestPriorityScrollData.contentSize.width - scrollElement.dimensions.width)) {
				if highestPriorityScrollData.scrollPosition.x < 0 {
					highestPriorityScrollData.scrollPosition.x = highestPriorityScrollData.scrollPosition.x
				} else {
					highestPriorityScrollData.scrollPosition.x = 0
				}
			} else {
				highestPriorityScrollData.scrollPosition.x = -(highestPriorityScrollData.contentSize.width - scrollElement.dimensions.width)
			}
		}
	}
}
func Clay_BeginLayout() {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__InitializeEphemeralMemory(context)
	context.generation++
	context.dynamicElementIndex = 0
	var rootDimensions Clay_Dimensions = Clay_Dimensions{width: context.layoutDimensions.width, height: context.layoutDimensions.height}
	if context.debugModeEnabled {
		rootDimensions.width -= float32(Clay__debugViewWidth)
	}
	context.booleanWarnings.maxElementsExceeded = false
	context.booleanWarnings.maxTextMeasureCacheExceeded = false
	context.booleanWarnings.maxRenderCommandsExceeded = false
	Clay__OpenElement()
	Clay__AttachId(Clay__HashString(Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay__RootContainer")}, 0, 0))
	// Clay__AttachLayoutConfig(Clay__StoreLayoutConfig(Clay__Clay_LayoutConfigWrapper{wrapped: Clay_LayoutConfig{sizing: Clay_Sizing{width: Clay_SizingAxis{size: struct {
	// 	// union
	// 	minMax  Clay_SizingMinMax
	// 	percent float32
	// }{minMax: Clay_SizingMinMax{min: rootDimensions.width, max: rootDimensions.width}}, type_: CLAY__SIZING_TYPE_FIXED}, height: Clay_SizingAxis{size: struct {
	// 	// union
	// 	minMax  Clay_SizingMinMax
	// 	percent float32
	// }{minMax: Clay_SizingMinMax{min: rootDimensions.height, max: rootDimensions.height}}, type_: CLAY__SIZING_TYPE_FIXED}}}}.wrapped))
	Clay__ElementPostConfiguration()
	Clay__int32_tArray_Add(&context.openLayoutElementStack, 0)
	Clay__LayoutElementTreeRootArray_Add(&context.layoutElementTreeRoots, Clay__LayoutElementTreeRoot{})
}

var Clay__DebugView_ErrorTextConfig Clay_TextElementConfig = Clay_TextElementConfig{textColor: Clay_Color{r: 255, g: 0, b: 0, a: 255}, fontSize: 16, wrapMode: CLAY_TEXT_WRAP_NONE}

func Clay_EndLayout() Clay_RenderCommandArray {
	var context *Clay_Context = Clay_GetCurrentContext()
	Clay__CloseElement()
	if context.debugModeEnabled {
		context.warningsEnabled = false
		context.warningsEnabled = true
	}
	if context.booleanWarnings.maxElementsExceeded {
		Clay__AddRenderCommand(Clay_RenderCommand{boundingBox: Clay_BoundingBox{x: context.layoutDimensions.width/2 - 59*4, y: context.layoutDimensions.height / 2, width: 0, height: 0}, config: Clay_ElementConfigUnion{textElementConfig: &Clay__DebugView_ErrorTextConfig}, text: Clay_String{length: int32(uint32((unsafe.Sizeof(string(0)) / unsafe.Sizeof(byte(0))) - unsafe.Sizeof(byte(0)))), chars: libc.CString("Clay Error: Layout elements exceeded Clay__maxElementCount")}, commandType: CLAY_RENDER_COMMAND_TYPE_TEXT})
	} else {
		Clay__CalculateFinalLayout()
	}
	return context.renderCommands
}
func Clay_GetElementId(idString Clay_String) Clay_ElementId {
	return Clay__HashString(idString, 0, 0)
}
func Clay_GetElementIdWithIndex(idString Clay_String, index uint32) Clay_ElementId {
	return Clay__HashString(idString, index, 0)
}
func Clay_Hovered() bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return false
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	if int(openLayoutElement.id) == 0 {
		Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	for i := int32(0); int(i) < int(context.pointerOverIds.length); i++ {
		if int(Clay__ElementIdArray_Get(&context.pointerOverIds, i).id) == int(openLayoutElement.id) {
			return true
		}
	}
	return false
}
func Clay_OnHover(onHoverFunction func(elementId Clay_ElementId, pointerInfo Clay_PointerData, userData int64), userData int64) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context.booleanWarnings.maxElementsExceeded {
		return
	}
	var openLayoutElement *Clay_LayoutElement = Clay__GetOpenLayoutElement()
	if int(openLayoutElement.id) == 0 {
		Clay__GenerateIdForAnonymousElement(openLayoutElement)
	}
	var hashMapItem *Clay_LayoutElementHashMapItem = Clay__GetHashMapItem(openLayoutElement.id)
	hashMapItem.onHoverFunction = onHoverFunction
	hashMapItem.hoverFunctionUserData = userData
}
func Clay_PointerOver(elementId Clay_ElementId) bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	for i := int32(0); int(i) < int(context.pointerOverIds.length); i++ {
		if int(Clay__ElementIdArray_Get(&context.pointerOverIds, i).id) == int(elementId.id) {
			return true
		}
	}
	return false
}
func Clay_GetScrollContainerData(id Clay_ElementId) Clay_ScrollContainerData {
	var context *Clay_Context = Clay_GetCurrentContext()
	for i := int32(0); int(i) < int(context.scrollContainerDatas.length); i++ {
		var scrollContainerData *Clay__ScrollContainerDataInternal = Clay__ScrollContainerDataInternalArray_Get(&context.scrollContainerDatas, i)
		if int(scrollContainerData.elementId) == int(id.id) {
			return Clay_ScrollContainerData{scrollPosition: &scrollContainerData.scrollPosition, scrollContainerDimensions: Clay_Dimensions{width: scrollContainerData.boundingBox.width, height: scrollContainerData.boundingBox.height}, contentDimensions: scrollContainerData.contentSize, config: *Clay__FindElementConfigWithType(scrollContainerData.layoutElement, CLAY__ELEMENT_CONFIG_TYPE_SCROLL_CONTAINER).scrollElementConfig, found: true}
		}
	}
	return Clay_ScrollContainerData{}
}
func Clay_SetDebugModeEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.debugModeEnabled = enabled
}
func Clay_IsDebugModeEnabled() bool {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.debugModeEnabled
}
func Clay_SetCullingEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.disableCulling = !enabled
}
func Clay_SetExternalScrollHandlingEnabled(enabled bool) {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.externalScrollHandlingEnabled = enabled
}
func Clay_GetMaxElementCount() int32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.maxElementCount
}
func Clay_SetMaxElementCount(maxElementCount int32) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context != nil {
		context.maxElementCount = maxElementCount
	} else {
		Clay__defaultMaxElementCount = maxElementCount
	}
}
func Clay_GetMaxMeasureTextCacheWordCount() int32 {
	var context *Clay_Context = Clay_GetCurrentContext()
	return context.maxMeasureTextCacheWordCount
}
func Clay_SetMaxMeasureTextCacheWordCount(maxMeasureTextCacheWordCount int32) {
	var context *Clay_Context = Clay_GetCurrentContext()
	if context != nil {
		Clay__currentContext.maxMeasureTextCacheWordCount = maxMeasureTextCacheWordCount
	} else {
		Clay__defaultMaxMeasureTextWordCacheCount = maxMeasureTextCacheWordCount
	}
}
func Clay_ResetMeasureTextCache() {
	var context *Clay_Context = Clay_GetCurrentContext()
	context.measureTextHashMapInternal.length = 0
	context.measureTextHashMapInternalFreeList.length = 0
	context.measureTextHashMap.length = 0
	context.measuredWords.length = 0
	context.measuredWordsFreeList.length = 0
	for i := int32(0); int(i) < int(context.measureTextHashMap.capacity); i++ {
		*(*int32)(unsafe.Add(unsafe.Pointer(context.measureTextHashMap.internalArray), unsafe.Sizeof(int32(0))*uintptr(i))) = 0
	}
	context.measureTextHashMapInternal.length = 1
}
