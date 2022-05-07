<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# m202md10b

```go
import "github.com/mamemomonga/go-Futaba-M202MD10B"
```

Copyright 2022 mamemomonga https://github.com/mamemomonga/go-Futaba-M202MD10B/

Futaba M202MD10B用ライブラリ

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [type VFD](<#type-vfd>)
  - [func New() *VFD](<#func-new>)
  - [func (t *VFD) Brightness(value int) error](<#func-vfd-brightness>)
  - [func (t *VFD) CGRAM(code byte, data [7]byte) error](<#func-vfd-cgram>)
  - [func (t *VFD) CGRAMFromStrings(code byte, graphic [7]string) error](<#func-vfd-cgramfromstrings>)
  - [func (t *VFD) Clear() error](<#func-vfd-clear>)
  - [func (t *VFD) ClearAnimation() error](<#func-vfd-clearanimation>)
  - [func (t *VFD) ClearCursorNRTH() error](<#func-vfd-clearcursornrth>)
  - [func (t *VFD) Close()](<#func-vfd-close>)
  - [func (t *VFD) CursorBlink(blink bool) error](<#func-vfd-cursorblink>)
  - [func (t *VFD) CursorDisable() error](<#func-vfd-cursordisable>)
  - [func (t *VFD) CursorEnable(cursorType int) error](<#func-vfd-cursorenable>)
  - [func (t *VFD) CursorForward() error](<#func-vfd-cursorforward>)
  - [func (t *VFD) CursorHome() error](<#func-vfd-cursorhome>)
  - [func (t *VFD) CursorLine1() error](<#func-vfd-cursorline1>)
  - [func (t *VFD) CursorLine2() error](<#func-vfd-cursorline2>)
  - [func (t *VFD) CursorReverse() error](<#func-vfd-cursorreverse>)
  - [func (t *VFD) Open(port string, baud int) (err error)](<#func-vfd-open>)
  - [func (t *VFD) Print(str string) error](<#func-vfd-print>)
  - [func (t *VFD) Println(str string) error](<#func-vfd-println>)
  - [func (t *VFD) PutChar(c byte) error](<#func-vfd-putchar>)
  - [func (t *VFD) Reset() error](<#func-vfd-reset>)
  - [func (t *VFD) WriteByte(c byte) error](<#func-vfd-writebyte>)


## Constants

明るさ

```go
const (
    Brightness1 = 3 // 明るさ1(最も暗い)
    Brightness2 = 2 // 明るさ2
    Brightness3 = 1 // 明るさ3
    Brightness4 = 0 // 明るさ4(最も明るい)
)
```

カーソルの種類

```go
const (
    CursorTypeUnderline = iota // 下線
    CursorTypeTofu             // 豆腐
    CursorTypeXOR              // 反転
)
```

エフェクトモード

```go
const (
    AnimationDisable = iota // 文字表示効果無効
    AnimationEnable         // 文字表示効果有効
)
```

## Variables

```go
var (
    ErrBrightnessOutOfRange = errors.New("brightness out of range") // 明るさ設定が範囲外
    ErrCursorTypeNotDefined = errors.New("cursor type not defined") // カーソルの種類が範囲外
    ErrBufferOverflow       = errors.New("buffer overflow")         // バッファがいっぱい

)
```

## type VFD

VFDライブラリ

```go
type VFD struct {
    Wait               time.Duration // 次の文字を表示するまでの待機時間
    Animation          int           // 文字アニメーション有効 (AnimationDisable:無効 /  AnimationEnable:有効)
    AnimationCharStart byte          // 文字表示効果開始ポイント
    AnimationWait      time.Duration // 文字表示効果ウェイト
    // contains filtered or unexported fields
}
```

<details><summary>Example</summary>
<p>

```go
package main

import (
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {
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
}
```

#### Output

```

```

</p>
</details>

### func New

```go
func New() *VFD
```

初期化

### func \(\*VFD\) Brightness

```go
func (t *VFD) Brightness(value int) error
```

明るさの設定 Brightness1\, Brightness2\, Brightness3\, Brightness4 のいずれかを設定

### func \(\*VFD\) CGRAM

```go
func (t *VFD) CGRAM(code byte, data [7]byte) error
```

外字の設定 \(既存の文字コードに上書き、最大8個まで\)

### func \(\*VFD\) CGRAMFromStrings

```go
func (t *VFD) CGRAMFromStrings(code byte, graphic [7]string) error
```

Stringsの配列から外字の登録

<details><summary>Example</summary>
<p>

```go
package main

import (
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {

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

}
```

#### Output

```

```

</p>
</details>

### func \(\*VFD\) Clear

```go
func (t *VFD) Clear() error
```

画面クリア

<details><summary>Example</summary>
<p>

```go
package main

import (
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {
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
}
```

#### Output

```

```

</p>
</details>

### func \(\*VFD\) ClearAnimation

```go
func (t *VFD) ClearAnimation() error
```

アニメーション付き画面消去

<details><summary>Example</summary>
<p>

```go
package main

import (
	"errors"
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {
	vfd := M202MD10B.New()
	var err error

	err = vfd.Open(SERIAL_PORT, 9600) // SERIAL_PORT: e.g. /dev/ttyUSB0, /dev/cu.usbserial
	if err != nil {
		log.Fatal(err)
	}
	defer vfd.Close()

	errChk := func(err error) {
		if err != nil {
			if errors.Is(err, M202MD10B.ErrBufferOverflow) {
				log.Printf("warn: %v", err)
			} else {
				log.Fatalf("alert: %v", err)
			}
		}
	}

	vfd.CursorEnable(M202MD10B.CursorTypeUnderline)
	vfd.CursorBlink(true)
	vfd.Animation = M202MD10B.AnimationEnable

	err = vfd.Print("  Welcome To THE\nC Y B E R S P A C E")
	errChk(err)

	time.Sleep(WAIT_TIME)

	err = vfd.ClearAnimation()
	errChk(err)
}
```

#### Output

```

```

</p>
</details>

### func \(\*VFD\) ClearCursorNRTH

```go
func (t *VFD) ClearCursorNRTH() error
```

画面クリア\(カーソルを移動しない\)

### func \(\*VFD\) Close

```go
func (t *VFD) Close()
```

ポートを閉じる

### func \(\*VFD\) CursorBlink

```go
func (t *VFD) CursorBlink(blink bool) error
```

カーソルの点滅

### func \(\*VFD\) CursorDisable

```go
func (t *VFD) CursorDisable() error
```

カーソルの無効化

### func \(\*VFD\) CursorEnable

```go
func (t *VFD) CursorEnable(cursorType int) error
```

カーソルの表示\( CursorTypeUnderline\, CursorTypeTofu\, CursorTypeXOR \)

### func \(\*VFD\) CursorForward

```go
func (t *VFD) CursorForward() error
```

カーソルを右に移動

### func \(\*VFD\) CursorHome

```go
func (t *VFD) CursorHome() error
```

カーソルを右上に

### func \(\*VFD\) CursorLine1

```go
func (t *VFD) CursorLine1() error
```

カーソルを1行目に\(CursorHomeと同じ\)

### func \(\*VFD\) CursorLine2

```go
func (t *VFD) CursorLine2() error
```

カーソルを2行目に

### func \(\*VFD\) CursorReverse

```go
func (t *VFD) CursorReverse() error
```

カーソルを左に移動

### func \(\*VFD\) Open

```go
func (t *VFD) Open(port string, baud int) (err error)
```

ポートを開く port: シリアルポート名 baud: ボーレート\(9600\, 4800\, 2400\, 1200\)

### func \(\*VFD\) Print

```go
func (t *VFD) Print(str string) error
```

文字の表示

<details><summary>Example</summary>
<p>

```go
package main

import (
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {
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
}
```

#### Output

```

```

</p>
</details>

### func \(\*VFD\) Println

```go
func (t *VFD) Println(str string) error
```

文字の表示\(改行付き\)

<details><summary>Example</summary>
<p>

```go
package main

import (
	M202MD10B "github.com/mamemomonga/go-Futaba-M202MD10B"
	"log"
	"time"
)

const (
	SERIAL_PORT = "/dev/cu.usbserial-1101"
	WAIT_TIME   = time.Second * 1
)

func main() {
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
}
```

#### Output

```

```

</p>
</details>

### func \(\*VFD\) PutChar

```go
func (t *VFD) PutChar(c byte) error
```

1文字書き込み

### func \(\*VFD\) Reset

```go
func (t *VFD) Reset() error
```

ハードウェアリセット\(/DTRピンの接続が必要\)

### func \(\*VFD\) WriteByte

```go
func (t *VFD) WriteByte(c byte) error
```

1バイト書き込み



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)