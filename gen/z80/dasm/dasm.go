package main

//go:generate go run .
//go:generate go fmt ../../../rcs/z80/dasm.go
//go:generate go fmt ../../../rcs/z80/dasm_harston_test.go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

var (
	root      = filepath.Join("..", "..", "..")
	targetDir = filepath.Join(root, "rcs", "z80")
	sourceDir = filepath.Join(root, "ext", "harston")
)

const (
	lineStart = 7
	lineEnd   = 262
)

var breaks = []int{
	3,
	7,
	18,
	22,
	33,
	38,
	45,
	49,
	62,
	67,
	80,
}

func dasm() {
	var out bytes.Buffer

	out.WriteString(`
// Code generated by gen/z80/dasm/dasm.go. DO NOT EDIT.

package z80

import "github.com/blackchip-org/retro-cs/rcs"

`)

	listFile := filepath.Join(sourceDir, "z80oplist.txt")
	data, err := ioutil.ReadFile(listFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	// unprefixed
	out.WriteString("var dasmTable = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 0, "")
	}
	out.WriteString("}\n")

	// dd prefix
	out.WriteString("var dasmTableDD = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 2, "dd")
	}
	out.WriteString("}\n")

	// fd prefix
	out.WriteString("var dasmTableFD = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 2, "fd")
	}
	out.WriteString("}\n")

	// cb prefix
	out.WriteString("var dasmTableCB = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 4, "cb")
	}
	out.WriteString("}\n")

	// fd cb prefix
	out.WriteString("var dasmTableFDCB = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 6, "fdcb")
	}
	out.WriteString("}\n")

	// dd cb prefix
	out.WriteString("var dasmTableDDCB = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 6, "ddcb")
	}
	out.WriteString("}\n")

	// ed prefix
	out.WriteString("var dasmTableED = map[uint8]func(rcs.Eval){\n")
	for i := lineStart; i <= lineEnd; i++ {
		line := lines[i]
		parseTable(&out, line, 8, "ed")
	}
	out.WriteString("}\n")

	outFile := filepath.Join(targetDir, "dasm.go")
	err = ioutil.WriteFile(outFile, out.Bytes(), 0644)
	if err != nil {
		fmt.Printf("unable to write file: %v", err)
		os.Exit(1)
	}
}

func parseTable(out *bytes.Buffer, line string, firstBreak int, prefix string) {
	break1 := breaks[firstBreak]
	break2 := breaks[firstBreak+1]
	break3 := breaks[firstBreak+2]

	// Ensure the line is at least 80 characters long and then we don't
	// have to worry about different line lengths when slicing.
	line = fmt.Sprintf("%-80s", line)

	strOpcode := strings.TrimSpace(line[0:2])
	opcode, _ := strconv.ParseUint(strOpcode, 16, 8)

	switch {
	case prefix == "" && opcode == 0xcb:
		out.WriteString("0xcb: func(e rcs.Eval) { opCB(e) },\n")
		return
	case prefix == "" && opcode == 0xdd:
		out.WriteString("0xdd: func(e rcs.Eval) { opDD(e) },\n")
		return
	case prefix == "" && opcode == 0xed:
		out.WriteString("0xed: func(e rcs.Eval) { opED(e) },\n")
		return
	case prefix == "" && opcode == 0xfd:
		out.WriteString("0xfd: func(e rcs.Eval) { opFD(e) },\n")
		return
	case
		prefix == "dd" && opcode == 0xcb,
		prefix == "dd" && opcode == 0xdd,
		prefix == "dd" && opcode == 0xed,
		prefix == "dd" && opcode == 0xfd,
		prefix == "fd" && opcode == 0xcb,
		prefix == "fd" && opcode == 0xdd,
		prefix == "fd" && opcode == 0xed,
		prefix == "fd" && opcode == 0xfd:
		return
	}

	out.WriteString("0x")
	out.WriteString(fmt.Sprintf("%02x", opcode))
	if prefix == "ddcb" || prefix == "fdcb" {
		out.WriteString(": func(e rcs.Eval) { op2(e, ")
	} else {
		out.WriteString(": func(e rcs.Eval) { op1(e, ")
	}

	args := make([]string, 1)
	args[0] = strings.TrimSpace(line[break1:break2])

	switch {
	case strings.HasPrefix(args[0], "MOS_"):
		args[0] = "-"
	case strings.HasPrefix(args[0], "ED_"):
		args[0] = "-"
	case unicode.IsLower(rune(args[0][0])):
		args[0] = "-"
	case args[0][0] == '[':
		args[0] = "-"
	}

	if args[0] == "-" {
		out.WriteString(fmt.Sprintf(`"?%v%02x"`, prefix, opcode))
	} else {
		args[0] = `"` + args[0] + `"`
		fields := strings.Split(line[break2:break3], ",")
		for _, field := range fields {
			args = append(args, `"`+strings.TrimSpace(field)+`"`)
		}
		entry := strings.Join(args, ",")
		if prefix == "fd" {
			entry = strings.Replace(entry, "IX", "IY", -1)
		}
		if prefix == "ddcb" {
			entry = strings.Replace(entry, "IY", "IX", -1)
		}
		out.WriteString(strings.ToLower(entry))
	}

	out.WriteString(") },\n")
}

func harston() {
	var out bytes.Buffer

	out.WriteString(`
// Code generated by gen/z80/dasm/dasm.go. DO NOT EDIT.

package z80

var harstonTests = []harstonTest{
`)

	data, err := ioutil.ReadFile("expected.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		if line[0] == '=' {
			break
		}
		if line[0] == '#' {
			continue
		}
		data := strings.Split(line, " ")
		strdata := strings.Join(data, " ")
		hexdata := "0x" + strings.Join(data, ", 0x")
		i++
		op := lines[i]
		out.WriteString(fmt.Sprintf(`harstonTest{"%v", "%v", []uint8{%v}},`, strdata, op, hexdata))
		out.WriteString("\n")
	}
	out.WriteString("}\n")

	outFile := filepath.Join(targetDir, "dasm_harston_test.go")
	err = ioutil.WriteFile(outFile, out.Bytes(), 0644)
	if err != nil {
		fmt.Printf("unable to write file: %v", err)
		os.Exit(1)
	}
}

func main() {
	dasm()
	harston()
}
