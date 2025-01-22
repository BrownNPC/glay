#include <stdlib.h>
#define CLAY_IMPLEMENTATION
#include "clay.h"

Clay_Arena NewArena()
{
  uint64_t clayRequiredMemory = Clay_MinMemorySize();
  Clay_Arena clayMemory = Clay_CreateArenaWithCapacityAndMemory(
    clayRequiredMemory, malloc(clayRequiredMemory));
  return clayMemory;
};
