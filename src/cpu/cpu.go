package cpu

type Cpu struct {
	Registers [32]int
	Pc       uint32
}