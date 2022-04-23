package main

import (
	"chip8-emulator/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	/// ---------
	/// Parse commandline arguments
	/// ---------

	//parser := argparse.NewParser("chip8-emulator", "A Chip8 Emulator written in Go")
	//romFile := parser.String("r", "rom", &argparse.Options{Help: "Path to the ROM file", Default: "./roms/PONG"})
	//
	//err := parser.Parse(os.Args)
	//if err != nil {
	//	fmt.Print(parser.Usage(err))
	//	os.Exit(0)
	//}

	/// ---------
	/// Read ROM
	/// ---------

	//rom, err := ioutil.ReadFile(*romFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	/// ---------
	/// Start the emulator
	/// ---------

	screen := graphics.New()
	defer screen.Close()

	for i := 0; i <= 800*300; i++ {
		screen.Buffer[i] = 0xFF000000
	}

	screen.Update()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
