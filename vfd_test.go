package m202md10b_test

import (
	"fmt"
	"log"

	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
)

func Example() {
	v := M202MD10B.New()
	var err error

	err = v.Open("/dev/cu.usbserial-1101", 9600)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()

	err = v.Brightness(M202MD10B.Brightness4)
	if err != nil {
		log.Fatal(err)
	}

	err = v.CursorDisable()
	if err != nil {
		log.Fatal(err)
	}

	err = v.Clear()
	if err != nil {
		log.Fatal(err)
	}

	err = v.Print("Hello World!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
	// Output: OK

}
