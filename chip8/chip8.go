package chip8

import "chip8-emulator/graphics"

const (
	memorySize        = 4096
	variableRegisters = 0xF
)

type Chip8 struct {
	/* Peripherals */
	Memory   [memorySize]byte
	Graphics graphics.Graphics

	/* Registers */
	PC       uint16
	Index    uint16
	Stack    uint16
	Delay    uint8
	Sound    uint8
	Variable [variableRegisters]byte
}

//func New(rom []byte, graphics graphics.Graphics) *Chip8 {
//	return &Chip8{}
//}
