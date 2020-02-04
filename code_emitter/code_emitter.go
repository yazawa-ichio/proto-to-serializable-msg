package code_emitter

import (
	"bytes"
	"fmt"
	"strings"
)

type indentMode int

const (
	indentTab indentMode = iota
	indent4Space
	indent2Space
)

type CodeEmitter struct {
	indentCount int
	indentMode  indentMode
	buf         bytes.Buffer
}

func (ce *CodeEmitter) Reset() {
	ce.indentCount = 0
	ce.buf = *new(bytes.Buffer)
}

func (ce *CodeEmitter) f(text string, p ...interface{}) {
	if len(p) == 0 {
		fmt.Fprintf(&ce.buf, text)
	} else {
		fmt.Fprintf(&ce.buf, text, p...)
	}
}

func (ce *CodeEmitter) StartIndent() {
	ce.indentCount++
}

func (ce *CodeEmitter) EndIndent() {
	ce.indentCount--
}

func (ce *CodeEmitter) EmitTab() {
	tab := "\t"
	if &ce.indentMode != nil {
		switch ce.indentMode {
		case indent4Space:
			tab = "    "
			break
		case indent2Space:
			tab = "  "
			break
		case indentTab:
		default:
			tab = "\t"
			break
		}
	}
	ce.buf.WriteString(strings.Repeat(tab, ce.indentCount))
}

func (ce *CodeEmitter) EmitIndentLine(text string, p ...interface{}) {
	ce.indentCount++
	ce.EmitLine(text, p...)
	ce.indentCount--
}

func (ce *CodeEmitter) Emit(text string, p ...interface{}) {
	ce.EmitTab()
	ce.f(text, p...)
}

func (ce *CodeEmitter) EmitLine(text string, p ...interface{}) {
	ce.EmitTab()
	ce.f(text, p...)
	ce.buf.WriteString("\n")
}

func (ce *CodeEmitter) EmitAppend(text string, p ...interface{}) {
	ce.f(text, p...)
}

func (ce *CodeEmitter) NewLine() {
	ce.buf.WriteString("\n")
}

func (ce *CodeEmitter) Bracket(text string, f func(), p ...interface{}) {
	ce.StartBracket(text, p...)
	f()
	ce.EndBracket("")
}

func (ce *CodeEmitter) StartBracket(text string, p ...interface{}) {
	ce.EmitTab()
	ce.f(text, p...)
	ce.buf.WriteString("{\n")
	ce.StartIndent()
}

func (ce *CodeEmitter) EndBracket(text string, p ...interface{}) {
	ce.EndIndent()
	ce.EmitTab()
	ce.buf.WriteString("}")
	ce.f(text, p...)
	ce.buf.WriteString("\n")
}

func (ce *CodeEmitter) EndAndStartBracket(text string, p ...interface{}) {
	ce.EndIndent()
	ce.EmitTab()
	ce.buf.WriteString("}")
	ce.f(text, p...)
	ce.buf.WriteString("{\n")
	ce.StartIndent()
}

func (ce *CodeEmitter) String() string {
	return ce.buf.String()
}
