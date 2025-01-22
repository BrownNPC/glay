package main

import (
	gen "clay-ui/clay/generated"
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"sync"

	embind "github.com/jerbob92/wazero-emscripten-embind"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/emscripten"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {

	cls, err := gen.NewClassClay(engine, ctx, 120, 120)
	if err != nil {
		panic(err)
	}
	gen.DebugModeEnabled(engine, ctx, true)
	err = cls.BeginLayout(ctx)
	if err != nil {
		panic(err)
	}
	arr, err := cls.EndLayout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(arr.(*gen.ClassRenderCommandArray).GetPropertyLength(ctx))

}

//go:embed clay/lib/wasm/build/clay.wasm
var wasm []byte

var (
	ctx     context.Context
	engine  embind.Engine
	runtime wazero.Runtime
	module  wazero.CompiledModule

	initonce = sync.OnceFunc(initialize)
)

func initialize() {
	ctx = context.Background()
	runtimeConfig := wazero.NewRuntimeConfig()
	r := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Fatal(err)
	}

	compiledModule, err := r.CompileModule(ctx, wasm)
	if err != nil {
		log.Fatal(err)
	}

	builder := r.NewHostModuleBuilder("env")
	emscriptenExporter, err := emscripten.NewFunctionExporterForModule(compiledModule)
	if err != nil {
		log.Fatal(err)
	}

	emscriptenExporter.ExportFunctions(builder)

	engine = embind.CreateEngine(embind.NewConfig())

	embindExporter := engine.NewFunctionExporterForModule(compiledModule)

	err = embindExporter.ExportFunctions(builder)
	if err != nil {
		log.Fatal(err)
	}

	_, err = builder.Instantiate(ctx)
	if err != nil {
		log.Fatal(err)
	}

	moduleConfig := wazero.NewModuleConfig().
		WithStartFunctions("_initialize").
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithName("")

	ctx = engine.Attach(ctx)
	_, err = r.InstantiateModule(ctx, compiledModule, moduleConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = engine.RegisterEmvalSymbol("clay", &clay{})
	if err != nil {
		panic(err)
	}
	err = gen.Attach(engine)
	if err != nil {
		log.Fatal(err)
	}
	module = compiledModule

}

func init() {
	initonce()
}

type clay struct {
}

func (c clay) MeasureText(arg *gen.ClassStringSlice, arg1 *gen.ClassTextElementConfig, arg2 uint32) *gen.ClassDimensions {

	fmt.Println("measure text is here: ", arg, arg1, arg2)
	dim, err := gen.NewClassDimensions(engine, ctx)
	if err != nil {
		panic(err)
	}
	return dim
}
