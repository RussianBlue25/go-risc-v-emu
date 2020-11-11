package main

import (
	"fmt"
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
	"github.com/RussianBlue25/go-risc-v-emu/src/rv32i"
	"github.com/RussianBlue25/go-risc-v-emu/src/elf"
)

func main() {
	var Memory [65536]uint8

	Memory = elf.ElfLoad(Memory)

	var inst instruction.Instruction
	cpu := cpu.Cpu{}

	var code uint32

	for {
		code = uint32(Memory[cpu.Pc]) | uint32(Memory[cpu.Pc+1])<<8 | uint32(Memory[cpu.Pc+2])<<16 | uint32(Memory[cpu.Pc+3])<<24
		//TODO: consider memory's last
		if code == 0x0000 {
			break
		}
		fmt.Printf("%x\n", code)
		cpu.Pc += 4
		inst = decode(code)
		fmt.Println(inst)
		execute(inst, &cpu)

		fmt.Println(cpu.Registers)
	}
}

func decode(code uint32) (inst instruction.Instruction) {
	opcode := int(code & 0x0000007F)
	var rd int
	var funct3 int
	var funct7 int
	var rs1 int
	var rs2 int
	var imm int

	if opcode == 19 || opcode == 3 || opcode == 115 { //I format
		rd = int((code & 0x00000F80) >> 7)
		funct3 = int((code & 0x00007000) >> 12)
		rs1 = int((code & 0x000F8000) >> 15)
		imm = int((code & 0xFFF00000) >> 20)
	} else if opcode == 23 || opcode == 55 { //U format
		rd = int((code & 0x00000F80) >> 7)
		imm = int((code & 0xFFFFF000) >> 12) << 12
	} else if opcode == 35 { //S format
		imm1 := int((code & 0x00000F80) >> 7)
		funct3 = int((code & 0x00007000) >> 12)
		rs1 = int((code & 0x000F8000) >> 15)
		rs2 = int((code & 0x01F00000) >> 20)
		imm2 := int((code & 0xFE000000) >> 25)
		imm = (imm2 << 5) | imm1
	} else if opcode == 51 {//R format
		rd = int((code & 0x00000F80) >> 7)
		funct3 = int((code & 0x00007000) >> 12)
		rs1 = int((code & 0x000F8000) >> 15)
		rs2 = int((code & 0x01F00000) >> 20)
		funct7 = int((code & 0xFE000000) >> 25)
	} else if opcode == 99 { //B format
		imm11 := int((code & 0x00000080) >> 7)
		imm1_4 := int((code & 0x00000F00) >> 8)
		funct3 = int((code & 0x00007000) >> 12)
		rs1 = int((code & 0x000F8000) >> 15)
		rs2 = int((code & 0x01F00000) >> 20)
		imm5_10 := int((code & 0x7E000000) >> 25)
		imm12 := int((code & 0x80000000) >> 31)
		imm = (imm12 << 12) | (imm11 << 11) | (imm5_10) << 5 | (imm1_4) << 1
	} else if opcode == 111 { //J format
		rd = int((code & 0x00000F80) >> 7)
		imm19_12 := int((code & 0x000FF000) >> 12)
		imm11 := int((code & 0x00100000) >> 20)
		imm10_1 := int((code & 0x7FE00000) >> 21)
		imm20 := int((code & 0x80000000) >> 31)
		imm = (imm20 << 20) | (imm19_12 << 12) | (imm11 << 11) | (imm10_1 << 1)
	} else {
		fmt.Println("unknown format!!")
	}
	return instruction.Instruction{Opcode: opcode, Rd: rd, Rs1: rs1, Rs2: rs2, Funct3: funct3, Funct7: funct7, Imm: imm}
}

func execute(inst instruction.Instruction, cpu *cpu.Cpu) {
	switch inst.Opcode {
	case 3:
		switch inst.Funct3 {
		case 0:
			fmt.Println("lb")
		case 1:
			fmt.Println("lh")
		case 2:
			fmt.Println("lw")
		case 4:
			fmt.Println("lbu")
		case 5:
			fmt.Println("lhu")
		}
	case 19:
		switch inst.Funct3 {
		case 0:
			rv32i.Addi(inst, cpu)
			fmt.Println("addi")
		case 1:
			fmt.Println("slli")
		case 2:
			rv32i.Slti(inst, cpu)
			fmt.Println("slti")
		case 4:
			rv32i.Xori(inst, cpu)
			fmt.Println("xori")
		case 5:
			shamt := ((inst.Imm & 0xFC0) >> 4)
			if shamt == 0 {
				fmt.Println("srli")
			} else if shamt == 16 {
				fmt.Println("srai")
			} else {
				fmt.Println("unknown")
			}
		case 6:
			rv32i.Ori(inst, cpu)
			fmt.Println("ori")
		case 7:
			rv32i.Andi(inst, cpu)
			fmt.Println("andi")
		default:
			fmt.Println("unknown")
		}
	case 23:
		fmt.Println("lui")
	case 35:
		switch inst.Funct3 {
		case 0:
			fmt.Println("sb")
		case 1:
			fmt.Println("sh")
		case 2:
			fmt.Println("sw")
		}
	case 51:
		switch inst.Funct3 {
		case 0:
			if inst.Funct7 == 0 {
				fmt.Println("add")
			} else if inst.Funct7 == 32 {
				fmt.Println("sub")
			} else {
				fmt.Println("unknown")
			}
		case 1:
			fmt.Println("sll")
		case 2:
			fmt.Println("slt")
		case 3:
			fmt.Println("sltu")
		case 4:
			fmt.Println("xor")
		case 5:
			if inst.Funct7 == 0 {
				fmt.Println("srl")
			} else if inst.Funct7 == 32 {
				fmt.Println("sra")
			} else {
				fmt.Println("unknown")
			}
		case 6:
			fmt.Println("or")
		case 7:
			fmt.Println("and")
		default:
			fmt.Println("unknown")
		}
	case 55:
		fmt.Println("auipc")
	case 99:
		switch inst.Funct3 {
		case 0:
			fmt.Println("beq")
		case 1:
			fmt.Println("bne")
		case 4:
			fmt.Println("blt")
		case 5:
			fmt.Println("bge")
		case 6:
			fmt.Println("bltu")
		case 7:
			fmt.Println("bgeu")
		}
	case 111:
		fmt.Println("jal")
	case 115:
		fmt.Println("jalr")
	}
}