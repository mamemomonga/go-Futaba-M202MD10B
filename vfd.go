// Copyright 2022 mamemomonga
// https://github.com/mamemomonga/go-Futaba-M202MD10B/
/*
Futaba M202MD10B用ライブラリ
*/
package m202md10b

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	bitarray "github.com/tunabay/go-bitarray"
	"go.bug.st/serial"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

// 明るさ
const (
	Brightness1 = 3 // 明るさ1(最も暗い)
	Brightness2 = 2 // 明るさ2
	Brightness3 = 1 // 明るさ3
	Brightness4 = 0 // 明るさ4(最も明るい)
)

// カーソルの種類
const (
	CursorTypeUnderline = iota // 下線
	CursorTypeTofu             // 豆腐
	CursorTypeXOR              // 反転
)

// エフェクトモード
const (
	AnimationDisable = iota // 文字表示効果無効
	AnimationEnable         // 文字表示効果有効
)

var (
	ErrBrightnessOutOfRange = errors.New("brightness out of range") // 明るさ設定が範囲外
	ErrCursorTypeNotDefined = errors.New("cursor type not defined") // カーソルの種類が範囲外
	rexLF                   = regexp.MustCompile("(?:\r\n)|\n")
)

// VFDライブラリ
type VFD struct {
	port               serial.Port   // シリアルポート
	Wait               time.Duration // 次の文字を表示するまでの待機時間
	Animation          int           // 文字アニメーション有効 (AnimationDisable:無効 /  AnimationEnable:有効)
	AnimationCharStart byte          // 文字表示効果開始ポイント
	AnimationWait      time.Duration // 文字表示効果ウェイト
	bufText            [40]byte      // 文字データ
	bufPos             int           // カーソル
}

// 初期化
func New() *VFD {
	t := new(VFD)
	t.Wait = time.Millisecond * 1
	t.Animation = AnimationDisable
	t.AnimationCharStart = 0x08
	t.AnimationWait = time.Millisecond * 10
	t.bufClear()
	return t
}

func (t *VFD) bufClear() {
	for i := 0; i < 40; i++ {
		t.bufText[i] = 0x20 // スペースで埋める
	}
	t.bufPos = 0
}

// ポートを開く
//   port: シリアルポート名
//   baud: ボーレート(9600, 4800, 2400, 1200)
func (t *VFD) Open(port string, baud int) (err error) {
	t.port, err = serial.Open(port, &serial.Mode{BaudRate: baud})
	if err != nil {
		return err
	}
	t.Reset()
	return nil
}

// ポートを閉じる
func (t *VFD) Close() {
	t.port.Close()
}

// 1バイト書き込み
func (t *VFD) WriteByte(c byte) error {
	if _, err := t.port.Write([]byte{c}); err != nil {
		return err
	}
	time.Sleep(t.Wait)
	return nil
}

// 1文字書き込み
func (t *VFD) PutChar(c byte) error {
	if err := t.WriteByte(c); err != nil {
		return err
	}

	if t.bufPos >= 40 {
		return nil // FIXME
	}

	t.bufText[t.bufPos] = c
	t.bufPos++

	return nil
}

// 文字の変換
func (t *VFD) convertText(str string) ([]byte, error) {
	var err error
	str = rexLF.ReplaceAllString(str, string([]byte{0x0A})) // CRLFはLFに統一
	str = width.Narrow.String(norm.NFD.String(str))         // 半角にする
	var ra []byte
	{
		iostr := strings.NewReader(str)
		ior := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
		ra, err = ioutil.ReadAll(ior)
		if err != nil {
			return nil, err
		}
	} // Shift_JISにする

	nb := make([]byte, len(ra))
	for i, c := range ra {
		switch {
		case (c >= 0x20) && (c <= 0x7E): // 英数字
			nb[i] = c
		case (c >= 0xA1) && (c <= 0xDF): // カタカナ
			nb[i] = c
		case c == 0x0A: // LF
			nb[i] = 0x0A
		default:
			nb[i] = 0x20
		}
	}
	return nb, nil
}

// 文字の表示
func (t *VFD) Print(str string) error {
	buf, err := t.convertText(str)
	for _, c := range buf {

		if t.bufPos > 40 {
			break
		}

		switch {
		case c == 0x0A: // LF
			t.CursorLine2()
			continue

		case c < 0x20:
			t.bufText[t.bufPos] = 0x20
			t.bufPos++
			continue
		}

		switch t.Animation {
		case AnimationEnable:
			err = t.textAnimation(c)
			continue
		case AnimationDisable:
			err = t.PutChar(c)
			continue
		}
	}
	return err
}

// 文字の表示(改行付き)
func (t *VFD) Println(str string) error {
	return t.Print(str + "\n")
}

// テキストエフェクト
func (t *VFD) textAnimation(c byte) error {
	if t.bufPos >= 40 {
		return nil // FIXME
	}

	dir := 0
	switch {
	case c >= 0xE0:
		dir = 0
	case c >= 0x80:
		dir = 1
	case c > 0x20:
		dir = -1
	default:
		dir = 0
	}
	if t.bufPos == 39 { // 最後の文字はアニメーションしない(カーソルが戻せない)
		dir = 0
	}

	switch {
	case dir == 1:
		start := c - t.AnimationCharStart
		if start < 0x20 {
			start = 0x20
		}
		for i := start; i <= c; i++ {
			if err := t.WriteByte(i); err != nil {
				return err
			}
			if c == i {
				break
			}
			time.Sleep(t.AnimationWait)
			t.CursorReverse()
		}
	case dir == -1:
		start := c + t.AnimationCharStart
		if start > 0xE0 {
			start = 0xE0
		}
		for i := start; i >= c; i-- {
			if err := t.WriteByte(i); err != nil {
				return err
			}
			if c == i {
				break
			}
			time.Sleep(t.AnimationWait)
			t.CursorReverse()
		}
	default:
		if err := t.WriteByte(c); err != nil {
			return err
		}
	}
	t.bufText[t.bufPos] = c
	t.bufPos++
	return nil
}

// アニメーション付き画面消去
func (t *VFD) ClearAnimation() error {
	t.CursorDisable()
	t.CursorHome()
	for i := byte(0); i < 8; i++ {
		for j := 0; j < 40; j++ {
			if t.bufText[j] != 0x20 {
				nc := t.bufText[j] + i
				switch {
				case nc >= 0xE0:
					nc = 0x20
				case (nc >= 0x80) && (nc < 0xA0):
					nc = 0x20
				case nc < 0x20:
					nc = 0x20
				}
				t.WriteByte(nc)
			} else {
				t.WriteByte(0x20)
			}
		}
		time.Sleep(time.Millisecond * 5)
	}
	t.clearDisplay()
	for i := 0; i < 40; i++ {
		t.bufText[i] = 0x20
	}
	t.bufPos = 0
	return nil
}

// 画面クリア
func (t *VFD) Clear() error {
	t.bufClear()
	t.clearDisplay()
	return nil
}

// 画面を削除して左上に移動
func (t *VFD) clearDisplay() error {
	return t.WriteByte(byte(0x0c))
}

// 画面を削除するが左上に戻らない
func (t *VFD) clearDisplayNoReturnToHome() error {
	return t.WriteByte(byte(0x0a))
}

// カーソルを右に移動
func (t *VFD) CursorForward() error {
	return t.WriteByte(byte(0x09))
}

// カーソルを左に移動
func (t *VFD) CursorReverse() error {
	return t.WriteByte(byte(0x08))
}

// カーソルを右上に
func (t *VFD) CursorHome() error {
	t.bufPos = 0
	return t.WriteByte(byte(0x0d))
}

// カーソルを1行目に(CursorHomeと同じ)
func (t *VFD) CursorLine1() error {
	return t.CursorHome()
}

// カーソルを2行目に
func (t *VFD) CursorLine2() error {

	log.Println("CursorLine2")
	if err := t.CursorHome(); err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		if err := t.CursorForward(); err != nil {
			return err
		}
	}
	t.bufPos = 20
	return nil
}

// カーソルの無効化
func (t *VFD) CursorDisable() error {
	return t.WriteByte(byte(0x14))
}

// カーソルの点滅
func (t *VFD) CursorBlink(blink bool) error {
	switch blink {
	case true:
		return t.WriteByte(byte(0x15)) // 点滅オン
	case false:
		return t.WriteByte(byte(0x13)) // 点滅オフ
	}
	return nil
}

// カーソルの表示( CursorTypeUnderline, CursorTypeTofu, CursorTypeXOR )
func (t *VFD) CursorEnable(cursorType int) error {
	switch cursorType {
	case CursorTypeUnderline:
		if err := t.WriteByte(byte(0x16)); err != nil {
			return err
		}
	case CursorTypeTofu:
		if err := t.WriteByte(byte(0x17)); err != nil {
			return err
		}
	case CursorTypeXOR:
		if err := t.WriteByte(byte(0x18)); err != nil {
			return err
		}
	}
	return ErrCursorTypeNotDefined
}

// 明るさの設定
// Brightness1, Brightness2, Brightness3, Brightness4 のいずれかを設定
func (t *VFD) Brightness(value int) error {
	if value < 0 || value > 3 {
		return ErrBrightnessOutOfRange
	}
	if err := t.WriteByte(byte(0x01 + value)); err != nil {
		return err
	}
	return nil
}

// 外字の設定 (既存の文字コードに上書き、最大8個まで)
func (t *VFD) CGRAM(code byte, data [7]byte) error {
	if err := t.WriteByte(byte(0x1a)); err != nil {
		return err
	}
	if err := t.WriteByte(code); err != nil {
		return err
	}

	for i := 0; i < 7; i++ {
		if err := t.WriteByte(data[i]); err != nil {
			return err
		}
	}
	return nil
}

// Stringsの配列から外字の登録
func (t *VFD) CGRAMFromStrings(code byte, graphic [7]string) error {
	var data [7]byte
	for i := 0; i < 7; i++ {
		ba, err := bitarray.Parse(graphic[i])
		if err != nil {
			return err
		}
		ba = ba.Reverse()
		b, _ := ba.Bytes()
		data[i] = b[0] >> 3
	}
	return t.CGRAM(code, data)
}

// ハードウェアリセット(/DTRピンの接続が必要)
func (t *VFD) Reset() error {
	if err := t.port.SetDTR(true); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 10)
	if err := t.port.SetDTR(false); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 10)
	return nil
}
