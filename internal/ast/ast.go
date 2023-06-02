package ast

type Node interface{}

type Identifier struct {
	value string
	start int
	end   int
}

type symbol struct {
	value string
	start int
	end   int
}

type FunctionParam struct {
	name Identifier
	typ  Identifier
}

type function struct {
	params  []FunctionParam
	retType Identifier
	body    []Node
}

func NewFunction(retType Identifier, params ...FunctionParam) Node {
	return &function{params: params, retType: retType}
}
