package main

import (
	"flag"
	"log"

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
	err = vfd.Clear()
	if err != nil {
		log.Fatal(err)
	}
}
