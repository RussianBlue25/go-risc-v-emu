package instruction

type Instruction struct {
	Opcode uint8
	Rd     uint8
	Rs1    uint8
	Rs2	   uint8
	Funct3 uint8
	Funct7 uint8
	Imm    int32
}