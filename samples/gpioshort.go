// +build ignore

package main

import (
	"flag"

	"github.com/marian-craciunescu/embd"

	_ "github.com/marian-craciunescu/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(10, embd.Out)
	embd.DigitalWrite(10, embd.High)
}
