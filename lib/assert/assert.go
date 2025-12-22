package assert

import (
	"fmt"
)

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func True(v bool, msg ...any) {
	if !v {
		panic(fmt.Sprint(msg...))
	}
}
