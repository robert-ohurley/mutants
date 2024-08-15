package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

type Tfunc struct {
	funcName       string
	funcParams     [][2]string
	funcCode       string
	popes          []*Pope
	validReturn    string
	errorStack     map[ErrorType]*errorNode
	testCase       TestCase
	invalidReturns []string
	rootErrorNode  *errorNode
}

func NewTFunc(name string) *Tfunc {
	return &Tfunc{
		funcName:       name,
		funcParams:     [][2]string{},
		funcCode:       "",
		popes:          []*Pope{},
		errorStack:     make(map[ErrorType]*errorNode),
		validReturn:    "",
		testCase:       TestCase{},
		invalidReturns: []string{""},
		rootErrorNode:  NewErrorNode(NewErr(noErr, "", "")),
	}
}

// Replaces all popes with valid input but records the idxs as part of the pope to define the beginning and end
// of the substitution.

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
func (t *Tfunc) SetReturn(correctOut string, errReturns ...string) *Tfunc {
	t.validReturn = correctOut
	t.invalidReturns = errReturns

	return t
}

// Define a set of parameters and a result which is a correct test case.
// We are aiming to determine if this passing test case would also be returned from a function body with a defined POPE
func (t *Tfunc) AddPassingTestCase(correctOut string, correctOutType string, arguments ...string) *Tfunc {
	tc := TestCase{
		expectedOutput:     correctOut,
		expectedOutputType: correctOutType,
		arguments:          arguments,
	}
	t.testCase = tc
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
				node := NewErrorNode(NewErr(paramError, expectedParamString, recievedParamString))
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
				node := NewErrorNode(NewErr(returnError, t.validReturn, ret))
				v.CurrNode().AddNext(node)
			}

		}
	}

}

func (t *Tfunc) AddPopeErrorsToTree() {
	v := NewVisitor(t.rootErrorNode)
	p := NewParser(t.funcCode, t.popes)
	expectedBody := t.ExpectedFuncBody()

	for v.Walk() {
		if v.CurrNode().IsLeaf() {
			codeBodies := p.CreateAllCodePermutations()
			for _, code := range codeBodies {
				node := NewErrorNode(NewErr(logicError, expectedBody, code))
				v.CurrNode().AddNext(node)
			}
		}
	}
}

// returns the completed correct code body
func (t *Tfunc) ExpectedFuncBody() string {
	funcBody := t.funcCode

	for _, pope := range t.popes {
		funcBody = CreatePopeSubstitution(funcBody, pope.TemplateString(), pope.validInput)
	}
	return funcBody
}

func (t *Tfunc) CreateErrorTree() {
	t.AddParamErrorsToTree()
	t.AddReturnErrorsToTree()
	t.AddPopeErrorsToTree()
}

func (t *Tfunc) execute(node *errorNode, w io.Writer) {
	if node == nil {
		return
	}

	t.errorStack[node.err.errorType] = node

	if node.IsLeaf() {
		t.Write()
		node.Error()
		t.osExec(w)
		os.Remove("temp/main.go")
	}

	for _, n := range node.next {
		t.execute(n, w)
	}

}

func (t *Tfunc) osExec(w io.Writer) {
	cmd := exec.Command("go", "run", "temp/main.go")

	// if w == nil {
	// 	log, err := os.Create("output.log")
	// 	defer log.Close()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	cmd.Stdout = log
	// 	cmd.Stderr = log
	// }

	cmd.Stdout = w
	cmd.Stderr = w

	cmd.Start()
	cmd.Wait()
}

// Writes a Tfunc permutation to a .go file. The state is determined by the current stack of errors which is created as we traverse the tree.
func (t *Tfunc) Write() {
	sb := strings.Builder{}

	os.MkdirAll("temp", 0755)

	sb.WriteString("package main\nimport \"fmt\"\nfunc main() {\n\tfmt.Println(func(")
	sb.WriteString(getParamString(t.errorStack[paramError].err.recieved, t.testCase))
	sb.WriteString(getBodyString(t.errorStack[logicError].err.recieved, t.errorStack[returnError].err.recieved))
	sb.WriteString("\n\t}(")
	sb.WriteString(getArgsString(t.testCase))
	sb.WriteString("))\n}")

	os.WriteFile("temp/main.go", []byte(sb.String()), 0755)
}

// Creates all possible permutations of the test function and executes them all.
// Return values are written to the provided writer, if no writer is provided.
// Return values will be written to output.log
func (t *Tfunc) Execute(w io.Writer) {
	t.CreateErrorTree()
	//t.execute(t.rootErrorNode, w)
}
