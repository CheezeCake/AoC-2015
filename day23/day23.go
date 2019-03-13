package main

import (
	"fmt"
	"io"
)

type Instr struct {
	opcode string
	reg    byte
	offset int
}

type VM struct {
	registers map[byte]int
	pc        int
}

func newVM() VM {
	return VM{map[byte]int{'a': 0, 'b': 0}, 0}
}

func (vm *VM) exec(program []Instr) {
	vm.pc = 0

	for vm.pc >= 0 && vm.pc < len(program) {
		i := program[vm.pc]

		switch i.opcode {
		case "hlf":
			vm.registers[i.reg] /= 2
		case "tpl":
			vm.registers[i.reg] *= 3
		case "inc":
			vm.registers[i.reg]++
		case "jmp":
			vm.pc += i.offset - 1
		case "jie":
			if vm.registers[i.reg]%2 == 0 {
				vm.pc += i.offset - 1
			}
		case "jio":
			if vm.registers[i.reg] == 1 {
				vm.pc += i.offset - 1
			}
		}

		vm.pc++
	}
}

func main() {
	program := []Instr{}
	for {
		var i Instr
		fmt.Scanf("%s", &i.opcode)

		var err error
		switch i.opcode {
		case "jmp":
			_, err = fmt.Scanf("%d\n", &i.offset)
		default:
			_, err = fmt.Scanf("%c, %d\n", &i.reg, &i.offset)
		}
		if err == io.EOF {
			break
		}
		program = append(program, i)
	}

	vm := newVM()
	vm.exec(program)
	fmt.Println("part 1:", vm.registers['b'])

	vm.registers['a'] = 1
	vm.registers['b'] = 0
	vm.exec(program)
	fmt.Println("part 2:", vm.registers['b'])
}
