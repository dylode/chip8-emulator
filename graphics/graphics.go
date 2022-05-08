package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	screenWidth   = 64
	screenHeight  = 32
	scalingFactor = 10
	pixelOff      = 0x00_00_00_FF
	pixelOn       = 0xFF_FF_FF_FF
)

type Graphics struct {
	buffer   []uint32
	window   *sdl.Window
	rendered *sdl.Renderer
	texture  *sdl.Texture
}

func New() *Graphics {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}

	window, err := sdl.CreateWindow("Chip-8 Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth*scalingFactor, screenHeight*scalingFactor, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s\n", err)
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, screenWidth, screenHeight)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
	}

	graphics := &Graphics{
		buffer:   getInitialScreenBuffer(),
		window:   window,
		rendered: renderer,
		texture:  texture,
	}

	graphics.Update()
	return graphics
}

func (screen *Graphics) Update() {
	err := screen.texture.UpdateRGBA(nil, screen.buffer, screenWidth)
	if err != nil {
		log.Fatal(err)
	}

	err = screen.rendered.Clear()
	if err != nil {
		log.Fatal(err)
	}

	err = screen.rendered.Copy(screen.texture, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	screen.rendered.Present()
}

func (screen *Graphics) GetPixel(offset uint32) byte {
	if screen.buffer[offset] == pixelOff {
		return 0
	}

	return 1
}

func (screen *Graphics) SetPixel(offset uint32, value byte) {
	if value == 0 {
		screen.buffer[offset] = pixelOff
	} else {
		screen.buffer[offset] = pixelOn
	}
}

func (screen *Graphics) Close() {
	sdl.Quit()
	err := screen.window.Destroy()
	if err != nil {
		log.Println(err)
	}

	err = screen.rendered.Destroy()
	if err != nil {
		log.Println(err)
	}

	err = screen.texture.Destroy()
	if err != nil {
		log.Println(err)
	}
}

func getInitialScreenBuffer() []uint32 {
	buffer := make([]uint32, screenWidth*screenHeight)

	for i := 0; i < screenWidth*screenHeight; i++ {
		buffer[i] = pixelOff
	}

	return buffer
}
