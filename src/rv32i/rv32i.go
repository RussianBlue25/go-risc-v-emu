package rv32i

import (
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
)

func Addi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] + inst.Imm
	cpu.Pc += 4
}

func Slti(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < inst.Imm {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
	cpu.Pc += 4
}

func Andi(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] & inst.Imm
	cpu.Pc += 4
}

func Ori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | inst.Imm
	cpu.Pc += 4
}

func Xori(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] ^ inst.Imm
	cpu.Pc += 4
}

func Slli(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := inst.Imm & 0x01F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] << shift
	cpu.Pc += 4
}

func Srli(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := inst.Imm & 0x01F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> shift
	cpu.Pc += 4
}

func Srai(inst instruction.Instruction, cpu *cpu.Cpu) {
	signedBit := inst.Imm & 0x800
	shift := inst.Imm & 0x01F
	for i := 0; i < int(shift); i++ {
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> 1
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | (signedBit << 7)
	}
	cpu.Pc += 4
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
	cpu.Pc += 4
}

func Sub(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] - cpu.Registers[inst.Rs2]
	cpu.Pc += 4
}

func Sltu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) < uint32(cpu.Registers[inst.Rs2]) {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
	cpu.Pc += 4
}

func Slt(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < cpu.Registers[inst.Rs2] {
		cpu.Registers[inst.Rd] = 1
	} else {
		cpu.Registers[inst.Rd] = 0
	}
	cpu.Pc += 4
}

func And(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] & cpu.Registers[inst.Rs2]
	cpu.Pc += 4
}

func Or(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | cpu.Registers[inst.Rs2]
	cpu.Pc += 4
}

func Xor(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] ^ cpu.Registers[inst.Rs2]
	cpu.Pc += 4
}

func Sll(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := cpu.Registers[inst.Rs2] & 0x1F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] << shift
	cpu.Pc += 4
}

func Srl(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := cpu.Registers[inst.Rs2] & 0x1F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> shift
	cpu.Pc += 4
}

func Sra(inst instruction.Instruction, cpu *cpu.Cpu) {
	signedBit := cpu.Registers[inst.Rs2] & 0x800
	shift := cpu.Registers[inst.Rs2] & 0x8000
	for i := 0; i < int(shift); i++ {
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> 1
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | (signedBit << 7)
	}
	cpu.Pc += 4
}

func Lui(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = (inst.Imm << 12)
	cpu.Pc += 4
}

func Auipc(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + (inst.Imm << 12)
	cpu.Pc += 4
}

func Beq(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] == cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}

func Bne(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] != cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}

func Blt(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}

func Bge(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] >= cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}

func Bltu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) < uint32(cpu.Registers[inst.Rs2]) {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}

func Bgeu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) >= uint32(cpu.Registers[inst.Rs2]) {
		cpu.Pc += inst.Imm
	} else {
		cpu.Pc += 4
	}
}