package main

// An error template is an object which allows for formatting of a code body.
// Use templating syntax to provide a template e.g. {{p1:+}} in the code body.
// Where a represents an identifier and b is the correct value for this point in the code.

// Replacements are what might be mistakenly coded in it's place e.g. * instead of +

//We are trying to generate all possible incorrect code bodies.
//Then we give a positive test case and see if the inputs for the positive test case match any of the error templates

func main() {
	sum := NewTFunc()
	sum.SetFuncName("sum")
	sum.SetParam([2]string{"a", "int"})
	sum.SetParam([2]string{"b", "int"})
	sum.SetParam([2]string{"c", "int"})

	sum.SetCodeBody(`
		a = (a {{id1:+}} b) {{id2:*}} b
		c = a {{id3:-}} 5
    `)
	sum.SetReturn("c", []string{"a", "b"})

	sum.AddPope("id1", []string{"-", "*"})
	sum.AddPope("id2", []string{"-", "+"})
	sum.AddPope("id3", []string{"*", "+"})

	sum.CountPOPEPermutations()

	PrintTree(sum.rootErrorNode)
}
