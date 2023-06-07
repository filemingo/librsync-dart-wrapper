package main

import "C"
import librsyncdartwrapper "github.com/filemingo/librsync-dart-wrapper"

//export ComputeDelta
func ComputeDelta(signatureStr string, targetFilePath string, writeCallback func([]byte, int)) {
	librsyncdartwrapper.ComputeDelta(signatureStr, targetFilePath, writeCallback)
}

func main() {}
