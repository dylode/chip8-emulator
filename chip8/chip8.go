package chip8

import (
	"chip8-emulator/graphics"
	"fmt"
)

const (
	memorySize        = 4096
	variableRegisters = 0xF
	programStart      = 0x200
	fontStart         = 0x50
)

type Chip8 struct {
	/* Peripherals */
	memory []byte
	screen *graphics.Graphics

	/* Registers */
	pc       uint16
	index    uint16
	stack    uint16
	variable []byte

	/* Timers */
	delay uint8
	sound uint8
}

func New(rom []byte, screen *graphics.Graphics) *Chip8 {
	chip := &Chip8{
		memory:   make([]byte, memorySize),
		screen:   screen,
		variable: make([]byte, variableRegisters),
	}

	chip.injectIntoMemory(programStart, rom)
	chip.injectIntoMemory(fontStart, getFontData())

	return chip
}

func (chip *Chip8) injectIntoMemory(offset int, data []byte) {
	for position, value := range data {
		chip.memory[offset+position] = value
	}
}

func (chip *Chip8) PrintMemory(from int, to int, padding int) {
	start := from - padding
	end := from + to + padding

	for position, value := range chip.memory[start:end] {
		fmt.Printf("%#x = %#x\n", start+position, value)
	}
}
