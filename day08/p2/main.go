package main

import (
	"aoc2019/lib/util"
	"fmt"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := util.Input()[0]

	const width = 25
	const height = 6
	const pixelsPerLayer = width * height

	var layers [][]string
	for layer := range slices.Chunk(strings.Split(input, ""), pixelsPerLayer) {
		layers = append(layers, layer)
	}

	img := make([]string, pixelsPerLayer)

	for i := 0; i < pixelsPerLayer; i++ {
		for j := 0; j < len(layers); j++ {
			if layers[j][i] == "2" {
				continue
			}
			if layers[j][i] == "0" {
				img[i] = " "
			} else {
				img[i] = "#"
			}
			break
		}
	}

	for row := range slices.Chunk(img, width) {
		fmt.Println(row)
	}

	fmt.Println(time.Since(start))
}
