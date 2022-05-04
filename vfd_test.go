package m202md10b_test

import (
	"log"
	"time"

	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 3
)

func ExampleVFD() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	err = vfd.Brightness(M202MD10B.Brightness4)
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.CursorDisable()
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.Clear()
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.Print(" Futaba M202MD10B")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(WAIT_TIME)
	// Output:

}

func ExampleVFD_Print() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	err = vfd.Print("Hello World!\n コンニチワ")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(WAIT_TIME)
	// Output:
}

func ExampleVFD_Println() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	err = vfd.Println("Hello ")
	if err != nil {
		log.Fatal(err)
	}
	err = vfd.Print("  World!")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(WAIT_TIME)
	// Output:
}

func ExampleVFD_Clear() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	err = vfd.Clear()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func ExampleVFD_ClearAnimation() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	vfd.CursorEnable(M202MD10B.CursorTypeUnderline)
	vfd.CursorBlink(true)
	vfd.Animation = M202MD10B.AnimationEnable
	err = vfd.Print("  Welcome To THE\nC Y B E R S P A C E")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(WAIT_TIME)
	vfd.ClearAnimation()
	// Output:
}

func ExampleVFD_CGRAMFromStrings() {

	const (
		CGRAM_HEART byte = 0x88
		CGRAM_STAR  byte = 0x89
	)

	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

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

	vfd.Animation = M202MD10B.AnimationEnable

	err = vfd.Print("  ")
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.PutChar(CGRAM_HEART)
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.Print("マーク ダッテ\n    デルンダ ゾッ")
	if err != nil {
		log.Fatal(err)
	}

	err = vfd.PutChar(CGRAM_STAR)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(WAIT_TIME)
	vfd.ClearAnimation()

	// Output:
}
