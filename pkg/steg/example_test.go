package steg_test

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/steg"
)

func ExampleEncode() {
	plain := []byte("hello world")
	encode, err := steg.Encode(plain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Raw encode: %q", encode)
	// Output: Raw encode: "\u200c\u200d\u2064\u2060\u200c\u2060\u200c\u200d\u2062\u200c\u200d\u200c\u2060\u200d\u2064\u200d\u2060\u200d\u2064\u2062\u2063\u2060\u2062\u2060\u2062\u2060\u2064\u2064\u2064\u2064\u2064\u2064\u200c\u2060\u200c\u2060\u2060\u2063\u200d\u2060\u200d\u200c\u2062\u2060\u200d\u200c\u2062\u2060\u200d\u2062\u2062\u2060\u200d\u2064\u2060\u200c\u2062\u200d\u2062\u200d\u2062\u2062\u2060\u200d\u2060\u200c\u2062\u200d\u200c\u2062\u2060\u200d\u200c\u200d\u2060\u200d\u2064\u2064\u2064\u2064\u2064\u2064\u2064\u2064\u2060\u200c\u2060\u200c\u2060\u200d\u2060\u200d\u2062\u2060\u2062\u2060\u2060\u2062\u200c\u2062"
}

func ExampleDecode() {
	plain := []byte("hello world")
	encode, err := steg.Encode(plain)
	if err != nil {
		panic(err)
	}
	decode, err := steg.Decode(encode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode: %q", decode)
	// Output: Decode: "hello world"
}
