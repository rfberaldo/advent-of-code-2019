package assert

import (
	"fmt"
)

func NoErr(err error) {
	if err != nil {
		panic(fmt.Sprint("assert:", err))
	}
}

func True(v bool, msg ...any) {
	if !v {
		msg = append([]any{"assert: "}, msg...)
		panic(fmt.Sprint(msg))
	}
}
