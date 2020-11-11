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

	//TODO: implement type-aware processing
	// this is I type

	//TODO: implement cpu
	var inst instruction.Instruction
	cpu := cpu.Cpu{}

	var fetchedBinary uint32

	for {
		for i := 0; i < 4; i++ {
			fetchedBinary = fetchedBinary << 8
			fetchedBinary = fetchedBinary | uint32(Memory[cpu.Pc+uint32(i)])
		}
		//TODO: consider memory's last
		if fetchedBinary == 0x0000 {
			break
		}
		fmt.Printf("%x\n", fetchedBinary)
		cpu.Pc += 4
		inst = interpretInst(fetchedBinary)
		execute(inst, cpu)
	}

	fmt.Println(cpu.Registers[inst.Rd])
}

func interpretInst(fetchedBinary uint32) (inst instruction.Instruction) {
	opcode := int(fetchedBinary & 0x0000007F)
	var rd int
	var funct3 int
	var rs1 int
	var imm int

	if opcode == 19 {
		rd = int((fetchedBinary & 0x00000F80) >> 7)
		funct3 = int((fetchedBinary & 0x00007000) >> 12)
		rs1 = int((fetchedBinary & 0x000F8000) >> 15)
		imm = int((fetchedBinary & 0xFFF00000) >> 20)
	}
	return instruction.Instruction{Opcode: opcode, Rd: rd, Rs1: rs1, Funct3: funct3, Imm: imm}
}

func execute(inst instruction.Instruction, cpu cpu.Cpu) {
	switch inst.Opcode {
	case 19:
		switch inst.Funct3 {
			case 0:
				rv32i.Addi(inst, &cpu)
				fmt.Println("addi")
			case 1:
				fmt.Println("slli")
			case 2:
				rv32i.Slti(inst, &cpu)
				fmt.Println("slti")
			case 4:
				rv32i.Xori(inst, &cpu)
				fmt.Println("xori")
			case 5:
				fmt.Println("srli")
			case 6:
				rv32i.Ori(inst, &cpu)
				fmt.Println("ori")
			case 7:
				rv32i.Andi(inst, &cpu)
				fmt.Println("andi")
			default:
				fmt.Println("unknown")
		}
	}
}