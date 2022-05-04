package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
)

var vfd *M202MD10B.VFD

func main() {
	flg_p := flag.String("p", "/dev/cu.usbserial-1101", "serial port")
	flg_b := flag.Int("b", 9600, "baud rate")
	flag.Parse()

	var err error
	vfd = M202MD10B.New()
	err = vfd.Open(*flg_p, *flg_b)
	if err != nil {
		log.Fatal(err)
	}

	for {
		charactors()
	}
}

func charactors() {
	pos := byte(0x20)
	for {
		for i := 0; i < 4; i++ {
			vfd.Print(fmt.Sprintf("  0x%X [", pos))
			vfd.PutChar(pos)
			vfd.Print("]")
			if pos == 0xFF {
				vfd.CursorHome()
				return
			}
			pos++
		}
		time.Sleep(time.Millisecond * 500)
		vfd.CursorHome()
	}
}
