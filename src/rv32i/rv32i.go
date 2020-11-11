package rv32i

import (
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
)

func Addi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] + inst.Imm
}

func Slti(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < inst.Imm {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
}

func Andi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] & inst.Imm
}

func Ori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | inst.Imm
}

func Xori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] ^ inst.Imm
}

func Slli(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := inst.Imm & 0x01F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] << shift
}

func Srli(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := inst.Imm & 0x01F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> shift
}

func Srai(inst instruction.Instruction, cpu *cpu.Cpu) {
	signedBit := inst.Imm & 0x800
	shift := inst.Imm & 0x01F
	for i := 0; i < shift; i++ {
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rd] >> 1
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rd] | (signedBit << 7)
	}
}

func Jal(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + 4
	cpu.Pc += inst.Imm
}

func Jalr(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + 4
	cpu.Pc += (cpu.Registers[inst.Rs1] + inst.Imm) & 0xFFE //12bit
}

func Add(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] + cpu.Registers[inst.Rs2]
}

func Sub(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] - cpu.Registers[inst.Rs2]
}

func Sltu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) < uint32(cpu.Registers[inst.Rs2]) {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
}

func Slt(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < cpu.Registers[inst.Rs2] {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
}

func And(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] & cpu.Registers[inst.Rs2]
}

func Or(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | cpu.Registers[inst.Rs2]
}

func Xor(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] ^ cpu.Registers[inst.Rs2]
}




