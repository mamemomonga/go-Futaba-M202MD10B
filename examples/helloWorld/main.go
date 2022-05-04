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

	for {
		example1()
		example2()
		example3()
		example4()
	}
}

func waitAndClear() {
	time.Sleep(time.Second * 3)
	err := vfd.Clear()
	if err != nil {
		log.Fatal(err)
	}
}

func example1() {
	err := vfd.Print("[1] Hello World!")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}

func example2() {
	err := vfd.Print("[2] Hello\n World!")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}

func example3() {
	err := vfd.Println("[3] Hello")
	if err != nil {
		log.Fatal(err)
	}
	err = vfd.Println("   World!")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}

func example4() {
	err := vfd.Println("[4] コンニチワ")
	if err != nil {
		log.Fatal(err)
	}
	err = vfd.Println("   セカイ")
	if err != nil {
		log.Fatal(err)
	}
	waitAndClear()
}
