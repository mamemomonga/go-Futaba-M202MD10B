package main

import (
	"flag"
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

	vfd.Animation = M202MD10B.AnimationEnable
	vfd.Brightness(M202MD10B.Brightness3)
	for {
		example1()
		example2()
	}
}

func waitAndClear() {
	time.Sleep(time.Second * 3)
	err := vfd.ClearAnimation()
	if err != nil {
		log.Fatal(err)
	}
}

func example1() {
	vfd.CursorEnable(M202MD10B.CursorTypeUnderline)
	vfd.CursorBlink(true)
	err := vfd.Print("  Welcome To THE\nC Y B E R S P A C E")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}

func example2() {
	vfd.CursorEnable(M202MD10B.CursorTypeUnderline)
	vfd.CursorBlink(true)
	err := vfd.Print(" M202MD10B デス\n  ドウゾヨロシク")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}
