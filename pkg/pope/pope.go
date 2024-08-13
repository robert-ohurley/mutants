package main

type Pope struct {
	identifier   string
	validInput   string
	invalidInput []string
	startIdx     int
	endIdx       int
}

// in template pope will look like {{ident:valid}}
func NewPope(ident string, invalid []string) *Pope {
	return &Pope{
		identifier:   ident,
		invalidInput: invalid,
	}
}

func (p *Pope) TemplateString() string {
	return "{{" + p.identifier + ":" + p.validInput + "}}"
}

func (p *Pope) SetIdx(begin, end int) {
	p.startIdx = begin
	p.endIdx = end
}

func (p *Pope) SetValidInput(s string) {
	p.validInput = s
}
