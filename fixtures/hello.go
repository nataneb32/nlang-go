package main

import (
	"os"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)
func main () {
	module := ir.NewModule()

	printf := module.NewFunc("printf", types.NewPointer(types.NewInt(32)), ir.NewParam("format", types.NewPointer(types.I8)))

	main := module.NewFunc("main", types.NewInt(32))
	entry := main.NewBlock("entry")
	entry.NewRet(constant.NewInt(types.I32, 0))
	hello := constant.NewCharArrayFromString("Hello, World!\n")
	helloStr := module.NewGlobalDef("hello", hello)
	zero := constant.NewInt(types.I64, 0)
	gep := constant.NewGetElementPtr(hello.Typ, helloStr, zero, zero)

	entry.NewCall(printf, gep)

	module.WriteTo(os.Stdout)
}
