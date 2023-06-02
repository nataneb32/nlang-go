package internal


type Program struct {
	GlobalNamespace *Namespace

}

type Namespace struct {
	Parent *Namespace
	Children []*Namespace
	Functions map[string]*Function
	Variables map[string]*Variable
	TypeDefs map[string]*TypeDef
}

type Function struct {
	Name string
	Namespace *Namespace
	ParamsType *Type
}

type Variable struct {
	Name string
	Type *Type
}

type TypeDef struct {
	Type *Type
}

type Type struct {
	Name string
}
