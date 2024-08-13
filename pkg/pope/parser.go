package main

import (
	"fmt"
	"strings"
)

// in template pope will look like {{ident:valid}}
type parser struct {
	code  string
	popes []*Pope
}

func NewParser(code string, p []*Pope) *parser {
	newParser := &parser{code, p}
	newParser.DetermineTemplateRanges()
	newParser.ParseTemplateStrings()

	for _, pope := range p {
		fmt.Println("templ strings", code[pope.startIdx:pope.endIdx])
	}
	// newParser.SetCorrectCodeBody()
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

// Returns all possible function bodies given a single pope.
// All popes valid and invalid values are substuted.
// All other popes are set to valid value.
func (p *parser) GetPopeSubstitutions(pope *Pope) []string {
	funcStrs := []string{}

	//create valid substition first
	validStr := strings.Replace(p.code, pope.TemplateString(), pope.validInput, 1)
	validStr = p.SetAllRemainingPopes(validStr, pope)
	funcStrs = append(funcStrs, validStr)

	//create invalid substitions
	for _, sub := range pope.invalidInput {
		//set desired pope to invalid input
		invalidStr := strings.Replace(p.code, pope.TemplateString(), sub, 1)

		//set the remaining popes to their valid input
		invalidStr = p.SetAllRemainingPopes(invalidStr, pope)
		funcStrs = append(funcStrs, invalidStr)
	}
	return funcStrs
}

// Skips the pope we're currently examining. Subs all remaining popes to their valid value.
func (p *parser) SetAllRemainingPopes(code string, exluding *Pope) string {
	clone := code
	for _, pope := range p.popes {
		if pope.identifier == exluding.identifier {
			continue
		}
		clone = strings.Replace(clone, pope.TemplateString(), pope.validInput, 1)
	}
	return clone
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
