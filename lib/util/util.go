package util

import (
	"aoc2019/lib/assert"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Input returns the data of the file "input.txt".
func Input() []string {
	_, callerfile, _, ok := runtime.Caller(1)
	assert.True(ok)
	data, err := os.ReadFile(filepath.Dir(callerfile) + "/input.txt")
	assert.NoErr(err)
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
