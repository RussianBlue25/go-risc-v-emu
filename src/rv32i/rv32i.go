package rv32i

import (
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
	"fmt"
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
	for i := 0; i < int(shift); i++ {
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> 1
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | (signedBit << 7)
	}
}

func Jal(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + 4
	cpu.Pc += inst.Imm
}

func Jalr(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + 4
	cpu.Pc = (cpu.Registers[inst.Rs1] + inst.Imm) & 0x00000FFE //12bit
	fmt.Printf("%x\n", cpu.Pc)
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

func Sll(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := cpu.Registers[inst.Rs2] & 0x1F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] << shift
}

func Srl(inst instruction.Instruction, cpu *cpu.Cpu) {
	shift := cpu.Registers[inst.Rs2] & 0x1F
	cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> shift
}

func Sra(inst instruction.Instruction, cpu *cpu.Cpu) {
	signedBit := cpu.Registers[inst.Rs2] & 0x800
	shift := cpu.Registers[inst.Rs2] & 0x8000
	for i := 0; i < int(shift); i++ {
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] >> 1
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs1] | (signedBit << 7)
	}
}

func Lui(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = inst.Imm
}

func Auipc(inst instruction.Instruction, cpu *cpu.Cpu) {
	cpu.Registers[inst.Rd] = cpu.Pc + inst.Imm
}

func Beq(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] == cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	}
}

func Bne(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] != cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	}
}

func Blt(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] < cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	}
}

func Bge(inst instruction.Instruction, cpu *cpu.Cpu) {
	if cpu.Registers[inst.Rs1] >= cpu.Registers[inst.Rs2] {
		cpu.Pc += inst.Imm
	}
}

func Bltu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) < uint32(cpu.Registers[inst.Rs2]) {
		cpu.Pc += inst.Imm
	}
}

func Bgeu(inst instruction.Instruction, cpu *cpu.Cpu) {
	if uint32(cpu.Registers[inst.Rs1]) >= uint32(cpu.Registers[inst.Rs2]) {
		cpu.Pc += inst.Imm
	}
}

//TODO: are these right implementation?
func Lb(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	cpu.Registers[inst.Rd] = int32(int8(mem[cpu.Registers[inst.Rs1]]))
}


func Lh(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	b := (mem[cpu.Registers[inst.Rs1]+1] << 8) | mem[cpu.Registers[inst.Rs1]]
	cpu.Registers[inst.Rd] = int32(int16(b))
}

func Lw(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	b := (mem[cpu.Registers[inst.Rs1]+3] << 24) | (mem[cpu.Registers[inst.Rs1]+2] << 16) | (mem[cpu.Registers[inst.Rs1]+1] << 8) | mem[cpu.Registers[inst.Rs1]]
	cpu.Registers[inst.Rd] = int32(b)
}

func Lbu(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	cpu.Registers[inst.Rd] = int32(uint32(int8(mem[cpu.Registers[inst.Rs1]])))
}

func Lhu(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	b := (mem[cpu.Registers[inst.Rs1]+1] << 8) | mem[cpu.Registers[inst.Rs1]]
	cpu.Registers[inst.Rd] = int32(uint32(int16(b)))
}

func Sb(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	mem[cpu.Registers[inst.Rs1]] = uint8(cpu.Registers[inst.Rs2] & 0x000000FF)
}

func Sh(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	mem[cpu.Registers[inst.Rs1]] = uint8(cpu.Registers[inst.Rs2] & 0x000000FF)
	mem[cpu.Registers[inst.Rs1]+1] = uint8((cpu.Registers[inst.Rs2] >> 8) & 0x000000FF )
}

func Sw(inst instruction.Instruction, cpu *cpu.Cpu, mem [1024*1024]uint8) {
	mem[cpu.Registers[inst.Rs1]] = uint8(cpu.Registers[inst.Rs2] & 0x000000FF)
	mem[cpu.Registers[inst.Rs1]+1] = uint8((cpu.Registers[inst.Rs2]  >> 8) & 0x000000FF)
	mem[cpu.Registers[inst.Rs1]+2] = uint8((cpu.Registers[inst.Rs2] >> 16) & 0x000000FF)
	mem[cpu.Registers[inst.Rs1]+3] = uint8((cpu.Registers[inst.Rs2] >> 24) & 0x000000FF)
}
