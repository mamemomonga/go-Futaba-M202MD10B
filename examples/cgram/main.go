package main

import (
	"flag"
	"log"

	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
)

const (
	CGRAM_HEART byte = 0x88
	CGRAM_STAR  byte = 0x89
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
	loadCGRAM()
	example1()
}

func example1() {
	err := vfd.Print("  ")
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.WriteByte(CGRAM_HEART)
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.Print("マーク ダッテ\n    デルンダ ゾッ")
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.WriteByte(CGRAM_STAR)
	if err != nil {
		log.Fatal(err)
	}
}

func loadCGRAM() {
	{
		var g [7]string
		g[0] = "00000"
		g[1] = "11011"
		g[2] = "11111"
		g[3] = "11111"
		g[4] = "11111"
		g[5] = "01110"
		g[6] = "01100"
		err := vfd.CGRAMFromStrings(CGRAM_HEART, g)
		if err != nil {
			log.Fatal(err)
		}

	}
	{
		var g [7]string
		g[0] = "00100"
		g[1] = "00100"
		g[2] = "11111"
		g[3] = "01110"
		g[4] = "01110"
		g[5] = "01010"
		g[6] = "10001"
		err := vfd.CGRAMFromStrings(CGRAM_STAR, g)
		if err != nil {
			log.Fatal(err)
		}
	}
}
