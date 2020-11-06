package rv32i

import (
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
)

func Addi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] + inst.Imm
}

func Slti(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Register[inst.Rs1] < inst.Imm {
		cpu.Register[inst.Rd] = 1
	} else {
		cpu.Register[inst.Rd] = 0
	}
}

func Andi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] & inst.Imm
}

func Ori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] | inst.Imm
}

func Xori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] ^ inst.Imm
}

