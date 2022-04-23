package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Graphics struct {
	Buffer   []uint32
	window   *sdl.Window
	rendered *sdl.Renderer
	texture  *sdl.Texture
}

func New() *Graphics {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}

	window, err := sdl.CreateWindow("Chip-8 Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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

	return &Graphics{
		Buffer:   make([]uint32, screenWidth*screenHeight),
		window:   window,
		rendered: renderer,
		texture:  texture,
	}
}

func (screen *Graphics) Update() {
	err := screen.texture.UpdateRGBA(nil, screen.Buffer, screenWidth)
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
