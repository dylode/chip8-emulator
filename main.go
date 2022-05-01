package main

import (
	"chip8-emulator/chip8"
	"chip8-emulator/graphics"
	"fmt"
	"github.com/akamensky/argparse"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	//emulator.PrintMemory(0x50, 5*16, 3)

	cpuClock := time.NewTicker(time.Second / 700)
	timerClock := time.NewTicker(time.Second / 60)

	for {
		select {
		case <-cpuClock.C:
			emulator.Step()
		case <-timerClock.C:
			emulator.UpdateTimers()
		}
	}
	//for i := 0; i <= 800*300; i++ {
	//	screen.Buffer[i] = 0xFF000000
	//}
	//
	//screen.Update()
	//
	//running := true
	//for running {
	//	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	//		switch event.(type) {
	//		case *sdl.QuitEvent:
	//			println("Quit")
	//			running = false
	//			break
	//		}
	//	}
	//}
}
