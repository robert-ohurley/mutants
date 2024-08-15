package main

import "strings"

type parser struct {
	code  string
	popes []*Pope
}

func NewParser(code string, p []*Pope) *parser {
	newParser := &parser{code, p}
	newParser.DetermineTemplateRanges()
	newParser.ParseTemplateStrings()

	return newParser
}

// determines and sets the valid inputs from a pope
func (p *parser) ParseTemplateStrings() {
	for _, pope := range p.popes {
		templString := p.code[pope.startIdx+2 : pope.endIdx-2]
		parts := strings.Split(templString, ":")
		assert(len(parts), 2, "There should only be two parts in a template string")
		pope.SetValidInput(parts[1])
	}
}

// Sets the start and end idx of each pope in the code body
func (p *parser) DetermineTemplateRanges() {
	for _, pope := range p.popes {
		startIdx := strings.Index(p.code, pope.identifier)
		startIdx = startIdx - 2

		if startIdx == -1 {
			assertGreaterThan(startIdx, 0, "Start idx should be greater than zero")
		}

		parsing := true
		endIdx := startIdx

		for parsing == true {
			if p.code[endIdx] == '}' && p.code[endIdx+1] == '}' {
				endIdx = endIdx + 1 //Ends the char after }
				parsing = false
			}
			endIdx += 1
		}

		pope.SetIdx(startIdx, endIdx)
	}
}

func (p *parser) CreateAllCodePermutations() []string {
	codeBodies := []string{}

	validStr := CreatePopeSubstitution(p.code, p.popes[0].TemplateString(), p.popes[0].validInput)
	codeBodies = append(codeBodies, validStr)

	for _, invInput := range p.popes[0].invalidInput {
		invalidStr := CreatePopeSubstitution(p.code, p.popes[0].TemplateString(), invInput)
		codeBodies = append(codeBodies, invalidStr)

	}

	for _, pope := range p.popes[1:] {
		cloneSlice := []string{}
		for _, code := range codeBodies {
			validStr := CreatePopeSubstitution(code, pope.TemplateString(), pope.validInput)
			cloneSlice = append(cloneSlice, validStr)

			for _, invInput := range pope.invalidInput {
				invalidStr := CreatePopeSubstitution(code, pope.TemplateString(), invInput)
				cloneSlice = append(cloneSlice, invalidStr)
			}
		}
		codeBodies = cloneSlice
	}

	return codeBodies
}

// Creates a code body where a specific pope is set to its valid value
// Get tmplStr from pope.TemplateString()
func CreatePopeSubstitution(codeBody, templStr, sub string) string {
	validStr := strings.Replace(codeBody, templStr, sub, 1)
	return validStr
}
