package assert

import (
	"log"
)

func NoErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func True(v bool, msg ...any) {
	if !v {
		log.Fatal(msg...)
	}
}
