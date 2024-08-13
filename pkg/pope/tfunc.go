package main

type Tfunc struct {
	funcName         string
	funcParams       [][2]string
	funcCode         string
	popes            []*Pope
	errorFuncReturns []any
	validReturn      string
	invalidReturns   []string
	rootErrorNode    *errorNode
}

func NewTFunc() *Tfunc {
	return &Tfunc{
		funcName:         "",
		funcParams:       [][2]string{},
		funcCode:         "",
		popes:            []*Pope{},
		errorFuncReturns: []any{},
		validReturn:      "",
		invalidReturns:   []string{""},
		rootErrorNode:    NewErrorNode(NewErr(noErr, "", "")),
	}
}

// Replaces all popes with valid input but records the idxs as part of the pope to define the beginning and end
// of the substitution.

func (t *Tfunc) SetFuncName(s string) *Tfunc {
	t.funcName = s
	return t
}

func (t *Tfunc) SetParam(p [2]string) *Tfunc {
	t.funcParams = append(t.funcParams, p)
	return t
}

// Use templating syntax e.g. {{}} to add a point of potential error.
func (t *Tfunc) SetCodeBody(c string) *Tfunc {
	t.funcCode = c
	return t
}

// Adds a point of potential error (POPE) as defined in the code body using templating syntax.
func (t *Tfunc) AddPope(ident string, invalid []string) *Tfunc {
	t.popes = append(t.popes, NewPope(ident, invalid))
	return t
}

// Define the variable to be correctly returned from the function body. Include an array of values which might mistakenly be returned.
func (t *Tfunc) SetReturn(r string, errReturns []string) *Tfunc {
	t.validReturn = r
	t.invalidReturns = errReturns

	return t
}

// Define a set of parameters and a result which is a correct test case.
// We are aiming to determine if this passing test case would also be returned from a function body with a defined POPE
func (t *Tfunc) AddPassingTestCase() *Tfunc {
	return t
}

func (t *Tfunc) AddParamErrorsToTree() {
	v := NewVisitor(t.rootErrorNode)
	expectedParamString := ""

	for _, p := range t.funcParams {
		expectedParamString = createParamString(" ", expectedParamString, p[0], p[1])
	}

	for v.Walk() {
		if v.CurrNode().IsLeaf() {
			Permute(t.funcParams, func(a [][2]string) {
				recievedParamString := ""
				for _, p := range t.funcParams {
					recievedParamString = createParamString(" ", recievedParamString, p[0], p[1])
				}
				node := NewErrorNode(NewErr(paramErr, expectedParamString, recievedParamString))
				v.CurrNode().AddNext(node)
			})

		}
	}
}

func (t *Tfunc) AddReturnErrorsToTree() {
	allReturnVars := []string{}
	allReturnVars = append(allReturnVars, t.validReturn)
	allReturnVars = append(allReturnVars, t.invalidReturns...)

	v := NewVisitor(t.rootErrorNode)

	for v.Walk() {
		if v.CurrNode().IsLeaf() {
			for _, ret := range allReturnVars {
				node := NewErrorNode(NewErr(returnErr, t.validReturn, ret))
				v.CurrNode().AddNext(node)
			}

		}
	}

}

func (t *Tfunc) AddPopeErrorsToTree() {
	v := NewVisitor(t.rootErrorNode)
	p := NewParser(t.funcCode, t.popes)

	//set each pope to is correct value in the code body.
	//for every pope we walk the tree
	for _, pope := range p.popes {
		for v.Walk() {
			if v.CurrNode().IsLeaf() {
				codeBodies := p.GetPopeSubstitutions(pope)
				for _, code := range  

				node := NewErrorNode(NewErr(paramErr, "", ""))
				v.CurrNode().AddNext(node)
			}
		}
	}
}

func (t *Tfunc) CountPOPEPermutations() {
	t.AddParamErrorsToTree()
	t.AddReturnErrorsToTree()
	t.AddPopeErrorsToTree()
}
