# go-Futaba-M202MD10B

Futaba M202MD10B Go ライブラリ

# リファレンス

[![GoDoc](https://godoc.org/github.com/mamemomonga/go-Futaba-M202MD10B?status.svg)](https://godoc.org/github.com/mamemomonga/go-Futaba-M202MD10B)

# Futaba M202MD10B

* 20x2, UART入力のドットマトリクスVFPモジュールです
* 電源およびIOは 5V です

## コネクタ

**PH 7ピン**

Pin | Name | Desc.
---|---|---
1 | VCC | 電源(5V 約650mA)
2 | VCC | 電源(5V 約650mA)
3 | #RST | Lでリセット
4 | DATA | UART入力
5 | BUSY | Hのときビジー
6 | GND | グランド
7 | GND | グランド

## DIPスイッチ

SW | NAME | Desc.
---|---|---
1 | CODE | ON: カタカナ / OFF: キリル文字
2 | BAUD0 | ボーレート設定0
3 | BAUD1 | ボーレート設定1
4 | TEST | ON: デモモード

BAUD0,1: ボーレート設定

Baud | BAUD0 | BAUD1
---|---|---
9600 | OFF | OFF
4800 | OFF | ON
2400 | ON | OFF
1200 | ON | ON

## キャラクター一覧

[![aキャラクター一覧](http://img.youtube.com/vi/s-9mbCNlsLk/0.jpg)](https://www.youtube.com/watch?v=s-9mbCNlsLk)

# 配線

* USB-UARTモジュールを使用して、macOSにて製作しました。
* このライブラリでは、TxをDATA, #RSTをDTRに接続します。
* ハードウェアリセットが不要な場合は#RSTを10kΩ抵抗を介して5Vに接続します。
* 起動直後はハードウェアリセットを行わないと正しく動作しない場合があります。
* IOは5Vです。必要に応じてレベル変換を行って下さい。
* USB-UARTモジュールが5Vトレラント対応や、オープンドレインの場合はレベル変換が不要な場合があります。

# サンプルコードの実行

	$ go mod vendor
	$ go run ./examples/helloWorld
	$ go run ./examples/animation
	$ go run ./examples/cgram

# 参考情報

* [5ch](https://rio2016.5ch.net/test/read.cgi/denki/1640165380/107)

# Lisence

MIT
