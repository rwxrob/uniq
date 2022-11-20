package uniq_test

import (
	"fmt"
	"log"

	"github.com/rwxrob/uniq"
)

func ExampleBytes() {

	bytes := fmt.Sprintf("%x", uniq.Bytes(4))
	fmt.Println(len(bytes))

	// Output:
	// 8
}

func ExampleUUID() {

	uuid := uniq.UUID()
	fmt.Println(len(uuid))

	// Output:
	// 36
}

func ExampleBase32() {

	uid32 := uniq.Base32()
	fmt.Println(len(uid32))

	// Output:
	// 32
}

func ExampleHex() {

	hex1 := uniq.Hex(1)
	hex3 := uniq.Hex(3)
	hex18 := uniq.Hex(18)
	fmt.Println(len(hex1))
	fmt.Println(len(hex3))
	fmt.Println(len(hex18))

	// Output:
	// 2
	// 6
	// 36
}

func ExampleSecond() {

	sec := uniq.Second()
	fmt.Println(len(sec))

	// Output:
	// 10
}

func ExampleIsosec() {

	sec := uniq.Isosec()
	log.Print(sec)
	fmt.Println(len(sec))

	// Output:
	// 14
}

func ExampleIsonan() {

	sec := uniq.Isonan()
	log.Print(sec)
	fmt.Println(len(sec))

	// Output:
	// 23
}

func ExampleIsodate() {

	sec := uniq.Isodate()
	log.Print(sec)
	fmt.Println(len(sec))

	// Output:
	// 20
}
