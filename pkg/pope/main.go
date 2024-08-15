package main

import "os"

func main() {
	myFunc := NewTFunc("sumTwoNumbers")
	myFunc.SetParam([2]string{"a", "int"})
	myFunc.SetParam([2]string{"b", "int"})

	myFunc.SetCodeBody(`
            a = (a {{id1:+}} b) {{id2:-}} b`)

	myFunc.AddPope("id1", []string{"-", "*", "/"})
	myFunc.AddPope("id2", []string{"+", "*", "/"})

	myFunc.SetReturn("a", "b")
	myFunc.AddPassingTestCase("4", "int", "2", "2")

	myFunc.Execute(os.Stdout)
	PrintTree(myFunc.rootErrorNode)

	//PrintTree(myFunc.rootErrorNode)
}
