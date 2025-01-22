#!/bin/bash

# Check if a platform argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <platform>"
    echo "Platforms: linux, windows, macos"
    exit 1
fi

# Get the platform from the first argument
PLATFORM=$1

# Compile based on the platform
case $PLATFORM in
    linux)
        echo "Compiling for Linux..."
        cc -shared -fPIC claybind.c -o ../../libclay.so
        ;;
    windows)
        echo "Compiling for Windows..."
        x86_64-w64-mingw32-gcc -shared -fPIC claybind.c -o ../../clay.dll
        ;;
    macos)
        echo "Compiling for macOS..."
        cc -shared -fPIC claybind.c -o ../../libclay.dylib
        ;;
    *)
        echo "Unknown platform: $PLATFORM"
        echo "Supported platforms: linux, windows, macos"
        exit 1
        ;;
esac
