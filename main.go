package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/drone/envsubst"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "-", "input file name, '-' means read from stdin")
}

func checkError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	var input io.ReadCloser
	if file == "-" {
		input = os.Stdin
	} else {
		f, err := os.Open(file)
		checkError(err)
		defer f.Close()
		input = f
	}
	inputData, err := io.ReadAll(input)
	checkError(err)
	result, err := envsubst.EvalEnv(string(inputData))
	checkError(err)
	fmt.Print(result)
}
