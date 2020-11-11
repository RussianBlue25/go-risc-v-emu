package cpu

type Cpu struct {
	Registers [32]int32
	Pc       int32
}