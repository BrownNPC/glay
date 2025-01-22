#include "clayman.hpp"
#include <emscripten/bind.h>
using namespace emscripten;
//This class initializes Clay.h layout library, manages it's context, and provides functions for convenience

EMSCRIPTEN_BINDINGS(my_module)
{
  class_<Clay_RenderCommandArray>("RenderCommandArray")
    .property("length", &Clay_RenderCommandArray::length)
    .function("get", Clay_RenderCommandArray_Get, allow_raw_pointers());

  class_<Clay_StringSlice>("StringSlice")
    .property("length", &Clay_StringSlice::length);
    class_<Clay_TextElementConfig>("TextElementConfig");

  class_<Clay_Dimensions>("Dimensions")
    .constructor()
    .property("width", &Clay_Dimensions::width)
    .property("height", &Clay_Dimensions::height);
  class_<ClayMan>("Clay")
    .constructor<uint32_t, uint32_t>()
    .function("BeginLayout", &ClayMan::beginLayout)
    .function("EndLayout", &ClayMan::endLayout);
  function("DebugModeEnabled", Clay_SetDebugModeEnabled);
}
