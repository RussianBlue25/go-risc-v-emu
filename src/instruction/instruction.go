package instruction

type Instruction struct {
	Opcode int
	Rd     int
	Rs1    int
	Rs2	   int
	Funct3 int
	Funct7 int
	Imm    int
}