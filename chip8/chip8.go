package chip8

import (
	"chip8-emulator/graphics"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

const (
	memorySize        = 4096
	variableRegisters = 0xF
	programStart      = 0x200
	fontStart         = 0x50
	stackStart        = 0x100
	stackEnd          = stackStart + (16 * 2)
	totalKeys         = 0xF
	cpuSpeed          = 700 // Mhz
	timerSpeed        = 60  // Mhz
)

type Chip8 struct {
	/* Peripherals */
	memory   []byte
	screen   *graphics.Graphics
	Keyboard keyboard

	/* Registers */
	pc       uint16
	index    uint16
	stack    uint16
	variable []byte

	/* Timers */
	delay uint8
	sound uint8

	/* Clocks */
	cpuClock   *time.Ticker
	timerClock *time.Ticker
}

func New(rom []byte, screen *graphics.Graphics) *Chip8 {
	chip := &Chip8{
		memory: make([]byte, memorySize),
		screen: screen,
		Keyboard: keyboard{
			Input:   make(chan sdl.Scancode, totalKeys*2),
			state:   make([]bool, totalKeys+1),
			mapping: getKeyboardMapping(),
		},

		variable: make([]byte, variableRegisters),

		cpuClock:   time.NewTicker(time.Second / cpuSpeed),
		timerClock: time.NewTicker(time.Second / timerSpeed),
	}

	chip.injectIntoMemory(programStart, rom)
	chip.injectIntoMemory(fontStart, getFontData())

	return chip
}

func (chip *Chip8) Run() {
	for {
		select {
		case <-chip.cpuClock.C:
			chip.step()
		case <-chip.timerClock.C:
			chip.updateTimers()
		}
	}
}

func (chip *Chip8) Close() {
	chip.cpuClock.Stop()
	chip.timerClock.Stop()
}

func (chip *Chip8) PrintMemory(from int, to int, padding int) {
	start := from - padding
	end := from + to + padding

	for position, value := range chip.memory[start:end] {
		fmt.Printf("%#x = %#x\n", start+position, value)
	}
}

func (chip *Chip8) injectIntoMemory(offset int, data []byte) {
	for position, value := range data {
		chip.memory[offset+position] = value
	}
}

func (chip *Chip8) step() {
	chip.Keyboard.update()

}

func (chip *Chip8) updateTimers() {
	if chip.delay > 0 {
		chip.delay--
	}

	if chip.sound > 0 {
		chip.sound--
	}
}

func (chip *Chip8) push(value uint16) {
	if chip.stack+2 == stackEnd {
		log.Fatal("Stack overflow")
		return
	}

	chip.memory[stackStart+chip.stack] = byte(value >> 8)
	chip.memory[stackStart+chip.stack+1] = byte(value & 0xFF)
	chip.stack += 2
}

func (chip *Chip8) pop() uint16 {
	if chip.stack == stackStart {
		log.Fatal("Stack underflow")
		return 0
	}

	chip.stack -= 2
	return uint16(chip.memory[stackStart+chip.stack])<<8 | uint16(chip.memory[stackStart+chip.stack+1])
}
