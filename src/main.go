package main

import (
	"encoding/binary"
	"fmt"
	"github.com/RussianBlue25/go-risc-v-emu/src/cpu"
	"github.com/RussianBlue25/go-risc-v-emu/src/instruction"
	"github.com/RussianBlue25/go-risc-v-emu/src/rv32i"
	"os"
)

func main() {
	//TODO: implement type-aware processing
	// this is I type
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open file")
		panic(err)
	}

	//TODO: implement cpu
	var fetchedBinary uint32
	var inst instruction.Instruction

	for {
		errb := binary.Read(file, binary.BigEndian, &fetchedBinary)
		if errb != nil {
			fmt.Println("can't read binary")
			break
		}
		fmt.Printf("%x\n", fetchedBinary)
		inst = interpretInst(fetchedBinary)
	}
	cpu := cpu.Cpu{}

	switch inst.Opcode {
	case 19:
		switch inst.Funct3 {
		case 0:
			rv32i.Addi(inst, &cpu)
		case 2:
			rv32i.Slti(inst, &cpu)
		case 4:
			rv32i.Xori(inst, &cpu)
		case 6:
			rv32i.Ori(inst, &cpu)
		case 7:
			rv32i.Andi(inst, &cpu)
		}
	}
	fmt.Println(cpu.Register[inst.Rd])
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
