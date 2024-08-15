package main

import "strings"

func createParamString(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

func assert(expected, recieved any, desc string) {
	if expected != recieved {
		panic(desc)
	}
}

func assertGreaterThan(expected, recieved int, desc string) {
	if expected < recieved {
		panic(desc)
	}
}

func getArgsString(tc TestCase) string {
	sb := strings.Builder{}

	for i, arg := range tc.arguments {
		sb.WriteString(arg)
		if i != len(tc.arguments)-1 {
			sb.WriteString(",")
		}
	}

	return sb.String()
}

// Creates a function body string by trimming empty space and appending the return variable.
func getBodyString(body, ret string) string {
	sb := strings.Builder{}
	sb.WriteString(strings.Trim(body, " \n\t"))
	sb.WriteString("\n\treturn ")
	sb.WriteString(ret)

	return sb.String()
}

func getParamString(params string, tc TestCase) string {
	sb := strings.Builder{}
	splitParams := strings.Split(params, " ")

	for i, p := range splitParams {
		sb.WriteString(p)
		if i%2 == 0 && i != 0 && i != len(splitParams)-1 {
			sb.WriteString(",")
		}
		sb.WriteString(" ")
	}
	sb.WriteString(") ")
	sb.WriteString(tc.expectedOutputType)
	sb.WriteString("{\n\t")

	return sb.String()
}
