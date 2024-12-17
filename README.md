# Pope: Point of Potential Error

**Pope** is a tool for dynamically managing and testing code that can have various points of potential error. It provides a flexible way to define, execute, and validate code logic with customizable parameters and error handling. 

## Objective
Separate good test cases from bad test cases. Define points of potential error in a function body using a simple templating syntax. Define paramters, return variables and test cases. In an ideal world we are looking to find if any permutation of predictable errors can also yeild the same output as a passing test case. 

## Key Features
Pope allows you to create and test functions with specific points of potential error (POPE) where different operations or logic variations can be tested. This can be particularly useful for ensuring code robustness and identifying how changes in logic can affect the outcome.

- **Testing**: Add test cases to automatically validate if function behavior can be generated through programming error.
- **Error Visualization**: Print your error tree and be mortified at all the possible ways you can stuff things up.

## Example Usage
Below is an example demonstrating how to define a function, set parameters, find potential errors, and test the function.

```go
package main

import (
    "os"
    "github.com/robert-ohurley/pope"  
)v


func main() {
    // Create a new function with the name "sumTwoNumbers"
    myFunc := pope.NewTFunc("sumTwoNumbers")

    // Define parameters for the function
    myFunc.SetParam("a", "int"})
    myFunc.SetParam("b", "int"})

    // Define the code body of the function with placeholders for potential errors.
    // You define a pope like this {{identifier:expectedValue}}
    myFunc.SetCodeBody(`
        a = (a {{id1:+}} b) {{id2:-}} b`)

    // Define the popes with their alternative values
    myFunc.AddPope("id1", []string{"-", "*", "/"})
    myFunc.AddPope("id2", []string{"+", "*", "/"})

    // Add the expected return values in the form of (expectedReturn, ...invalidReturns)
    myFunc.SetReturn("a", "b")

    // Add test cases for the function in the form of (expectedReturn, returnType, ...TestCaseParameters)
    myFunc.AddPassingTestCase("4", "int", "2", "2")

    // Execute the function by provided a writer to recieve your results of all permutations of the defined function.
    myFunc.Execute(os.Stdout)

    // Optionally, print and visualise the tree. 
    pope.PrintTree(myFunc.RootErrorNode)
}
