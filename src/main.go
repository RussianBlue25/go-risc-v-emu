package main

import (
  "fmt"
  "os"
  "encoding/binary"
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
  cp int
}

func main() {
  //TODO: implement type-aware processing
  // this is I type
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("can't open file")
    panic(err)
  }

  //TODO: implement cpu
  var fetched_binary uint32
  var inst Instruction

  for {
    errb := binary.Read(file, binary.BigEndian, &fetched_binary)
    if errb != nil {
      fmt.Println("can't read binary")
      break
    }
    fmt.Printf("%x\n", fetched_binary)
    inst = interpret_inst(fetched_binary)
  }
  cpu := Cpu{}

  switch inst.opcode {
    case 19:
      switch inst.funct3 {
        case 0:
          addi(inst, &cpu)
      }
  }
  fmt.Println(cpu.register[inst.rd])
}

func interpret_inst(fetched_binary uint32) (inst Instruction) {
  opcode := int(fetched_binary & 0x0000007F)
  var rd int
  var funct3 int
  var rs1 int
  var imm int

  if opcode == 19 {
    rd = int((fetched_binary & 0x00000F80) >> 7)
    funct3 = int((fetched_binary & 0x00007000) >> 12)
    rs1 = int((fetched_binary & 0x000F8000) >> 15)
    imm = int((fetched_binary & 0xFFF00000) >> 20)
  }
  return Instruction{opcode: opcode, rd: rd, rs1: rs1, funct3: funct3, imm: imm}
}


func addi(inst Instruction, cpu *Cpu) {
  cpu.register[inst.rd] = cpu.register[inst.rs1] + inst.imm
  fmt.Println(cpu.register[inst.rd])
}
