package chip8

import "github.com/veandco/go-sdl2/sdl"

type keyboard struct {
	Input   chan sdl.Scancode
	state   []bool
	mapping map[sdl.Scancode]uint32
}

func (keyboard *keyboard) update() {
	if len(keyboard.Input) == 0 {
		return
	}

	key := <-keyboard.Input
	if offset, found := keyboard.mapping[key]; found {
		keyboard.state[offset] = !keyboard.state[offset]
	}
}

func getKeyboardMapping() map[sdl.Scancode]uint32 {
	mapping := make(map[sdl.Scancode]uint32)
	mapping[sdl.SCANCODE_1] = 0x0
	mapping[sdl.SCANCODE_2] = 0x1
	mapping[sdl.SCANCODE_3] = 0x2
	mapping[sdl.SCANCODE_4] = 0x3
	mapping[sdl.SCANCODE_Q] = 0x4
	mapping[sdl.SCANCODE_W] = 0x5
	mapping[sdl.SCANCODE_E] = 0x6
	mapping[sdl.SCANCODE_R] = 0x7
	mapping[sdl.SCANCODE_A] = 0x8
	mapping[sdl.SCANCODE_S] = 0x9
	mapping[sdl.SCANCODE_D] = 0xA
	mapping[sdl.SCANCODE_F] = 0xB
	mapping[sdl.SCANCODE_Z] = 0xC
	mapping[sdl.SCANCODE_X] = 0xD
	mapping[sdl.SCANCODE_C] = 0xE
	mapping[sdl.SCANCODE_V] = 0xF

	return mapping
}
