package main

import (
	"flag"
	"fmt"
	"os"

	json "github.com/bitly/go-simplejson"
)

var colors = flag.Bool("color", false, "colorize the output")

func init() {
	flag.BoolVar(colors, "c", false, "colorize the output")
	flag.Parse()
}

func getFileHandle() *os.File {
	var f *os.File
	args := flag.Args()
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe != 0 {
		f = os.Stdin
	} else if len(args) > 0 {
		f, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("missing file name")
		os.Exit(1)
	}
	return f
}

func main() {
	f := getFileHandle()
	j, err := json.NewFromReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := j.EncodePretty()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *colors {
		fmt.Printf("%s\n", colorize(res))
	} else {
		fmt.Printf("%s\n", res)
	}
}
