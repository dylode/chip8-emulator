package main

import (
	"fmt"
	"github.com/akamensky/argparse"
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

	fmt.Println(rom)
}
