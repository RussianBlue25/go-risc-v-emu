package instruction

type Instruction struct {
	Opcode int
	Rd     int
	Rs1    int
	Funct3 int
	Imm    int
}