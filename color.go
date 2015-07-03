package main

import (
	"log"
	"regexp"

	"github.com/fatih/color"
)

var (
	numbers = color.New(color.FgRed).SprintFunc()
	strings = color.New(color.FgBlue).SprintFunc()
	bools   = color.New(color.FgMagenta).SprintFunc()
	arrays  = color.New(color.FgGreen).SprintFunc()
	hashes  = color.New(color.FgCyan).SprintFunc()
	nulls   = color.New(color.FgYellow).SprintFunc()
)

func colorize(blob []byte) []byte {
	blob = replaceArrays(blob) //must be first because colors contain [,]
	blob = replaceStrings(blob)
	blob = replaceNumbers(blob)
	blob = replaceBooleans(blob)
	blob = replaceHashes(blob)
	return replaceNulls(blob)
}

func replaceStrings(blob []byte) []byte {
	re, err := regexp.Compile(`"(?U)(.*)"`)
	if err != nil {
		log.Fatal(err)
	}
	color := strings(`"$1"`)
	return re.ReplaceAll(blob, []byte(color))
}

func replaceNumbers(blob []byte) []byte {
	re, err := regexp.Compile("(-?\\d+\\.?\\d*)(,?\\s*\n)")
	if err != nil {
		log.Fatal(err)
	}
	color := numbers("$1")
	return re.ReplaceAll(blob, []byte(color+"$2"))
}

func replaceBooleans(blob []byte) []byte {
	re, err := regexp.Compile("(true|false)(,?\\s*\n)")
	if err != nil {
		log.Fatal(err)
	}
	color := bools("$1")
	return re.ReplaceAll(blob, []byte(color+"$2"))
}

func replaceArrays(blob []byte) []byte {
	re, err := regexp.Compile("(\\s*)(\\[|\\])(,?\\s*(:?\n|$))")
	if err != nil {
		log.Fatal(err)
	}
	color := arrays("$2")
	return re.ReplaceAll(blob, []byte("$1"+color+"$3"))
}

func replaceHashes(blob []byte) []byte {
	re, err := regexp.Compile("(\\s*)({|})(,?\\s*(?:\n|$))")
	if err != nil {
		log.Fatal(err)
	}
	color := hashes("$2")
	return re.ReplaceAll(blob, []byte("$1"+color+"$3"))
}

func replaceNulls(blob []byte) []byte {
	re, err := regexp.Compile("null(,?\\s*\n)")
	if err != nil {
		log.Fatal(err)
	}
	color := nulls("null")
	return re.ReplaceAll(blob, []byte(color+"$1"))
}
