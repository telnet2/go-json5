package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/adhocore/jsonc"
)

var (
	indentedOutput bool
)

func main() {
	flag.BoolVar(&indentedOutput, "indented-output", false, "print out indented JSON output")
	flag.Parse()

	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	jc := jsonc.New()
	raw = jc.Strip(raw)

	buf := bytes.Buffer{}
	if indentedOutput {
		json.Indent(&buf, raw, "", "  ")
	} else {
		buf.Write(raw)
	}
	_, _ = io.Copy(os.Stdout, &buf)
}
