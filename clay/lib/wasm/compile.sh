#!/bin/bash

# Compilation command for Emscripten
compile_cmd="emcc -sERROR_ON_UNDEFINED_SYMBOLS=0 \
-std=c++20 \
-sEXPORTED_FUNCTIONS=[\"_free\",\"_malloc\",\"___getTypeName\"] \
-sEXPORTED_RUNTIME_METHODS=ccall,cwrap \
-Wl,--initial-memory=6553600\
 -g claybind.cpp -o build/clay.wasm -lembind --no-entry"

# Run the compilation process with error handling
echo "Running the compilation command..."
$compile_cmd
if [ $? -ne 0 ]; then
    echo "Compilation failed."
    exit 1
else
    echo "Compilation succeeded."
fi

# Template for the Go package
generate_template="//go:generate go run github.com/jerbob92/wazero-emscripten-embind/generator -v=true -wasm=../lib/wasm/build/clay.wasm
package internal
"


# Directory and file cleanup
generated_dir="../../generated"

if [ -d "$generated_dir" ]; then
    files=$(find "$generated_dir" -type f -name "*.go")
    for f in $files; do
        rm "$f"
        echo "Removed file: $f"
    done
else
    echo "Directory does not exist: $generated_dir"
fi

# Write the Go template to a file
mkdir -p "$generated_dir"
go_file_path="$generated_dir/generate.go"
echo "$generate_template" > "$go_file_path"
echo "Go template written to: $go_file_path"
echo "Running 'go generate' in the generated directory..."
cd "$generated_dir" && go generate
if [ $? -ne 0 ]; then
    echo "'go generate' failed."
    exit 1
else
echo "'go generate' succeeded."
fi
