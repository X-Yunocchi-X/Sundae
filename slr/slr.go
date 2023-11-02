package slr

type (
	Symbol string // terminal or non-terminal
	// Production cannot use symbol S', because it is reserved for augmented production
	Production struct {
		Head Symbol
		Body []Symbol
	}
)

type Item struct {
	Prod   Production
	DotPos int // "." position
}

type Action int

const (
	SHIFT Action = iota
	REDUCE
	ACCEPT
	ERROR
)

type SLRAction struct {
	Action Action
	Value  int
}

type SLRTable struct {
	ActionTable map[int]map[Symbol]SLRAction
	GotoTable   map[int]map[Symbol]int
}

func NewSLRTable(productions []Production) SLRTable {
	table := SLRTable{
		ActionTable: map[int]map[Symbol]SLRAction{},
		GotoTable:   map[int]map[Symbol]int{},
	}

	productions = augment(productions)

	return table
}

func augment(productions []Production) []Production {
	first := productions[0].Head
	productions = append([]Production{{Head: "S'", Body: []Symbol{first}}}, productions...)
	return productions
}
