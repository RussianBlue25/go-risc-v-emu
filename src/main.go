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

	for i := 0; i < 65536; i++{
		code = uint32(Memory[cpu.Pc]) | uint32(Memory[cpu.Pc+1])<<8 | uint32(Memory[cpu.Pc+2])<<16 | uint32(Memory[cpu.Pc+3])<<24
		//TODO: consider memory's last
		if code == 0x00000000 {
			continue
		}
		//fmt.Printf("%x\n", code)
		inst = decode(code)
		fmt.Printf("%x\n", cpu.Pc)
		fmt.Printf("Rs1 is %x\n", inst.Rs1)
		fmt.Printf("Rs2 is %x\n", inst.Rs2)
		fmt.Printf("Rd is %x\n", inst.Rd)
		fmt.Printf("Imm is %x\n", inst.Imm)
		fmt.Println(inst)
		execute(inst, &cpu, Memory)
		// zero register
		if cpu.Registers[0] != 0 {
			cpu.Registers[0] = 0
		}

		fmt.Println(cpu.Registers)
		cpu.Pc += 4
	}
}

func decode(code uint32) (inst instruction.Instruction) {
	opcode := uint8(code & 0x0000007F)
	var rd uint8
	var funct3 uint8
	var funct7 uint8
	var rs1 uint8
	var rs2 uint8
	var imm int32

	if opcode == 19 || opcode == 3 || opcode == 115 { //I format
		rd = uint8((code & 0x00000F80) >> 7)
		funct3 = uint8((code & 0x00007000) >> 12)
		rs1 = uint8((code & 0x000F8000) >> 15)
		imm = int32(code & 0xFFF00000) >> 20
	} else if opcode == 23 || opcode == 55 { //U format
		rd = uint8((code & 0x00000F80) >> 7)
		imm = (int32(code & 0xFFFFF000) >> 12) << 12
	} else if opcode == 35 { //S format
		imm1 := (code & 0x00000F80) >> 7
		funct3 = uint8((code & 0x00007000) >> 12)
		rs1 = uint8((code & 0x000F8000) >> 15)
		rs2 = uint8((code & 0x01F00000) >> 20)
		imm2 := (code & 0xFE000000) >> 25
		imm = int32((imm2 << 5) | imm1)
	} else if opcode == 51 {//R format
		rd = uint8((code & 0x00000F80) >> 7)
		funct3 = uint8((code & 0x00007000) >> 12)
		rs1 = uint8((code & 0x000F8000) >> 15)
		rs2 = uint8((code & 0x01F00000) >> 20)
		funct7 = uint8((code & 0xFE000000) >> 25)
	} else if opcode == 99 { //B format
		imm11 := (code & 0x00000080) >> 7
		imm1_4 := (code & 0x00000F00) >> 8
		funct3 = uint8((code & 0x00007000) >> 12)
		rs1 = uint8((code & 0x000F8000) >> 15)
		rs2 = uint8((code & 0x01F00000) >> 20)
		imm5_10 := (code & 0x7E000000) >> 25
		imm12 := (code & 0x80000000) >> 31
		imm = int32(imm12 << 12) | int32(imm11 << 11) | int32(imm5_10) << 5 | int32(imm1_4) << 1
	} else if opcode == 111 { //J format
		rd = uint8((code & 0x00000F80) >> 7)
		imm19_12 := (code & 0x000FF000) >> 12
		imm11 := (code & 0x00100000) >> 20
		imm10_1 := (code & 0x7FE00000) >> 21
		imm20 := (code & 0x80000000) >> 31
		imm = int32(imm20 << 20) | int32(imm19_12 << 12) | int32(imm11 << 11) | int32(imm10_1 << 1)
	} else {
		fmt.Println("unknown format!!")
	}
	return instruction.Instruction{Opcode: opcode, Rd: rd, Rs1: rs1, Rs2: rs2, Funct3: funct3, Funct7: funct7, Imm: imm}
}

func execute(inst instruction.Instruction, cpu *cpu.Cpu, mem [65536]uint8) {
	switch inst.Opcode {
	case 3:
		switch inst.Funct3 {
		case 0:
			rv32i.Lb(inst, cpu, mem)
			fmt.Println("lb")
		case 1:
			rv32i.Lh(inst, cpu, mem)
			fmt.Println("lh")
		case 2:
			rv32i.Lw(inst, cpu, mem)
			fmt.Println("lw")
		case 4:
			rv32i.Lbu(inst, cpu, mem)
			fmt.Println("lbu")
		case 5:
			rv32i.Lhu(inst, cpu, mem)
			fmt.Println("lhu")
		}
	case 19:
		switch inst.Funct3 {
		case 0:
			rv32i.Addi(inst, cpu)
			fmt.Println("addi")
		case 1:
			rv32i.Slli(inst, cpu)
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
				rv32i.Srli(inst, cpu)
				fmt.Println("srli")
			} else if shamt == 16 {
				rv32i.Srai(inst, cpu)
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
		rv32i.Lui(inst, cpu)
		fmt.Println("lui")
	case 35:
		switch inst.Funct3 {
		case 0:
			rv32i.Sb(inst, cpu, mem)
			fmt.Println("sb")
		case 1:
			rv32i.Sh(inst, cpu, mem)
			fmt.Println("sh")
		case 2:
			rv32i.Sw(inst, cpu, mem)
			fmt.Println("sw")
		}
	case 51:
		switch inst.Funct3 {
		case 0:
			if inst.Funct7 == 0 {
				rv32i.Add(inst, cpu)
				fmt.Println("add")
			} else if inst.Funct7 == 32 {
				rv32i.Sub(inst, cpu)
				fmt.Println("sub")
			} else {
				fmt.Println("unknown")
			}
		case 1:
			rv32i.Sll(inst, cpu)
			fmt.Println("sll")
		case 2:
			rv32i.Slt(inst, cpu)
			fmt.Println("slt")
		case 3:
			rv32i.Sltu(inst, cpu)
			fmt.Println("sltu")
		case 4:
			rv32i.Xor(inst, cpu)
			fmt.Println("xor")
		case 5:
			if inst.Funct7 == 0 {
				rv32i.Srl(inst, cpu)
				fmt.Println("srl")
			} else if inst.Funct7 == 32 {
				rv32i.Sra(inst, cpu)
				fmt.Println("sra")
			} else {
				fmt.Println("unknown")
			}
		case 6:
			rv32i.Or(inst, cpu)
			fmt.Println("or")
		case 7:
			rv32i.And(inst, cpu)
			fmt.Println("and")
		default:
			fmt.Println("unknown")
		}
	case 55:
		rv32i.Auipc(inst, cpu)
		fmt.Println("auipc")
	case 99:
		switch inst.Funct3 {
		case 0:
			rv32i.Beq(inst, cpu)
			fmt.Println("beq")
		case 1:
			rv32i.Bne(inst, cpu)
			fmt.Println("bne")
		case 4:
			rv32i.Blt(inst, cpu)
			fmt.Println("blt")
		case 5:
			rv32i.Bge(inst, cpu)
			fmt.Println("bge")
		case 6:
			rv32i.Bltu(inst, cpu)
			fmt.Println("bltu")
		case 7:
			rv32i.Bgeu(inst, cpu)
			fmt.Println("bgeu")
		}
	case 111:
		rv32i.Jal(inst, cpu)
		fmt.Println("jal")
	case 115:
		rv32i.Jalr(inst, cpu)
		fmt.Println("jalr")
	}
}