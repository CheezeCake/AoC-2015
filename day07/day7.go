package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	operator      string
	op1, op2      string
	valueComputed bool
	value         uint16
}

func newOperation(cmd string) (string, *operation) {
	words := strings.Split(cmd, " ")
	wire := words[len(words)-1]

	if words[1] == "->" {
		return wire, &operation{op1: words[0]}
	} else if words[0] == "NOT" {
		return wire, &operation{operator: words[0], op1: words[1]}
	} else {
		return wire, &operation{op1: words[0], operator: words[1], op2: words[2]}
	}

	panic("error while parsing command: " + cmd)
}

type circuit map[string]*operation

func newCircuit() *circuit {
	c := make(circuit)
	return &c
}

func (c *circuit) addWire(wire string, operation *operation) {
	(*c)[wire] = operation
}

func (c *circuit) reset() {
	for _, operation := range *c {
		operation.valueComputed = false
	}
}

func (c *circuit) setWireValue(wire string, value uint16) {
	(*c)[wire].value = value
	(*c)[wire].valueComputed = true
}

func (c *circuit) wireValue(wire string) uint16 {
	operation, ok := (*c)[wire]
	if !ok {
		imm, err := strconv.Atoi(wire)
		if err != nil {
			panic("invalid wire: " + wire)
		}
		return uint16(imm)
	}

	if operation.valueComputed {
		return operation.value
	}

	var value uint16
	switch operation.operator {
	case "":
		value = c.wireValue(operation.op1)
	case "NOT":
		value = ^c.wireValue(operation.op1)
	case "AND":
		value = (c.wireValue(operation.op1) & c.wireValue(operation.op2))
	case "OR":
		value = (c.wireValue(operation.op1) | c.wireValue(operation.op2))
	case "LSHIFT":
		value = (c.wireValue(operation.op1) << c.wireValue(operation.op2))
	case "RSHIFT":
		value = (c.wireValue(operation.op1) >> c.wireValue(operation.op2))
	}

	c.setWireValue(wire, value)
	return value
}

func main() {
	circuit := newCircuit()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wire, op := newOperation(scanner.Text())
		circuit.addWire(wire, op)
	}

	a := circuit.wireValue("a")
	fmt.Println("part 1:", a)

	circuit.reset()
	circuit.setWireValue("b", a)
	fmt.Println("part 2:", circuit.wireValue("a"))
}
