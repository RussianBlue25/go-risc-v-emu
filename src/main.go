package main

import (
  "fmt"
)

type Instruction struct {
  opcode int
  rd int
  rs1 int
  funct3 int
  imm int
}

type Cpu struct {
  register[32] int
}

func main() {
  //TODO: implement type-aware processing
  // this is I type
  inst := Instruction{opcode: 19, rd: 5, rs1: 6, funct3: 0, imm: 1}
  cpu := Cpu{}

  switch inst.opcode {
    case 19:
      switch inst.funct3 {
        case 0:
          addi(inst, &cpu)
      }
  }
}


func addi(inst Instruction, cpu *Cpu) {
  cpu.register[inst.rd] = cpu.register[inst.rs1] + inst.imm
  fmt.Println(cpu.register[inst.rd])
}
