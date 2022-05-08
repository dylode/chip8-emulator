package main

import (
	"chip8-emulator/chip8"
	"chip8-emulator/graphics"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	/// ---------
	/// Parse commandline arguments
	/// ---------

	parser := argparse.NewParser("chip8-emulator", "A Chip8 Emulator written in Go")
	romFile := parser.String("r", "rom", &argparse.Options{Help: "Path to the ROM file", Default: "./roms/PONG"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	/// ---------
	/// Read ROM
	/// ---------

	rom, err := ioutil.ReadFile(*romFile)
	if err != nil {
		log.Fatal(err)
	}

	/// ---------
	/// Start the emulator
	/// ---------

	screen := graphics.New()
	defer screen.Close()

	emulator := chip8.New(rom, screen)
	defer emulator.Close()
	//emulator.PrintMemory(0x50, 5*16, 3)

	go emulator.Run()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				if e.Repeat == 0 {
					emulator.Keyboard.Input <- e.Keysym.Scancode
				}
			}
		}
	}
}
